load("@rules_pact//private:provider.bzl", "ExampleInfo", "ContractInfo")

script_template="""\
#!/bin/bash
_healthCheck () {{
    echo "starting health check of $2"
    healthy="503"
    if [ $1 == "nop" ]; then healthy="200" && echo "health check ignored"; fi
    attempt=0
    until [ $healthy = "200" ]
    do
     healthy=$(curl -s -o /dev/null -w "%{{http_code}}" $1)
     echo "health check of $2 not ok, will recheck in 1 sec.."
     sleep 1
    done
}}
pwd
cp {libpact_ffi} $(dirname $(dirname {run_consumer_test}))
cp {libpact_ffi} .
ls .
echo "### Running Consumers Tests ###"
mkdir -p protobuf-0.3.5
cp {manifest} protobuf-0.3.5
cp {plugin} pact-protobuf-plugin
mv pact-protobuf-plugin protobuf-0.3.5
export PACT_PLUGIN_DIR=$(pwd)
./{run_consumer_test}
ls shape-app/api/pacts/
echo "### Running Providers Tests ###"
contract=$(dirname $(dirname {run_consumer_test}))/pacts/{contract}.json
pact_verifier_cli_args=$(cat {pact_verifier_cli_opts} || echo "--help")
side_car_cli_args=$(cat {side_car_opts} || echo "")
cli_args="$side_car_cli_args -f $contract $pact_verifier_cli_args"
echo $cli_args
while read first_line; read second_line
do
    export "$first_line"="$second_line"
done < {env_side_car}
nohup {provider_bin} &
echo "Provider started.."
nohup {side_car_bin} &
echo "State manager started.."
_healthCheck $(cat {health_check_side_car}) "side_car"
echo "Now running provider on $contract"
./{pact_verifier_cli} $cli_args
"""


def _pact_test_impl(ctx):
    pact_plugins = ctx.toolchains["@rules_pact//:pact_protobuf_plugin_toolchain_type"]
    pact_reference = ctx.toolchains["@rules_pact//:pact_reference_toolchain_type"]
    consumer = ctx.attr.consumer[DefaultInfo].default_runfiles.files.to_list()
    provider = ctx.attr.provider[DefaultInfo].default_runfiles.files.to_list()
    dict = {}
    for p in provider:
        dict.update({p.basename: p.short_path})
        if ctx.attr.provider[ExampleInfo].file == p.basename:
            dict.update({ctx.attr.provider[ExampleInfo].file: p.short_path})
        dict.update({"contract": ctx.attr.consumer[ContractInfo].name + "-" + ctx.attr.provider[ContractInfo].name})
    script_content = script_template.format(
        manifest = pact_plugins.manifest.short_path,
        plugin = pact_plugins.protobuf_plugin.short_path,
        run_consumer_test = consumer[0].short_path,
        libpact_ffi = pact_reference.libpact_ffi.short_path,
        pact_verifier_cli = pact_reference.pact_verifier_cli.short_path,
        pact_verifier_cli_opts = dict.setdefault("cli_args", "nop"),
        side_car_opts = dict.setdefault("side_car_cli_args", "nop"),
        provider_bin = dict.setdefault("cmd", "nop"),
        side_car_bin = dict.setdefault(ctx.attr.provider[ExampleInfo].file, "nop"),
        env_side_car = dict.setdefault("env_side_car","nop"),
        health_check_side_car = dict.setdefault("health_check_side_car", "nop"),
        contract = dict.setdefault("contract", "nop")
    )
    ctx.actions.write(ctx.outputs.executable, script_content)
    runfiles = ctx.runfiles(files = consumer + [pact_plugins.manifest, pact_plugins.protobuf_plugin, pact_reference.pact_verifier_cli, pact_reference.libpact_ffi, consumer[1]] + provider)

    return [DefaultInfo(runfiles = runfiles)]

pact_test = rule(
    implementation = _pact_test_impl,
    attrs = {
        "consumer": attr.label(),
        "provider": attr.label()
    },
    toolchains = ["@rules_pact//:pact_reference_toolchain_type", "@rules_pact//:pact_protobuf_plugin_toolchain_type"],
    test = True,
)

def _consumer_impl(ctx):
    srcs = ctx.attr.srcs[DefaultInfo].files_to_run.executable
    runfiles = ctx.runfiles(files = [srcs])
    runfiles = runfiles.merge(ctx.attr.data[0].data_runfiles)
    return [DefaultInfo(runfiles = runfiles),
            ContractInfo(name = ctx.attr.name)]

consumer = rule(
    implementation = _consumer_impl,
    doc = """Rule that wrap consumer interaction.
    It executes the test provided in srcs attribute through the toolchain.
    This rule will be executed from the pact_test rule.
    """,
    attrs = {
        "srcs": attr.label(
            allow_files = True,
            providers = [DefaultInfo],
            doc = "a test target"
        ),
        "data": attr.label_list(
            allow_files = True,
            doc = "data useful to provide with test target"
        ),
    },
)

def _provider_impl(ctx):

    args = ctx.actions.args()
    cli_args = ctx.actions.declare_file("cli_args")
    for k, v in ctx.attr.opts.items():
        args.add("--"+k, v)
    ctx.actions.write(cli_args, args)

    runfiles = ctx.runfiles(files = [cli_args] +
        ctx.attr.srcs[DefaultInfo].default_runfiles.files.to_list()
    )

    for dep in ctx.attr.deps:
        runfiles = runfiles.merge(dep[DefaultInfo].default_runfiles)

    return [DefaultInfo(runfiles = runfiles),
            ExampleInfo(file = dep[ExampleInfo].file),
            ContractInfo(name = ctx.attr.name)]
provider = rule(
    implementation = _provider_impl,
    doc = "Rule that describe provider interaction",
    attrs = {
        "srcs": attr.label(allow_files = True,
            providers = [DefaultInfo],
            doc = "the provider to run"
        ),
        "opts": attr.string_dict(
            doc = "options to provide to pact_verifier_cli"
        ),
        "deps": attr.label_list(
            allow_files = True,
            providers = [ExampleInfo],
            doc="any useful dep to run with the provider like a state-manager, a proxy or a side-car"
        ),
    },
)

def _side_car_impl(ctx):
    args = ctx.actions.args()
    cli_args = ctx.actions.declare_file("side_car_cli_args")
    for k, v in ctx.attr.opts.items():
       args.add("--"+k, v) if k != "state-change-teardown" else args.add("--"+k)

    ctx.actions.write(cli_args, args)
    bin = ctx.attr.srcs[DefaultInfo].files_to_run.executable

    env_args_file = ctx.actions.declare_file("env_side_car")
    env_args = ctx.actions.args()
    for k, v in ctx.attr.env.items():
        path = ctx.expand_location(v, ctx.attr.data)
        env_args.add(k, path)
    ctx.actions.write(env_args_file, env_args)

    health_check_file = ctx.actions.declare_file("health_check_side_car")
    ctx.actions.write(health_check_file, ctx.attr.health_check)

    runfiles = ctx.runfiles(files = [bin, cli_args, env_args_file, health_check_file] + ctx.files.data)
    return [DefaultInfo(runfiles = runfiles),
            ExampleInfo(file = bin.basename)
]

side_car = rule(
    implementation = _side_car_impl,
    attrs = {
        "srcs": attr.label(allow_files = True, providers = [DefaultInfo], doc = "the side-car to run"),
        "opts": attr.string_dict(
            doc = "the option specific to the side-car"
        ),
        "env": attr.string_dict(
            doc = "any environment variable to provide with the side_car"
        ),
        "data": attr.label_list(
            allow_files = True,
            doc = "any data useful to run with the side-car, like a configuration file for instance"
        ),
        "health_check": attr.string(default = "nop", doc = "uri to curl before launching provider test")
    },
)
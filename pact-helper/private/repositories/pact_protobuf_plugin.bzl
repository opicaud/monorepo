load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")
load("@rules_pact//private:versions.bzl", "CONSTRAINTS", "PLATFORMS", "PACT_PROTOBUF_PLUGINS_VERSIONS", "PACT_PROTOBUF_PLUGIN_JSON_VERSIONS", "DEFAULT_PACT_PROTOBUF_PLUGIN_VERSISON")

_PACT_TOOLCHAIN_BUILD_CONTENT = """\
load("@rules_pact//:defs.bzl", "pact_protobuf_plugin_toolchain")

pact_protobuf_plugin_toolchain(
    name = "toolchain_impl",
    protobuf_plugin = ":pact_protobuf_plugin_toolchain_bin_{platform}",
    manifest = ":pact_plugin_json_archive"
)

genrule(
    name = "pact_plugin_json_archive",
    outs = ["pact-plugin.json"],
    srcs = ["@pact_plugin_json_archive//file"],
    cmd = "cp $< $@",
)

genrule(
    name = "pact_protobuf_plugin_toolchain_bin_{platform}",
    outs = ["pact-protobuf-plugin_{platform}"],
    srcs = ["@pact_protobuf_plugin_archive_{platform}//file"],
    cmd = "gzip -d - < $< > $@",
)

toolchain(
    name = "toolchain",
    toolchain = ":toolchain_impl",
    toolchain_type = "@rules_pact//:pact_protobuf_plugin_toolchain_type",
    exec_compatible_with = {exec_compatible_with},
    target_compatible_with = {target_compatible_with}
)

"""

_PACT_WORKSPACE_CONTENT = """\
workspace(name = {})
"""

def repos(default_version = DEFAULT_PACT_PROTOBUF_PLUGIN_VERSISON):


    for platform in PACT_PROTOBUF_PLUGINS_VERSIONS[default_version].keys():
        value = PACT_PROTOBUF_PLUGINS_VERSIONS[default_version][platform]
        maybe(
            http_file,
            name = "pact_protobuf_plugin_archive_{platform}".format(platform = platform),
            sha256 = "{sha256}".format(sha256 = value.sha256),
            urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-{version}/pact-protobuf-plugin-{os}-{cpu}.gz".format(
                os = PLATFORMS[platform].os,
                cpu = PLATFORMS[platform].cpu,
                version = default_version
            )],
        )

        pact_plugins_toolchain_repository(
            name = "pact_protobuf_plugin_{os}_{cpu}_toolchain".format(
                os = PLATFORMS[platform].os,
                cpu = PLATFORMS[platform].cpu.replace("-", "_"),
            ),
            platform = platform,
            exec_compatible_with = CONSTRAINTS[platform],
            target_compatible_with = CONSTRAINTS[platform]
        )

    maybe(
        http_file,
        name = "pact_plugin_json_archive",
        sha256 = PACT_PROTOBUF_PLUGIN_JSON_VERSIONS[default_version].sha256,
        urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-{version}/pact-plugin.json".format(

            version = default_version
        )],
    )


def _pact_plugins_toolchain_repository_impl(repository_ctx):
    repository_ctx.file("BUILD.bazel", _PACT_TOOLCHAIN_BUILD_CONTENT.format(
        platform = repository_ctx.attr.platform,
        exec_compatible_with = repository_ctx.attr.exec_compatible_with,
        target_compatible_with = repository_ctx.attr.target_compatible_with
    ));
    repository_ctx.file("WORKSPACE.bazel", _PACT_WORKSPACE_CONTENT.format(repository_ctx.name));

pact_plugins_toolchain_repository = repository_rule(
    implementation = _pact_plugins_toolchain_repository_impl,
    attrs = {
        "platform": attr.string(
            doc = "Platform the pact-protobuf-plugin executable was built for.",
            mandatory = True,
        ),
        "exec_compatible_with": attr.string_list(mandatory = True),
        "target_compatible_with": attr.string_list(mandatory = True)
    },
)

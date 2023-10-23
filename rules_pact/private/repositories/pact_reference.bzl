load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")
load("@rules_pact//private:versions.bzl", "CONSTRAINTS", "PLATFORMS", "PACT_VERIFIER_CLI_VERSIONS", "PACT_VERIFIER_LIB_PACTFFI_VERSIONS")

_PACT_REFERENCE_BUILD_CONTENT = """\
load("@rules_pact//:defs.bzl", "pact_reference_toolchain")

pact_reference_toolchain(
    name = "pact_reference_toolchain_impl",
    pact_verifier_cli = ":pact_verifier_cli_toolchain_{platform}",
    libpact_ffi = ":pact_ffi_{platform}"
)

genrule(
    name = "pact_verifier_cli_toolchain_{platform}",
    outs = ["pact_verifier_cli_bin_{platform}"],
    srcs = ["@pact_verifier_cli_archive_{platform}//file"],
    cmd = "gzip -d - < $< > $@",
)

genrule(
    name = "pact_ffi_{platform}",
    outs = ["libpact_ffi.{ext}"],
    srcs = ["@pact_ffi_archive_{platform}//file"],
    cmd = "gzip -d - < $< > $@",
    visibility = ["//visibility:public"],
)

toolchain(
    name = "toolchain",
    toolchain = ":pact_reference_toolchain_impl",
    toolchain_type = "@rules_pact//:pact_reference_toolchain_type",
    exec_compatible_with = {exec_compatible_with},
    target_compatible_with = {target_compatible_with}
)
"""
_PACT_WORKSPACE_CONTENT = """\
workspace(name = {})
"""
DEFAULT_PACT_VERIFIER_CLI_VERSION = "1.0.1"
DEFAULT_PACTFFI_LIB_VERSION = "0.4.9"
def repos(pact_verifier_cli_version = DEFAULT_PACT_VERIFIER_CLI_VERSION, pactffi_lib_version = DEFAULT_PACTFFI_LIB_VERSION):

    for platform in PACT_VERIFIER_CLI_VERSIONS[pact_verifier_cli_version].keys():
        value = PACT_VERIFIER_CLI_VERSIONS[pact_verifier_cli_version][platform]
        maybe(
            http_file,
            name = "pact_verifier_cli_archive_{platform}".format(platform = platform),
            sha256 = "{sha256}".format(sha256 = value.sha256),
            urls = ["https://github.com/pact-foundation/pact-reference/releases/download/pact_verifier_cli-v{version}/pact_verifier_cli-{os}-{cpu}.gz".format(
                os = PLATFORMS[platform].os,
                cpu = PLATFORMS[platform].cpu,
                version = pact_verifier_cli_version)],
        )

    for platform in PACT_VERIFIER_LIB_PACTFFI_VERSIONS[pactffi_lib_version].keys():
        value = PACT_VERIFIER_LIB_PACTFFI_VERSIONS[pactffi_lib_version][platform]
        maybe(
            http_file,
            name = "pact_ffi_archive_{platform}".format(platform = platform),
            sha256 ="{sha256}".format(sha256 = value.sha256),
            urls = ["https://github.com/pact-foundation/pact-reference/releases/download/libpact_ffi-v{version}/libpact_ffi-{os}-{cpu}.{ext}.gz".format(
                os = PLATFORMS[platform].os,
                cpu = PLATFORMS[platform].cpu,
		ext = value.ext,
                version = pactffi_lib_version)],
        )
        if platform.startswith("darwin"):
            ext = "dylib"

        pact_reference_toolchain_repository(
            name = "pact_reference_{os}_{cpu}_toolchain".format(
                os = PLATFORMS[platform].os,
                cpu = PLATFORMS[platform].cpu.replace("-", "_"),
            ),
            ext = value.ext,
            platform = platform,
            exec_compatible_with = CONSTRAINTS[platform],
            target_compatible_with = CONSTRAINTS[platform]
        );

def _pact_reference_toolchain_repository_impl(repository_ctx):
    repository_ctx.file("BUILD.bazel", _PACT_REFERENCE_BUILD_CONTENT.format(
         platform = repository_ctx.attr.platform,
         ext = repository_ctx.attr.ext,
         exec_compatible_with = repository_ctx.attr.exec_compatible_with,
         target_compatible_with = repository_ctx.attr.target_compatible_with
     ));
    repository_ctx.file("WORKSPACE.bazel", _PACT_WORKSPACE_CONTENT.format(repository_ctx.name));

pact_reference_toolchain_repository = repository_rule(
    implementation = _pact_reference_toolchain_repository_impl,
    attrs = {
        "platform": attr.string(
            doc = "Platform the pact-reference executable was built for.",
            mandatory = False,
        ),
        "ext": attr.string(mandatory = True, default = "so"),
        "exec_compatible_with": attr.string_list(mandatory = True),
        "target_compatible_with": attr.string_list(mandatory = True)
    },
)

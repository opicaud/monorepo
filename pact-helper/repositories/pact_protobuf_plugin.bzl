load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")

_PACT_TOOLCHAIN_BUILD_CONTENT = """\
load("@pact-helper//:toolchains.bzl", "pact_protobuf_plugin_toolchain")

pact_protobuf_plugin_toolchain(
    name = "toolchain_impl",
    protobuf_plugin = ":pact_prototuf_plugin_toolchain_bin",
    manifest = ":pact_plugin_json_archive"
)

genrule(
    name = "pact_plugin_json_archive",
    outs = ["pact-plugin.json"],
    srcs = ["@pact_plugin_json_archive//file"],
    cmd = "cp $< $@",
)

genrule(
    name = "pact_prototuf_plugin_toolchain_bin",
    outs = ["pact-protobuf-plugin"],
    srcs = ["@pact_prototuf_plugin_archive//file"],
    cmd = "gzip -d - < $< > $@",
)

toolchain(
    name = "toolchain",
    toolchain = ":toolchain_impl",
    toolchain_type = "@pact-helper//:pact_protobuf_plugin_toolchain_type",
    exec_compatible_with = [
        "@platforms//os:macos",
        "@platforms//cpu:x86_64",
    ],
    target_compatible_with = [
        "@platforms//os:macos",
        "@platforms//cpu:x86_64",
    ],
)

"""

_PACT_WORKSPACE_CONTENT = """\
workspace(name = {})
"""

def repos():
    maybe(
        http_file,
        name = "pact_prototuf_plugin_archive",
        sha256 = "e3b09357c4ca793b7f0d78716ffe18916a7e72ed346ca549dfed79a4ff85cfc3",
        urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-0.3.5/pact-protobuf-plugin-osx-x86_64.gz"],
    )

    maybe(
        http_file,
        name = "pact_plugin_json_archive",
        sha256 = "70fa091ec6728d0077470d7ab1125be02b9b8211b73a552ea37f14e0276b7a52",
        urls = ["https://github.com/pactflow/pact-protobuf-plugin/releases/download/v-0.3.5/pact-plugin.json"],
    )

    helm_toolchain_repository(name = "protobuf_plugin_macos_x86_64_toolchain")

def _helm_toolchain_repository_impl(repository_ctx):
    repository_ctx.file("BUILD.bazel", _PACT_TOOLCHAIN_BUILD_CONTENT);
    repository_ctx.file("WORKSPACE.bazel", _PACT_WORKSPACE_CONTENT.format(repository_ctx.name));

helm_toolchain_repository = repository_rule(
    implementation = _helm_toolchain_repository_impl,
    attrs = {},
)

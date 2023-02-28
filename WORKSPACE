load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "56d8c5a5c91e1af73eca71a6fab2ced959b67c86d12ba37feedb0a2dfea441a6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version = "1.19")

go_rules_dependencies()

http_archive(
    name = "bazel_gazelle",
    sha256 = "448e37e0dbf61d6fa8f00aaa12d191745e14f07c31cabfa731f0c8e8a4f41b97",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:shape-app.bzl", "go_dependencies")

# gazelle:repository_macro shape-app.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl","git_repository")

#### PACT_PLUGINS ####
git_repository(
       name = "pact_plugins",
       remote = "https://github.com/opicaud/pact-protobuf-plugin",
       commit = "7f17a1697ee838b3059f56f326e8d905f9714d43",
       shallow_since = "1677603896 +0100"
)

load("@pact_plugins//:repositories.bzl", "repos")
repos()

load("@pact_plugins//:deps.bzl", "deps")
deps()

load("@pact_plugins//:create_crate.bzl", "create_crate_repositories")
create_crate_repositories()

#### PACT_FFI ####
git_repository(
       name = "pact_reference",
        remote = "https://github.com/opicaud/pact-reference",
        commit = "11db1a78c708a537f5a55f5cbc734a1de5618715",
        shallow_since = "1676577927 +0100",
        strip_prefix = "rust",
)

load("@pact_reference//:repositories.bzl", "repos")
repos()

load("@pact_reference//:deps.bzl", "deps")
deps()

load("@pact_reference//:create_crate.bzl", "create_crate_repositories")
create_crate_repositories()

load("@pact_reference//:create_pact_binaries.bzl", "create_pact_binaries")
create_pact_binaries("pact_bin", "pact_verifier_cli")



new_local_repository(
    name = "pact-helper",
    build_file = "pact-helper/BUILD.bazel",
    path = "pact-helper"
)

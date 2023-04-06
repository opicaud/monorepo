load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

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

git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "bazel_gazelle",
    sha256 = "448e37e0dbf61d6fa8f00aaa12d191745e14f07c31cabfa731f0c8e8a4f41b97",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:monorepo-deps.bzl", "go_dependencies")

# gazelle:repository_macro monorepo-deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()

#### PACT_PLUGINS ####
git_repository(
    name = "pact_plugins",
    commit = "2466ad7833a9ad6646ad8e0aabfb4ef32e086192",
    remote = "https://github.com/opicaud/pact-protobuf-plugin",
    shallow_since = "1677618560 +0100",
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
    commit = "93d658f566d62e07b8dd8f397a6d5f63348d14a3",
    remote = "https://github.com/opicaud/pact-reference",
    shallow_since = "1677795308 +0100",
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

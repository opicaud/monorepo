load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(version = "1.19")

go_rules_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "1add10f9bd92775b91f326da259f243881e904dd509367d5031d4c782ba82810",
    strip_prefix = "protobuf-3.21.9",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.21.9.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.21.9.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()



http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:monorepo-deps.bzl", "go_dependencies")

# gazelle:repository_macro monorepo-deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

http_archive(
    name = "io_bazel_rules_k8s",
    strip_prefix = "rules_k8s-0.5",
    urls = ["https://github.com/bazelbuild/rules_k8s/archive/v0.5.tar.gz"],
    sha256 = "773aa45f2421a66c8aa651b8cecb8ea51db91799a405bd7b913d77052ac7261a",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_defaults", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s_go_deps.bzl", k8s_go_deps = "deps")

k8s_go_deps()

[k8s_defaults(
    name = "k8s_" + kind,
    cluster = "docker-desktop",
    kind = kind,
) for kind in [
    "deployment",
    "service",
    "configmap",
    "job",
]]


http_archive(
    name = "rules_rust",
    sha256 = "d125fb75432dc3b20e9b5a19347b45ec607fabe75f98c6c4ba9badaab9c193ce",
    urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.17.0/rules_rust-v0.17.0.tar.gz"],
)

http_archive(
    name = "bazel-zig-cc",
    sha256 = "73afa7e1af49e3dbfa1bae9362438cdc51cb177c359a6041a7a403011179d0b5",
    strip_prefix = "bazel-zig-cc-v0.9.2",
    urls = ["https://git.sr.ht/~motiejus/bazel-zig-cc/archive/v0.9.2.tar.gz"]
)

load("@bazel-zig-cc//toolchain:defs.bzl", zig_toolchains = "toolchains")

zig_toolchains()

register_toolchains(
    "@zig_sdk//toolchain:linux_amd64_gnu.2.19",
)

load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains", "rust_repository_set")
rules_rust_dependencies()

rust_register_toolchains(
    edition = "2021",
    extra_target_triples = ["x86_64-unknown-linux-gnu"],
)


http_archive(
    name = "aspect_bazel_lib",
    sha256 = "3534a27621725fbbf1d3e53daa0c1dda055a2732d9031b8c579f917d7347b6c4",
    strip_prefix = "bazel-lib-1.16.1",
    url = "https://github.com/aspect-build/bazel-lib/archive/refs/tags/v1.16.1.tar.gz",
)

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")

aspect_bazel_lib_dependencies()

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

register_toolchains(
   "@zig_sdk//toolchain:linux_amd64_gnu.2.19",
)

load("@pact_plugins//:create_crate.bzl", "create_crate_repositories")

create_crate_repositories()

#### PACT_FFI ####
git_repository(
    name = "pact_reference",
    commit = "467a44233894d177c5df444d98873d3c920007f4",
    remote = "https://github.com/opicaud/pact-reference",
    shallow_since = "1683567089 +0200",
    strip_prefix = "rust",
)

load("@pact_reference//:repositories.bzl", "repos")

repos()

load("@pact_reference//:deps.bzl", "deps")

deps()

load("@pact_reference//:create_crate.bzl", "create_crate_repositories")

create_crate_repositories()

load("@pact_reference//:create_pact_binaries.bzl", "create_pact_binaries")

load("@io_bazel_rules_docker//container:container.bzl","container_pull")

container_pull(
    name = "debian_base",
    registry = "docker.io",
    repository = "debian:stable-slim",
    digest = "sha256:1529cbfd67815df9c001ed90a1d8fe2d91ef27fcaa5b87f549907202044465cb",
)

# See releases for urls and checksums
http_archive(
    name = "rules_helm",
    sha256 = "4593f521b30b976c1f02932211b705c220615d8940f0c6d35daa07ab060f97d8",
    urls = ["https://github.com/abrisco/rules_helm/releases/download/0.0.3/rules_helm-v0.0.3.tar.gz"],
)

load("@rules_helm//helm:repositories.bzl", "helm_register_toolchains", "rules_helm_dependencies")

rules_helm_dependencies()

helm_register_toolchains()

load("@rules_helm//helm:repositories_transitive.bzl", "rules_helm_transitive_dependencies")

rules_helm_transitive_dependencies()
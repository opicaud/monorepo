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
        "https://github.com/protocolbuffers/protobuf/archive/refs/tags/v3.21.9.tar.gz",
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
    name = "rules_rust",
    sha256 = "4a9cb4fda6ccd5b5ec393b2e944822a62e050c7c06f1ea41607f14c4fdec57a2",
    urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.25.1/rules_rust-v0.25.1.tar.gz"],
)

http_archive(
    name = "hermetic_cc_toolchain",
    sha256 = "57f03a6c29793e8add7bd64186fc8066d23b5ffd06fe9cc6b0b8c499914d3a65",
    urls = [
        "https://mirror.bazel.build/github.com/uber/hermetic_cc_toolchain/releases/download/v2.0.0/hermetic_cc_toolchain-v2.0.0.tar.gz",
        "https://github.com/uber/hermetic_cc_toolchain/releases/download/v2.0.0/hermetic_cc_toolchain-v2.0.0.tar.gz",
    ],
)

load("@hermetic_cc_toolchain//toolchain:defs.bzl", zig_toolchains = "toolchains")

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


git_repository(
    name = "aspect_bazel_lib",
    commit = "794df714d7efbf5f2b986470428bea311f4fd772",
    shallow_since = "1687478984 -0700",
    remote = "https://github.com/aspect-build/bazel-lib.git"
)

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")

aspect_bazel_lib_dependencies()

load("//pact-helper:repositories.bzl", "repos")
repos()
register_toolchains("//pact-helper:toolchain")

register_toolchains(
   "@zig_sdk//toolchain:linux_amd64_gnu.2.19",
)


#### PACT_FFI ####
http_archive(
    name = "pact_reference",
    strip_prefix = "pact-reference-pact-reference-rust-v1.3.1/rust",
    sha256 = "2c53b9da8bb8ca8f55ac4b2405e676ec41fe3150fa4beb2686772015ec9fcce4",
    url = "https://github.com/opicaud/pact-reference/archive/refs/tags/pact-reference-rust-v1.3.1.tar.gz",
)

load("@pact_reference//:repositories.bzl", "repos")

repos()

load("@pact_reference//:deps.bzl", "deps")

deps("cargo-bazel-lock-pact-reference.json")
register_toolchains("@pact_reference//:toolchain")

load("@pact_reference//:create_crate.bzl", "create_crate_repositories")

create_crate_repositories()

load("@pact_reference//:create_pact_binaries.bzl", "create_pact_binaries")

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

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "rules_pkg",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
    ],
    sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
)
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")
rules_pkg_dependencies()

http_archive(
    name = "aspect_rules_js",
    sha256 = "d8827db3c34fe47607a0668e86524fd85d5bd74f2bfca93046d07f890b5ad4df",
    strip_prefix = "rules_js-1.27.0",
    url = "https://github.com/aspect-build/rules_js/releases/download/v1.27.0/rules_js-v1.27.0.tar.gz",
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

load("@rules_nodejs//nodejs:repositories.bzl", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "nodejs",
    node_version = "18.13.0",
)

load("@aspect_rules_js//npm:repositories.bzl", "npm_translate_lock")

npm_translate_lock(
     name = "npm",
     data = ["//hack:package.json"],
     pnpm_lock = "//hack:pnpm-lock.yaml",
     update_pnpm_lock = True,
     #verify_node_modules_ignored = "//:.bazelignore",
)

load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_oci",
    sha256 = "db57efd706f01eb3ce771468366baa1614b5b25f4cce99757e2b8d942155b8ec",
    strip_prefix = "rules_oci-1.0.0",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.0.0/rules_oci-v1.0.0.tar.gz",
)

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "LATEST_ZOT_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    #zot_version = LATEST_ZOT_VERSION,
)


load("@rules_oci//oci:pull.bzl", "oci_pull")

oci_pull(
    name = "distroless_debian",
    digest = "sha256:1529cbfd67815df9c001ed90a1d8fe2d91ef27fcaa5b87f549907202044465cb",
    image = "debian",
)

oci_pull(
    name = "distroless_go",
    digest = "sha256:0530d193888bcd7bd0376c8b34178ea03ddb0b2b18caf265135b6d3a393c8d05",
    image = "gcr.io/distroless/base",
)


http_archive(
    name = "bazel_skylib",
    sha256 = "b8a1527901774180afc798aeb28c4634bdccf19c4d98e7bdd1ce79d1fe9aaad7",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.4.1/bazel-skylib-1.4.1.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.1/bazel-skylib-1.4.1.tar.gz",
    ],
)

git_repository(
    name = "container_structure_test",
    commit = "104a53ede5f78fff72172639781ac52df9f5b18f",
    shallow_since = "1683241066 -0400",
    remote = "https://github.com/GoogleContainerTools/container-structure-test.git",
)

load("@container_structure_test//:repositories.bzl", "container_structure_test_register_toolchain")

container_structure_test_register_toolchain(name = "cst")

load("@aspect_bazel_lib//lib:repositories.bzl", "register_yq_toolchains")

register_yq_toolchains()

git_repository(
    name = "environment_secrets",
    commit = "103b222eba64355b93649b06ecfe3844466b5a65",
    shallow_since = "1537893432 -0600",
    remote = "https://github.com/solarhess/rules_build_secrets.git",
)

load("@environment_secrets//lib:secrets.bzl","environment_secrets")

environment_secrets(
    name="env",
    entries = {
        "GH_TOKEN": "default",
    },
)
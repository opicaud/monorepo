"""# rules_pact
Bazel rules to test your services interactions via [pacts][pacts_web_site]
[pact_web_site]: https://docs.pact.io/
"""
load("@pact-helper//:toolchains.bzl", _pact_reference_toolchain = "pact_reference_toolchain", _pact_protobuf_plugin_toolchain = "pact_protobuf_plugin_toolchain")

pact_reference_toolchain = _pact_reference_toolchain
pact_protobuf_plugin_toolchain = _pact_protobuf_plugin_toolchain
repos = _repos
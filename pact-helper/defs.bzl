"""# rules_pact
Bazel rules to test services interactions with [pacts][pactsws]

[pactsws]: https://docs.pact.io/
"""
load("@pact-helper//private:toolchains.bzl", _pact_reference_toolchain = "pact_reference_toolchain", _pact_protobuf_plugin_toolchain = "pact_protobuf_plugin_toolchain")

pact_reference_toolchain = _pact_reference_toolchain
pact_protobuf_plugin_toolchain = _pact_protobuf_plugin_toolchain
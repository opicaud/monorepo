"""# rules_pact
Bazel rules to test services interactions with [pacts][pactsws]

[pactsws]: https://docs.pact.io/

- [consumer](#consumer)
- [provider](#provider)
- [side_car](#side_car)
- [pact_test](#pact_test)
- [pact_reference_toolchain](#pact_reference_toolchain)
- [pact_protobuf_plugin_toolchain](#pact_protobuf_plugin_toolchain)
"""
load("@rules_pact//private:toolchains.bzl", _pact_reference_toolchain = "pact_reference_toolchain", _pact_protobuf_plugin_toolchain = "pact_protobuf_plugin_toolchain")
load("@rules_pact//private:consumer.bzl", _pact_test = "pact_test", _consumer = "consumer", _provider = "provider", _side_car = "side_car")

pact_reference_toolchain = _pact_reference_toolchain
pact_protobuf_plugin_toolchain = _pact_protobuf_plugin_toolchain
pact_test = _pact_test
consumer = _consumer
provider = _provider
side_car = _side_car
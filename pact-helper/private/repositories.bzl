load("@rules_pact//private/repositories:pact_protobuf_plugin.bzl", _repos_pact_protobuf_plugin = "repos")
load("@rules_pact//private/repositories:pact_reference.bzl", _repos_pact_reference = "repos",)
def repos():
    _repos_pact_protobuf_plugin()
    _repos_pact_reference()
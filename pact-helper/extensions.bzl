load("@pact-helper//:repositories.bzl", "repos")

def _impl(ctx):
    repos()

options = tag_class(attrs={})
pact_helper = module_extension(
    implementation = _impl,
    tag_classes = {"options": options},
)
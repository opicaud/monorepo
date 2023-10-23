load("@rules_pact//private:repositories.bzl", "repos")

def _impl(ctx):
    repos()

options = tag_class(attrs={})
rules_pact = module_extension(
    implementation = _impl,
    tag_classes = {"options": options},
)
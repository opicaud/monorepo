"""pact_protobuf_plugin toolchain implementation"""

def _pact_protobuf_plugin_toolchain_impl(ctx):
    return platform_common.ToolchainInfo(
        protobuf_plugin = ctx.file.protobuf_plugin,
        manifest = ctx.file.manifest
    )

pact_protobuf_plugin_toolchain = rule(
    implementation = _pact_protobuf_plugin_toolchain_impl,
    doc = "A pact protobuf plugin toolchain",
    attrs = {
        "protobuf_plugin": attr.label(
            doc = "A pact protobuf plugin binary",
            allow_single_file = True,
            mandatory = True,
        ),
        "manifest": attr.label(
            doc = "A json manifest",
            allow_single_file = True,
            mandatory = True,
        )
    },
)
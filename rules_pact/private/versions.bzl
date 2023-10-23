DEFAULT_PACT_PROTOBUF_PLUGIN_VERSISON="0.3.5"
PACT_PROTOBUF_PLUGINS_VERSIONS = {
    "0.3.5": {
        "darwin_amd64": struct(sha256 = "e3b09357c4ca793b7f0d78716ffe18916a7e72ed346ca549dfed79a4ff85cfc3"),
        "linux_amd64": struct(sha256 = "9a0058f10e2eae9375c019453491897ac3eb87bb87b017c311d5d1ec3a4eb979"),
    }
}
PACT_PROTOBUF_PLUGIN_JSON_VERSIONS = {
     "0.3.5": struct(sha256 = "70fa091ec6728d0077470d7ab1125be02b9b8211b73a552ea37f14e0276b7a52"),
}

PACT_VERIFIER_CLI_VERSIONS = {
    "1.0.1": {
        "darwin_amd64": struct(sha256 = "77ffc38f4564cfef42f64b9eb33bebfc4d787e65ef7ff7213640a3d63d2cf5a7"),
        "linux_amd64": struct(sha256 = "57c8ae7c95f46e4a48d3d6a251853dd5dd58917e866266ced665fc48a3fdecdd"),
    }
}
PACT_VERIFIER_LIB_PACTFFI_VERSIONS = {
    "0.4.9": {
        "darwin_amd64": struct(sha256 = "b8c87e2cc2f83ae9e79678d3288f2f9f7cea27d023576f565d8a203441600a59", ext = "dylib"),
        "linux_amd64": struct(sha256 = "86d8b82ab0843909642bec8f3a1bea702bbe65f3665de18f024fdfdf62b8cf0c", ext = "so"),
    }
}

CONSTRAINTS = {
    "darwin_amd64": ["@platforms//os:macos", "@platforms//cpu:x86_64"],
    "linux_amd64": ["@platforms//os:linux", "@platforms//cpu:x86_64"],
}

PLATFORMS = {
    "darwin_amd64": struct(os = "osx", cpu = "x86_64"),
    "linux_amd64": struct(os = "linux", cpu = "x86_64")
}
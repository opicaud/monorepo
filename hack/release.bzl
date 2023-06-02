def release_me(**kwargs):
    native.sh_binary(
        name = "release_me",
        srcs = ["//hack:release.sh"],
        args = ["$(location :package.json)"],
        data = ["//hack:semantic_release_binary", ":package.json" ],
        **kwargs

)
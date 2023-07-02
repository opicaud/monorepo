load("@aspect_bazel_lib//lib:jq.bzl", "jq")

def release_me(**kwargs):
    native.sh_binary(
        name = "release_me",
        srcs = ["//hack:release.sh"],
        args = ["$(location :package.json)"],
        data = ["//hack:semantic_release_binary", ":package.json" ],
        **kwargs
    )
    jq(
        name = "no_srcs",
        srcs = ["package.json"],
        filter = ".name",
    )


    native.genrule(
        name = "find-what-next-releases-versions-are",
        srcs = ["//hack:semantic_release_binary", ":no_srcs"],
        outs = ["next-version-to-release"],
        cmd = "./$(location //hack:find-what-next-releases-are.sh) $(location //hack:semantic_release_binary) $(location :no_srcs) > \"$@\"",
        tools = [
            "//hack:find-what-next-releases-are.sh",
        ],
        visibility = ["//visibility:private"],
    )
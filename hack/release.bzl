load("@aspect_bazel_lib//lib:jq.bzl", "jq")
load("@aspect_bazel_lib//lib:run_binary.bzl", "run_binary")

def release_me(**kwargs):
    native.sh_binary(
        name = "release_me",
        srcs = ["//hack:release.sh"],
        args = ["$(location :package.json)", "$(location :do-i-need-to-be-released)"],
        data = ["//hack:semantic_release_binary", ":package.json", ":do-i-need-to-be-released" ],
        **kwargs
    )
    jq(
        name = "no_srcs",
        srcs = ["package.json"],
        filter = ".name",
    )


    run_binary(
        name = "find-what-next-releases-versions-are",
        env = {
            "GH_TOKEN": "$(GH_TOKEN)",
            "OUT": "$(location next-version-to-release)",
            },
        srcs = ["//hack:semantic_release_binary", ":no_srcs", ":package.json","//hack:get-latest-tags"],
        outs = ["next-version-to-release"],
        args = ["$(location //hack:semantic_release_binary)","$(location :no_srcs)","$(location :package.json)", "$(location //hack:get-latest-tags)" ],
        tool = "//hack:find-what-next-releases-are.sh",
        visibility = ["//visibility:private"],
    )

    native.genrule(
        name = "do-i-need-to-be-released",
        srcs = ["//hack:do-i-need-to-be-released.sh", "//hack:get-latest-tags", ":no_srcs", ":find-what-next-releases-versions-are"],
        outs = ["will-be-released"],
        cmd = "./$(location //hack:do-i-need-to-be-released.sh) $(location //hack:get-latest-tags) $(location :no_srcs) $(location :find-what-next-releases-versions-are) > \"$@\"",
        tools = [
            "//hack:do-i-need-to-be-released.sh",
        ],
        visibility = ["//visibility:private"],
    )

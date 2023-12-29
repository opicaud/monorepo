load("@aspect_bazel_lib//lib:jq.bzl", "jq")
load("@aspect_bazel_lib//lib:run_binary.bzl", "run_binary")
load("@env//:secrets.bzl","GH_TOKEN")

def release_me(**kwargs):
    if (GH_TOKEN != "default"):
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
            srcs = [":no_srcs", ":package.json"],
            outs = ["next-version-to-release"],
            cmd = "GH_TOKEN={0} ./$(locations //hack:find-what-next-releases-are.sh) $(locations //hack:semantic_release_binary) $(location :no_srcs) $(location :package.json) > \"$@\"".format(GH_TOKEN),
            tools = ["//hack:find-what-next-releases-are.sh", "//hack:semantic_release_binary"],
            visibility = ["//visibility:private"],
        )
    else:
        print("Release management is deactivated")

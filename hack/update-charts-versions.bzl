load("@aspect_bazel_lib//lib:yq.bzl", "yq")
load("@aspect_bazel_lib//lib:write_source_files.bzl", "write_source_files")

def update_me(appVersion, version, **kwargs):
    yq(
        name = "prepare_upgrade_chart_app_version",
        srcs = ["Chart.yaml"],
        expression = "|".join([
            "load(strenv(STAMP)) as $stamp",
            "(.appVersion) = ($stamp.{0} // \"<unstamped>\")".format(appVersion),
        ]),
        visibility = ["//visibility:private"],
    )

    yq(
        name = "prepare_upgrade_chart_version",
        srcs = [":prepare_upgrade_chart_app_version"],
        expression = "|".join([
            "load(strenv(STAMP)) as $stamp",
            "(.version) = ($stamp.{0} // \"0.1.0\")".format(version),
        ]),
        visibility = ["//visibility:private"],
    )

    write_source_files(
        name = "upgrade_chart_app_version",
        in_file = ":prepare_upgrade_chart_app_version",
        out_file = "Chart.yaml",
        visibility = ["//visibility:private"],
    )

    write_source_files(
        name = "upgrade_chart_version",
        in_file = ":prepare_upgrade_chart_version",
        out_file = "Chart.yaml",
        visibility = ["//visibility:private"],
    )

    write_source_files(
        name = "apply_chart",
        additional_update_targets = [
            ":upgrade_chart_app_version",
            ":upgrade_chart_version",
        ]
    )
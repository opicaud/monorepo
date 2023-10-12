CONSTRAINTS = {
    "darwin_amd64": ["@platforms//os:macos", "@platforms//cpu:x86_64"],
    "linux_amd64": ["@platforms//os:linux", "@platforms//cpu:x86_64"],
}

PLATFORMS = {
    "darwin_amd64": struct(os = "osx", cpu = "x86_64"),
    "linux_amd64": struct(os = "linux", cpu = "x86_64")
}
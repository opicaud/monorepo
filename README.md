[![monorepo](https://github.com/opicaud/monorepo/actions/workflows/main.yml/badge.svg)](https://github.com/opicaud/monorepo/actions/workflows/main.yml)
### Start with monorepo
1. Clone the repository
2. Install bazel
3. Build with `bazel build //...`  

### Goal
Providing a technical way to enhance collaboration within a team or within multi-team by implementing a monorepo.

### Constraints
- build should be agnostic from plaforms and architecture to ensure reproductible build
- 0 extra tools should be necessary to enhance fast collaboration and onboarding within teams
- build/test/relase should take less than 10 minutes to provide fast feedback to the team

### Currently
- build has been succesfully ran on macos-darwin64 and linux-amd64 platforms
- no extra tools needs to be installed to build (a part Bazel of course)
- the build is taking around 10 minutes, the most difficult part is to ensure cache usages (a part the first build)

### Stats
| #Build                                                                              | Cache | Time   | Build type         | Cache time |
|-------------------------------------------------------------------------------------|-------|--------|--------------------|------------|
| [#133](https://github.com/opicaud/monorepo/actions/runs/5685417265/job/15410258206) | No    | 32m18s | Build from scratch | 0%         |
| [#134](https://github.com/opicaud/monorepo/actions/runs/5688779558/job/15419186110) | Yes   | 4m50   | Patch changes      | 50%        |

- _Patch changes_ are defined by changes that did not trigger any new releases of any components of the monorepo
- _Cache time_ is defined by the proportion of time to fetch the cache over the total time of the build
### Story
`shape-app` is sending events about area calculation to `eventstore-app`, two grpc monorepo components deployed via helm charts.
Their collaboration and integration is tested via [Pact](https://docs.pact.io/) during the build, thanks to a mononorepo component called pact-helper
Each monorepo components are released via a Bazel macro called `release_me()` such as OCI images are also tagged and pushed during the release process
Also, Helm charts are components and are also released each time a new image or a update is made in their manifests
In the end, the updated apps folder is used as an input to a GitOps platform in order to be deployed

### Next objective
- Migrate to Bazel 6 with bzlmod
- Provide a cli for shape-app

### Thanks :
- [#rules_go](https://github.com/bazelbuild/rules_go)
- [#bazel_gazelle](https://github.com/bazelbuild/bazel-gazelle)
- [#rules_rust](https://github.com/bazelbuild/rules_rust)
- [#hermetic_cc_toolchain](https://github.com/uber/hermetic_cc_toolchain)
- [#aspect-bazel-lib](https://github.com/aspect-build/bazel-lib)
- [#pact-plugins](https://docs.pact.io/plugins/quick_start)
- [#pact-reference](https://github.com/pact-foundation/pact-reference)
- [#rules_helm](https://github.com/abrisco/rules_helm)
- [#rules_pkg](https://github.com/bazelbuild/rules_pkg)
- [#aspect_rules_js](https://github.com/aspect-build/rules_js)
- [#rules_oci](https://github.com/bazel-contrib/rules_oci)

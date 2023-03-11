###Start with monorepo
1. Clone the repository
2. Install bazel
3. Build with `bazel build //...`  

Listed below module will be built (first build is taking a while)
####Module pact-helper
This module is used to provide `pact_ffi` library, `pact-verifier` cli and `pact-protobuf`pact-plugins to run pact-tests in client modules
####Module shape-app
This go module provide a grpc server to calculate shapes areas  
It is divided in packages `api`, `domain`, `infra`, `test`
#####api
* `api/cmd`: package containing the grpc server to run, pact is used to check the consumer contract
* `api/pacts`: package containing pact description, use to execute consumer test
* `api/proto`: package containing protobuf description
#####domain
* `domain/adapter`: package containing bridge with infra, should be migrate to infra module
* `domain/shape`: package containing the shape domain, running with CQRS and event-store (in memory or grpc)
#####infra
This module needs to be reworked
#####test
This module contains BDD Tests related to the monorepo

###Adding features
The monorepo is using `pre-commit`, `commit-msg`, `pre-push` hooks.
Each module can define its own hook, to be run at the right moment in time in the development of a feature
Eg: `shape-app` is running unit-test and go lint before each commit and pact-test before pushing new commit
Also, `shape-app`, like all other modules of the monorepo is using a commit-lint to respect commit convention  
Those services are giving by `mookme`
###Build
The monorepo is built via `bazel`, through the unique `WORKSPACE` file
* `shape-app` is using `gazelle` to automatically generate `BUILD` files
* `pact-helper`: home-made `BUILD` file using forks of `pact_ffi` and `pact-protobuf-plugins`

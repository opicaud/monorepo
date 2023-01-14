package pacts

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"log"
	"os"
)

func setEnvVarPactPluginDir() {
	r, err := runfiles.New()
	if err != nil {
		log.Printf("Bazel not present, use PACT_PLUGIN_DIR: %s\n", os.Getenv("PACT_PLUGIN_DIR"))
		return
	}
	path := os.Getenv("PACT_PLUGINS")
	pactPluginDr, err := r.Rlocation(path)
	log.Printf("PACT_PLUGIN_DIR: %s\n", pactPluginDr)
	if err != nil {
		log.Fatalf("path %s not found", path)
	}
	os.Setenv("PACT_PLUGIN_DIR", "/Users/opicaud/dev/microservices/go/monorepo/example2/bazel-out/darwin-fastbuild/bin/api/pacts/pacts_test_/pacts_test.runfiles/__main__/external/pact-plugins/plugins")

}

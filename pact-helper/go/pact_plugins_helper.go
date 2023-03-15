package pacts

import (
	"github.com/bazelbuild/rules_go/go/runfiles"
	"log"
	"os"
	"path/filepath"
)

func setEnvVarPactPluginDir() {
	r, err := runfiles.New()
	if err != nil {
		log.Printf("Bazel not present, use PACT_PLUGIN_DIR: %s\n", os.Getenv("PACT_PLUGIN_DIR"))
		return
	}

	path := os.Getenv("PACT_PLUGINS")
	pactPluginDr, err := r.Rlocation(path)
	_ = os.Setenv("PACT_PLUGIN_DIR", filepath.Dir(pactPluginDr))

	log.Printf("PACT_PLUGIN_DIR: %s", filepath.Dir(pactPluginDr))
	if err != nil {
		log.Fatalf("path %s not found", path)
	}

}

package pacts

import (
	"context"
	"fmt"
	"github.com/bazelbuild/rules_go/go/runfiles"
	ac "github.com/opicaud/monorepo/shape-app/api/proto"
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

var dir, _ = filepath.Abs("../proto/app_shape.proto")

func TestCreateShape(t *testing.T) {
	rlocation, err2 := runfiles.Rlocation(os.Getenv("SHAPEAPP_PROTO_FILE"))
	if err2 != nil {
		rlocation = dir
	}
	grpcInteraction := `{
		"pact:proto": "` + rlocation + `",
		"pact:proto-service": "Shapes/create",
		"pact:content-type": "application/protobuf",
		"request": {
			"shapes": 
				{
					"shape": "rectangle",
					"dimensions": [ "matching(number, 3)", "matching(number, 4)"]
				}
			
		},
		"response": {
			"message": { "code": "matching(number, 0)"}
		}
	}`
	SetEnvVarPactPluginDir()
	F := func(transport message.TransportConfig, m message.SynchronousMessage) error {
		dimensions := []float32{3, 4}
		shape := "rectangle"
		rectangle := ac.ShapeMessage{Shape: shape, Dimensions: dimensions}
		request := &ac.ShapeRequest{Shapes: &rectangle}
		area, err := getRectangleAndSquareArea(fmt.Sprintf("localhost:%d", transport.Port), request)

		assert.NoError(t, err)
		assert.Equal(t, uint32(0), area.GetCode())

		return err
	}

	var mockProvider, _ = message.NewSynchronousPact(message.Config{
		Consumer: "grpc-consumer-go",
		Provider: "area-calculator-provider",
	})

	_ = mockProvider.AddSynchronousMessage("calculate rectangle area request").
		UsingPlugin(message.PluginConfig{
			Plugin: "protobuf",
		}).
		WithContents(grpcInteraction, "application/grpc").
		StartTransport("grpc", "127.0.0.1", nil).
		ExecuteTest(t, F)

}

func getRectangleAndSquareArea(address string, request *ac.ShapeRequest) (*ac.Message, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := ac.NewShapesClient(conn)

	log.Println("Sending calculate rectangle and square request")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, request)

	if err != nil {
		return nil, err
	}
	return r.GetMessage(), nil
}

func SetEnvVarPactPluginDir() {
	r, err := runfiles.New()
	if err != nil {
		log.Printf("Bazel not present, use PACT_PLUGIN_DIR: %s\n", os.Getenv("PACT_PLUGIN_DIR"))
		return
	}

	zipPath := os.Getenv("PACT_PLUGINS_ZIP")
	zipFileLocation, _ := r.Rlocation(zipPath)
	pluginDir := unzip(zipFileLocation)
	_ = os.Setenv("PACT_PLUGIN_DIR", pluginDir)
	log.Printf("PACT_PLUGIN_DIR: %s", pluginDir)

	if err != nil {
		log.Fatalf("path %s not found", zipPath)
	}

}

func unzip(zipFile string) string {
	version := module.Version{
		Path:    "pact.plugins.protobuf",
		Version: "v0.3.1",
	}

	if _, err := os.Stat("./protobuf-0.3.1"); os.IsNotExist(err) {
		log.Println("Unzipping plugins..")
		uz := zip.Unzip("protobuf-0.3.1", version, zipFile)
		if uz != nil {
			log.Panic(uz.Error())
		}
		err := os.Chmod("./protobuf-0.3.1/pact-protobuf-plugin", 0777)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Printf("Plugins already present, skipping unzip process..")
	}
	getwd, _ := os.Getwd()
	return getwd

}

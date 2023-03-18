#!/usr/bin/env bash

nohup ./shape-app/api/cmd/cmd_/cmd &
echo $! > pid
export PACT_PLUGIN_DIR=$(dirname "${RUNFILES_DIR}"/"${PACT_PLUGINS}")
 ./external/pact_bin/pact_verifier_cli -f shape-app/api/pacts/pacts/grpc-consumer-go-area-calculator-provider.json --transport grpc -p 50051

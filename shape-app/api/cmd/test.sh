#!/usr/bin/env bash
unzip -o pact-helper/pact_plugins_zip.zip
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/$(dirname ${PACT_PLUGIN_DIR}))

nohup $1 &
sleep 10
./external/pact_reference/pact_verifier_cli/pact_verifier_cli -f "$2" --transport grpc -p 50051

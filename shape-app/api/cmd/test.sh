#!/usr/bin/env bash
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/${PACT_PLUGIN_DIR})
nohup $1 &
sleep 10
./external/pact_reference/pact_verifier_cli/pact_verifier_cli -f "$2" --transport grpc -p 50051

#!/usr/bin/env bash

echo $1, $2
nohup $1 &
echo $! > pid
export PACT_PLUGIN_DIR=$(${RUNFILES_DIR}/${PACT_PLUGINS}pact-helper)
 ./external/pact_bin/pact_verifier_cli -f $2 --transport grpc -p 50051

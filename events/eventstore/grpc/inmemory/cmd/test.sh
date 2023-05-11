#!/usr/bin/env bash
unzip -o pact-helper/pact_plugins_zip.zip
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/$(dirname ${PACT_PLUGIN_DIR}))

consumer_contract=$3
state_manager=$2
event_store=$1


nohup "$event_store" &
sleep 3
nohup "$state_manager" &
sleep 3
./external/pact_reference/pact_verifier_cli/pact_verifier_cli --state-change-url http://localhost:8080/event --state-change-teardown -f "${consumer_contract}" --transport grpc -p 50052

#!/usr/bin/env bash
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/${PACT_PLUGIN_DIR})

consumer_contract=$3
state_manager=$2
event_store=$1


nohup "$state_manager" &
nohup "$event_store" &
sleep 2
./external/pact_bin/pact_verifier_cli --state-change-url http://localhost:8080/event --state-change-teardown -f "${consumer_contract}" --transport grpc -p 50052

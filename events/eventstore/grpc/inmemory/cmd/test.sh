#!/usr/bin/env bash

consumer_contract=$3
state_manager=$2
event_store=$1

nohup "$state_manager" &
nohup "$event_store" &

./external/pact_bin/pact_verifier_cli --state-change-url http://localhost:8080/event -f "${consumer_contract}" --transport grpc -p 50051


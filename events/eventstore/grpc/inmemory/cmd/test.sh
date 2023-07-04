#!/usr/bin/env bash
ls pact-helper
unzip -o pact-helper/pact_plugins_zip.zip
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/$(dirname ${PACT_PLUGIN_DIR}))

consumer_contract=$3
state_manager=$2
event_store=$1


nohup "$event_store" &
nohup "$state_manager" &
healthy="503"
attempt=0
until [ $healthy = "200" ] || [ $attempt = 20 ]
do
 healthy=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8081/healthz)
 echo "not_ok, attempt: $attempt"
 sleep 1
 attempt=$(( attempt + 1))
done

./external/pact_reference/pact_verifier_cli/pact_verifier_cli --state-change-url http://localhost:8081/event --state-change-teardown -f "${consumer_contract}" --transport grpc -p 50052

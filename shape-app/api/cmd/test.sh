#!/usr/bin/env bash
unzip -o pact-helper/pact_plugins_zip.zip
export PACT_PLUGIN_DIR=$(dirname ${RUNFILES_DIR}/$(dirname ${PACT_PLUGIN_DIR}))

nohup $1 &
nohup $2 &
healthy="503"
attempt=0
until [ $healthy = "200" ] || [ $attempt = 10 ]
do
 healthy=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/healthz)
 echo "not_ok, attempt: $attempt"
 sleep 1
 attempt=$(( attempt + 1))
done
./external/pact_reference/pact_verifier_cli/pact_verifier_cli -f "$3" --transport grpc -p 50051

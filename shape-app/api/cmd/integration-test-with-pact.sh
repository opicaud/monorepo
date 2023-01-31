#!/usr/bin/env bash
./shape-app/api/pacts/pacts_test_/pacts_test
nohup ./shape-app/api/cmd/cmd_/cmd &
echo $! > pid
./external/pact-helper/verifier/pact_verifier -f shape-app/api/pacts/pacts/grpc-consumer-go-area-calculator-provider.json --transport grpc -p 50051

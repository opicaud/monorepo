#!/usr/bin/env bash
./api/pacts/pacts_test_/pacts_test
nohup ./api/cmd/cmd_/cmd &
echo $! > pid
./external/pact-verifier/pact_verifier -f api/pacts/pacts/grpc-consumer-go-area-calculator-provider.json --transport grpc -p 50051

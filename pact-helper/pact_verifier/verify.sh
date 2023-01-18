#!/usr/bin/env bash
echo "contract is: $1"
echo "port is: $PORT"
if test -f "$1"; then
  ./pact_verifier -f "$1" --transport grpc -p "${PORT}"
else
  echo "WARN: tests are skipped, contract $1 has not been found"
fi


#!/usr/bin/env bash

. ./events/eventstore/grpc/test/helper/k8s.apply
. ./events/eventstore/grpc/inmemory/k8s/k8s.apply
. ./events/eventstore/grpc/test/pact/k8s.apply

/usr/local/bin/kubectl wait --for=condition=complete job.batch/test-grpc-eventstore-as-provider --timeout=15s

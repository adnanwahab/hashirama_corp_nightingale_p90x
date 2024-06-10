#!/usr/bin/env sh

#!/bin/bash

~/go/bin/air &

AIR_PID=$!

cleanup() {
    kill -SIGINT $AIR_PID
    exit 0
}

trap cleanup EXIT SIGINT SIGTERM

wait $AIR_PID

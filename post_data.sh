#!/bin/bash

for i in {1..10}
do
  timestamp=$(date +%s)$(printf "%03d" $(($(date +%N) / 1000000)))
  curl -kq -X POST -H "x-numaflow-event-time: $timestamp" -d "101 103 105 107" https://localhost:8444/vertices/in
  sleep 1
done
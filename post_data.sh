#!/bin/bash

for i in {1..600}
do
  timestamp=$(date +%s)$(printf "%03d" $(($(date +%N) / 1000000)))
  curl -kq -X POST -H "x-numaflow-event-time: $timestamp" -d "Hi how are you" https://localhost:8444/vertices/in
  sleep 1
done
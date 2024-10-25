#!/bin/bash

# Trap Ctrl+C (SIGINT)
trap 'echo ".....Ctrl+C received"; exit' SIGINT

i=1
while true; do
  echo "Hello Process 1, $i"
  sleep 1
  ((i++))
done

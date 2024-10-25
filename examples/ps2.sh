#!/bin/bash

# Trap Ctrl+C (SIGINT)
trap 'echo "Ctrl+C received"; exit' SIGINT

i=1
while true; do
  echo "Hello Process 2, $i"
  sleep 1
  ((i++))
  if [ $i -eq 5 ]; then
    echo "Process 2 is done"
    break
  fi
done

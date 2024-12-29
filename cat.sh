#!/bin/bash

own_pid=$$

touch /tmp/logfile

i=0
while true; do
    echo "stdout: Hello World" $i $own_pid
    echo "stderr: Hello World" $i $own_pid >&2
    echo "stderr: Hello World" $i $own_pid >>/tmp/logfile

    sleep 1
    i=$((i + 1))
done

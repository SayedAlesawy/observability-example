#!/bin/bash

# Check if the number of instances is passed as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <number_of_instances>"
  exit 1
fi

# Number of instances to run
NUM_INSTANCES=$1

# Array to hold process IDs
PIDS=()

# Function to run the script and store the PID
run_script() {
  ./gen-traffic.sh &
  PIDS+=($!)
}

# Function to kill all child processes
cleanup() {
  echo "Cleaning up..."
  for PID in "${PIDS[@]}"; do
    kill $PID 2>/dev/null
  done
  exit 1
}

# Trap SIGINT signal (Ctrl+C)
trap cleanup SIGINT

# Run the script multiple times
for ((i=0; i<NUM_INSTANCES; i++)); do
  run_script
done

# Wait for all instances to complete
for PID in "${PIDS[@]}"; do
  wait $PID
done

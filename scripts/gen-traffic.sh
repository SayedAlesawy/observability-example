#!/bin/bash

# Endpoint to send GET request
URL1="http://localhost:8080/ping/1/status"
URL2="http://localhost:8080/ping/1/info"

# Interval in seconds
INTERVAL=2

while true
do
  # Send GET request
  echo "sending request"
  curl -X GET $URL1 > /dev/null 2>&1
  curl -X GET $URL2 > /dev/null 2>&1

  # Wait for the specified interval
  sleep $INTERVAL
done

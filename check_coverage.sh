#!/bin/bash

# Check if coverage.out file exists
if [ ! -f "coverage.out" ]; then
  echo "coverage.out file not found."
  exit 1
fi

# Calculate test coverage percentage
coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Check if coverage is less than 50%
if (( $(echo "$coverage < 50" | bc -l) )); then
  echo "Error: Coverage is less than 50%: $coverage%"
  exit 1
fi

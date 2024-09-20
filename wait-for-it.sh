#!/usr/bin/env bash

# Use this script to wait for a service to be available.
# Usage: wait-for-it.sh <host>:<port> [--timeout=<timeout>] [--] <command>

TIMEOUT=15
WAITFORIT_CMD=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    *:* )
      HOST_PORT="$1"
      shift
      ;;
    --timeout)
      TIMEOUT="${2:-15}"
      shift 2
      ;;
    --)
      shift
      WAITFORIT_CMD="$@"
      break
      ;;
    *)
      echo "Invalid argument: $1"
      exit 1
      ;;
  esac
done

# Split HOST and PORT
IFS=: read -r HOST PORT <<< "$HOST_PORT"

echo "Waiting for $HOST:$PORT..."

for (( i=0; i<TIMEOUT; i++ )); do
  nc -z "$HOST" "$PORT" && break
  sleep 1
done

if ! nc -z "$HOST" "$PORT"; then
  echo "Timeout: $HOST:$PORT not available after $TIMEOUT seconds."
  exit 1
fi

echo "$HOST:$PORT is available."
exec "${WAITFORIT_CMD[@]}"

#!/bin/bash

set -e

function timestamp() {
    date +"%Y/%m/%d %T"
}

function log() {
    local type="$1"
    local msg="$2"
    local script_name="create-eks.sh"
    echo "$(timestamp) [$script_name] [$type] $msg"
}

function retry {
  local retries="$1"
  shift

  echo "$@"

  local count=0
  local wait=5
  until "$@"; do
      exit="$?"
      if [ $count -lt $retries ]; then
          log "INFO" "Attempt $count/$retries. Command exited with exit_code: $exit. Retrying after $wait seconds..."
          sleep $wait
      else
          log "INFO" "Command failed in all $retries attempts with exit_code: $exit. Stopping trying any further...."
          return $exit
      fi
      count=$(($count + 1))
  done
  return 0
}

cmd="la -l"
retry 1 ${cmd} > test.txt

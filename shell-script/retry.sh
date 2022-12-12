#!/bin/bash

function timestamp() {
    date +"%Y/%m/%d %T"
}

function log() {
    local type="$1"
    local msg="$2"
    echo "$(timestamp) [$script_name] [$type] $msg"
}


function retry {
    local retries="$1"
    shift

    local count=0
    local wait=1
    echo "$@"
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

#export CLUSTER_NAME='cluster'
#export KUBERNETES_VERSION="v1.23.12"
#cmd1="clusterctl generate cluster ${CLUSTER_NAME} --kubernetes-version ${KUBERNETES_VERSION} --flavor aks"
#retry 2 ${cmd1}  > capi-aks-cluster.yaml

#cmd1="kubectl get pods -A"
#retry 5 ${cmd1} | grep capz-system


#kubectl get pod capi-visualizer-6bbdd7d7b5-79jqg -n vis -o yaml > pod.yaml
cmd1="kubectl get pod capi-visualizer-6bbdd7d7b5-79jqg -n vis -o yaml"
retry 2 ${cmd1} > pod.yaml
#!/bin/bash

case $(uname -m) in
x86_64)
    sys_arch=amd64
    ;;
arm64|aarch64)
    sys_arch=arm64
    ;;
ppc64le)
    sys_arch=ppc64le
    ;;
s390x)
    sys_arch=s390x
    ;;
*)
    sys_arch=amd64
    ;;
esac


cmd="hello"
if [[ "$sys_arch" = "amd64" ]]; then
  cmd="rasel"
fi

echo $cmd

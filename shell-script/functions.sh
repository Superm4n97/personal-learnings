#!/bin/bash
function foo() {
    echo "Hello"
}
bar() {
  echo $1
}

#foo
#bar "World"

loop_iteration(){
  for st in "$@" ; do
      echo $st
  done
}
#loop_iteration "hello" world "foo"

loop_iteration2(){
  for st in $* ; do
      echo $st
  done
}
#loop_iteration2 "hello" world "foo"

#retry () {
#  local count=$1
#  local wait=0
#  until "$@"; do
#    exit="$?"
#    echo exit
#  done
#  return 0
#}
#
#cmd={echo "hello world"}
#
#echo $cmd
#retry ${cmd}

#var=$(pwd)
#echo ${echo hello}
#echo ${var}
#echo $?

fun2() {
  echo "$@" #output: 1 /home/rasel/go/src/github.com/Superm4n97/personal-learnings/shell-script
  echo @* #output: @*
}
#var=$(pwd)
#fun2 1 $var

fun3() {
  until "$@" & echo "success"; do
    echo "failed"
    exit="$?"
    echo $exit
    if ( $exit -eq 0 ); then
      echo "retry"
      return $exit
    else
      sleep 1
    fi
  done

}
#var1=$(pwd)
#var2=$(echo "foo")
#fun3 $(export "TEST"="TEST")

echo ""'{"spec":{"containers":[{"name": "manager","image": "pkbhowmick/cluster-api-azure-controller-amd64:dev"}]}}'""

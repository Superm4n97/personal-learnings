#!/bin/bash

set -e

foo() {
  text1="rasel"
  local text2="hossain"
  text2="world"
  bar
}
bar() {
  echo $text1
  echo $text2
}

#foo
#bar

temp() {
  echo $1
}
#temp "hi there"

file_name() {
  script_name=${0##*/}
  echo $script_name
}
#file_name

architecture() {
  case $(uname -m) in
  x86_64)
      arch=amd64
      ;;
  arm64|aarch64)
      arch=arm64
      ;;
  ppc64le)
      arch=ppc64le
      ;;
  s390x)
      arch=s390x
      ;;
  *)
      arch=amd64
      ;;
  esac

  echo $arch
}
architecture
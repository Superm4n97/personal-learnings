#!/usr/bin/bash
function emptyString() {
    # empty string
    RESULT=anything

    #compare with [] bracket
    if [ $RESULT ] ; then
      echo "Non-empty string is true"
    fi
}

function ifCondision () {
  read NAME
  if [ $NAME == ADMIN ]; then
      echo "Hello ADMIN"
  fi
}

function arithmeticExpression() {
  read AGE

  # remember this double () bracket
  if (( $AGE >= 18 )); then
    echo matured
  fi

  if (( $AGE < 18 )) ; then
    echo "under aged"
  fi
}

function loop() {
  cd $ROOT
    for name in c*-* ; do
      echo $name
    done

}

#ifCondision
#emptyString
#arithmeticExpression
#loop

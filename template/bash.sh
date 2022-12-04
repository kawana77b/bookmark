function __bm() {
  p=$(bookmark get)
  if [ -d $p ]; then
    cd $p
  fi
}

function __bma() {
  bookmark add
}

alias bm='__bm'
alias bma='__bma'
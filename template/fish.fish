function __bm
  set -l p (bookmark get)
  if test -d $p
    cd $p
  end
end

function __bma
  bookmark add
end

alias bm='__bm'
alias bma='__bma'
function __bm {
  param (
  )

  $p = (bookmark get)
  if ($p) {
    if (Test-Path $p) {
      Set-Location $p
    }
  }
}

function __bma {
  param (
  )

  bookmark add
}

Set-Alias bm __bm
Set-Alias bma __bma
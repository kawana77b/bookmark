{{/* bash */}}
{{- define "bash-prof" -}}
eval "$(bookmark init bash)"
{{- end }}

{{/* fish */}}
{{- define "fish-prof" -}}
bookmark init fish | source
{{- end}}

{{/* powershell */}}
{{- define "powershell-prof" -}}
Invoke-Expression (& {
    (bookmark init powershell | Out-String)
})
{{- end}}
module "{{ .KeyvaultName }}" {
  source = "../../../../modules/azure/keyvault"
  name   = "{{ .Values.Name }}"
  resource_group = {
    name     = "{{ .Values.ResourceGroup }}"
    location = "{{ .Values.Location }}"
  }
  access_policies = [
    {{- range .Values.Properties.AccessPolicies }}
    {
      object_id = "{{ .ObjectId }}"
      certificate_permissions = [
        {{- range $CertificatePermissionIndex, $CertificatePermission := .Permissions.Certificates -}}
        {{- if eq $CertificatePermissionIndex 0 -}}
        "{{ $CertificatePermission }}"
        {{- else -}}
      , "{{ $CertificatePermission }}"
        {{- end -}}
        {{- end -}}
      ]
      secret_permissions      = [
        {{- range $SecretPermissionIndex, $SecretPermission := .Permissions.Secrets -}}
        {{- if eq $SecretPermissionIndex 0 -}}
        "{{ $SecretPermission }}"
        {{- else -}}
      , "{{ $SecretPermission }}"
        {{- end -}}
        {{- end -}}
      ]
      key_permissions         = [
        {{- range $KeyPermissionIndex, $KeyPermission := .Permissions.Keys -}}
        {{- if eq $KeyPermissionIndex 0 -}}
        "{{ $KeyPermission }}"
        {{- else -}}
      , "{{ $KeyPermission }}"
        {{- end -}}
        {{- end -}}
      ]
    },
    {{- end }}
  ]
}

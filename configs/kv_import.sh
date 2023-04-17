#!/bin/bash

terraform import 'module.{{.KeyvaultName}}.azurerm_key_vault.this' '{{.Values.Id}}'
{{ range .Values.Properties.AccessPolicies -}}
terraform import 'module.{{$.KeyvaultName}}.azurerm_key_vault_access_policy.this["{{.ObjectId}}"]' '{{$.Values.Id}}/objectId/{{.ObjectId}}'
{{ end -}}
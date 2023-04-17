#!/bin/bash

terraform import 'azuread_application.{{.ResourceName}}' '{{.Sp.AppID}}'
terraform import 'azuread_service_principal.{{.ResourceName}}' '{{.Sp.ID}}'

{{- if ne .App.AppId "" -}}
resource "azuread_application" "{{.ResourceName}}" {
  display_name = "{{ .App.DisplayName }}"
  oauth2_permission_scope = {
    admin_consent_description  =
    admin_consent_display_name =
    enabled                    =
    id                         =
    type                       =
    user_consent_description   =
    user_consent_display_name  =
    value                      =
  }
  app_role = {
    allowed_member_types =
    description          =
    display_name         =
    enabled              =
    id                   =
    value                =
  }
  description                    = "{{ .App.Description }}"
  device_only_auth_enabled       = "{{ .App.IsDeviceOnlyAuthSupported }}"
  fallback_public_client_enabled = "{{ .App.IsFallbackPublicClient }}"
  feature_tags = {
    custom_single_sign_on =
    enterprise            =
    gallery               =
    hide                  =
  }
  group_membership_claims       = "{{ .App.GroupMembershipClaims }}"
  identifier_uris               = {{ .App.IdentifierUris }}
  logo_image                    =
  marketing_url                 = "{{ .App.Info.MarketingUrl }} "
  notes                         = "{{ .App.Notes }}"
  oauth2_post_response_required =
  optional_claims = {
    access_token = {
      additional_properties =
      essential             =
      name                  =
      source                =
    }
    id_token = {
      additional_properties =
      essential             =
      name                  =
      source                =
    }
    saml2_token = {
      additional_properties =
      essential             =
      name                  =
      source                =
    }
  }
  owners                  =
  prevent_duplicate_names =
  privacy_statement_url   = "{{ .App.Info.PrivacyStatementUrl }}"
  public_client = {
    redirect_uris = {{ .App.PublicClient.RedirectUris }}
  }
  required_resource_access = {
    resource_access = {
      id   =
      type =
    }
    resource_app_id =
  }
  sign_in_audience =
  single_page_application = {
    redirect_uris =
  }
  support_url          =
  tags                 =
  template_id          =
  terms_of_service_url =
  web = {
    homepage_url =
    implicit_grant = {
      access_token_issuance_enabled =
      id_token_issuance_enabled     =
    }
    logout_url    =
    redirect_uris =
  }
  api = {
    known_client_applications      =
    mapped_claims_enabled          =
    oauth2_permission_scope        =
    requested_access_token_version =
  }
}

{{ end -}}

resource "azuread_service_principal" "{{ .ResourceName }}" {
{{- if eq .App.AppId "" }}
  application_id                = "{{ .Sp.AppID }}"
{{- else }}
  application_id                = azuread_application.{{ .ResourceName }}
{{- end }}
  account_enabled               = {{ .Sp.AccountEnabled }}
  app_role_assignment_required  = {{ .Sp.AppRoleAssignmentRequired }}
{{- if .Sp.Description }}
  description                   = "{{ .Sp.Description }}"
{{- end }}
{{- if .Sp.LoginURL }}
  login_url                     = "{{ .Sp.LoginURL }}"
{{- end }}
{{- if .Sp.Notes }}
  notes                         = "{{ .Sp.Notes }}"
{{- end }}
  use_existing                  = true
  notification_email_addresses  = [
    {{- range $idx, $value := .Sp.NotificationEmailAddresses -}}
      {{- if eq $idx 0 -}}
        "{{ $value }}"
      {{- else -}}
        , "{{ $value }}"
      {{- end -}}
    {{- end -}}
  ]
{{- if .Sp.PreferredSingleSignOnMode }}
  preferred_single_sign_on_mode = "{{ .Sp.PreferredSingleSignOnMode }}"
{{- end }}
  owners                        = [data.azuread_client_config.current.object_id]
  alternative_names = [
  {{- range .Sp.AlternativeNames }}
    "{{ . }}",
  {{- end }}
  ]

  tags = {
    managed-by = "terraform"
  }
}
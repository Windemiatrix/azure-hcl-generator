package tpl

import "time"

type AppParcedItem struct {
	OdataContext string        `json:"@odata.context"`
	AddIns       []interface{} `json:"addIns"`
	Api          struct {
		AcceptMappedClaims          interface{}   `json:"acceptMappedClaims"`
		KnownClientApplications     []interface{} `json:"knownClientApplications"`
		Oauth2PermissionScopes      []interface{} `json:"oauth2PermissionScopes"`
		PreAuthorizedApplications   []interface{} `json:"preAuthorizedApplications"`
		RequestedAccessTokenVersion int           `json:"requestedAccessTokenVersion"`
	} `json:"api"`
	AppId                     string        `json:"appId"`
	AppRoles                  []interface{} `json:"appRoles"`
	ApplicationTemplateId     interface{}   `json:"applicationTemplateId"`
	Certification             interface{}   `json:"certification"`
	CreatedDateTime           time.Time     `json:"createdDateTime"`
	DefaultRedirectUri        interface{}   `json:"defaultRedirectUri"`
	DeletedDateTime           interface{}   `json:"deletedDateTime"`
	Description               interface{}   `json:"description"`
	DisabledByMicrosoftStatus interface{}   `json:"disabledByMicrosoftStatus"`
	DisplayName               string        `json:"displayName"`
	GroupMembershipClaims     interface{}   `json:"groupMembershipClaims"`
	Id                        string        `json:"id"`
	IdentifierUris            []interface{} `json:"identifierUris"`
	Info                      struct {
		LogoUrl             interface{} `json:"logoUrl"`
		MarketingUrl        interface{} `json:"marketingUrl"`
		PrivacyStatementUrl interface{} `json:"privacyStatementUrl"`
		SupportUrl          interface{} `json:"supportUrl"`
		TermsOfServiceUrl   interface{} `json:"termsOfServiceUrl"`
	} `json:"info"`
	IsDeviceOnlyAuthSupported interface{}   `json:"isDeviceOnlyAuthSupported"`
	IsFallbackPublicClient    interface{}   `json:"isFallbackPublicClient"`
	KeyCredentials            []interface{} `json:"keyCredentials"`
	Notes                     interface{}   `json:"notes"`
	OptionalClaims            interface{}   `json:"optionalClaims"`
	ParentalControlSettings   struct {
		CountriesBlockedForMinors []interface{} `json:"countriesBlockedForMinors"`
		LegalAgeGroupRule         string        `json:"legalAgeGroupRule"`
	} `json:"parentalControlSettings"`
	PasswordCredentials []interface{} `json:"passwordCredentials"`
	PublicClient        struct {
		RedirectUris []interface{} `json:"redirectUris"`
	} `json:"publicClient"`
	PublisherDomain                   string        `json:"publisherDomain"`
	RequestSignatureVerification      interface{}   `json:"requestSignatureVerification"`
	RequiredResourceAccess            []interface{} `json:"requiredResourceAccess"`
	SamlMetadataUrl                   interface{}   `json:"samlMetadataUrl"`
	ServiceManagementReference        interface{}   `json:"serviceManagementReference"`
	ServicePrincipalLockConfiguration interface{}   `json:"servicePrincipalLockConfiguration"`
	SignInAudience                    string        `json:"signInAudience"`
	Spa                               struct {
		RedirectUris []interface{} `json:"redirectUris"`
	} `json:"spa"`
	Tags                 []interface{} `json:"tags"`
	TokenEncryptionKeyId interface{}   `json:"tokenEncryptionKeyId"`
	VerifiedPublisher    struct {
		AddedDateTime       interface{} `json:"addedDateTime"`
		DisplayName         interface{} `json:"displayName"`
		VerifiedPublisherId interface{} `json:"verifiedPublisherId"`
	} `json:"verifiedPublisher"`
	Web struct {
		HomePageUrl           interface{} `json:"homePageUrl"`
		ImplicitGrantSettings struct {
			EnableAccessTokenIssuance bool `json:"enableAccessTokenIssuance"`
			EnableIdTokenIssuance     bool `json:"enableIdTokenIssuance"`
		} `json:"implicitGrantSettings"`
		LogoutUrl           interface{}   `json:"logoutUrl"`
		RedirectUriSettings []interface{} `json:"redirectUriSettings"`
		RedirectUris        []interface{} `json:"redirectUris"`
	} `json:"web"`
}

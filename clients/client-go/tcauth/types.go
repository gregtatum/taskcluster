// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcauth

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster/v54/clients/client-go"
)

type (
	// Response for a request to get access to an S3 bucket.
	AWSS3CredentialsResponse struct {

		// Temporary STS credentials for use when operating on S3
		Credentials TemporarySecurityCredentials `json:"credentials"`

		// Date and time of when the temporary credentials expires.
		Expires tcclient.Time `json:"expires"`
	}

	// Response from a request to authenticate a hawk request.
	AuthenticationFailedResponse struct {

		// Message saying why the authentication failed.
		Message string `json:"message"`

		// If `auth-failed`, the request failed to authenticate (e.g., invalid credentials).
		//
		// Possible values:
		//   * "auth-failed"
		Status string `json:"status"`
	}

	// Response from a request to authenticate a hawk request.
	AuthenticationSuccessfulResponse struct {

		// The `clientId` that made this request.  This may be the `id` supplied in
		// the Authorization header, or in the case of a named temporary credential
		// may be embedded in the payload.  In any case, this clientId can be used
		// for logging, auditing, and identifying the credential but **must** not be
		// used for access control.  That's what scopes are for.
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		ClientID string `json:"clientId"`

		// The expiration time that the returned scopes are valid for.
		// This should be treated as the latest time at which the authorization
		// is valid. For most cases, where the access being authorized occurs
		// immediately, this field can be ignored, as the value will always be
		// in the future.
		Expires tcclient.Time `json:"expires"`

		// Payload as extracted from `Authentication` header. This property is
		// only present if a hash is available. You are not required to validate
		// this hash, but if you do, please check `scheme` to ensure that it's
		// on a scheme you support.
		Hash json.RawMessage `json:"hash,omitempty"`

		// Authentication scheme the client used. Generally, you don't need to
		// read this property unless `hash` is provided and you want to validate
		// the payload hash. Additional values may be added in the future.
		//
		// Possible values:
		//   * "hawk"
		Scheme string `json:"scheme"`

		// List of scopes the client is authorized to access.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`

		// If `auth-success`, the request has successfully authenticated.
		// This does not imply authorization (i.e. scope satisfaction not guaranteed).
		//
		// Possible values:
		//   * "auth-success"
		Status string `json:"status"`
	}

	// Response to a request for an Shared-Access-Signature to access an Azure
	// Blob Storage container.
	AzureBlobSharedAccessSignature struct {

		// Date and time of when the Shared-Access-Signature expires.
		Expiry tcclient.Time `json:"expiry"`

		// Shared-Access-Signature string. This is the querystring parameters to
		// be appened after `?` or `&` depending on whether or not a querystring is
		// already present in the URL.
		Sas string `json:"sas"`
	}

	// A list of Azure accounts managed by taskcluster-auth
	AzureListAccountResponse struct {

		// A list of accountIds that are managed by auth. These are
		// the accounts that can have SAS credentials fetched for tables
		// within them.
		//
		// Array items:
		Accounts []string `json:"accounts"`
	}

	// A list of Azure containers in an account
	AzureListContainersResponse struct {

		// A list of containers that are in an account.  Credentials are available for
		// these containers from the `azureBlobSAS` method.
		//
		// Array items:
		Containers []string `json:"containers"`

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of containers.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called this method with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// A list of Azure tables in an account
	AzureListTableResponse struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of tables.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `azureAccountTables` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// A list of tables that are in an account. These are
		// the tables that can have SAS credentials fetched for them.
		//
		// Array items:
		Tables []string `json:"tables"`
	}

	// Response to a request for an Shared-Access-Signature to access and Azure
	// Table Storage table.
	AzureTableSharedAccessSignature struct {

		// Date and time of when the Shared-Access-Signature expires.
		Expiry tcclient.Time `json:"expiry"`

		// Shared-Access-Signature string. This is the querystring parameters to
		// be appened after `?` or `&` depending on whether or not a querystring is
		// already present in the URL.
		Sas string `json:"sas"`
	}

	// Properties to create a client.
	CreateClientRequest struct {

		// If `true`, the service may delete this client after it has expired.  If
		// `false` (the default), the client will remain after expiration, although
		// it cannot be used for authentication in that state.
		//
		// Default:    false
		DeleteOnExpiration bool `json:"deleteOnExpiration,omitempty"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		Description string `json:"description"`

		// Date and time where the clients access is set to expire
		Expires tcclient.Time `json:"expires"`

		// List of scopes the client has (unexpanded).
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes,omitempty"`
	}

	// All details about a client including the `accessToken`
	CreateClientResponse struct {

		// AccessToken used for authenticating requests, you should store this
		// you won't be able to retrive it again!
		//
		// Syntax:     ^[a-zA-Z0-9_-]{22,66}$
		AccessToken string `json:"accessToken"`

		// ClientId of the client
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		ClientID string `json:"clientId"`

		// Date and time when this client was created
		Created tcclient.Time `json:"created"`

		// If `true`, the service may delete this client after it has expired.  If
		// `false`, the client will remain after expiration, although it cannot be
		// used for authentication in that state.
		DeleteOnExpiration bool `json:"deleteOnExpiration"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, this client is disabled and cannot be used.  This usually occurs when the
		// scopes available to the user owning the client no longer satisfy the client.
		Disabled bool `json:"disabled"`

		// List of scopes granted to this client by matching roles, including the
		// client's scopes and the implicit role `client-id:<clientId>`.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time where the clients access is set to expire
		Expires tcclient.Time `json:"expires"`

		// Date of last time this client was used. Will only be updated every 6 hours
		// or so this may be off by up-to 6 hours. But it still gives a solid hint
		// as to whether or not this client is in use.
		LastDateUsed tcclient.Time `json:"lastDateUsed"`

		// Date and time of last modification
		LastModified tcclient.Time `json:"lastModified"`

		// Date and time of when the `accessToken` was reset last time.
		LastRotated tcclient.Time `json:"lastRotated"`

		// List of scopes the client has (unexpanded).
		//
		// Default:    []
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// Data to create or update a role.
	CreateRoleRequest struct {

		// Description of what this role is used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		Description string `json:"description"`

		// List of scopes the role grants access to.  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// Access credentials and urls for the Sentry project.
	// Credentials will expire in 24-48 hours, you should refresh them within
	// 24 hours.
	Dsn struct {

		// Access credential and URL for public error reports.
		// These credentials can be used for up-to 24 hours.
		// This is for use in client-side applications only.
		Public string `json:"public"`

		// Access credential and URL for private error reports.
		// These credentials can be used for up-to 24 hours.
		// This is for use in serser-side applications and should **not** be
		// leaked.
		Secret string `json:"secret"`
	}

	// Response for a request to get a GCP temporary credential.
	GCPCredentialsResponse struct {

		// Temporary oauth2 access token to access the given service account
		AccessToken string `json:"accessToken"`

		// The access token expire time
		ExpireTime tcclient.Time `json:"expireTime"`
	}

	// List of roles
	GetAllRolesNoPagination []GetRoleResponse

	// If no limit is given, all roles are returned. Since this
	// list may become long, callers can use the `limit` and `continuationToken`
	// query arguments to page through the responses.
	GetAllRolesResponse struct {

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// A list of requested roles
		Roles []GetRoleResponse `json:"roles"`
	}

	// Get all details about a client, useful for tools modifying a client
	GetClientResponse struct {

		// ClientId of the client scopes is requested about
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		ClientID string `json:"clientId"`

		// Date and time when this client was created
		Created tcclient.Time `json:"created"`

		// If `true`, the service may delete this client after it has expired.  If
		// `false`, the client will remain after expiration, although it cannot be
		// used for authentication in that state.
		DeleteOnExpiration bool `json:"deleteOnExpiration"`

		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, this client is disabled and cannot be used.  This usually occurs when the
		// scopes available to the user owning the client no longer satisfy the client.
		Disabled bool `json:"disabled"`

		// List of scopes granted to this client by matching roles.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Scope that client is granted by a role
		//
		// Syntax:     ^[ -~]*$
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time where the clients access is set to expire
		Expires tcclient.Time `json:"expires"`

		// Date of last time this client was used. Will only be updated every 6 hours
		// or so this may be off by up-to 6 hours. But it still gives a solid hint
		// as to whether or not this client is in use.
		LastDateUsed tcclient.Time `json:"lastDateUsed"`

		// Date and time of last modification
		LastModified tcclient.Time `json:"lastModified"`

		// Date and time of when the `accessToken` was reset last time.
		LastRotated tcclient.Time `json:"lastRotated"`

		// List of scopes the client has (unexpanded).  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// If no limit is given, the roleIds of all roles are returned. Since this
	// list may become long, callers can use the `limit` and `continuationToken`
	// query arguments to page through the responses.
	GetRoleIdsResponse struct {

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// A list of requested roleIds
		//
		// Array items:
		RoleIds []string `json:"roleIds"`
	}

	// Get all details about a role
	GetRoleResponse struct {

		// Date and time when this role was created
		Created tcclient.Time `json:"created"`

		// Description of what this role is used for in markdown.
		// Should include who is the owner, point of contact.
		//
		// Max length: 10240
		Description string `json:"description"`

		// List of scopes granted anyone who assumes this role, including anything
		// granted by roles that can be assumed when you have this role.
		// Hence, this includes any scopes in-directly granted as well.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		ExpandedScopes []string `json:"expandedScopes"`

		// Date and time of last modification
		LastModified tcclient.Time `json:"lastModified"`

		// roleId of the role requested
		//
		// Syntax:     ^[\x20-\x7e]+$
		RoleID string `json:"roleId"`

		// List of scopes the role grants access to.  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// Array items:
		// A single scope. A scope must be composed of
		// printable ASCII characters and spaces.  Scopes ending in more than
		// one `*` character are forbidden.
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// Request to authenticate a hawk request.
	HawkSignatureAuthenticationRequest struct {

		// Authorization header, **must** only be specified if request being
		// authenticated has a `Authorization` header.
		Authorization string `json:"authorization,omitempty"`

		// Host for which the request came in, this is typically the `Host` header
		// excluding the port if any.
		Host string `json:"host"`

		// HTTP method of the request being authenticated.
		//
		// Possible values:
		//   * "get"
		//   * "post"
		//   * "put"
		//   * "head"
		//   * "delete"
		//   * "options"
		//   * "trace"
		//   * "copy"
		//   * "lock"
		//   * "mkcol"
		//   * "move"
		//   * "purge"
		//   * "propfind"
		//   * "proppatch"
		//   * "unlock"
		//   * "report"
		//   * "mkactivity"
		//   * "checkout"
		//   * "merge"
		//   * "m-search"
		//   * "notify"
		//   * "subscribe"
		//   * "unsubscribe"
		//   * "patch"
		//   * "search"
		//   * "connect"
		Method string `json:"method"`

		// Port on which the request came in, this is typically `80` or `443`.
		// If you are running behind a reverse proxy look for the `x-forwarded-port`
		// header.
		//
		// Mininum:    0
		// Maximum:    65535
		Port int64 `json:"port"`

		// Resource the request operates on including querystring. This is the
		// string that follows the HTTP method.
		// **Note,** order of querystring elements is important.
		Resource string `json:"resource"`

		// Source IP of the authentication request or request that requires
		// authentication. This is only used for audit logging.
		//
		// One of:
		//   * SourceIP
		//   * SourceIP1
		SourceIP string `json:"sourceIp,omitempty"`
	}

	// Response from a request to authenticate a hawk request.
	//
	// One of:
	//   * AuthenticationSuccessfulResponse
	//   * AuthenticationFailedResponse
	//   * NoAuthenticationResponse
	HawkSignatureAuthenticationResponse json.RawMessage

	// List of clients
	ListClientResponse struct {
		Clients []GetClientResponse `json:"clients"`

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// Response from a request to authenticate a hawk request.
	NoAuthenticationResponse struct {

		// The expiration time that the returned scopes are valid for.
		// This should be treated as the latest time at which the authorization
		// is valid. For most cases, where the access being authorized occurs
		// immediately, this field can be ignored, as the value will always be
		// in the future.
		Expires tcclient.Time `json:"expires"`

		// Authentication scheme the client used. Generally, you don't need to
		// read this property unless `hash` is provided and you want to validate
		// the payload hash. Additional values may be added in the future.
		//
		// Possible values:
		//   * "none"
		Scheme string `json:"scheme"`

		// List of scopes the client is authorized to access.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`

		// If `no-auth`, the request provided no credentials or bewit.
		//
		// Possible values:
		//   * "no-auth"
		Status string `json:"status"`
	}

	// Sentry DSN for submitting errors.
	SentryDSNResponse struct {

		// Access credentials and urls for the Sentry project.
		// Credentials will expire in 24-48 hours, you should refresh them within
		// 24 hours.
		Dsn Dsn `json:"dsn"`

		// Expiration time for the credentials. The credentials should not be used
		// after this time. They might not be revoked immediately, but will be at
		// some arbitrary point after this date-time.
		Expires tcclient.Time `json:"expires"`

		// Project name that the DSN grants access to.
		Project string `json:"project"`
	}

	// A set of scopes
	SetOfScopes struct {

		// List of scopes.  Scopes must be composed of printable ASCII characters and spaces.
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// Source IP of the authentication request or request that requires
	// authentication. This is only used for audit logging.
	SourceIP string

	// Source IP of the authentication request or request that requires
	// authentication. This is only used for audit logging.
	SourceIP1 string

	// Temporary STS credentials for use when operating on S3
	TemporarySecurityCredentials struct {

		// Access key identifier that identifies the temporary security
		// credentials.
		AccessKeyID string `json:"accessKeyId"`

		// Secret access key used to sign requests
		SecretAccessKey string `json:"secretAccessKey"`

		// A token that must passed with request to use the temporary
		// security credentials.
		SessionToken string `json:"sessionToken"`
	}

	// Details on how the test request should be authenticated.
	TestAuthenticateRequest struct {

		// List of scopes that should be client used should be given.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		ClientScopes []string `json:"clientScopes"`

		// List of scopes the request should require.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		RequiredScopes []string `json:"requiredScopes"`
	}

	// Details on how the test request was authenticated.
	TestAuthenticateResponse struct {

		// ClientId from the request as it will be logged
		//
		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		ClientID string `json:"clientId"`

		// List of scopes the request was authorized.
		//
		// Default:    []
		//
		// Array items:
		// Scope
		//
		// Syntax:     ^[ -~]*$
		Scopes []string `json:"scopes"`
	}

	// Token for connecting a worker to websocktunnel proxy
	WebsocktunnelTokenResponse struct {

		// The time at which the JWT will expire.
		Expires tcclient.Time `json:"expires"`

		// The JWT to be used as `Bearer <token>` when connecting to the service.
		Token string `json:"token"`

		// Audience identifying the websocktunnel servers that will honor this token; this will be the
		// same as the requested `wstAudience`.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{1,38}$
		WstAudience string `json:"wstAudience"`

		// Id for the websocktunnel client connection; this will be the same as the requested `wstClient`.
		//
		// Syntax:     ^[a-zA-Z0-9_~.%-]+$
		WstClient string `json:"wstClient"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// HawkSignatureAuthenticationResponse is of type json.RawMessage...
func (this *HawkSignatureAuthenticationResponse) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *HawkSignatureAuthenticationResponse) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("HawkSignatureAuthenticationResponse: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

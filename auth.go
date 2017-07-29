package quickbooks

import "net/url"

// GetConnectURL gets quickbooks login url for OAuth2
func GetConnectURL(clientID, scope, redirectURI, csrfToken string, isSandbox bool) (string, error) {
	discovery, err := NewDiscovery(isSandbox)
	if err != nil {
		return "", err
	}

	authorizationEndpoint := discovery.AuthorizationEndpoint
	URL, err := url.Parse(authorizationEndpoint)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Add("client_id", clientID)
	q.Add("response_type", "code")
	q.Add("scope", scope)
	q.Add("redirect_uri", redirectURI)
	q.Add("state", csrfToken)
	URL.RawQuery = q.Encode()

	return URL.String(), nil
}

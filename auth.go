package quickbooks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/jinmatt/go-quickbooks.v2/sdk/consts"
)

// BearerToken token response type from quickbooks
type BearerToken struct {
	RefreshToken           string `json:"refresh_token"`
	AccessToken            string `json:"access_token"`
	TokenType              string `json:"token_type"`
	IDToken                string `json:"id_token"`
	ExpiresIn              int64  `json:"expires_in"`
	XRefreshTokenExpiresIn int64  `json:"x_refresh_token_expires_in"`
}

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
	q.Set("client_id", clientID)
	q.Add("response_type", "code")
	q.Add("scope", scope)
	q.Add("redirect_uri", redirectURI)
	q.Add("state", csrfToken)
	URL.RawQuery = q.Encode()

	return URL.String(), nil
}

// GetBearerToken exchanges authorization code for bearer tokens
func GetBearerToken(clientID, clientSecret, code, redirectURI string, isSandbox bool) (*BearerToken, error) {
	discovery, err := NewDiscovery(isSandbox)
	if err != nil {
		return nil, err
	}

	tokenEndpoint := discovery.TokenEndpoint

	q := url.Values{}
	q.Set("grant_type", "authorization_code")
	q.Add("code", code)
	q.Add("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString(q.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, clientSecret))

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		sdkError := SDKError{}
		return nil, sdkError.New(consts.QBAuthorizationCodeFailure, consts.QBAuthorizationCodeFailureCode, consts.QBAuthorizationFailureCodeMessage)

	}

	bearerToken := BearerToken{}
	err = json.NewDecoder(res.Body).Decode(&bearerToken)
	if err != nil {
		return nil, err
	}

	return &bearerToken, nil
}

func basicAuth(clientID, clientSecret string) string {
	auth := clientID + ":" + clientSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

package quickbooks

import (
	"encoding/json"
	"net/http"

	"github.com/jinmatt/go-quickbooks.v2/sdk"
)

// Discovery quickbooks discovery API response type
type Discovery struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	RevocationEndpoint    string `json:"revocation_endpoint"`
	JwksURI               string `json:"jwks_uri"`
}

// NewDiscovery makes call to quickbooks discovery API and returns discovery object
func NewDiscovery(isSandbox bool) (*Discovery, error) {
	var discoveryHost string
	if isSandbox {
		discoveryHost = sdk.SandboxDiscoveryURL
	} else {
		discoveryHost = sdk.ProductionDiscoveryURL
	}

	req, err := http.NewRequest("GET", discoveryHost, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	discovery := Discovery{}
	err = json.NewDecoder(res.Body).Decode(&discovery)
	if err != nil {
		return nil, err
	}

	return &discovery, nil
}

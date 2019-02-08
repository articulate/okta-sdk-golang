package okta

import (
	"fmt"

	"github.com/okta/okta-sdk-golang/okta/query"
)

type AuthorizationServerResource resource

type AuthorizationServer struct {
	Audiences   []string                 `json:"audiences"`
	Credentials []*AuthServerCredentials `json:"credentials"`
	Description string                   `json:"descriptions"`
	Name        string                   `json:"name"`
}

type AuthServerCredentials struct {
	Signing *ApplicationCredentialsSigning `json:"signing"`
}

func NewAuthorizationServer() *AuthorizationServer {
	return &AuthorizationServer{}
}

func (m *AuthorizationServerResource) DeleteAuthorizationServer(id string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *AuthorizationServerResource) ListAuthorizationServers(id string) ([]AuthorizationServer, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var auth []AuthorizationServer
	resp, err := m.client.requestExecutor.Do(req, &auth)
	if err != nil {
		return nil, resp, err
	}
	return auth, resp, nil
}
func (m *AuthorizationServerResource) AddAuthorizationServer(id string, body AuthorizationServer, qp *query.Params) (interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := body
	resp, err := m.client.requestExecutor.Do(req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}
	return authorizationServer, resp, nil
}

func (m *AuthorizationServerResource) GetAuthorizationServer(id string, authorizationServerInstance AuthorizationServer) (interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := authorizationServerInstance
	resp, err := m.client.requestExecutor.Do(req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}
	return authorizationServer, resp, nil
}
func (m *AuthorizationServerResource) ActivateAuthorizationServer(id string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/lifecycle/activate", id)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *AuthorizationServerResource) DeactivateAuthorizationServer(id string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/lifecycle/deactivate", id)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

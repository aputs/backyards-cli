// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"emperror.dev/errors"
	"k8s.io/client-go/rest"
)

type Client interface {
	Login() (*ResponseBody, error)
}

type client struct {
	config *rest.Config
	url    string
}

type AuthenticationMode string

const (
	TokenAuth AuthenticationMode = "token"
	CertAuth  AuthenticationMode = "cert"
)

type RequestBody struct {
	Mode       AuthenticationMode `json:"mode"`
	Token      string             `json:"token,omitempty"`
	ClientCert struct {
		// base64 encoded client key
		Key string `json:"key"`
		// base64 encoded client cert
		Cert string `json:"cert"`
	} `json:"cert,omitempty"`
}

type ResponseBody struct {
	User struct {
		Name   string   `json:"name"`
		Groups []string `json:"groups"`
		// Token is an ID token containing user info and capabilities loaded at login
		Token string `json:"token"`
		// WrappedToken is a very short lifetime encrypted token that wraps the ID token.
		// It's for cases where the token must be exposed as HTTP GET parameters over a
		// secure connection where the token will available in access logs and/or browser
		// history which would mean a potential security risk.
		WrappedToken string `json:"wrappedToken"`
	} `json:"user"`
}

func NewClient(config *rest.Config, url string) Client {
	return &client{
		config: config,
		url:    url,
	}
}

func (c *client) Login() (*ResponseBody, error) {
	rb, err := c.requestBody()
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(rb)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode request")
	}

	response, err := http.Post(c.url, "application/json", b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode < 500 {
			return nil, errors.Wrap(err, "invalid request")
		} else {
			return nil, errors.Wrap(err, "server error")
		}
	}

	parsedResponse := &ResponseBody{}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return parsedResponse, errors.Wrap(err, "failed to read response body")
	}
	err = json.NewDecoder(bytes.NewBuffer(responseBody)).Decode(parsedResponse)
	if err != nil {
		return parsedResponse, errors.WrapWithDetails(err, "failed to decode response", "response", string(responseBody))
	}
	if parsedResponse.User.Name == "" {
		return nil, errors.New("invalid response")
	}
	return parsedResponse, nil
}

func (c *client) requestBody() (*RequestBody, error) {
	rb := &RequestBody{}
	if c.config.BearerToken != "" {
		rb.Mode = TokenAuth
		rb.Token = c.config.BearerToken
		return rb, nil
	} else if c.config.BearerTokenFile != "" {
		bearerToken, err := ioutil.ReadFile(c.config.BearerTokenFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load bearer token from %s", c.config.BearerTokenFile)
		}
		rb.Mode = TokenAuth
		rb.Token = string(bearerToken)
		return rb, nil
	} else if c.config.TLSClientConfig.CertFile != "" && c.config.TLSClientConfig.KeyFile != "" {
		cert, err := ioutil.ReadFile(c.config.TLSClientConfig.CertFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load client cert from %s", c.config.TLSClientConfig.CertFile)
		}
		key, err := ioutil.ReadFile(c.config.TLSClientConfig.KeyFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load client key from %s", c.config.TLSClientConfig.KeyFile)
		}
		rb.Mode = CertAuth
		rb.ClientCert.Cert = base64.StdEncoding.EncodeToString(cert)
		rb.ClientCert.Key = base64.StdEncoding.EncodeToString(key)
		return rb, nil
	} else if len(c.config.TLSClientConfig.CertData) > 0 && len(c.config.TLSClientConfig.KeyData) > 0 {
		rb.Mode = CertAuth
		rb.ClientCert.Cert = base64.StdEncoding.EncodeToString(c.config.TLSClientConfig.CertData)
		rb.ClientCert.Key = base64.StdEncoding.EncodeToString(c.config.TLSClientConfig.KeyData)
		return rb, nil
	}
	return nil, errors.NewWithDetails("no credentials found in the provided config")
}

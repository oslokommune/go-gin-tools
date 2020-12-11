package servicetesting

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// GetDatabaseBackendURI returns the domain and port of the created database backend
func (env *Environment) GetDatabaseBackendURI() string {
	return env.databaseBackend.URI
}

// DoRequests makes a request to the test environment
func (env *Environment) DoRequest(relativeURL string, method string, body []byte) (*httptest.ResponseRecorder, error) {
	requestURL, _ := url.Parse(fmt.Sprintf("http://localhost%s", relativeURL))

	w := httptest.NewRecorder()
	request := http.Request{
		Method: method,
		URL:    requestURL,
		Header: map[string][]string{},
	}

	if env.testToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", env.testToken))
	}
	request.Header.Add("Content-Type", "application/json")

	if body != nil {
		bodyCloser := ioutil.NopCloser(bytes.NewReader(body))

		request.Body = bodyCloser
	}

	env.TestServer.ServeHTTP(w, &request)

	return w, nil
}

func (env *Environment) Teardown() error {
	err := env.databaseBackend.dockerPool.Purge(env.databaseBackend.dockerResource)
	if err != nil {
		return fmt.Errorf("error tearing down docker resource: %w", err)
	}

	return nil
}

func NewGinTestEnvironment(dbOpts *DataBackendOptions, bearerToken string) (*Environment, error) {
	env := &Environment{
		testToken: bearerToken,
	}

	if dbOpts != nil {
		dbBackendResult, err := createDataBackend(dbOpts)
		if err != nil {
			return nil, fmt.Errorf("error creating database backend: %w", err)
		}

		env.databaseBackend = dbBackendResult
	}

	return env, nil
}

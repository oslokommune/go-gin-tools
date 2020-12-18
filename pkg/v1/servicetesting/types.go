package servicetesting

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest"
)

type DatabaseBackend struct {
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource

	URI string
}

type DataBackendOptions struct {
	Repository           string
	Tag                  string
	EnvironmentVariables []string
	Cmd					 string

	RelevantPort         string
}

// Environment contains all information and all references needed to create and use a test environment
type Environment struct {
	databaseBackend *DatabaseBackend
	testToken       string

	TestServer *gin.Engine
}

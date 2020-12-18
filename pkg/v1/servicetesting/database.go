package servicetesting

import (
	"fmt"
	"github.com/ory/dockertest"
	"strings"
	"time"
)

func convertOptsToRunOptions(opts *DataBackendOptions) (result *dockertest.RunOptions) {
	result = &dockertest.RunOptions{
		Repository: opts.Repository,
		Tag:        opts.Tag,
		Env: 		opts.EnvironmentVariables,
	}

	if opts.Cmd != "" {
		result.Cmd = strings.Split(opts.Cmd, " ")
	}
	
	return result
}

func createDataBackend(opts *DataBackendOptions) (result *DatabaseBackend, err error) {
	result = &DatabaseBackend{}

	result.dockerPool, err = dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("error connecting to docker pool: %w", err)
	}

	result.dockerResource, err = result.dockerPool.RunWithOptions(convertOptsToRunOptions(opts))
	if err != nil {
		return nil, fmt.Errorf("error creating docker resource: %w", err)
	}

	time.Sleep(2 * time.Second)

	result.URI = fmt.Sprintf("localhost:%s", result.dockerResource.GetPort(opts.RelevantPort))

	return result, nil
}

func CreateRedisDatabaseBackendOptions() *DataBackendOptions {
	return &DataBackendOptions{
		Repository:           "redis",
		Tag:                  "6.0.9-alpine",
		EnvironmentVariables: []string{},
		RelevantPort:         "6379/tcp",
	}
}

func CreatePostgresDatabaseBackendOptions(postgresPassword string) *DataBackendOptions {
	return &DataBackendOptions{
		Repository:           "postgres",
		Tag:                  "13.1-alpine",
		EnvironmentVariables: []string{fmt.Sprintf("POSTGRES_PASSWORD=%s", postgresPassword)},
		RelevantPort:         "5432/tcp",
	}
}

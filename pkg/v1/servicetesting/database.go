package servicetesting

import (
	"fmt"
	"github.com/ory/dockertest"
	"time"
)

func createDataBackend(opts *DataBackendOptions) (result *DatabaseBackend, err error) {
	result = &DatabaseBackend{}

	result.dockerPool, err = dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("error connecting to docker pool: %w", err)
	}

	result.dockerResource, err = result.dockerPool.Run(opts.Repository, opts.Tag, opts.EnvironmentVariables)
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

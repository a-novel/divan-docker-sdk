package d2sdk

import (
	"context"
	"github.com/docker/docker/client"
)

type DivanManager struct {
	ConfigPath    string
	Version       string
	ContainerName string
	VolumeName    string
	Env           []string

	containerID string
	status string

	ctx context.Context
	cli *client.Client
}

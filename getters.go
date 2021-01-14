package d2sdk

import (
	"context"
	"github.com/docker/docker/client"
)

func (dm *DivanManager) ID() string {
	return dm.status
}

func (dm *DivanManager) ExecutionError() error {
	return dm.executionError
}

func (dm *DivanManager) Status() string {
	return dm.status
}

func (dm *DivanManager) Cli() *client.Client {
	return dm.cli
}

func (dm *DivanManager) Ctx() context.Context {
	return dm.ctx
}

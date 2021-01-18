package d2sdk

import (
	"context"
	"github.com/docker/docker/client"
)

func (dm *DivanManager) Cli() *client.Client {
	return dm.cli
}

func (dm *DivanManager) Ctx() context.Context {
	return dm.ctx
}

package d2sdk

import (
	"context"
	"github.com/a-novel/errors"
	"github.com/docker/docker/client"
)

/*
	Setup manager environment.

	May return:
		- ErrCannotRunDocker
*/
func (dm *DivanManager) Init() *errors.Error {
	if dm.ctx == nil {
		dm.ctx = context.Background()
	}

	if dm.cli == nil {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			return errors.New(ErrCannotRunDocker, err.Error())
		}

		dm.cli = cli
	}

	return nil
}

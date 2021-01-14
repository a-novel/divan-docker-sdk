package d2sdk

import (
	"fmt"
	"github.com/a-novel/errors"
	"github.com/docker/docker/api/types"
)

func (dm *DivanManager) Remove() *errors.Error {
	if err := dm.cli.ContainerRemove(dm.ctx, dm.containerID, types.ContainerRemoveOptions{Force: true}); err != nil {
		return errors.New(ErrCannotRemoveContainer, fmt.Sprintf("unable to remove container : %s", err.Error()))
	}

	return nil
}

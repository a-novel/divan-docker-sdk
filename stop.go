package d2sdk

import (
	"fmt"
	"github.com/a-novel/errors"
)

func (dm *DivanManager) Stop() *errors.Error {
	if dm.status != StatusContainerRunning {
		return errors.New(ErrNoContainerRunning, "no container running")
	}

	if err := dm.cli.ContainerKill(dm.ctx, dm.containerID, ""); err != nil {
		return errors.New(ErrCannotKillContainer, fmt.Sprintf("unable to kill container : %s", err.Error()))
	}

	return nil
}

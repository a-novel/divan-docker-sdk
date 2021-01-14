package d2sdk

import (
	"bufio"
	"fmt"
	"github.com/a-novel/errors"
	"github.com/docker/docker/api/types"
)

func (dm *DivanManager) Exec(cmd string) (*bufio.Reader, *errors.Error) {
	if dm.status != StatusContainerRunning {
		return nil, errors.New(ErrNoContainerRunning, "no container running")
	}

	execID, err := dm.cli.ContainerExecCreate(dm.ctx, dm.containerID, types.ExecConfig{
		AttachStderr: true,
		Cmd:          []string{"sh", "-c", cmd},
	})
	if err != nil {
		return nil, errors.New(
			ErrCannotCreateCommand,
			fmt.Sprintf("cannot create cmd process on container %s : %s", dm.containerID, err.Error()),
		)
	}

	resp, err := dm.cli.ContainerExecAttach(dm.ctx, execID.ID, types.ExecStartCheck{
		Detach: false,
	})
	if err != nil {
		return nil, errors.New(
			ErrCannotRunCommand,
			fmt.Sprintf("cannot run cmd process on container %s : %s", dm.containerID, err.Error()),
		)
	}

	return resp.Reader, nil
}

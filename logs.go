package d2sdk

import (
	"fmt"
	"github.com/a-novel/errors"
	"github.com/docker/docker/api/types"
	"io/ioutil"
)

func (dm *DivanManager) Logs() (string, *errors.Error) {
	if dm.status == "" {
		return "", errors.New(ErrNoContainerRunning, "no container is running")
	}

	reader, err := dm.cli.ContainerLogs(dm.ctx, dm.containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Details:    true,
	})
	if err != nil {
		return "", errors.New(ErrCannotFetchLogs, fmt.Sprintf("cannot fetch logs : %s", err.Error()))
	}

	defer reader.Close()

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", errors.New(ErrCannotReadLogs, fmt.Sprintf("cannot read logs : %s", err.Error()))
	}

	return string(content), nil
}

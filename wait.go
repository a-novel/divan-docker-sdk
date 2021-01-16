package d2sdk

import (
	"github.com/a-novel/errors"
	"time"
)

/*
	Wait for configuration to complete.

	May return:
		- ErrExecutionError
		- ErrUnknownExecutionError
*/
func (dm *DivanManager) WaitForReady(timeout int) *errors.Error {
	tx := time.Now().UnixNano()

	for (time.Now().UnixNano() - tx) < int64(timeout*1000000000) {
		if dm.status != StatusContainerProcessing && dm.status != StatusContainerRunning {
			break
		}

		time.Sleep(time.Second)
	}

	if dm.status == StatusContainerConfigurationError {
		if dm.executionError != nil {
			return errors.New(
				ErrExecutionError,
				dm.executionError.Error(),
			)
		}

		return errors.New(
			ErrUnknownExecutionError,
			"unable to run configuration (unknown error)",
		)
	}

	return nil
}

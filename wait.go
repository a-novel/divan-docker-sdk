package d2sdk

import (
	"github.com/a-novel/errors"
)

/*
	Wait for configuration to complete.

	May return:
		- ErrExecutionError
		- ErrUnknownExecutionError
*/
func (dm *DivanManager) WaitForReady(timeout int) *errors.Error {
	status, err := dm.ClusterStatus(timeout)

	for err == nil && status == StatusContainerProcessing {
		status, err = dm.ClusterStatus(timeout)
	}

	if err != nil {
		return errors.New(
			ErrExecutionError,
			err.Error(),
		)
	}

	return nil
}

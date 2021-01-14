package d2sdk

import (
	"github.com/a-novel/errors"
)

/*
	Updates a running container.

	May return:
		- ErrNoContainerRunning
		- ErrCannotCreateCommand
		- ErrCannotRunCommand
*/
func (dm *DivanManager) Update() *errors.Error {
	if dm.status != StatusContainerRunning {
		return errors.New(ErrNoContainerRunning, "no container running")
	}

	if _, err := dm.Exec("cd /root/DIVAN_scripts && go run main.go"); err != nil {
		return err
	}

	return nil
}

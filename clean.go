package d2sdk

import "github.com/a-novel/errors"

/*
	Removes the container.

	May return:
		- ErrNoContainerRunning
		- ErrCannotKillContainer
		- ErrCannotRemoveContainer
*/
func (dm *DivanManager) Clean() *errors.Error {
	_ = dm.Stop()

	if err := dm.Remove(); err != nil {
		return err
	}

	return nil
}

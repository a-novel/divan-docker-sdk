package d2sdk

import "github.com/a-novel/errors"

/*
	Check if configuration is valid.

	May return:
		- ErrMissingConfigPath
*/
func (dm *DivanManager) Verify() *errors.Error {
	if dm.ConfigPath == "" {
		return errors.New(
			ErrMissingConfigPath,
			"missing path for config.json file",
		)
	}

	return nil
}

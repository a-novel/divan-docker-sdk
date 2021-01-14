package d2sdk

import "github.com/a-novel/errors"

type ConfigOptions struct {
	// Specify a version of Divan Docker. Leave blank to use latest stable version.
	Version       string `json:"version"`
	// Custom name for your container.
	ContainerName string `json:"containerName"`
	// Create a named volume to keep data consistent between shutdowns.
	VolumeName    string `json:"volumeName"`
	// Environment variable to pass to the container.
	Env           []string
}

/*
	Create a new manager object.

	May return:
		- ErrMissingConfigPath
		- ErrCannotRunDocker
*/
func New(configPath string, options *ConfigOptions) (*DivanManager, *errors.Error) {
	if options == nil {
		options = &ConfigOptions{}
	}

	output := &DivanManager{
		ConfigPath:    configPath,
		Version:       options.Version,
		ContainerName: options.ContainerName,
		VolumeName:    options.VolumeName,
	}

	if err := output.Verify(); err != nil {
		return nil, err
	}

	if err := output.Init(); err != nil {
		return nil, err
	}

	return output, nil
}

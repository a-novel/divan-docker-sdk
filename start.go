package d2sdk

import (
	"fmt"
	"github.com/a-novel/divan-docker-sdk/utils"
	"github.com/a-novel/errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
)

/*
	Runs the container.

	May return:
		- ErrContainerAlreadyRunning
		- ErrCannotPullImage
		- ErrCannotCreatePortMap
		- ErrCannotCreateContainer
		- ErrCannotRunContainer
*/
func (dm *DivanManager) Start(timeout int) *errors.Error {
	version := dm.Version
	if version == "" {
		version = "latest"
	}

	image := fmt.Sprintf("kushuh/divan:%s", version)

	// Pull container image.
	if _, err := dm.cli.ImagePull(dm.ctx, image, types.ImagePullOptions{}); err != nil {
		return errors.New(ErrCannotPullImage, fmt.Sprintf("cannot pull divan image : %s", err.Error()))
	}

	ports, err := utils.PortMapper("8091-8096:8091-8096", "11210-11211:11210-11211", "6666:8080")
	if err != nil {
		return err
	}

	volumes := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: dm.ConfigPath,
			Target: "/root/DIVAN_config/config.json",
		},
	}

	if dm.VolumeName != "" {
		volumes = append(volumes, mount.Mount{
			Type:   mount.TypeVolume,
			Source: dm.VolumeName,
			Target: "/opt/couchbase/var",
		})
	}

	resp, err2 := dm.cli.ContainerCreate(dm.ctx, &container.Config{
		Image: image,
		Env:   dm.Env,
	}, &container.HostConfig{
		PortBindings: ports,
		Mounts:       volumes,
	}, nil, nil, dm.ContainerName)

	if err2 != nil {
		return errors.New(ErrCannotCreateContainer, err2.Error())
	}

	dm.containerID = resp.ID

	if err := dm.cli.ContainerStart(dm.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return errors.New(ErrCannotRunContainer, err.Error())
	}

	if err := dm.WaitForReady(timeout); err != nil {
		return err
	}

	dm.status = StatusContainerReady

	return nil
}

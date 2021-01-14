# Divan Docker SDK

Go SDK providing a Docker manager for Divan Docker. Intended for tests and deploy scripts.

```cgo
go get github.com/a-novel/divan-docker-sdk
```

- [Divan Manager](#divan-manager)
    - [Manager options](#manager-options)
- [Starting and killing your container](#starting-and-killing-your-container)
- [Updating your container](#updating-your-container)
- [Supported Env variables](#supported-env-variables)

# Divan Manager

```go
package myPackage

import "github.com/a-novel/divan-docker-sdk"

func main() {
	manager, err := d2sdk.New("/path/to/config.json", nil)
	// Handle error here.
	
	// Run container.
	_ = manager.Start()
}
```

## Manager options

Divan manager can be created with the New function. This function takes a path to a [Divan Config file](
https://github.com/a-novel/divan-docker#create-a-config-file
), and an optional `*d2sdk.ConfigOptions` struct. Available options are:

| Key           | Type     | Description                                                                                                                         |
| :---          | :---     | :---                                                                                                                                |
| Version       | string   | Specify a version of Divan Docker to run                                                                                            |
| ContainerName | string   | Custom name for your container                                                                                                      |
| VolumeName    | string   | Name of an existing volume to keep Couchbase data                                                                                   |
| Env           | []string | A list of Env variables to pass to the container. See a list of supported ENV variables in [this section](#supported-env-variables) |

The `New` function may return any of the following errors:

- `d2sdk.ErrMissingConfigPath`
- `d2sdk.ErrCannotRunDocker`

# Starting and killing your container

Once you created your manager, you can run your container by using:

```go
manager.Start()
```

It may return any of the following errors:

- `d2sdk.ErrContainerAlreadyRunning`
- `d2sdk.ErrCannotPullImage`
- `d2sdk.ErrCannotCreatePortMap`
- `d2sdk.ErrCannotCreateContainer`
- `d2sdk.ErrCannotRunContainer`

You can wait for your configuration to complete with the following method:

```go
manager.WaitForReady(60)
```

The parameter is a timeout value in seconds.

It may return any of the following errors:

- `d2sdk.ErrExecutionError`
- `d2sdk.ErrUnknownExecutionError`

You can then delete your container by running:

```go
manager.Clean()
```

It may return any of the following errors:

- `d2sdk.ErrNoContainerRunning`
- `d2sdk.ErrCannotKillContainer`
- `d2sdk.ErrCannotRemoveContainer`

# Updating your container

Once running, update your configuration file and run:

```go
manager.Update()
```

Then if needed, run again the `WaitForReady` command.

It may return any of the following errors:

- `d2sdk.ErrNoContainerRunning`
- `d2sdk.ErrCannotCreateCommand`
- `d2sdk.ErrCannotRunCommand`

# Supported Env variables

`$NAME=$VALUE`

| Name | Values                                                                      |
| :--- | :---                                                                        |
| ENV  | Any value, `production` and `staging` will be recognized, otherwise ignored |

# License
2021, A-Novel [Apache 2.0 License](https://github.com/a-novel/divan-docker-sdk/blob/master/LICENSE).
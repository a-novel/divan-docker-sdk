package d2sdk

import "github.com/a-novel/divan-docker-sdk/utils"

const (
	ErrMissingConfigPath            = "err_missing_config_path"
	ErrContainerAlreadyRunning      = "err_container_already_running"
	ErrCannotRunDocker              = "err_cannot_run_docker"
	ErrCannotPullImage              = "err_cannot_pull_image"
	ErrCannotRunContainer           = "err_cannot_run_container"
	ErrCannotCreateContainer        = "err_cannot_create_container"
	ErrNoContainerRunning           = "err_no_container_running"
	ErrCannotFetchLogs              = "err_cannot_fetch_logs"
	ErrCannotReadLogs               = "err_cannot_read_logs"
	ErrCannotKillContainer          = "err_cannot_kill_container"
	ErrCannotRemoveContainer        = "err_cannot_remove_container"
	ErrCannotCreateCommand          = "err_cannot_create_command"
	ErrCannotRunCommand             = "err_cannot_run_command"
	ErrCannotReadConfigurationError = "err_cannot_read_configuration_error"
	ErrExecutionError               = "err_execution_error"
	ErrUnknownExecutionError        = "err_unknown_execution_error"
)

const (
	ErrCannotCreatePortMap = utils.ErrCannotCreatePortMap
)

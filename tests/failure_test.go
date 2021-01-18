package tests

import (
	"github.com/a-novel/divan-docker-sdk"
	"testing"
)

func TestFailure(t *testing.T) {
	_, err := d2sdk.New("", nil)
	if err == nil {
		t.Error("should fail with empty config path")
	} else if err.ID != d2sdk.ErrMissingConfigPath {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrMissingConfigPath)
	}

	divan, err := d2sdk.New("fakepath.json", nil)
	if err != nil {
		t.Fatalf("cannot create config : %s", err.Error())
	}

	defer divan.Clean()

	if err := divan.Start(120); err == nil {
		t.Error("should fail with wrong config path")
	} else if err.ID != d2sdk.ErrCannotCreateContainer {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrCannotCreateContainer)
	}

	if err := divan.Update(); err == nil {
		t.Error("should fail to update when no container is up")
	} else if err.ID != d2sdk.ErrNoContainerRunning {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrNoContainerRunning)
	}

	if err := divan.Stop(); err == nil {
		t.Error("should fail to stop when no container is up")
	} else if err.ID != d2sdk.ErrNoContainerRunning {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrNoContainerRunning)
	}

	if _, err := divan.Logs(); err == nil {
		t.Error("should fail to get logs when no container is up")
	} else if err.ID != d2sdk.ErrNoContainerRunning {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrNoContainerRunning)
	}

	if _, err := divan.Exec("ls -al"); err == nil {
		t.Error("should fail to exec command on container when no container is up")
	} else if err.ID != d2sdk.ErrNoContainerRunning {
		t.Errorf("unexpected error ID : got %s, expected %s", err.ID, d2sdk.ErrNoContainerRunning)
	}
}

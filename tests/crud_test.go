package tests

import (
	"fmt"
	"github.com/a-novel/divan-docker-sdk"
	"path"
	"testing"
)

func TestCRUD(t *testing.T) {
	ePath, err2 := getExecPath()
	if err2 != nil {
		t.Fatalf(err2.Error())
	}

	divan, err := d2sdk.New(fmt.Sprintf(path.Join(ePath, "configSample.json")), &d2sdk.ConfigOptions{
		ContainerName: "divan-test",
	})
	if err != nil {
		t.Fatalf("cannot create config : %s", err.Error())
	}

	if err := divan.Start(); err != nil {
		t.Fatalf("cannot start docker container : %s", err.Error())
	}

	if err := divan.WaitForReady(60); err != nil {
		t.Fatalf("cluster configuration failed : %s", err.Error())
	}

	divan.ConfigPath = fmt.Sprintf(path.Join(ePath, "configSampleUpdated.json"))

	if err := divan.Update(); err != nil {
		t.Errorf("cannot update docker container : %s", err.Error())
	}

	if err := divan.WaitForReady(60); err != nil {
		t.Fatalf("cluster configuration update failed : %s", err.Error())
	}

	if err := divan.Stop(); err != nil {
		t.Fatalf("cannot stop container : %s", err.Error())
	}

	if err := divan.Remove(); err != nil {
		t.Fatalf("cannot remove container : %s", err.Error())
	}
}

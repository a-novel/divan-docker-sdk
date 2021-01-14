package d2sdk

import (
	"encoding/json"
	"fmt"
	"github.com/a-novel/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type responseManager struct {
	Message string `json:"message"`
}

func get(u string) (string, int, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", 10, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if resp != nil {
		bd, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", resp.StatusCode, errors.New(
				ErrCannotReadConfigurationError,
				fmt.Sprintf("cannot parse configuration error : %s", err.Error()),
			)
		}

		parsedRes := responseManager{}
		if err := json.Unmarshal(bd, &parsedRes); err != nil {
			return "", 500, errors.New(
				ErrCannotReadConfigurationError,
				fmt.Sprintf("cannot read configuration error : %s\nresp: %s", err.Error(), string(bd)),
			)
		}

		return parsedRes.Message, resp.StatusCode, nil
	} else {
		return "nil response", -1, nil
	}
}

func (dm *DivanManager) Listen() {
	dm.status = StatusContainerRunning

	go func() {
		for {
			if dm.status == "" {
				return
			}

			res, status, err := get("http://localhost:7777")

			if err != nil {
				dm.status = StatusContainerConfigurationError
				dm.executionError = err
			}

			if err == nil && status == 102 {
				dm.status = StatusContainerProcessing
			}

			if err == nil && status == 500 {
				dm.status = StatusContainerConfigurationError
				dm.executionError = fmt.Errorf("unexpected response : %s", res)
			}

			if err == nil && status == 200 {
				dm.status = StatusContainerReady
			}

			if status >= 0 && status != 102 {
				return
			}

			time.Sleep(time.Second)
		}
	}()
}

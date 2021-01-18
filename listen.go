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
		return "", 0, err
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

func (dm *DivanManager) ClusterStatus(timeout int) (string, error) {
	_, status, _ := get("http://localhost:6666/divan_status")
	retry := 0
	for status == -1 && retry < timeout {
		retry++
		time.Sleep(time.Second)
		_, status, _ = get("http://localhost:6666/divan_status")
	}

	if status == -1 {
		return StatusContainerConfigurationError, fmt.Errorf("cannot reach backend (timeout)")
	}

	res, status, err := get("http://localhost:6666/divan_status")

	if err != nil {
		return StatusContainerConfigurationError, err
	}

	if status == 503 {
		return StatusContainerProcessing, nil
	}

	if status == 500 {
		return StatusContainerConfigurationError, fmt.Errorf("unexpected response : %s", res)
	}

	if status == 200 {
		return StatusContainerReady, nil
	}

	return "", fmt.Errorf("unexpected status %v", status)
}

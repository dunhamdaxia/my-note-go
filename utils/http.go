package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpPostJson(url string, data ...M) (res []byte, statusCode int, err error) {
	dataJson, err := json.Marshal(data[0])
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	if len(data) > 1 {
		if _, ok := data[1]["token"]; ok {
			req.Header.Set("token", data[1]["token"].(string))
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)

	statusCode = resp.StatusCode

	return
}

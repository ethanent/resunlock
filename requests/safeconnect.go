package requests

import (
	"errors"
	"net/http"
	"strconv"
)

func FetchSafeConnect(c *http.Client) error {
	resp, err := c.Get("https://resreg.ucsc.edu:9443/downloads/SafeConnectMacInstaller.zip")

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("unexpected status code " + strconv.Itoa(resp.StatusCode))
	}

	return resp.Body.Close()
}

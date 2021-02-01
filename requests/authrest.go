package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func SubmitAuthRest(c *http.Client) error {
	d, err := json.Marshal(map[string]string{
		"appversion": "5.0 (Macintosh)",
		"password":   "tnuser",
		"platform":   "MacIntel",
		"username":   "nguser",
	})

	if err != nil {
		return err
	}

	resp, err := c.Post("https://resreg.ucsc.edu:9443/api/authRest", "application/json", bytes.NewReader(d))

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("unexpected status code " + strconv.Itoa(resp.StatusCode))
	}

	return resp.Body.Close()
}

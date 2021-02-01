package requests

import (
	"net/http"
	"strings"
)

// CheckBlocked returns a boolean representing whether the device is logged out
func CheckBlocked(c *http.Client) (bool, error) {
	req, err := http.NewRequest("GET", "https://www.google.com/", nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:85.0) Gecko/20100101 Firefox/85.0")

	resp, err := c.Do(req)

	if err != nil {
		return false, err
	}

	return strings.Contains(resp.Header.Get("Server"), "Apache-Coyote"), nil
}

package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const createURLwithHTTPS = "https://%v"

// SendRequest adds 'https://' prefix to URL, reads request Body
// and returns result as []byte or err if something was wrong
func SendRequest(url *url.URL) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(createURLwithHTTPS, url.String()))
	if err != nil {
		return nil, fmt.Errorf("sendRequest Get error %w", err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

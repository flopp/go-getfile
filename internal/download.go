package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Download(httpClient *http.Client, userAgent string, url string) ([]byte, error) {
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(userAgent) > 0 {
		req.Header.Add("user-agent", userAgent)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		return nil, fmt.Errorf("Non-OK HTTP status: %d", response.StatusCode)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, response.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

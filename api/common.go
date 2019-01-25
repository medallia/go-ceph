package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type CustomLogger struct{}

func (cl *CustomLogger) Printf(template string, args ...interface{}) {
	// Only log [ERR] messages
	if strings.Contains(template, "[ERR]") {
		log.Printf(template, args...)
	}
}

type CephClient struct {
	BaseUrl string // e.g. http://<ceph-rest-api>:5000/v1/api/
}

func (cc *CephClient) callApi(endpoint string, method string) (string, error) {
	var body string
	endpoint = cc.BaseUrl + endpoint

	// Backoff configuration: 5 retries from 5 second to 2 minute
	// Put an arbitrarily timeout of 30 seconds to every request
	client := retryablehttp.NewClient()
	client.RetryWaitMin = 5 * time.Second
	client.RetryWaitMax = 2 * time.Minute
	client.RetryMax = 5
	client.HTTPClient.Timeout = 30 * time.Second
	client.Logger = &CustomLogger{}

	req, err := retryablehttp.NewRequest(method, endpoint, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	log.Printf("Sending request to ceph-rest-api with endpoint %s", endpoint)

	resp, err := client.Do(req)
	log.Printf("Got request response to ceph-rest-api with endpoint %s", endpoint)
	if err != nil {
		return body, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("Received unexpected status code from server: %d", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return string(bodyBytes), nil
}

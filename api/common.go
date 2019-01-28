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

var patterns = []string{
	"[ERR]",       // Error logs
	"retrying in", // Logs about retries
}

func (cl *CustomLogger) Printf(template string, args ...interface{}) {
	for _, pattern := range patterns {
		if strings.Contains(template, pattern) {
			log.Printf(template, args...)
			break
		}
	}
}

type CephClient struct {
	BaseUrl string // e.g. http://<ceph-rest-api>:5000/v1/api/
}

func (cc *CephClient) callApi(endpoint string, method string) (string, error) {
	var body string
	endpoint = cc.BaseUrl + endpoint

<<<<<<< HEAD
	// Backoff configuration: 7 retries from 5 second to 1 minute
	// 1º Retry: 5 seconds
	// 2º Retry: 10 seconds
	// 3º Retry: 20 seconds
	// 4º Retry: 40 seconds
	// 5º Retry: 1 minute
	// 6º Retry: 1 minute
	client.RetryWaitMin = 5 * time.Second
	client.RetryWaitMax = 1 * time.Minute
	client.RetryMax = 7
	client.HTTPClient.Timeout = 5 * time.Minute
	// Add a custom logger to only write ERR logs
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

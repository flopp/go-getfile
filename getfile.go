package getfile

import (
	"net/http"
	"time"

	"github.com/flopp/go-getfile/internal"
)

type Client struct {
	userAgent   string
	rateLimiter *time.Ticker
	httpClient  *http.Client
}

func NewClient(delay time.Duration) *Client {
	return &Client{"", nil, &http.Client{}}
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) SetRateLimit(delay time.Duration) {
	if c.rateLimiter != nil {
		c.rateLimiter.Stop()
		c.rateLimiter = nil
	}

	if delay > 0 {
		c.rateLimiter = time.NewTicker(delay)
	}
}

func (d *Client) GetIfOutdated(url, targetFilePath string, maxFileAge time.Duration) error {
	if internal.FileExists(targetFilePath) {
		if fileAge, err := internal.FileAge(targetFilePath); err != nil {
			if fileAge <= maxFileAge {
				return nil
			}
		}
	}

	return d.Get(url, targetFilePath)
}

func (d *Client) GetIfNotExists(url, targetFilePath string) error {
	if internal.FileExists(targetFilePath) {
		return nil
	}

	return d.Get(url, targetFilePath)
}

func (d *Client) Get(url, targetFilePath string) error {
	if d.rateLimiter != nil {
		<-d.rateLimiter.C
	}

	data, err := internal.Download(d.httpClient, d.userAgent, url)
	if err != nil {
		return err
	}

	return internal.WriteFile(targetFilePath, data)
}

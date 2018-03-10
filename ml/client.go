package ml

import (
	"net/url"
	"net/http"
	"github.com/pkg/errors"
	"context"
	"io"
	"path"
	"encoding/json"
)

type Client struct {
	URL *url.URL
	HTTPClient *http.Client
}

func NewClient(urlStr string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url %s", urlStr)
	}

	client := &http.Client{}

	return &Client{
		URL: parsedURL,
		HTTPClient: client,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method string, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
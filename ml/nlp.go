package ml

import (
	"context"
	"RestaurantSearcherAPI/models"
	"net/url"
	"strings"
	"github.com/pkg/errors"
)

func (c *Client) ParseKNP(ctx context.Context, review *models.Review) (*models.ParsedText, error) {
	spath := "/parse"
	values := url.Values{}

	values.Add("id", string(review.Id))
	values.Add("body", review.Body)

	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.Wrapf(err, "status code is not 200 but %s", res.StatusCode)
	}

	var parsedText models.ParsedText
	if err := decodeBody(res, &parsedText); err != nil {
		return nil, err
	}

	return &parsedText, nil
}

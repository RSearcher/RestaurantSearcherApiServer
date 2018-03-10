package ml

import (
	"context"
	"RestaurantSearcherAPI/models"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin/json"
	"bytes"
)

func (c *Client) ParseKNP(ctx context.Context, review *models.Review) (*models.ParsedText, error) {
	spath := "/parse"

	input, err := json.Marshal(review)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, "POST", spath, bytes.NewBuffer(input))
	req.Header.Add("Content-Type", "application/json")
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

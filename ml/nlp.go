package ml

import (
	"context"
	"RestaurantSearcherAPI/models"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin/json"
	"bytes"
	"net/http"
)

func (c *Client) SimilarTerms(ctx context.Context, query *models.Query) (*models.Terms, error) {
	spath := "/terms"

	input, err := json.Marshal(query)
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

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "status code is not 200 but %s", res.StatusCode)
	}

	var terms models.Terms
	if err := decodeBody(res, &terms); err != nil {
		return nil, err
	}

	return &terms, nil
}

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

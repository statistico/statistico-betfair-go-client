package betfair

import (
	"bytes"
	"context"
)

const (
	listCompetitions = "/listCompetitions/"
)

func (c *Client) ListCompetitions() ([]CompetitionResult, error) {
	var response []CompetitionResult

	body := []byte(`{"filter":{}}`)

	if err := c.getResource(context.Background(), bettingURL+listCompetitions, bytes.NewBuffer(body), &response); err != nil {
		return response, err
	}

	return response, nil
}
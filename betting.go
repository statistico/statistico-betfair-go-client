package betfair

import (
	"bytes"
	"encoding/json"
)

const (
	listCompetitions = "/listCompetitions/"
	listEventTypes = "/listEventTypes/"
)

func (c *Client) ListCompetitions(request ListCompetitionsRequest) ([]CompetitionResult, error) {
	var response []CompetitionResult

	body, err := json.Marshal(request)

	if err != nil {
		return response, err
	}

	if err := c.getResource(request.Context, bettingURL+listCompetitions, bytes.NewBuffer(body), &response); err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) ListEventTypes(request ListEventTypesRequest) ([]EventTypeResult, error) {
	var response []EventTypeResult

	body, err := json.Marshal(request)

	if err != nil {
		return response, err
	}

	if err := c.getResource(request.Context, bettingURL+listEventTypes, bytes.NewBuffer(body), &response); err != nil {
		return response, err
	}

	return response, nil
}

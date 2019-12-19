package betfair

import (
	"context"
)

const (
	listCompetitions    = "/listCompetitions/"
	listEvents          = "/listEvents/"
	listEventTypes      = "/listEventTypes/"
	listMarketCatalogue = "/listMarketCatalogue/"
)

func (c *Client) ListCompetitions(ctx context.Context, req ListCompetitionsRequest) ([]CompetitionResult, error) {
	var response []CompetitionResult

	if err := c.getResource(ctx, bettingURL+listCompetitions, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListEventTypes(ctx context.Context, req ListEventTypesRequest) ([]EventTypeResult, error) {
	var response []EventTypeResult

	if err := c.getResource(ctx, bettingURL+listEventTypes, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListEvents(ctx context.Context, req ListEventsRequest) ([]EventResult, error) {
	var response []EventResult

	if err := c.getResource(ctx, bettingURL+listEvents, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListMarketCatalogue(ctx context.Context, req ListMarketCatalogueRequest) ([]MarketCatalogue, error) {
	var response []MarketCatalogue

	if err := c.getResource(ctx, bettingURL+listMarketCatalogue, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

package betfair

import (
	"context"
)

const (
	listCompetitions    = "/listCompetitions/"
	listEvents          = "/listEvents/"
	listEventTypes      = "/listEventTypes/"
	listMarketBook      = "/listMarketBook/"
	listMarketCatalogue = "/listMarketCatalogue/"
	listRunnerBook      = "/listRunnerBook/"
	placeOrders         = "/placeOrders/"
)

// ListCompetition returns a slice of Competition struct (i.e., World Cup 2013) associated with the markets
// selected by the MarketFilter. Currently only Football markets have an associated competition.
func (c *client) ListCompetitions(ctx context.Context, req ListCompetitionsRequest) ([]CompetitionResult, error) {
	var response []CompetitionResult

	if err := c.getResource(ctx, BettingURL+listCompetitions, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListEventTypes returns a slice of EventTypeResult struct (i.e. Sports) associated with the markets selected by
// the MarketFilter.
func (c *client) ListEventTypes(ctx context.Context, req ListEventTypesRequest) ([]EventTypeResult, error) {
	var response []EventTypeResult

	if err := c.getResource(ctx, BettingURL+listEventTypes, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListEvents returns a slice of Event struct (i.e, Reading vs. Man United) associated with the markets selected
// by the MarketFilter.
func (c *client) ListEvents(ctx context.Context, req ListEventsRequest) ([]EventResult, error) {
	var response []EventResult

	if err := c.getResource(ctx, BettingURL+listEvents, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListMarketCatalogue returns a list of information about published (ACTIVE/SUSPENDED) markets that does not change
// (or changes very rarely). You use listMarketCatalogue to retrieve the name of the market, the names of selections
// and other information about markets.  Market Data Request Limits apply to requests made to listMarketCatalogue.
//
// Please note: listMarketCatalogue does not return markets that are CLOSED.
func (c *client) ListMarketCatalogue(ctx context.Context, req ListMarketCatalogueRequest) ([]MarketCatalogue, error) {
	var response []MarketCatalogue

	if err := c.getResource(ctx, BettingURL+listMarketCatalogue, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListMarketBook returns a list of dynamic data about markets. Dynamic data includes prices, the status of the market,
// the status of selections, the traded volume, and the status of any orders you have placed in the market.
//
// Please note: Separate requests should be made for OPEN & CLOSED markets. Request that include both OPEN & CLOSED
// markets will only return those markets that are OPEN.
func (c *client) ListMarketBook(ctx context.Context, req ListMarketBookRequest) ([]MarketBook, error) {
	var response []MarketBook

	if err := c.getResource(ctx, BettingURL+listMarketBook, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListRunnerBook return a list of dynamic data about a market and a specified runner. Dynamic data includes prices,
// the status of the market, the status of selections, the traded volume, and the status of any orders you have
// placed in the market.
func (c *client) ListRunnerBook(ctx context.Context, req ListRunnerBookRequest) ([]MarketBook, error) {
	var response []MarketBook

	if err := c.getResource(ctx, BettingURL+listRunnerBook, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// PlaceOrder executes a PlaceOrderRequest and returns a PlaceExecutionReport detailing the outcome of the
// transaction
func (c *client) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (*PlaceExecutionReport, error) {
	var response PlaceExecutionReport

	if err := c.getResource(ctx, BettingURL+placeOrders, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

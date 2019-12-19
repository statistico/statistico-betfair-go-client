package betfair

import (
	"context"
)

const (
	getAccountDetails = "/getAccountDetails/"
	getAccountFunds   = "/getAccountFunds/"
)

func (c *Client) AccountFunds(ctx context.Context) (*AccountFunds, error) {
	var response AccountFunds

	req, err := c.buildRequest(ctx, accountsURL+getAccountFunds, nil)

	if err != nil {
		return nil, err
	}

	if err := c.do(req, &response); err != nil {
		return &response, err
	}

	return &response, nil
}

func (c *Client) AccountDetails(ctx context.Context) (*AccountDetails, error) {
	var response AccountDetails

	req, err := c.buildRequest(ctx, accountsURL+getAccountDetails, nil)

	if err != nil {
		return nil, err
	}

	if err := c.do(req, &response); err != nil {
		return &response, err
	}

	return &response, nil
}

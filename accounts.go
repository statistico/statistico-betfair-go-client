package betfair

import (
	"context"
)

const (
	getAccountDetails = "/getAccountDetails/"
	getAccountFunds = "/getAccountFunds/"
)

func (c *Client) AccountFunds() (*AccountFunds, error) {
	var response AccountFunds

	if err := c.getResource(context.Background(), accountsURL+getAccountFunds, nil, &response); err != nil {
		return &response, err
	}

	return &response, nil
}

func (c *Client) AccountDetails() (*AccountDetails, error) {
	var response AccountDetails

	if err := c.getResource(context.Background(), accountsURL+getAccountDetails, nil, &response); err != nil {
		return &response, err
	}

	return &response, nil
}
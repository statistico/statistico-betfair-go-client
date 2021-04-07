package betfair

import (
	"context"
)

const (
	getAccountDetails = "/getAccountDetails/"
	getAccountFunds   = "/getAccountFunds/"
)

// AccountFunds returns the available to bet amount, exposure and commission information.
func (c *client) AccountFunds(ctx context.Context) (*AccountFunds, error) {
	var response AccountFunds

	if err := c.getResource(ctx, AccountsURL+getAccountFunds, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AccountDetails returns the details relating your account, including your discount rate and Betfair point balance.
func (c *client) AccountDetails(ctx context.Context) (*AccountDetails, error) {
	var response AccountDetails

	if err := c.getResource(ctx, AccountsURL+getAccountDetails, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

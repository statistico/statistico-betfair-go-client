package betfair

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	AccountsURL = "https://api.betfair.com/exchange/account/rest/v1.0"
	BettingURL  = "https://api.betfair.com/exchange/betting/rest/v1.0"
	LoginURL    = "https://identitysso.betfair.com/api/login"
)

type Client interface {
	// AccountFunds returns the available to bet amount, exposure and commission information.
	AccountFunds(ctx context.Context) (*AccountFunds, error)

	// AccountDetails returns the details relating your account, including your discount rate and Betfair point balance.
	AccountDetails(ctx context.Context) (*AccountDetails, error)

	// ListCompetitions returns a slice of Competition struct (i.e., World Cup 2013) associated with the markets
	// selected by the MarketFilter. Currently only Football markets have an associated competition.
	ListCompetitions(ctx context.Context, req ListCompetitionsRequest) ([]CompetitionResult, error)

	// ListEventTypes returns a slice of EventTypeResult struct (i.e. Sports) associated with the markets selected by
	// the MarketFilter.
	ListEventTypes(ctx context.Context, req ListEventTypesRequest) ([]EventTypeResult, error)

	// ListEvents returns a slice of Event struct (i.e, Reading vs. Man United) associated with the markets selected
	// by the MarketFilter.
	ListEvents(ctx context.Context, req ListEventsRequest) ([]EventResult, error)

	// ListMarketCatalogue returns a list of information about published (ACTIVE/SUSPENDED) markets that does not change
	// (or changes very rarely). You use listMarketCatalogue to retrieve the name of the market, the names of selections
	// and other information about markets.  Market Data Request Limits apply to requests made to listMarketCatalogue.
	//
	// Please note: listMarketCatalogue does not return markets that are CLOSED.
	ListMarketCatalogue(ctx context.Context, req ListMarketCatalogueRequest) ([]MarketCatalogue, error)

	// ListMarketBook returns a list of dynamic data about markets. Dynamic data includes prices, the status of the market,
	// the status of selections, the traded volume, and the status of any orders you have placed in the market.
	//
	// Please note: Separate requests should be made for OPEN & CLOSED markets. Request that include both OPEN & CLOSED
	// markets will only return those markets that are OPEN.
	ListMarketBook(ctx context.Context, req ListMarketBookRequest) ([]MarketBook, error)

	// ListRunnerBook return a list of dynamic data about a market and a specified runner. Dynamic data includes prices,
	// the status of the market, the status of selections, the traded volume, and the status of any orders you have
	// placed in the market.
	ListRunnerBook(ctx context.Context, req ListRunnerBookRequest) ([]MarketBook, error)

	// PlaceOrder executes a PlaceOrderRequest and returns a PlaceExecutionReport detailing the outcome of the
	// transaction
	PlaceOrder(ctx context.Context, req PlaceOrderRequest) (*PlaceExecutionReport, error)
}

type client struct {
	HTTPClient  *http.Client
	Credentials InteractiveCredentials
}

type InteractiveCredentials struct {
	Username string
	Password string
	Key      string
	Token    string
}

type session struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

func (c *client) createSession() error {
	body := fmt.Sprintf("username=%s&password=%s", c.Credentials.Username, c.Credentials.Password)

	req, err := http.NewRequest(http.MethodPost, LoginURL, strings.NewReader(body))

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Application", c.Credentials.Key)

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	var session session

	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return err
	}

	if session.Status != "SUCCESS" {
		return fmt.Errorf("error calling client %s", session.Error)
	}

	c.Credentials.Token = session.Token

	return nil
}

func (c *client) getResource(ctx context.Context, url string, req interface{}, res interface{}) error {
	request, err := c.buildRequest(ctx, url, req)

	if err != nil {
		return err
	}

	return c.do(request, res)
}

func (c *client) buildRequest(ctx context.Context, url string, body interface{}) (*http.Request, error) {
	parsed, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsed))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	if err := c.addHeaders(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (c *client) do(req *http.Request, resp interface{}) error {
	response, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if err = checkStatusCode(response); err != nil {
		return err
	}

	return parseJSONResponseBody(response.Body, resp)
}

func (c *client) addHeaders(req *http.Request) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", c.Credentials.Key)

	if c.Credentials.Token == "" {
		if err := c.createSession(); err != nil {
			return err
		}
	}

	req.Header.Set("X-Authentication", c.Credentials.Token)

	return nil
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		err := ErrorResponse{}

		e := parseJSONResponseBody(resp.Body, &err)

		if e != nil {
			return e
		}

		return fmt.Errorf(
			"error: FaultCode '%s' Fault '%s' ErrorCode '%s'",
			err.FaultCode,
			err.FaultString,
			err.Detail.AccountAPINGException.ErrorCode,
		)
	}

	return nil
}

func parseJSONResponseBody(body io.ReadCloser, response interface{}) error {
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return err
	}

	return nil
}

func NewClient(h *http.Client, c InteractiveCredentials) Client {
	return &client{
		HTTPClient:  h,
		Credentials: c,
	}
}

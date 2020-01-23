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

type Client struct {
	HTTPClient  *http.Client
	Credentials InteractiveCredentials
	BaseURLs    BaseURLs
}

type InteractiveCredentials struct {
	Username string
	Password string
	Key      string
	Token    string
}

type BaseURLs struct {
	Accounts string
	Betting  string
	Login    string
}

type session struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

func (c *Client) createSession() error {
	body := fmt.Sprintf("username=%s&password=%s", c.Credentials.Username, c.Credentials.Password)

	req, err := http.NewRequest(http.MethodPost, c.BaseURLs.Login, strings.NewReader(body))

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

func (c *Client) getResource(ctx context.Context, url string, req interface{}, res interface{}) error {
	request, err := c.buildRequest(ctx, url, req)

	if err != nil {
		return err
	}

	return c.do(request, res)
}

func (c *Client) buildRequest(ctx context.Context, url string, body interface{}) (*http.Request, error) {
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

func (c *Client) do(req *http.Request, resp interface{}) error {
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

func (c *Client) addHeaders(req *http.Request) error {
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

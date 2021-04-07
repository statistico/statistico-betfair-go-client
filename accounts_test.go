package betfair

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_AccountDetails(t *testing.T) {
	t.Run("returns account details struct", func(t *testing.T) {
		url := "https://api.betfair.com/exchange/account/rest/v1.0/getAccountDetails/"
		server := mockResponseServer(t, accountDetailsResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
		}

		details, err := client.AccountDetails(context.Background())

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, "GBP", details.CurrencyCode)
		assert.Equal(t, "Joe", details.FirstName)
		assert.Equal(t, "Sweeny", details.LastName)
		assert.Equal(t, "en", details.Locale)
		assert.Equal(t, "GBR", details.Region)
		assert.Equal(t, "Europe/London", details.Timezone)
		assert.Equal(t, float32(0.0), details.DiscountRate)
		assert.Equal(t, 20, details.PointsBalance)
		assert.Equal(t, "GB", details.CountryCode)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://api.betfair.com/exchange/account/rest/v1.0/getAccountDetails/"
		server := mockResponseServer(t, errorAccountsResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
		}

		details, err := client.AccountDetails(context.Background())

		if details != nil {
			t.Fatalf("Test failed, expected nil, got %+v", details)
		}

		str := "error: FaultCode 'Client' Fault 'AANGX-0002' ErrorCode 'INVALID_SESSION_INFORMATION'"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_AccountFunds(t *testing.T) {
	t.Run("returns account funds struct", func(t *testing.T) {
		url := "https://api.betfair.com/exchange/account/rest/v1.0/getAccountFunds/"
		server := mockResponseServer(t, accountFundsResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
		}

		funds, err := client.AccountFunds(context.Background())

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, float32(10.0), funds.Balance)
		assert.Equal(t, float32(-59.82), funds.Exposure)
		assert.Equal(t, float32(0.0), funds.RetainedCommission)
		assert.Equal(t, float32(-5000.00), funds.ExposureLimit)
		assert.Equal(t, float32(0.0), funds.DiscountRate)
		assert.Equal(t, 20, funds.PointBalance)
		assert.Equal(t, "UK", funds.Wallet)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://api.betfair.com/exchange/account/rest/v1.0/getAccountFunds/"
		server := mockResponseServer(t, errorAccountsResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
		}

		funds, err := client.AccountFunds(context.Background())

		if funds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", funds)
		}

		str := "error: FaultCode 'Client' Fault 'AANGX-0002' ErrorCode 'INVALID_SESSION_INFORMATION'"

		assert.Equal(t, str, err.Error())
	})
}

var accountFundsResponse = `
{
  "availableToBetBalance": 10.0,
  "exposure": -59.82,
  "retainedCommission": 0.0,
  "exposureLimit": -5000.0,
  "discountRate": 0.0,
  "pointsBalance": 20,
  "wallet": "UK"
}
`

var accountDetailsResponse = `
{
  "currencyCode": "GBP",
  "firstName": "Joe",
  "lastName": "Sweeny",
  "localeCode": "en",
  "region": "GBR",
  "timezone": "Europe/London",
  "discountRate": 0.0,
  "pointsBalance": 20,
  "countryCode": "GB"
}
`

var errorAccountsResponse = `
{
 "faultcode": "Client",
 "faultstring": "AANGX-0002",
 "detail": {
   "exceptionname": "AccountAPINGException",
   "AccountAPINGException": {
     "requestUUID": "null",
     "errorCode": "INVALID_SESSION_INFORMATION",
     "errorDetails": ""
   }
 }
}`

package betfair

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_ListCompetitions(t *testing.T) {
	t.Run("returns a slice of competition result struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listCompetitions/"
		server := mockResponseServer(t, competitionsResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		competitions, err := client.ListCompetitions(context.Background(), ListCompetitionsRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(competitions))
		assert.Equal(t, "12801", competitions[0].Competition.ID)
		assert.Equal(t, "Spanish Copa del Rey", competitions[0].Competition.Name)
		assert.Equal(t, 1, competitions[0].MarketCount)
		assert.Equal(t, "ESP", competitions[0].Region)

		assert.Equal(t, "23", competitions[1].Competition.ID)
		assert.Equal(t, "Danish Superliga", competitions[1].Competition.Name)
		assert.Equal(t, 13, competitions[1].MarketCount)
		assert.Equal(t, "DNK", competitions[1].Region)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listCompetitions/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		competitions, err := client.ListCompetitions(context.Background(), ListCompetitionsRequest{})

		if competitions != nil {
			t.Fatalf("Test failed, expected nil, got %+v", competitions)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_ListEventTypes(t *testing.T) {
	t.Run("returns a slice of event type result", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listEventTypes/"
		server := mockResponseServer(t, eventTypesResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		types, err := client.ListEventTypes(context.Background(), ListEventTypesRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(types))
		assert.Equal(t, "1", types[0].EventType.ID)
		assert.Equal(t, "Soccer", types[0].EventType.Name)
		assert.Equal(t, 12546, types[0].MarketCount)

		assert.Equal(t, "2", types[1].EventType.ID)
		assert.Equal(t, "Tennis", types[1].EventType.Name)
		assert.Equal(t, 112, types[1].MarketCount)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listEventTypes/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		types, err := client.ListEventTypes(context.Background(), ListEventTypesRequest{})

		if types != nil {
			t.Fatalf("Test failed, expected nil, got %+v", types)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_ListEvents(t *testing.T) {
	t.Run("returns a slice of event result struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listEvents/"
		server := mockResponseServer(t, eventsResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		events, err := client.ListEvents(context.Background(), ListEventsRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(events))
		assert.Equal(t, "29623290", events[0].Event.ID)
		assert.Equal(t, "Everton v Burnley", events[0].Event.Name)
		assert.Equal(t, "GB", events[0].Event.CountryCode)
		assert.Equal(t, "GMT", events[0].Event.Timezone)
		assert.Equal(t, "2019-12-26T15:00:00.000Z", events[0].Event.OpenDate)
		assert.Equal(t, 5, events[0].MarketCount)

		assert.Equal(t, "29615932", events[1].Event.ID)
		assert.Equal(t, "Everton v Arsenal", events[1].Event.Name)
		assert.Equal(t, "GB", events[1].Event.CountryCode)
		assert.Equal(t, "GMT", events[1].Event.Timezone)
		assert.Equal(t, "2019-12-21T12:30:00.000Z", events[1].Event.OpenDate)
		assert.Equal(t, 73, events[1].MarketCount)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listEvents/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		events, err := client.ListEvents(context.Background(), ListEventsRequest{})

		if events != nil {
			t.Fatalf("Test failed, expected nil, got %+v", events)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_ListMarketCatalogue(t *testing.T) {
	t.Run("returns a slice of market catalogue struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listMarketCatalogue/"
		server := mockResponseServer(t, marketCatalogueResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		catalogue, err := client.ListMarketCatalogue(context.Background(), ListMarketCatalogueRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(catalogue))
		assert.Equal(t, "1.160740937", catalogue[0].MarketID)
		assert.Equal(t, "Exit Dates - Boris Johnson 2", catalogue[0].MarketName)
		assert.Equal(t, float32(14349.0), catalogue[0].TotalMatched)

		assert.Equal(t, "1.129097136", catalogue[1].MarketID)
		assert.Equal(t, "Trump Exit Date", catalogue[1].MarketName)
		assert.Equal(t, float32(912082.02), catalogue[1].TotalMatched)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listMarketCatalogue/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		catalogue, err := client.ListMarketCatalogue(context.Background(), ListMarketCatalogueRequest{})

		if catalogue != nil {
			t.Fatalf("Test failed, expected nil, got %+v", catalogue)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_ListMarketBook(t *testing.T) {
	t.Run("returns a slice of market book struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listMarketBook/"
		server := mockResponseServer(t, marketBookResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		book, err := client.ListMarketBook(context.Background(), ListMarketBookRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, len(book))
		assert.Equal(t, "1.166260957", book[0].MarketID)
		assert.Equal(t, true, book[0].IsMarketDataDelayed)
		assert.Equal(t, "OPEN", book[0].Status)
		assert.Equal(t, uint32(0), book[0].BetDelay)
		assert.Equal(t, false, book[0].BSPReconciled)
		assert.Equal(t, true, book[0].Complete)
		assert.Equal(t, false, book[0].InPlay)
		assert.Equal(t, uint32(1), book[0].NumberOfWinners)
		assert.Equal(t, uint32(2), book[0].NumberOfRunners)
		assert.Equal(t, uint32(2), book[0].NumberOfActiveRunners)
		assert.Equal(t, "2019-12-20T20:14:58.282Z", book[0].LastMatchTime)
		assert.Equal(t, float32(2769.45), book[0].TotalMatched)
		assert.Equal(t, float32(76229.91), book[0].TotalAvailable)
		assert.Equal(t, true, book[0].CrossMatching)
		assert.Equal(t, false, book[0].RunnersVoidable)
		assert.Equal(t, 3089381622, book[0].Version)
		assert.Equal(t, 2, len(book[0].Runners))
		assert.Equal(t, 47972, book[0].Runners[0].SelectionID)
		assert.Equal(t, float32(0), book[0].Runners[0].Handicap)
		assert.Equal(t, "ACTIVE", book[0].Runners[0].Status)
		assert.Equal(t, float32(2.64), book[0].Runners[0].LastPriceTraded)
		assert.Equal(t, float32(0), book[0].Runners[0].TotalMatched)
		assert.Equal(t, 47973, book[0].Runners[1].SelectionID)
		assert.Equal(t, float32(0), book[0].Runners[1].Handicap)
		assert.Equal(t, "ACTIVE", book[0].Runners[1].Status)
		assert.Equal(t, float32(1.61), book[0].Runners[1].LastPriceTraded)
		assert.Equal(t, float32(0), book[0].Runners[1].TotalMatched)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listMarketBook/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		book, err := client.ListMarketBook(context.Background(), ListMarketBookRequest{})

		if book != nil {
			t.Fatalf("Test failed, expected nil, got %+v", book)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

func TestClient_ListRunnerBook(t *testing.T) {
	t.Run("returns a slice of market book struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listRunnerBook/"
		server := mockResponseServer(t, marketBookResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		book, err := client.ListRunnerBook(context.Background(), ListRunnerBookRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, len(book))
		assert.Equal(t, "1.166260957", book[0].MarketID)
		assert.Equal(t, true, book[0].IsMarketDataDelayed)
		assert.Equal(t, "OPEN", book[0].Status)
		assert.Equal(t, uint32(0), book[0].BetDelay)
		assert.Equal(t, false, book[0].BSPReconciled)
		assert.Equal(t, true, book[0].Complete)
		assert.Equal(t, false, book[0].InPlay)
		assert.Equal(t, uint32(1), book[0].NumberOfWinners)
		assert.Equal(t, uint32(2), book[0].NumberOfRunners)
		assert.Equal(t, uint32(2), book[0].NumberOfActiveRunners)
		assert.Equal(t, "2019-12-20T20:14:58.282Z", book[0].LastMatchTime)
		assert.Equal(t, float32(2769.45), book[0].TotalMatched)
		assert.Equal(t, float32(76229.91), book[0].TotalAvailable)
		assert.Equal(t, true, book[0].CrossMatching)
		assert.Equal(t, false, book[0].RunnersVoidable)
		assert.Equal(t, 3089381622, book[0].Version)
		assert.Equal(t, 2, len(book[0].Runners))
		assert.Equal(t, 47972, book[0].Runners[0].SelectionID)
		assert.Equal(t, float32(0), book[0].Runners[0].Handicap)
		assert.Equal(t, "ACTIVE", book[0].Runners[0].Status)
		assert.Equal(t, float32(2.64), book[0].Runners[0].LastPriceTraded)
		assert.Equal(t, float32(0), book[0].Runners[0].TotalMatched)
		assert.Equal(t, 47973, book[0].Runners[1].SelectionID)
		assert.Equal(t, float32(0), book[0].Runners[1].Handicap)
		assert.Equal(t, "ACTIVE", book[0].Runners[1].Status)
		assert.Equal(t, float32(1.61), book[0].Runners[1].LastPriceTraded)
		assert.Equal(t, float32(0), book[0].Runners[1].TotalMatched)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/listRunnerBook/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		book, err := client.ListRunnerBook(context.Background(), ListRunnerBookRequest{})

		if book != nil {
			t.Fatalf("Test failed, expected nil, got %+v", book)
		}

		str := "error: FaultCode 'Client' Fault 'DSC-0018' ErrorCode ''"

		assert.Equal(t, str, err.Error())
	})
}

var competitionsResponse = `
[
	{
    	"competition": {
      		"id": "12801",
      		"name": "Spanish Copa del Rey"
		},
		"marketCount": 1,
    	"competitionRegion": "ESP"
  	},
	{
    	"competition": {
      		"id": "23",
      		"name": "Danish Superliga"
    	},
    	"marketCount": 13,
    	"competitionRegion": "DNK"
  	}
]`

var eventTypesResponse = `
[
	{
    	"eventType": {
      		"id": "1",
      		"name": "Soccer"
		},
    	"marketCount": 12546
  	},
  	{
    	"eventType": {
      		"id": "2",
      		"name": "Tennis"
		},
    	"marketCount": 112
  	}
]
`

var eventsResponse = `
[
  {
    "event": {
      "id": "29623290",
      "name": "Everton v Burnley",
      "countryCode": "GB",
      "timezone": "GMT",
      "openDate": "2019-12-26T15:00:00.000Z"
    },
    "marketCount": 5
  },
  {
    "event": {
      "id": "29615932",
      "name": "Everton v Arsenal",
      "countryCode": "GB",
      "timezone": "GMT",
      "openDate": "2019-12-21T12:30:00.000Z"
    },
    "marketCount": 73
  }
]
`

var marketBookResponse = `
[
  {
    "marketId": "1.166260957",
    "isMarketDataDelayed": true,
    "status": "OPEN",
    "betDelay": 0,
    "bspReconciled": false,
    "complete": true,
    "inplay": false,
    "numberOfWinners": 1,
    "numberOfRunners": 2,
    "numberOfActiveRunners": 2,
    "lastMatchTime": "2019-12-20T20:14:58.282Z",
    "totalMatched": 2769.45,
    "totalAvailable": 76229.91,
    "crossMatching": true,
    "runnersVoidable": false,
    "version": 3089381622,
    "runners": [
      {
        "selectionId": 47972,
        "handicap": 0.0,
        "status": "ACTIVE",
        "lastPriceTraded": 2.64,
        "totalMatched": 0.0
      },
      {
        "selectionId": 47973,
        "handicap": 0.0,
        "status": "ACTIVE",
        "lastPriceTraded": 1.61,
        "totalMatched": 0.0
      }
    ]
  }
]
`

var marketCatalogueResponse = `
[
	{
    	"marketId": "1.160740937",
    	"marketName": "Exit Dates - Boris Johnson 2",
    	"totalMatched": 14349.0
  	},
  	{
    	"marketId": "1.129097136",
    	"marketName": "Trump Exit Date",
    	"totalMatched": 912082.02
  	}
]`

var errorBettingResponse = `
{
  "faultcode": "Client",
  "faultstring": "DSC-0018",
  "detail": {}
}
`

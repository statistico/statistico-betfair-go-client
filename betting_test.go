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
		server := mockResponseServer(t, listRunnerResponse, 200, url)

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
		assert.Equal(t, "1.167425856", book[0].MarketID)
		assert.Equal(t, true, book[0].IsMarketDataDelayed)
		assert.Equal(t, "OPEN", book[0].Status)
		assert.Equal(t, uint32(0), book[0].BetDelay)
		assert.Equal(t, false, book[0].BSPReconciled)
		assert.Equal(t, true, book[0].Complete)
		assert.Equal(t, false, book[0].InPlay)
		assert.Equal(t, uint32(1), book[0].NumberOfWinners)
		assert.Equal(t, uint32(2), book[0].NumberOfRunners)
		assert.Equal(t, uint32(2), book[0].NumberOfActiveRunners)
		assert.Equal(t, "2020-01-23T17:44:45.559Z", book[0].LastMatchTime)
		assert.Equal(t, float32(77184.93), book[0].TotalMatched)
		assert.Equal(t, float32(334223.54), book[0].TotalAvailable)
		assert.Equal(t, true, book[0].CrossMatching)
		assert.Equal(t, false, book[0].RunnersVoidable)
		assert.Equal(t, 3117607446, book[0].Version)
		assert.Equal(t, 1, len(book[0].Runners))
		assert.Equal(t, 47972, book[0].Runners[0].SelectionID)
		assert.Equal(t, float32(0), book[0].Runners[0].Handicap)
		assert.Equal(t, "ACTIVE", book[0].Runners[0].Status)
		assert.Equal(t, float32(2.02), book[0].Runners[0].LastPriceTraded)
		assert.Equal(t, float32(0), book[0].Runners[0].TotalMatched)
		assert.Equal(t, 3, len(book[0].Runners[0].EX.AvailableToBack))
		assert.Equal(t, float32(2.02), book[0].Runners[0].EX.AvailableToBack[0].Price)
		assert.Equal(t, float32(873.64), book[0].Runners[0].EX.AvailableToBack[0].Size)
		assert.Equal(t, float32(2.0), book[0].Runners[0].EX.AvailableToBack[1].Price)
		assert.Equal(t, float32(4444.81), book[0].Runners[0].EX.AvailableToBack[1].Size)
		assert.Equal(t, float32(1.99), book[0].Runners[0].EX.AvailableToBack[2].Price)
		assert.Equal(t, float32(675.28), book[0].Runners[0].EX.AvailableToBack[2].Size)
		assert.Equal(t, 3, len(book[0].Runners[0].EX.AvailableToLay))
		assert.Equal(t, float32(2.04), book[0].Runners[0].EX.AvailableToLay[0].Price)
		assert.Equal(t, float32(245.98), book[0].Runners[0].EX.AvailableToLay[0].Size)
		assert.Equal(t, float32(2.06), book[0].Runners[0].EX.AvailableToLay[1].Price)
		assert.Equal(t, float32(2437.16), book[0].Runners[0].EX.AvailableToLay[1].Size)
		assert.Equal(t, float32(2.08), book[0].Runners[0].EX.AvailableToLay[2].Price)
		assert.Equal(t, float32(789.3), book[0].Runners[0].EX.AvailableToLay[2].Size)
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

func TestClient_PlaceOrder(t *testing.T) {
	t.Run("returns a PlaceExecutionReport struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/placeOrders/"
		server := mockResponseServer(t, placeOrderSuccessResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		report, err := client.PlaceOrder(context.Background(), PlaceOrderRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		instruction := PlaceInstruction{
			OrderType:          "LIMIT",
			SelectionID:        47973,
			Side:               "BACK",
			LimitOrder:         LimitOrder{
				Size: 8.24,
				Price: 1.62,
				PersistenceType: "PERSIST",
			},
		}

		a.Equal("1.173957583", report.MarketID)
		a.Equal("SUCCESS", report.Status)
		a.Empty(report.ErrorCode)
		a.Empty(report.CustomerRef)
		a.Equal("SUCCESS", report.InstructionReports[0].Status)
		a.Empty(report.InstructionReports[0].ErrorCode)
		a.Equal("214568461952", report.InstructionReports[0].BetID)
		a.Equal("2020-10-22T10:46:57.000Z", report.InstructionReports[0].PlacedDate)
		a.Equal(float32(1.62), report.InstructionReports[0].AveragePriceMatched)
		a.Equal(float32(8.24), report.InstructionReports[0].SizeMatched)
		a.Equal("EXECUTION_COMPLETE", report.InstructionReports[0].OrderStatus)
		a.Equal(instruction, report.InstructionReports[0].Instruction)
	})

	t.Run("returns a failing PlaceExecutionReport struct", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/placeOrders/"
		server := mockResponseServer(t, placeOrderFailureResponse, 200, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		report, err := client.PlaceOrder(context.Background(), PlaceOrderRequest{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		instruction := PlaceInstruction{
			OrderType:          "LIMIT",
			SelectionID:        47972,
			Side:               "BACK",
			LimitOrder:         LimitOrder{
				Size: 2.0,
				Price: 39.65,
				PersistenceType: "PERSIST",
			},
		}

		a.Equal("1.173885527", report.MarketID)
		a.Equal("FAILURE", report.Status)
		a.Equal("BET_ACTION_ERROR", report.ErrorCode)
		a.Empty(report.CustomerRef)
		a.Equal("FAILURE", report.InstructionReports[0].Status)
		a.Equal("INVALID_ODDS", report.InstructionReports[0].ErrorCode)
		a.Empty(report.InstructionReports[0].BetID)
		a.Empty(report.InstructionReports[0].PlacedDate)
		a.Empty(report.InstructionReports[0].AveragePriceMatched)
		a.Empty(report.InstructionReports[0].SizeMatched)
		a.Empty(report.InstructionReports[0].OrderStatus)
		a.Equal(instruction, report.InstructionReports[0].Instruction)
	})

	t.Run("gracefully handles error response", func(t *testing.T) {
		url := "https://test/betting/rest/v1.0/placeOrders/"
		server := mockResponseServer(t, errorBettingResponse, 400, url)

		client := Client{
			HTTPClient:  server,
			Credentials: creds,
			BaseURLs:    base,
		}

		report, err := client.PlaceOrder(context.Background(), PlaceOrderRequest{})

		if report != nil {
			t.Fatalf("Test failed, expected nil, got %+v", report)
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

var listRunnerResponse = `[
  {
    "marketId": "1.167425856",
    "isMarketDataDelayed": true,
    "status": "OPEN",
    "betDelay": 0,
    "bspReconciled": false,
    "complete": true,
    "inplay": false,
    "numberOfWinners": 1,
    "numberOfRunners": 2,
    "numberOfActiveRunners": 2,
    "lastMatchTime": "2020-01-23T17:44:45.559Z",
    "totalMatched": 77184.93,
    "totalAvailable": 334223.54,
    "crossMatching": true,
    "runnersVoidable": false,
    "version": 3117607446,
    "runners": [
      {
        "selectionId": 47972,
        "handicap": 0.0,
        "status": "ACTIVE",
        "lastPriceTraded": 2.02,
        "totalMatched": 0.0,
        "ex": {
          "availableToBack": [
            {
              "price": 2.02,
              "size": 873.64
            },
            {
              "price": 2.0,
              "size": 4444.81
            },
            {
              "price": 1.99,
              "size": 675.28
            }
          ],
          "availableToLay": [
            {
              "price": 2.04,
              "size": 245.98
            },
            {
              "price": 2.06,
              "size": 2437.16
            },
            {
              "price": 2.08,
              "size": 789.3
            }
          ],
          "tradedVolume": []
        }
      }
    ]
  }
]`

var placeOrderSuccessResponse = `{
  "status": "SUCCESS",
  "marketId": "1.173957583",
  "instructionReports": [
    {
      "status": "SUCCESS",
      "instruction": {
        "selectionId": 47973,
        "limitOrder": {
          "size": 8.24,
          "price": 1.62,
          "persistenceType": "PERSIST"
        },
        "orderType": "LIMIT",
        "side": "BACK"
      },
      "betId": "214568461952",
      "placedDate": "2020-10-22T10:46:57.000Z",
      "averagePriceMatched": 1.62,
      "sizeMatched": 8.24,
      "orderStatus": "EXECUTION_COMPLETE"
    }
  ]
}`

var placeOrderFailureResponse = `{
  "status": "FAILURE",
  "errorCode": "BET_ACTION_ERROR",
  "marketId": "1.173885527",
  "instructionReports": [
    {
      "status": "FAILURE",
      "errorCode": "INVALID_ODDS",
      "instruction": {
        "selectionId": 47972,
        "limitOrder": {
          "size": 2.0,
          "price": 39.65,
          "persistenceType": "PERSIST"
        },
        "orderType": "LIMIT",
        "side": "BACK"
      }
    }
  ]
}`

var errorBettingResponse = `
{
  "faultcode": "Client",
  "faultstring": "DSC-0018",
  "detail": {}
}
`

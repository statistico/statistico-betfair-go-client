package betfair

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_Login(t *testing.T) {
	creds := InteractiveCredentials{
		Username: "joesweeny84@hotmail.com",
		Password: "WestHam66!",
		Key:      "tNEsm0ybTdHmMP4L",
	}

	client := Client{
		HTTPClient:  &http.Client{},
		Credentials: creds,
	}

	request := ListCompetitionsRequest{
		Filter: MarketFilter{
			TextQuery: "English Premier League",
		},
		Locale: "en",
	}

	competitions, err := client.ListCompetitions(context.Background(), request)

	if err != nil {
		t.Fatalf(err.Error())
	}

	premier := competitions[0]

	t.Errorf("Competition: %+v", premier.Competition.ID)

	r := ListEventsRequest{
		Filter: MarketFilter{
			TextQuery: "Newcastle",
			CompetitionIDs: []string{premier.Competition.ID},
			MarketStartTime: TimeRange{
				To: "2019-12-22T16:00:00.000Z",
			},
		},
	}

	events, err := client.ListEvents(context.Background(), r)

	if err != nil {
		t.Fatalf(err.Error())
	}

	event := events[0]


	c := ListMarketCatalogueRequest{
		Filter: MarketFilter{
			EventIDs: []string{event.Event.ID},
			MarketTypeCodes: []string{"OVER_UNDER_25"},
		},
		MaxResults: 100,
	}

	catalogue, err := client.ListMarketCatalogue(context.Background(), c)

	if err != nil {
		t.Fatalf(err.Error())
	}

	cat := catalogue[0]

	t.Fatalf("Catalogue: %+v", cat)

}

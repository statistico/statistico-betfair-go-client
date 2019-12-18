package betfair

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_Login(t *testing.T) {
	creds := Credentials{
		Username: "joesweeny84@hotmail.com",
		Password: "WestHam66!",
		Key:      "tNEsm0ybTdHmMP4L",
	}

	client := Client{
		HTTPClient:  &http.Client{},
		Credentials: creds,
	}

	request := ListEventTypesRequest{
		Context: context.Background(),
		Filter: MarketFilter{
			CompetitionIDs: []string{"10932509"},
		},
		Locale: "en",
	}

	competitions, _ := client.ListEventTypes(request)

	for _, comp := range competitions {
		t.Errorf("Event %+v\n", comp)
	}
}

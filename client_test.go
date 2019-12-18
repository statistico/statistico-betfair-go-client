package betfair

import (
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

	balance, _ := client.ListCompetitions()

	t.Errorf("Competitions %+v", balance)
}

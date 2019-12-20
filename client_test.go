package betfair

import (
	"testing"
)

func TestClient_Login(t *testing.T) {
	//creds := InteractiveCredentials{
	//	Username: "joesweeny84@hotmail.com",
	//	Password: "WestHam66!",
	//	Key:      "tNEsm0ybTdHmMP4L",
	//}
	//
	//base := BaseURLs{
	//	Accounts: "https://api.betfair.com/exchange/account/rest/v1.0",
	//	Betting:  "https://api.betfair.com/exchange/betting/rest/v1.0",
	//	Login:    "https://identitysso.betfair.com/api/login",
	//}
	//
	//client := Client{
	//	HTTPClient:  &http.Client{},
	//	Credentials: creds,
	//	BaseURLs: base,
	//}
	//
	//accounts, err := client.AccountFunds(context.Background())
	//
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	//
	//t.Fatalf("Accounts: %+v", accounts)
	//
	////request := ListCompetitionsRequest{
	////	Filter: MarketFilter{
	////		CompetitionIDs: []string{"10932509"},
	////	},
	////	Locale: "en",
	////}
	////
	////competitions, err := client.ListCompetitions(context.Background(), request)
	////
	////if err != nil {
	////	t.Fatalf(err.Error())
	////}
	////
	////premier := competitions[0]
	////
	////r := ListEventsRequest{
	////	Filter: MarketFilter{
	////		TextQuery:      "Newcastle",
	////		CompetitionIDs: []string{premier.Competition.ID},
	////		MarketStartTime: TimeRange{
	////			To: "2019-12-22T16:00:00.000Z",
	////		},
	////	},
	////}
	////
	////events, err := client.ListEvents(context.Background(), r)
	////
	////if err != nil {
	////	t.Fatalf(err.Error())
	////}
	////
	////event := events[0]
	////
	////c := ListMarketCatalogueRequest{
	////	Filter: MarketFilter{
	////		EventIDs:        []string{event.Event.ID},
	////		MarketTypeCodes: []string{"OVER_UNDER_25"},
	////	},
	////	Sort:       "FIRST_TO_START",
	////	MaxResults: 100,
	////}
	////
	////catalogue, err := client.ListMarketCatalogue(context.Background(), c)
	////
	////if err != nil {
	////	t.Fatalf(err.Error())
	////}
	////
	////cat := catalogue[0]
	////
	////b := ListMarketBookRequest{
	////	MarketIDs: []string{cat.MarketID},
	////}
	////
	////books, err := client.ListMarketBook(context.Background(), b)
	////
	////t.Fatalf("Books %+v", books)
}

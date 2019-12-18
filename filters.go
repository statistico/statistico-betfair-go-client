package betfair

type MarketFilter struct {
	TextQuery string `json:"textQuery"`
	ExchangeIDs []string `json:"exchangeIds"`
	EventTypeIDs []string `json:"eventTypeIds"`
	EventIDs []string `json:"eventIds"`
	CompetitionIDs []string `json:"competitionIds"`
	MarketIDs []string `json:"marketIds"`
	Venues []string `json:"venues"`
	BSPOnly bool `json:"bspOnly"`
	TurnInPlayEnabled bool `json:"turnInPlayEnabled"`
	InPlayOnly bool `json:"inPlayOnly"`
	MarketBettingTypes []string `json:"marketBettingTypes"`
	MarketCountries []string `json:"marketCountries"`
	MarketTypeCodes []string `json:"marketTypeCodes"`
	MarketStartTime TimeRange `json:"marketStartTime"`
	WithOrders []string `json:"withOrders"`
	RaceTypes []string `json:"raceTypes"`
}

type TimeRange struct {
	From string `json:"from"`
	To string `json:"to"`
}

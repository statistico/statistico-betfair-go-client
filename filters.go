package betfair

type MarketFilter struct {
	TextQuery          string    `json:"textQuery"`
	ExchangeIDs        []string  `json:"exchangeIds,omitempty"`
	EventTypeIDs       []string  `json:"eventTypeIds,omitempty"`
	EventIDs           []string  `json:"eventIds,omitempty"`
	CompetitionIDs     []string  `json:"competitionIds,omitempty"`
	MarketIDs          []string  `json:"marketIds,omitempty"`
	Venues             []string  `json:"venues,omitempty"`
	BSPOnly            bool      `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  bool      `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         bool      `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []string  `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string  `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string  `json:"marketTypeCodes,omitempty"`
	MarketStartTime    TimeRange `json:"marketStartTime,omitempty"`
	WithOrders         []string  `json:"withOrders,omitempty"`
	RaceTypes          []string  `json:"raceTypes,omitempty"`
}

type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

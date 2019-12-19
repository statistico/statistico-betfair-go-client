package betfair

type (
	AccountDetails struct {
		CurrencyCode  string  `json:"currencyCode"`
		FirstName     string  `json:"firstName"`
		LastName      string  `json:"lastName"`
		Locale        string  `json:"localeCode"`
		Region        string  `json:"region"`
		Timezone      string  `json:"timezone"`
		DiscountRate  float32 `json:"discountRate"`
		PointsBalance int     `json:"pointsBalance"`
		CountryCode   string  `json:"countryCode"`
	}

	AccountFunds struct {
		Balance            float32 `json:"availableToBetBalance"`
		DiscountRate       float32 `json:"discountRate"`
		Exposure           float32 `json:"exposure"`
		ExposureLimit      float32 `json:"exposureLimit"`
		PointBalance       int     `json:"pointsBalance"`
		RetainedCommission float32 `json:"retainedCommission"`
		Wallet             string  `json:"wallet"`
	}

	Competition struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	CompetitionResult struct {
		Competition Competition `json:"competition"`
		MarketCount int         `json:"marketCount"`
		Region      string      `json:"competitionRegion"`
	}

	Event struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		CountryCode string `json:"countryCode"`
		Timezone    string `json:"timezone"`
		Venue       string `json:"venue"`
		OpenDate    string `json:"openDate"`
	}

	EventResult struct {
		Event       Event `json:"event"`
		MarketCount int   `json:"marketCount"`
	}

	EventType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	EventTypeResult struct {
		EventType   EventType `json:"eventType"`
		MarketCount int       `json:"marketCount"`
	}

	MarketCatalogue struct {
		MarketID        string            `json:"marketId"`
		MarketName      string            `json:"marketName"`
		MarketStartTime string            `json:"marketStartTime"`
		Description     MarketDescription `json:"description"`
		TotalMatched    float32           `json:"totalMatched"`
		Runners         []RunnerCatalogue `json:"runners"`
		EventType       EventType         `json:"eventType"`
		Competition     Competition       `json:"competition"`
		Event           Event             `json:"event"`
	}

	MarketDescription struct {
		PersistenceEnabled     bool                   `json:"persistenceEnabled"`
		BSPMarket              bool                   `json:"bspMarket"`
		MarketTime             string                 `json:"marketTime"`
		SuspendTime            string                 `json:"suspendTime"`
		SettleTime             string                 `json:"settleTime"`
		BettingType            string                 `json:"bettingType"`
		TurnInPlayEnabled      bool                   `json:"turnInPlayEnabled"`
		MarketType             string                 `json:"marketType"`
		Regulator              string                 `json:"regulator"`
		MarketBaseRate         float32                `json:"marketBaseRate"`
		DiscountAllowed        bool                   `json:"discountAllowed"`
		Wallet                 string                 `json:"wallet"`
		Rules                  string                 `json:"rules"`
		RulesHasDate           bool                   `json:"rulesHasDate"`
		EachWayDivisor         float32                `json:"eachWayDivisor"`
		Clarifications         string                 `json:"clarifications"`
		LineRangeInfo          MarketLineRangeInfo    `json:"lineRangeInfo"`
		RaceType               string                 `json:"raceType"`
		PriceLadderDescription PriceLadderDescription `json:"priceLadderDescription"`
	}

	MarketLineRangeInfo struct {
		MaxUnitValue float32 `json:"maxUnitValue"`
		MinUnitValue float32 `json:"minUnitValue"`
		Interval     float32 `json:"interval"`
		MarketUnit   string  `json:"marketUnit"`
	}

	PriceLadderDescription struct {
		Type string `json:"type"`
	}

	RunnerCatalogue struct {
		SelectionID  uint64            `json:"selectionId"`
		RunnerName   string            `json:"runnerName"`
		Handicap     string            `json:"handicap"`
		SortPriority int               `json:"sortPriority"`
		MetaData     map[string]string `json:"metadata"`
	}
)

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

	ExchangePrices struct {
		AvailableToBack []PriceSize `json:"availableToBack"`
		AvailableToLay  []PriceSize `json:"availableToLay"`
		TradedVolume    []PriceSize `json:"tradedVolume"`
	}

	KeyLineDescription struct {
		KeyLine []KeyLineSelection `json:"keyLine"`
	}

	KeyLineSelection struct {
		SelectionID string  `json:"selectionId"`
		Handicap    float32 `json:"handicap"`
	}

	MarketBook struct {
		MarketID              string             `json:"marketId"`
		IsMarketDataDelayed   bool               `json:"isMarketDataDelayed"`
		Status                string             `json:"status"`
		BetDelay              uint32             `json:"betDelay"`
		BSPReconciled         bool               `json:"bspReconciled"`
		Complete              bool               `json:"complete"`
		InPlay                bool               `json:"inplay"`
		NumberOfWinners       uint32             `json:"numberOfWinners"`
		NumberOfRunners       uint32             `json:"numberOfRunners"`
		NumberOfActiveRunners uint32             `json:"numberOfActiveRunners"`
		LastMatchTime         string             `json:"lastMatchTime"`
		TotalMatched          float32            `json:"totalMatched"`
		TotalAvailable        float32            `json:"totalAvailable"`
		CrossMatching         bool               `json:"crossMatching"`
		RunnersVoidable       bool               `json:"runnersVoidable"`
		Version               int                `json:"version"`
		Runners               []Runner           `json:"runners"`
		KeyLineDescription    KeyLineDescription `json:"keyLineDescription"`
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

	Match struct {
		BetID     string  `json:"betId"`
		MatchID   string  `json:"matchId"`
		Side      string  `json:"side"`
		Price     float32 `json:"price"`
		Size      float32 `json:"size"`
		MatchDate string  `json:"matchDate"`
	}

	Order struct {
		BetID               string  `json:"betid"`
		OrderType           string  `json:"orderType"`
		Status              string  `json:"status"`
		PersistenceType     string  `json:"persistenceType"`
		Side                string  `json:"side"`
		Price               float32 `json:"price"`
		Size                float32 `json:"size"`
		BSPLiability        float32 `json:"bspLiability"`
		PlacedDate          string  `json:"placedDate"`
		AvgPriceMatched     float32 `json:"avgPriceMatched"`
		SizeMatched         float32 `json:"sizeMatched"`
		SizeRemaining       float32 `json:"sizeRemaining"`
		SizeLapsed          float32 `json:"sizeLapsed"`
		SizeCancelled       float32 `json:"sizeCancelled"`
		SizeVoided          float32 `json:"sizeVoided"`
		CustomerOrderRef    string  `json:"customerOrderRef"`
		CustomerStrategyRef string  `json:"customerStrategyRef"`
	}

	PriceLadderDescription struct {
		Type string `json:"type"`
	}

	PriceSize struct {
		Price float32 `json:"price"`
		Size  float32 `json:"size"`
	}

	RunnerCatalogue struct {
		SelectionID  uint64            `json:"selectionId"`
		RunnerName   string            `json:"runnerName"`
		Handicap     float32           `json:"handicap"`
		SortPriority int               `json:"sortPriority"`
		MetaData     map[string]string `json:"metadata"`
	}

	Runner struct {
		SelectionID      int            `json:"selectionId"`
		Handicap         float32        `json:"handicap"`
		Status           string         `json:"status"`
		AdjustmentFactor float32        `json:"adjustmentFactor"`
		LastPriceTraded  float32        `json:"lastPriceTraded"`
		TotalMatched     float32        `json:"totalMatched"`
		RemovalDate      string         `json:"removalDate"`
		SP               StartingPrices `json:"sp"`
		EX               ExchangePrices `json:"ex"`
		Orders           []Order        `json:"orders"`
		Matches          []Match        `json:"matches"`
		KeyLineDescription
	}

	StartingPrices struct {
		NearPrice         float32     `json:"nearPrice"`
		FarPRice          float32     `json:"farPrice"`
		BackStakeTaken    []PriceSize `json:"backStakeTaken"`
		LayLiabilityTaken []PriceSize `json:"layLiabilityTaken"`
		ActualSP          float32     `json:"actualSP"`
	}
)

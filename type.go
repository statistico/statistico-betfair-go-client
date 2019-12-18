package betfair

type (
	AccountDetails struct {
		CurrencyCode string `json:"currencyCode"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Locale string `json:"localeCode"`
		Region string `json:"region"`
		Timezone string `json:"timezone"`
		DiscountRate float32 `json:"discountRate"`
		PointsBalance int `json:"pointsBalance"`
		CountryCode string `json:"countryCode"`
	}
	AccountFunds struct {
		Balance            float32 `json:"availableToBetBalance"`
		DiscountRate       float32 `json:"discountRate"`
		Exposure           float32 `json:"exposure"`
		ExposureLimit      float32 `json:"exposureLimit"`
		PointBalance       int `json:"pointsBalance"`
		RetainedCommission float32 `json:"retainedCommission"`
		Wallet             string  `json:"wallet"`
	}

	Competition struct {
		ID string `json:"id"`
		Name string `json:"name"`
	}

	CompetitionResult struct {
		Competition Competition `json:"competition"`
		MarketCount int `json:"marketCount"`
		Region string `json:"competitionRegion"`
	}
)

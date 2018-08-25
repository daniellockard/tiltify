package tiltify

type Cause struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		LegalName string `json:"legalName"`
		Slug      string `json:"slug"`
		Currency  string `json:"currency"`
		About     string `json:"about"`
		Video     string `json:"video"`
		Image     struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
		Avatar struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
		Logo struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"logo"`
		Banner struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"banner"`
		ContactEmail       string `json:"contactEmail"`
		PaypalEmail        string `json:"paypalEmail"`
		PaypalCurrencyCode string `json:"paypalCurrencyCode"`
		Social             struct {
			Twitter   string `json:"twitter"`
			Youtube   string `json:"youtube"`
			Facebook  string `json:"facebook"`
			Instagram string `json:"instagram"`
		} `json:"social"`
		Settings struct {
			Colors struct {
				Background string `json:"background"`
				Highlight  string `json:"highlight"`
			} `json:"colors"`
			HeaderIntro     string `json:"headerIntro"`
			HeaderTitle     string `json:"headerTitle"`
			FooterCopyright string `json:"footerCopyright"`
			FindOutMoreLink string `json:"findOutMoreLink"`
		} `json:"settings"`
		Status             string `json:"status"`
		StripeConnected    bool   `json:"stripeConnected"`
		MailchimpConnected bool   `json:"mailchimpConnected"`
		Address            struct {
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
			City         string `json:"city"`
			Region       string `json:"region"`
			PostalCode   string `json:"postalCode"`
			Country      string `json:"country"`
		} `json:"address"`
	} `json:"data"`
}

type CauseCampaigns struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Thumbnail   struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnail"`
		CauseID            int         `json:"causeId"`
		UserID             int         `json:"userId"`
		TeamID             interface{} `json:"teamId"`
		FundraisingEventID int         `json:"fundraisingEventId"`
		Currency           string      `json:"currency"`
		Goal               int         `json:"goal"`
		OriginalGoal       int         `json:"originalGoal"`
		AmountRaised       int         `json:"amountRaised"`
		TotalAmountRaised  int         `json:"totalAmountRaised"`
		StartsOn           string      `json:"startsOn"`
		EndsOn             string      `json:"endsOn"`
	} `json:"data"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type CauseDonations struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID           int     `json:"id"`
		Amount       float64 `json:"amount"`
		Name         string  `json:"name"`
		DonorComment string  `json:"donorComment"`
		CreatedAt    int64   `json:"createdAt"`
	} `json:"data"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type CauseEvents struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Video       struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"video"`
		Image struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
		Avatar struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
		Logo struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"logo"`
		Banner struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"banner"`
		BannerTitle  string `json:"bannerTitle"`
		BannerIntro  string `json:"bannerIntro"`
		Currency     string `json:"currency"`
		Goal         int    `json:"goal"`
		AmountRaised int    `json:"amountRaised"`
		StartsOn     string `json:"startsOn"`
		EndsOn       string `json:"endsOn"`
		CauseID      int    `json:"causeId"`
	} `json:"data"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type CauseLeaderboard struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		Individual []struct {
			UserID       int    `json:"userId"`
			Name         string `json:"name"`
			URL          string `json:"url"`
			AmountRaised int    `json:"amount_raised"`
		} `json:"individual"`
		Team []struct {
			TeamID       int    `json:"teamId"`
			Name         string `json:"name"`
			URL          string `json:"url"`
			AmountRaised int    `json:"amount_raised"`
		} `json:"team"`
	} `json:"data"`
}

type CauseVisibilityOptions struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		Donate struct {
			Visible bool `json:"visible"`
		} `json:"donate"`
		IndividualLeaderboard struct {
			Visible bool   `json:"visible"`
			Type    string `json:"type"`
		} `json:"individual_leaderboard"`
		TeamLeaderboard struct {
			Visible bool   `json:"visible"`
			Type    string `json:"type"`
		} `json:"team_leaderboard"`
	} `json:"data"`
}

type CausePermissions struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		Activated                               bool `json:"activated"`
		CauseLeaderboard                        bool `json:"causeLeaderboard"`
		DashboardEnabled                        bool `json:"dashboardEnabled"`
		DashboardChart                          bool `json:"dashboardChart"`
		FundraisingEventsEnabled                bool `json:"fundraisingEventsEnabled"`
		FundraisingEventsLeaderboard            bool `json:"fundraisingEventsLeaderboard"`
		FundraisingEventsIncentives             bool `json:"fundraisingEventsIncentives"`
		FundraisingEventsSchedule               bool `json:"fundraisingEventsSchedule"`
		FundraisingEventsReporting              bool `json:"fundraisingEventsReporting"`
		FundraisingEventsReportingDonations     bool `json:"fundraisingEventsReportingDonations"`
		FundraisingEventsReportingCampaigns     bool `json:"fundraisingEventsReportingCampaigns"`
		FundraisingEventsReportingFundraisers   bool `json:"fundraisingEventsReportingFundraisers"`
		FundraisingEventsGeneralEnabled         bool `json:"fundraisingEventsGeneralEnabled"`
		FundraisingEventsGeneralDetails         bool `json:"fundraisingEventsGeneralDetails"`
		FundraisingEventsGeneralColors          bool `json:"fundraisingEventsGeneralColors"`
		FundraisingEventsGeneralImages          bool `json:"fundraisingEventsGeneralImages"`
		FundraisingEventsRegistrationEnabled    bool `json:"fundraisingEventsRegistrationEnabled"`
		FundraisingEventsVisibilityEnabled      bool `json:"fundraisingEventsVisibilityEnabled"`
		FundraisingEventsVisibilityDetails      bool `json:"fundraisingEventsVisibilityDetails"`
		FundraisingEventsVisibilityLeaderboards bool `json:"fundraisingEventsVisibilityLeaderboards"`
		AdminEnabled                            bool `json:"adminEnabled"`
		AdminGeneralEnabled                     bool `json:"adminGeneralEnabled"`
		AdminFinanceEnabled                     bool `json:"adminFinanceEnabled"`
		AdminBrandingEnabled                    bool `json:"adminBrandingEnabled"`
		AdminBrandingDetails                    bool `json:"adminBrandingDetails"`
		AdminBrandingColors                     bool `json:"adminBrandingColors"`
		AdminBrandingImages                     bool `json:"adminBrandingImages"`
		AdminIntegrationsEnabled                bool `json:"adminIntegrationsEnabled"`
		AdminVisibilityEnabled                  bool `json:"adminVisibilityEnabled"`
		AdminAPIEnabled                         bool `json:"adminApiEnabled"`
		ReportingEnabled                        bool `json:"reportingEnabled"`
		ReportingDonations                      bool `json:"reportingDonations"`
		ReportingCampaigns                      bool `json:"reportingCampaigns"`
		ReportingFundraisers                    bool `json:"reportingFundraisers"`
		ReportingChart                          bool `json:"reportingChart"`
	} `json:"data"`
}

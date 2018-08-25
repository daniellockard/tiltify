package tiltify

type FundraisingEvent struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		StartsAt     string      `json:"startsAt"`
		EndsAt       string      `json:"endsAt"`
		Slug         string      `json:"slug"`
		Description  string      `json:"description"`
		Image        interface{} `json:"image"`
		ContactEmail interface{} `json:"contactEmail"`
		Message      string      `json:"message"`
		Logo         struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"logo"`
		Video  interface{} `json:"video"`
		Avatar struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
		Banner struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"banner"`
		Parent struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"parent"`
		Status           string      `json:"status"`
		ShortDescription interface{} `json:"shortDescription"`
		Settings         struct {
			VisibilityOptions struct {
				StartsAt struct {
					Visible bool `json:"visible"`
				} `json:"startsAt"`
				EndsAt struct {
					Visible bool `json:"visible"`
				} `json:"endsAt"`
				FundraiserGoalAmount struct {
					Visible bool   `json:"visible"`
					Type    string `json:"type"`
				} `json:"fundraiserGoalAmount"`
				IndividualLeaderboard struct {
					Visible bool `json:"visible"`
				} `json:"individualLeaderboard"`
				TeamLeaderboard struct {
					Visible bool   `json:"visible"`
					Type    string `json:"type"`
				} `json:"teamLeaderboard"`
				DonorLeaderboard struct {
					Visible bool   `json:"visible"`
					Type    string `json:"type"`
				} `json:"donorLeaderboard"`
				PreventDonationsBeforeStart struct {
					Visible bool `json:"visible"`
				} `json:"preventDonationsBeforeStart"`
				Registration struct {
					Visible bool `json:"visible"`
				} `json:"registration"`
				SustainedGiving struct {
					Visible bool `json:"visible"`
				} `json:"sustainedGiving"`
				SustainedGivingStats struct {
					Visible bool `json:"visible"`
				} `json:"sustainedGivingStats"`
				Donate struct {
					Visible bool `json:"visible"`
				} `json:"donate"`
			} `json:"visibilityOptions"`
			RegistrationFields struct {
				Address struct {
					Position int  `json:"position"`
					Enabled  bool `json:"enabled"`
					Required bool `json:"required"`
				} `json:"address"`
				PhoneNumber struct {
					Position int  `json:"position"`
					Enabled  bool `json:"enabled"`
					Required bool `json:"required"`
				} `json:"phoneNumber"`
				ShirtSize struct {
					Position int  `json:"position"`
					Enabled  bool `json:"enabled"`
					Required bool `json:"required"`
				} `json:"shirtSize"`
				ServiceHours struct {
					Position int  `json:"position"`
					Enabled  bool `json:"enabled"`
					Required bool `json:"required"`
				} `json:"serviceHours"`
			} `json:"registrationFields"`
		} `json:"settings"`
		TotalAmountRaised    float64 `json:"totalAmountRaised"`
		CauseID              int     `json:"causeId"`
		FundraiserGoalAmount float64 `json:"fundraiserGoalAmount"`
		Type                 string  `json:"type"`
	} `json:"data"`
}

type FundraisingEventCampaigns struct {
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

type FundraisingEventDonations struct {
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

type FundraisingEventIncentives struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
		CreatedAt int64 `json:"createdAt"`
	} `json:"data"`
}

type FundraisingEventLeaderboards struct {
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

type FundraisingEventRegistratons struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID           int    `json:"id"`
		UserID       int    `json:"userId"`
		Email        string `json:"email"`
		Subscribed   bool   `json:"subscribed"`
		RegisteredAt string `json:"registeredAt"`
		Address      struct {
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
			City         string `json:"city"`
			State        string `json:"state"`
			PostalCode   string `json:"postalCode"`
		} `json:"address"`
		ShirtSize    string `json:"shirtSize"`
		ServiceHours bool   `json:"serviceHours"`
	} `json:"data"`
}

type FundraisingEventRegistrationFields struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		Address struct {
			Enabled bool `json:"enabled"`
		} `json:"address"`
		ShirtSize struct {
			Enabled bool `json:"enabled"`
		} `json:"shirt_size"`
		ServiceHours struct {
			Enabled bool `json:"enabled"`
		} `json:"service_hours"`
	} `json:"data"`
}

type FundraisingEventSchedule struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		StartsAt    int64  `json:"startsAt"`
	} `json:"data"`
}

type FundraisingEventVisibilityOptions struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
		StartsAt struct {
			Visible bool `json:"visible"`
		} `json:"startsAt"`
		EndsAt struct {
			Visible bool `json:"visible"`
		} `json:"endsAt"`
		FundraiserGoalAmount struct {
			Visible bool   `json:"visible"`
			Type    string `json:"type"`
		} `json:"fundraiserGoalAmount"`
		IndividualLeaderboard struct {
			Visible bool `json:"visible"`
		} `json:"individualLeaderboard"`
		TeamLeaderboard struct {
			Visible bool   `json:"visible"`
			Type    string `json:"type"`
		} `json:"teamLeaderboard"`
		DonorLeaderboard struct {
			Visible bool   `json:"visible"`
			Type    string `json:"type"`
		} `json:"donorLeaderboard"`
		PreventDonationsBeforeStart struct {
			Visible bool `json:"visible"`
		} `json:"preventDonationsBeforeStart"`
		Registration struct {
			Visible bool `json:"visible"`
		} `json:"registration"`
		Donate struct {
			Visible bool `json:"visible"`
		} `json:"donate"`
	} `json:"data"`
}

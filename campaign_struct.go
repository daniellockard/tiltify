package tiltify

type Campaign struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         string `json:"url"`
	StartsAt    int64  `json:"startsAt"`
	EndsAt      int64  `json:"endsAt"`
	Description string `json:"description"`
	Avatar      struct {
		Src    string `json:"src"`
		Alt    string `json:"alt"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"avatar"`
	CauseID                int     `json:"causeId"`
	FundraisingEventID     int     `json:"fundraisingEventId"`
	FundraiserGoalAmount   int     `json:"fundraiserGoalAmount"`
	OriginalGoalAmount     int     `json:"originalGoalAmount"`
	AmountRaised           float64 `json:"amountRaised"`
	SupportingAmountRaised float64 `json:"supportingAmountRaised"`
	TotalAmountRaised      float64 `json:"totalAmountRaised"`
	Supportable            bool    `json:"supportable"`
	Status                 string  `json:"status"`
	User                   struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Slug     string `json:"slug"`
		URL      string `json:"url"`
		Avatar   struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
	} `json:"user"`
	Team struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Slug     string `json:"slug"`
		URL      string `json:"url"`
		Avatar   struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
	} `json:"team"`
	Livestream struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	} `json:"livestream"`
}

type Donations struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID          int     `json:"id"`
		Amount      float64 `json:"amount"`
		Name        string  `json:"name"`
		Comment     string  `json:"comment"`
		CompletedAt int64   `json:"completedAt"`
		Sustained   bool    `json:"sustained"`
	} `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
		Self string `json:"self"`
	} `json:"links"`
}

type Rewards struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID                      int         `json:"id"`
		Name                    string      `json:"name"`
		Description             string      `json:"description"`
		Amount                  int         `json:"amount"`
		Kind                    string      `json:"kind"`
		Quantity                interface{} `json:"quantity"`
		Remaining               interface{} `json:"remaining"`
		FairMarketValue         int         `json:"fairMarketValue"`
		Currency                string      `json:"currency"`
		ShippingAddressRequired bool        `json:"shippingAddressRequired"`
		ShippingNote            interface{} `json:"shippingNote"`
		Image                   struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image"`
		Active    bool  `json:"active"`
		StartsAt  int   `json:"startsAt"`
		CreatedAt int64 `json:"createdAt"`
		UpdatedAt int64 `json:"updatedAt"`
	} `json:"data"`
}

type Polls struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Active     bool   `json:"active"`
		CampaignID int    `json:"campaignId"`
		CreatedAt  int64  `json:"createdAt"`
		UpdatedAt  int64  `json:"updatedAt"`
		Options    []struct {
			ID                int    `json:"id"`
			PollID            int    `json:"pollId"`
			Name              string `json:"name"`
			TotalAmountRaised int    `json:"totalAmountRaised"`
			CreatedAt         int64  `json:"createdAt"`
			UpdatedAt         int64  `json:"updatedAt"`
		} `json:"options"`
	} `json:"data"`
}

type Challenges struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Amount            int    `json:"amount"`
		TotalAmountRaised int    `json:"totalAmountRaised"`
		Active            bool   `json:"active"`
		ActivatesOn       int64  `json:"activatesOn"`
		CampaignID        int    `json:"campaignId"`
		EndsAt            int64  `json:"endsAt"`
		CreatedAt         int64  `json:"createdAt"`
		UpdatedAt         int64  `json:"updatedAt"`
	} `json:"data"`
}

type Schedule struct {
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

type SupportingCampaigns struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID          int    `json:"id"`
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
		Goal               int         `json:"goal"`
		OriginalGoal       int         `json:"originalGoal"`
		AmountRaised       int         `json:"amountRaised"`
		TotalAmountRaised  int         `json:"totalAmountRaised"`
		StartsOn           string      `json:"startsOn"`
		EndsOn             string      `json:"endsOn"`
	} `json:"data"`
}

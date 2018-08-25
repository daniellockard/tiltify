package tiltify

type User struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
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
		About             string  `json:"about"`
		TotalAmountRaised float64 `json:"totalAmountRaised"`
		Social            struct {
			Facebook interface{} `json:"facebook"`
			Twitch   interface{} `json:"twitch"`
			Twitter  interface{} `json:"twitter"`
			Website  interface{} `json:"website"`
			Youtube  interface{} `json:"youtube"`
			Mixer    interface{} `json:"mixer"`
		} `json:"social"`
	} `json:"data"`
}

type Users struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Slug      string `json:"slug"`
		Thumbnail struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnail"`
		Status string `json:"status"`
	} `json:"data"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type UserCampaigns struct {
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
		CauseID              int         `json:"causeId"`
		UserID               int         `json:"userId"`
		TeamID               interface{} `json:"teamId"`
		FundraisingEventID   int         `json:"fundraisingEventId"`
		Currency             string      `json:"currency"`
		Goal                 int         `json:"goal"`
		SupportingCampaignID int         `json:"supportingCampaignId"`
		OriginalGoal         int         `json:"originalGoal"`
		AmountRaised         int         `json:"amountRaised"`
		TotalAmountRaised    int         `json:"totalAmountRaised"`
		StartsOn             string      `json:"startsOn"`
		EndsOn               string      `json:"endsOn"`
	} `json:"data"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type UserCampaign struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data struct {
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
	} `json:"data"`
}

type UserOwnedTeams struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		URL        string `json:"url"`
		Bio        string `json:"bio"`
		InviteOnly bool   `json:"inviteOnly"`
		Disbanded  bool   `json:"disbanded"`
		Social     struct {
			Facebook string `json:"facebook"`
			Twitch   string `json:"twitch"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Youtube  string `json:"youtube"`
		} `json:"social"`
		Avatar struct {
			Src    string `json:"src"`
			Alt    string `json:"alt"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
	} `json:"data"`
}

type UserTeams struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		URL    string `json:"url"`
		Avatar struct {
			Alt    string `json:"alt"`
			Height int    `json:"height"`
			Src    string `json:"src"`
			Width  int    `json:"width"`
		} `json:"avatar"`
	} `json:"data"`
}

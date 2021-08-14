package data

type (
	RegulatoinResponse struct {
			Message string `json:"message"`
	}

	UserRequest struct {
			Name string `json:"name"`
			Device string `json:"device"`
			Version string `json:"version"`
	}
	UserResponse struct {
			UUID string `json:"uuid"`
	}

	ImageRequest struct {
			Quiz int `json:"quiz"`
	}
	ImageResponse struct {
			URL string `json:"url"`
	}
	
	BeaconRequest struct {
			Quiz int `json:"quiz"`
			Beacon []int `json:"beacon"`
	}
	BeaconResponse struct {
			ID int `json:"id"`
			Quiz int `json:"quiz"`
			URL string `json:"url"`
	}

	AnswerRequest struct {
		Quiz int `json:"quiz"`
		Answer string `json:"answer"`
	}
	AnswerResponse struct {
		Quiz int `json:"quiz"`
		Correct bool `json:"correct"`
	}
)
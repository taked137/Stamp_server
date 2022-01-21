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

	GoalResponse struct {
		Accept bool `json:"accept"`
	}

	InfoTitleResponse struct {
		ID int `json:"id"`
		Message string `json:"message"`
	}

	InfoRequest struct {
		ID int `json:"id"`
	}
	InfoResponse struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Category string `json:"category"`
		Message string `json:"message"`
	}

	EventResponse struct {
		Events []string `json:"events"`
	}

	ScheduleResponse struct {
		Time int `json:"time"`
		Event map[string]string `json:"events"`
	}

	CheckPoint struct {
		Num int `json:"num"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	MapResponse struct {
		Point []CheckPoint `json:"checkpoint"`
	}
)
package models


type AnalyticsResponse struct {
	Booking struct {
		Flight struct {
			Total  int `json:"total"`
			Status struct {
				Canceled int `json:"canceled"`
				Pending  int `json:"pending"`
			} `json:"status"`
		} `json:"flight"`

		Hotel struct {
			Total  int `json:"total"`
			Status struct {
				Canceled int `json:"canceled"`
				Pending  int `json:"pending"`
			} `json:"status"`
		} `json:"hotel"`
	} `json:"booking"`

	User struct {
		Total int `json:"total"`
	} `json:"user"`
}
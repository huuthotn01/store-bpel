package branch_service

type GetResponse struct {
	StatusCode int         `xml:"StatusCode"`
	Message    string      `xml:"Message"`
	Data       interface{} `xml:"Data"`
}

type UpdateResponse struct {
	StatusCode int    `xml:"StatusCode"`
	Message    string `xml:"Message"`
}

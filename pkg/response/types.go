package response

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []ErrorItem `json:"errors,omitempty"`
	Meta    MetaData    `json:"meta,omitempty"`
}

type ErrorItem struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type MetaData struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

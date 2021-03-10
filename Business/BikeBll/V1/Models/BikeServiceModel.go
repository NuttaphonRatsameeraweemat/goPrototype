package models

// BikeResponseModel struct{message_body string "json:\"message_body\""; message_title string "json:\"message_title\""; result string "json:\"result\""; status_code string "json:\"status_code\""; system_error_message string "json:\"system_error_message\""}
type BikeResponseModel struct {
	Message_body         string                  `json:"message_body"`
	Message_title        string                  `json:"message_title"`
	Result               []BikeResponseItemModel `json:"result"`
	Status_code          string                  `json:"status_code"`
	System_error_message string                  `json:"system_error_message"`
}

// BikeResponseItemModel struct{task_id string "json:\"task_id\""}
type BikeResponseItemModel struct {
	Task_id string `json:"task_id"`
}

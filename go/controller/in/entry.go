package in

type EntryCreateRequest struct {
	UserID int `json:"user_id"`
	JobID  int `json:"job_id"`
}

type EntryUpdateRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	JobID  int `json:"job_id"`
}

type EntryDeleteRequest struct {
	IDs []int `json:"ids"`
}

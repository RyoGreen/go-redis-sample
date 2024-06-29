package out

import "time"

type EntryResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	JobID     int       `json:"job_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

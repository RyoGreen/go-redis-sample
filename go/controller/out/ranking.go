package out

import "time"

type RankingResponse struct {
	Rank            int       `json:"rank"`
	JobID           int       `json:"job_id"`
	ApplicantsCount int       `json:"applicants_count"`
	UpdatedAt       time.Time `json:"updated_at"`
}

package domain

import (
	"context"
	"time"
)

type Ranking struct {
	Rank            int
	JobID           int
	ApplicantsCount int
	UpdatedAt       time.Time
}

type RankingRepository interface {
	Update(ctx context.Context, rankings []*Ranking) error
	List(ctx context.Context) ([]*Ranking, error)
}

func NewRanking(rank, jobID, ApplicantsCount int, updatedAt time.Time) *Ranking {
	return &Ranking{
		Rank:            rank,
		JobID:           jobID,
		ApplicantsCount: ApplicantsCount,
		UpdatedAt:       updatedAt,
	}
}

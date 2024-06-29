package domain

import "context"

type Ranking struct {
	ID    int
	Rank  int
	JobID int
}

type RankingRepository interface {
	Update(ctx context.Context) error
	List(ctx context.Context) ([]*Ranking, error)
}

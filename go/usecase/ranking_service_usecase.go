package usecase

import (
	"context"
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
	"go-redis/infrastructure/redis"
)

type RankingService interface {
	Update(ctx context.Context) error
	List(ctx context.Context) ([]*out.RankingResponse, error)
}

type RankingServiceImpl struct {
	RankingRepo domain.RankingRepository
	UserRepo    domain.UserRepository
	JobRepo     domain.JobRepository
}

func NewRankingService() RankingService {
	return &RankingServiceImpl{
		RankingRepo: redis.NewRedisRepository(),
		UserRepo:    postgres.NewUserPostgresRepository(),
		JobRepo:     postgres.NewJobPostgresRepository(),
	}
}

func (u *RankingServiceImpl) Update(ctx context.Context) error {
	return nil
}

func (u *RankingServiceImpl) List(ctx context.Context) ([]*out.RankingResponse, error) {
	rankings, err := u.RankingRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var rankingResponses []*out.RankingResponse
	for _, ranking := range rankings {
		var ranking = &out.RankingResponse{
			ID:    ranking.ID,
			Rank:  ranking.Rank,
			JobID: ranking.JobID,
		}
		rankingResponses = append(rankingResponses, ranking)
	}
	return rankingResponses, nil
}

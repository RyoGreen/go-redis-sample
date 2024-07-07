package usecase

import (
	"context"
	"go-redis/controller/out"
	"go-redis/domain"
	"go-redis/infrastructure/postgres"
	"go-redis/infrastructure/redis"
	"sort"
	"time"
)

type RankingService interface {
	Update(ctx context.Context) error
	List(ctx context.Context) ([]*out.RankingResponse, error)
}

type RankingServiceImpl struct {
	RankingRepo domain.RankingRepository
	EntryRepo   domain.EntryRepository
}

func NewRankingService() RankingService {
	return &RankingServiceImpl{
		RankingRepo: redis.NewRedisRepository(),
		EntryRepo:   postgres.NewEntryPostgresRepository(),
	}
}

const rankingMax = 3

func (u *RankingServiceImpl) Update(ctx context.Context) error {
	entries, err := u.EntryRepo.FindAll()
	if err != nil {
		return err
	}
	r := make(map[int]int)
	for _, entry := range entries {
		r[entry.JobID]++
	}

	type kv struct {
		Key   int
		Value int
	}
	var sortedData []kv
	for k, v := range r {
		sortedData = append(sortedData, kv{k, v})
	}

	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i].Value > sortedData[j].Value
	})

	now := time.Now()
	var rankings []*domain.Ranking
	for i := 0; i < rankingMax; i++ {
		rank := i + 1
		jobID := sortedData[i].Key
		ranking := domain.NewRanking(rank, jobID, sortedData[i].Value, now)
		rankings = append(rankings, ranking)
	}

	err = u.RankingRepo.Update(ctx, rankings)
	if err != nil {
		return err
	}

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
			Rank:            ranking.Rank,
			JobID:           ranking.JobID,
			ApplicantsCount: ranking.ApplicantsCount,
			UpdatedAt:       ranking.UpdatedAt,
		}
		rankingResponses = append(rankingResponses, ranking)
	}
	return rankingResponses, nil
}

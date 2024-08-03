package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis/domain"
	"sort"
	"time"
)

func NewRedisRepository() domain.RankingRepository {
	return &RedisRepository{}
}

type RedisRepository struct{}

type redisRanking struct {
	Rank            int       `json:"rank"`
	JobID           int       `json:"job_id"`
	ApplicantsCount int       `json:"applicants_count"`
	UpdatedAt       time.Time `json:"updated_at"`
}

const RankingKey = "ranking"

func (r *RedisRepository) Update(ctx context.Context, rankings []*domain.Ranking) error {
	pipe := client.TxPipeline()
	delCmd := pipe.Del(ctx, RankingKey)
	if delCmd.Err() != nil {
		return fmt.Errorf("failed to update ranking: %w", delCmd.Err())
	}
	var rdsRankings = make([]redisRanking, 0, len(rankings))
	for _, v := range rankings {
		rdsRankings = append(rdsRankings, redisRanking{
			Rank:            v.Rank,
			JobID:           v.JobID,
			ApplicantsCount: v.ApplicantsCount,
			UpdatedAt:       v.UpdatedAt,
		})
	}

	var data []interface{}
	for _, v := range rdsRankings {
		dataJson, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal ranking: %w", err)
		}
		data = append(data, string(dataJson))

	}
	saddCmd := pipe.SAdd(ctx, RankingKey, data...)
	if saddCmd.Err() != nil {
		return fmt.Errorf("failed to update ranking: %w", saddCmd.Err())
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to update ranking: %w", err)
	}

	return nil
}

func (r *RedisRepository) List(ctx context.Context) ([]*domain.Ranking, error) {
	cmd := client.SMembers(ctx, RankingKey)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("failed to get ranking: %w", cmd.Err())
	}
	rankings := make([]*redisRanking, 0)
	for _, v := range cmd.Val() {
		r := new(redisRanking)
		err := json.Unmarshal([]byte(v), r)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal ranking: %w", err)
		}
		rankings = append(rankings, r)
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Rank < rankings[j].Rank
	})

	var rankingsDomain = make([]*domain.Ranking, 0, len(rankings))
	for _, v := range rankings {
		rankingsDomain = append(rankingsDomain, &domain.Ranking{
			Rank:            v.Rank,
			JobID:           v.JobID,
			ApplicantsCount: v.ApplicantsCount,
			UpdatedAt:       v.UpdatedAt,
		})
	}

	return rankingsDomain, nil
}

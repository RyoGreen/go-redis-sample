package redis

import (
	"context"
	"go-redis/domain"
)

func NewRedisRepository() domain.RankingRepository {
	return &RedisRepository{}
}

type RedisRepository struct {
}

func (r *RedisRepository) Update(ctx context.Context) error {
	return nil
}

func (r *RedisRepository) List(ctx context.Context) ([]*domain.Ranking, error) {
	return nil, nil
}

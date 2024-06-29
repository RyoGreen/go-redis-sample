package redis

import (
	"context"
	"go-redis/domain"
)

func NewRedisRepository() domain.EntryCache {
	return &RedisRepository{}
}

type RedisRepository struct {
}

func (r *RedisRepository) GetAll(ctx context.Context) ([]*domain.Entry, error) {
	return []*domain.Entry{}, nil
}

func (r *RedisRepository) GetByID(ctx context.Context, id int) (*domain.Entry, error) {
	return &domain.Entry{}, nil
}

func (r *RedisRepository) Update(ctx context.Context) error {
	return nil
}

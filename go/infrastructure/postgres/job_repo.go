package postgres

import "go-redis/domain"

func NewJobPostgresRepository() domain.JobRepository {
	return &JobRepositoryImpl{}
}

type JobRepositoryImpl struct {
}

func (r *JobRepositoryImpl) FindAll() ([]domain.Job, error) {
	return nil, nil
}

func (r *JobRepositoryImpl) FindById(id int) (domain.Job, error) {
	return domain.Job{}, nil
}

func (r *JobRepositoryImpl) Create(user domain.Job) (domain.Job, error) {
	return domain.Job{}, nil
}

func (r *JobRepositoryImpl) Update(user domain.Job) (domain.Job, error) {
	return domain.Job{}, nil
}

func (r *JobRepositoryImpl) Delete(id int) error {
	return nil
}

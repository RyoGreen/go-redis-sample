package repository

type JobRepository interface {
}

type JobRepositoryImpl struct {
}

func NewJobRepository() JobRepository {
	return &JobRepositoryImpl{}
}

package postgres

import (
	"go-redis/domain"
	"time"
)

func NewJobPostgresRepository() domain.JobRepository {
	return &JobRepositoryImpl{}
}

type JobRepositoryImpl struct{}

func (r *JobRepositoryImpl) FindAll() ([]*domain.Job, error) {
	sql := `SELECT * FROM jobs order by id desc`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobs []*domain.Job
	for rows.Next() {
		job := domain.Job{}
		err := rows.Scan(&job.ID, &job.Name, &job.Description, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func (r *JobRepositoryImpl) FindById(id int) (*domain.Job, error) {
	sql := `SELECT * FROM jobs WHERE id = $1`
	row := db.QueryRow(sql, id)
	job := &domain.Job{}
	err := row.Scan(&job.ID, &job.Name, &job.Description, &job.CreatedAt, &job.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (r *JobRepositoryImpl) Create(job *domain.Job) (*domain.Job, error) {
	sql := `INSERT INTO jobs (name, description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	now := time.Now()
	err := db.QueryRow(sql, job.Name, job.Description, now, now).Scan(&job.ID, &job.CreatedAt, &job.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &domain.Job{}, nil
}

func (r *JobRepositoryImpl) Update(job *domain.Job) (*domain.Job, error) {
	sql := `UPDATE jobs SET name=$1, description=$2, updated_at=$3 WHERE id=$4 RETURNING name,description,updated_at`
	now := time.Now()
	err := db.QueryRow(sql, job.Name, job.Description, now, job.ID).Scan(&job.Name, &job.Description, &job.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &domain.Job{}, nil
}

func (r *JobRepositoryImpl) Delete(ids []int) error {
	sql := `DELETE FROM jobs WHERE id = ANY($1)`
	_, err := db.Exec(sql, ids)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"go-redis/controller/in"
	"go-redis/controller/out"
	"go-redis/domain"
)

type JobUsecase interface {
	List() ([]*out.JobResponse, error)
	Get(id int) (*out.JobResponse, error)
	Create(job *in.JobCreateRequest) (*out.JobResponse, error)
	Update(job *in.JobUpdateRequest) (*out.JobResponse, error)
	Delete(ids *in.JobDeleteRequest) error
}

type JobUsecaseImpl struct {
	JobRepo domain.JobRepository
}

func NewJobUsecase() JobUsecase {
	return &JobUsecaseImpl{}
}

func (u *JobUsecaseImpl) List() ([]*out.JobResponse, error) {
	jobs, err := u.JobRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var jobResponses []*out.JobResponse
	for _, job := range jobs {
		jobResponses = append(jobResponses, &out.JobResponse{
			ID:          job.ID,
			Name:        job.Name,
			Description: job.Description,
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   job.UpdatedAt,
		})
	}
	return jobResponses, nil
}

func (u *JobUsecaseImpl) Get(id int) (*out.JobResponse, error) {
	job, err := u.JobRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &out.JobResponse{
		ID:          job.ID,
		Name:        job.Name,
		Description: job.Description,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}, nil
}

func (u *JobUsecaseImpl) Create(job *in.JobCreateRequest) (*out.JobResponse, error) {
	j := domain.NewJob(0, job.Name, job.Description)
	createdJob, err := u.JobRepo.Create(j)
	if err != nil {
		return nil, err
	}
	return &out.JobResponse{
		ID:          createdJob.ID,
		Name:        createdJob.Name,
		Description: createdJob.Description,
		CreatedAt:   createdJob.CreatedAt,
		UpdatedAt:   createdJob.UpdatedAt,
	}, nil
}

func (u *JobUsecaseImpl) Update(job *in.JobUpdateRequest) (*out.JobResponse, error) {
	j := domain.NewJob(job.ID, job.Name, job.Description)
	updatedJob, err := u.JobRepo.Update(j)
	if err != nil {
		return nil, err
	}
	return &out.JobResponse{
		ID:          updatedJob.ID,
		Name:        updatedJob.Name,
		Description: updatedJob.Description,
		CreatedAt:   updatedJob.CreatedAt,
		UpdatedAt:   updatedJob.UpdatedAt,
	}, nil
}

func (u *JobUsecaseImpl) Delete(ids *in.JobDeleteRequest) error {
	err := u.JobRepo.Delete(ids.ID)
	if err != nil {
		return err
	}
	return nil
}

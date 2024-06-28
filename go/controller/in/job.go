package in

type JobCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type JobUpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type JobDeleteRequest struct {
	ID []int `json:"ids"`
}

package in

type DeleteUserRequest struct {
	ID []int `json:"ids"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type UpdateUserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

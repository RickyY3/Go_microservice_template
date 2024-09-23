package model

type UpsertUserRequest struct {
	Username *string
	Pass     *string
	Email    *string
}

type GetUserResponse struct {
	
}

type User struct {
	Username string
	Pass     string
	Email    string
}

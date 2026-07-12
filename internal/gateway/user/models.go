package user

type UserRequest struct {
	ID       int
	Name     string
	Age      int32
	Email    string
	Phone    string
	Password string
}

type UserResponse struct {
	ID       int
	Name     string
	Age      int32
	Email    string
	Phone    string
	Password string
}

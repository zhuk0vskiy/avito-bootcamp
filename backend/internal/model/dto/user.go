package dto

type SignUpRequest struct {
	Email    string
	Password string
	Role     string
}

// type SignUpResponse struct {

// }

type LogInRequest struct {
	Email    string
	Password string
}

type LogInResponse struct {
	Token string
}



package domain

type CreateUserErrorCode int

const (
	CreateErrorConflict CreateUserErrorCode = 409
)

type UserQrCodeInput struct {
}

type UserQrCodeOutput struct {
	Created   bool
	ErrorCode CreateUserErrorCode
	Details   string
}

type CreateUserInput struct {
}

type CreateUserOutput struct {
}

type ReadUserInput struct {
}

type ReadUserOutput struct {
}

type UpdateUserInput struct {
}

type UpdateUserOutput struct {
}

type DeleteUserInput struct {
}

type DeleteUserOutput struct {
}

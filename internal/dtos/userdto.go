package dtos

type UserCreateDto struct {
	Username string
	Password string
	Email    string
}

type UserUpdateDto struct {
	Password string
	Email    string
}

type SearchCriteria interface {
	Match(user UserCreateDto) bool
}

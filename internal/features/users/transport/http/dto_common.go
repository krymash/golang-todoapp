package users_transport_http

import "github.com/krymash/golang-todoapp/internal/core/domain"

type UserDtoResponse struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDTOFromDomain(user domain.User) UserDtoResponse {
	return UserDtoResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDtoResponse {
	usersDTO := make([]UserDtoResponse, len(users))
	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}
	return usersDTO

}

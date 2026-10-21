package users_postgres_repository

import "github.com/krymash/golang-todoapp/internal/core/domain"

type UserModel struct {
	ID           int
	Version      int
	Full_name    string
	Phone_Number string
}

func UserDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Full_name,
			user.Phone_Number,
		)
	}
	return userDomains
}

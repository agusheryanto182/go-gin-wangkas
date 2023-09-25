package user

type Service interface {
	LoginUser(input User) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LoginUser(input User) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Password != password {
		return user, err
	}

	return user, nil
}

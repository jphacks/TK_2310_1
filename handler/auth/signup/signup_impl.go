package signup

import DBRepository "github.com/giraffe-org/backend/repository/db"

type signupImpl struct {
	db DBRepository.DB
}

func New() Signup {
	return &signupImpl{
		db: DBRepository.New(),
	}
}

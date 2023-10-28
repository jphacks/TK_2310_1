package signup

import DBRepository "github.com/jphacks/TK_2310_1/repository/db"

type signupImpl struct {
	db DBRepository.DB
}

func New() Signup {
	return &signupImpl{
		db: DBRepository.New(),
	}
}

package auth

import "chetam/internal/model"

func (a *Auth) CreateUser(req model.RegisterRequest) (string, error) {
	err := a.repo.CreateUser(req.Email, req.Login, req.Password)
	if err != nil {
		return "", err
	}

	token, err := a.generateJWT(req.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

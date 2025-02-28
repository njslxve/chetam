package repository

import "context"

func (r *Repository) CreateUser(email, login, password string) error {
	querry := qb.Insert("users").
		Columns(
			"login",
			"email",
			"password",
		).
		Values(
			login,
			email,
			password,
		).
		Suffix("returnig id")

	sql, args, err := querry.ToSql()
	if err != nil {
		return err
	}

	var id int

	err = r.db.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

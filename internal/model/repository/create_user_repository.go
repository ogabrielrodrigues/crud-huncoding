package repository

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

var create_query = `INSERT INTO users (name, email, age, password) VALUES ($1, $2, $3, $4) RETURNING id`

func (ur *userRepository) CreateUser(user_domain model.UserDomainInterface) (model.UserDomainInterface, *rest.RestErr) {
	conn, err := ur.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	defer conn.Close()

	row := conn.QueryRow(
		create_query,
		user_domain.GetName(),
		user_domain.GetEmail(),
		user_domain.GetAge(),
		user_domain.GetPassword(),
	)

	if row.Err() != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var id string

	if err := row.Scan(&id); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	user_domain.SetID(id)

	return user_domain, nil
}

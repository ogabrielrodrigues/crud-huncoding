package repository

import (
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

var find_by_id_query = `SELECT name, email, age, password FROM users WHERE id=$1`
var find_by_email_query = `SELECT id, name, age, password FROM users WHERE email=$1`

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest.RestErr) {
	conn, err := ur.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	defer conn.Close()

	row := conn.QueryRow(find_by_id_query, id)

	if row.Err() != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var name, email, password string
	var age int8

	if err = row.Scan(
		&name,
		&email,
		&age,
		&password,
	); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	user_domain := model.NewUserDomain(name, email, password, age)
	user_domain.SetID(id)

	return user_domain, nil
}

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest.RestErr) {
	conn, err := ur.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	defer conn.Close()

	row := conn.QueryRow(find_by_email_query, email)

	if row.Err() != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var id, name, password string
	var age int8

	if err = row.Scan(
		&id,
		&name,
		&age,
		&password,
	); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	user_domain := model.NewUserDomain(name, email, password, age)
	user_domain.SetID(id)

	return user_domain, nil
}

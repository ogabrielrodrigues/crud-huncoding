package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	name     string
	email    string
	age      int8
	password string
}

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetAge() int8
	GetPassword() string

	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name,
		email,
		age,
		password,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

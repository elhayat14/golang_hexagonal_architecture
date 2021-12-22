package user

import (
	portContractUser "golang_hexagonal_architecture/ports/contract/user"
)

type IRepository interface {
	Create(userObj portContractUser.Object) (*portContractUser.Object, error)
	Update(userId string, userObj portContractUser.Object) (*portContractUser.Object, error)
	Delete(userId string) (*bool, error)
	GetById(userId string) (*portContractUser.Object, error)
}

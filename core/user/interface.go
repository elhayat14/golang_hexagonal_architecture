package user

import portContractUser "golang_hexagonal_architecture/ports/contract/user"

type IService interface {
	DoCreateUser(userObj portContractUser.Object) (*portContractUser.Object, error)
	DoEditUser(userId string, userObj portContractUser.Object) (*portContractUser.Object, error)
	DoDeleteUser(userId string) (*bool, error)
	DoGetUserById(userId string) (*portContractUser.Object, error)
}

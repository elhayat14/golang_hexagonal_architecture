package user

import (
	userRepo "golang_hexagonal_architecture/adapters/repositories/mongodb/user"
	portContractUser "golang_hexagonal_architecture/ports/contract/user"
)

type Service struct {
	repository userRepo.IRepository
}

//put all the business logic here
func New(repository userRepo.IRepository) IService {
	return &Service{
		repository: repository,
	}
}

func (service *Service) DoCreateUser(userObj portContractUser.Object) (*portContractUser.Object, error) {
	//add default role
	userObj.Role = "admin"
	//insert to database
	return service.repository.Create(userObj)

}
func (service *Service) DoEditUser(userId string, userObj portContractUser.Object) (*portContractUser.Object, error) {
	//put logic here, will simulate logic to edit role
	userObj.Role = "superadmin"
	//update to database
	return service.repository.Update(userId, userObj)
}

func (service *Service) DoDeleteUser(userId string) (*bool, error) {
	return service.repository.Delete(userId)
}

func (service *Service) DoGetUserById(userId string) (*portContractUser.Object, error) {
	return service.repository.GetById(userId)
}

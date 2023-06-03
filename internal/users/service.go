package users

type UserServiceInterface interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
}

type UserService struct {
	UserRepository UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (us *UserService) GetAll() ([]User, error) {
	return us.UserRepository.GetAll()
}

func (us *UserService) GetByID(id string) (User, error) {
	return us.UserRepository.GetByID(id)
}

package main

type User struct {
	Password string
	Email    string
	Phone    string
}

type UserService interface {
	SignUp(user User)
}

type Notifier interface {
	NotifyUserRegistration(user User)
}

type UserRepository interface {
	Save(user User)
}

type notifier struct{}

func (n *notifier) NotifyUserRegistration(user User) {
	println("notify for user " + user.Email)
}

func NewNotifier() Notifier {
	return &notifier{}
}

type userRepository struct{}

func (r *userRepository) Save(user User) {
	println("saving user with email : " + user.Email)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userService struct {
	Notifier
	UserRepository
}

func (r *userService) SignUp(user User) {
	println("signing up user " + user.Email)
	r.UserRepository.Save(user)
	println("user has saved")
	r.Notifier.NotifyUserRegistration(user)
	println("--- User signed up successfully ---")
}

func NewUserService(n Notifier, r UserRepository) UserService {
	return &userService{
		Notifier:       n,
		UserRepository: r,
	}
}

func main() {
	// we have a notification service
	notifier := NewNotifier()
	// also we need a user repository
	repo := NewUserRepository()
	// now we can create an instant of UserService
	service := NewUserService(notifier, repo)

	// we can hide complex actions we have to run for signing up a new user
	service.SignUp(User{"12345", "one@somthing.com", "1234567890"})
	service.SignUp(User{"54321", "two@somthing.com", "0987654321"})
}

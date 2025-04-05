package user

// Service provides methods to manage products
type Service struct {
	Repository *Repository
}

// NewService creates a new Product service
func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

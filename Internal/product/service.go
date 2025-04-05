package product

// Service provides methods to manage products
type Service struct {
	Repository *Repository
}

// NewService creates a new Product service
func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// GetAllProducts gets all products from the repository
func (s *Service) GetAllProducts() ([]Product, error) {
	return s.Repository.GetAllProducts()
}

// GetProductByID gets a product by its ID
func (s *Service) GetProductByID(id uint) (Product, error) {
	return s.Repository.GetProductByID(id)
}

// CreateProduct creates a new product in the repository
func (s *Service) CreateProduct(p Product) (Product, error) {
	return s.Repository.CreateProduct(p)
}

// UpdateProduct updates an existing product
func (s *Service) UpdateProduct(p Product) (Product, error) {
	return s.Repository.UpdateProduct(p)
}

// DeleteProduct deletes a product by its ID
func (s *Service) DeleteProduct(id uint) error {
	return s.Repository.DeleteProduct(id)
}

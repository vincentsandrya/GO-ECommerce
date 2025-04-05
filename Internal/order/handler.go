package order

// Handler handles HTTP requests for products
type Handler struct {
	Service *Service
}

// NewHandler creates a new product handler
func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

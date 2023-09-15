package application

type Checker interface {
	Check() bool
}

type HealthService struct {
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) Check() bool {
	// @todo: add external services check
	return true
}

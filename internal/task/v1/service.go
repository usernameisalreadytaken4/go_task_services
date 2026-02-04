package task

type service struct {
	repo Repository
}

type Service interface {
	// Create(*Task) error
	// Get(int) (*Task, error)
	// GetByUserID(int) ([]*Task, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

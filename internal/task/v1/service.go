package task

type TaskService struct {
	repo *TaskRepository
}

func NewService(repo *TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

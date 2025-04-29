package todo

import "errors"

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type TaskService struct {
	tasks []Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make([]Task, 0),
	}
}

func (s *TaskService) AddTask(title string) Task {
	newTask := Task{
		ID:        len(s.tasks) + 1,
		Title:     title,
		Completed: false,
	}
	s.tasks = append(s.tasks, newTask)
	return newTask
}

func (s *TaskService) CompleteTask(id int) (Task, error) {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Completed = true
			return s.tasks[i], nil
		}
	}
	return Task{}, ErrTaskNotFound
}

func (s *TaskService) GetTask(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, ErrTaskNotFound
}

func (s *TaskService) ListTasks() []Task {
	return s.tasks
}

var ErrTaskNotFound = errors.New("task not found")

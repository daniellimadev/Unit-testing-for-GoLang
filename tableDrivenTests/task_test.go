package todo

import (
	"errors"
	"reflect"
	"testing"
)

func TestAddTask(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected Task
	}{
		{
			name:     "Add simple task",
			title:    "Buy milk",
			expected: Task{ID: 1, Title: "Buy milk", Completed: false},
		},
		{
			name:     "Add task with empty title",
			title:    "",
			expected: Task{ID: 1, Title: "", Completed: false},
		},
		{
			name:     "Add task with special characters",
			title:    "Make report @#$%",
			expected: Task{ID: 1, Title: "Make report @#$%", Completed: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewTaskService()
			result := service.AddTask(tt.title)

			if result.ID != tt.expected.ID {
				t.Errorf("expected ID %d, got %d", tt.expected.ID, result.ID)
			}

			if result.Title != tt.expected.Title {
				t.Errorf("expected Title %q, got %q", tt.expected.Title, result.Title)
			}

			if result.Completed != tt.expected.Completed {
				t.Errorf("expected Completed %v, got %v", tt.expected.Completed, result.Completed)
			}

			// Checks if the task was actually added to the service
			if len(service.ListTasks()) != 1 {
				t.Errorf("expected 1 task in service, got %d", len(service.ListTasks()))
			}
		})
	}
}

func TestCompleteTask(t *testing.T) {
	tests := []struct {
		name        string
		initialTask Task
		taskID      int
		expected    Task
		expectError bool
	}{
		{
			name:        "Complete existing task",
			initialTask: Task{ID: 1, Title: "Task 1", Completed: false},
			taskID:      1,
			expected:    Task{ID: 1, Title: "Task 1", Completed: true},
			expectError: false,
		},
		{
			name:        "Complete non-existent task",
			initialTask: Task{ID: 1, Title: "Task 1", Completed: false},
			taskID:      2,
			expected:    Task{},
			expectError: true,
		},
		{
			name:        "Complete already completed task",
			initialTask: Task{ID: 1, Title: "Task 1", Completed: true},
			taskID:      1,
			expected:    Task{ID: 1, Title: "Task 1", Completed: true},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewTaskService()
			service.AddTask(tt.initialTask.Title) // Add the initial task

			result, err := service.CompleteTask(tt.taskID)

			if tt.expectError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if !errors.Is(err, ErrTaskNotFound) {
					t.Errorf("expected error %v, got %v", ErrTaskNotFound, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %+v, got %+v", tt.expected, result)
				}

				// Checks if the state has been persisted
				taskFromService, _ := service.GetTask(tt.taskID)
				if !taskFromService.Completed {
					t.Error("task should be completed in service")
				}
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	tests := []struct {
		name         string
		initialTasks []Task
		taskID       int
		expected     Task
		expectError  bool
	}{
		{
			name: "Get existing task",
			initialTasks: []Task{
				{ID: 1, Title: "Task 1", Completed: false},
				{ID: 2, Title: "Task 2", Completed: true},
			},
			taskID:      2,
			expected:    Task{ID: 2, Title: "Task 2", Completed: true},
			expectError: false,
		},
		{
			name: "Get non-existent task",
			initialTasks: []Task{
				{ID: 1, Title: "Task 1", Completed: false},
			},
			taskID:      3,
			expected:    Task{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewTaskService()
			// Add initial tasks
			for _, task := range tt.initialTasks {
				service.AddTask(task.Title)
				if task.Completed {
					service.CompleteTask(task.ID)
				}
			}

			result, err := service.GetTask(tt.taskID)

			if tt.expectError {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("expected %+v, got %+v", tt.expected, result)
				}
			}
		})
	}
}

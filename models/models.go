package models

import (
   "fmt"
)

var DefaultTaskList *TaskManager 

type Task struct {
    ID    int  
    Title string 
}

func NewTask(title string) (*Task, error) {
    if title == "" {
        return nil, fmt.Errorf("empty title")
    }
    return &Task{0, title}, nil
}



type TaskManager struct {
    tasks  []*Task
    lastID int
}


func (m *TaskManager) Save(task *Task) error {
    if task.ID == 0 {
        m.lastID++
        task.ID = m.lastID
        m.tasks = append(m.tasks, cloneTask(task))
        return nil
    }

    for i, t := range m.tasks {
        if t.ID == task.ID {
            m.tasks[i] = cloneTask(task)
            return nil
        }
    }
    return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
    c := *t
    return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*Task {
    return m.tasks
}









func NewTaskManager() *TaskManager {
    return &TaskManager{}
}

func init() {
    DefaultTaskList  = NewTaskManager()
} 

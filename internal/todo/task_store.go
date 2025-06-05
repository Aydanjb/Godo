package todo

import (
	"encoding/json"
	"os"
)

type TaskStore interface {
	Save(*TaskList) error
	Load() (*TaskList, error)
}

type JSONTaskStore struct {
	Filepath string
}

func (j JSONTaskStore) Save(tl *TaskList) error {
	data, err := json.MarshalIndent(tl, "", "   ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.Filepath, data, 0644)
}

func (j JSONTaskStore) Load() (*TaskList, error) {
	data, err := os.ReadFile(j.Filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return &TaskList{
				Tasks:  []*Task{},
				NextID: 0,
			}, nil
		}
		return nil, err
	}

	var tl TaskList
	err = json.Unmarshal(data, &tl)
	if err != nil {
		return nil, err
	}

	return &tl, nil
}

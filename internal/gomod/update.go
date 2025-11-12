package gomod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// DependencyUpdate информация о обновлениях
type DependencyUpdate struct {
	Path    string
	Version string
	Update  *struct {
		Version string // Новая доступная версия
	}
}

// CheckUpdates проверяет обновления зависимостей
func CheckUpdates(repoPath string) ([]*DependencyUpdate, error) {
	// go list -m -u -json all
	cmd := exec.Command("go", "list", "-m", "-u", "-json", "all")
	cmd.Dir = repoPath

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить 'go list -m -u all': %v\n%s", err, stderr.String())
	}

	dec := json.NewDecoder(&out)
	var updates []*DependencyUpdate

	for dec.More() {
		var dep DependencyUpdate
		err := dec.Decode(&dep)
		if err != nil {
			return nil, fmt.Errorf("ошибка при разборе json от 'go list': %w", err)
		}

		if dep.Update != nil {
			updates = append(updates, &dep)
		}
	}

	return updates, nil
}

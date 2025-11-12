package gitrepo

import (
	"fmt"
	"net/url"
	"os"

	git "github.com/go-git/go-git/v6"
)

// Clone клонирует репозиторий во временную папку
func Clone(repoURL string) (string, error) {
	if err := validateURLHelper(repoURL); err != nil {
		return "", err
	}

	tmpDir, err := os.MkdirTemp("", "upgit-*")
	if err != nil {
		return "", fmt.Errorf("не удалось создать временную папку: %w", err)
	}

	_, err = git.PlainClone(tmpDir, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", fmt.Errorf("ошибка при клонировании репозитория: %w", err)
	}

	return tmpDir, nil
}

// validateURLHelper небольшая валидация url
func validateURLHelper(repoURL string) error {
	u, err := url.Parse(repoURL)
	if err != nil {
		return fmt.Errorf("некорректный URL: %v", err)
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("адрес репозитория не прошел валидацию: %s", repoURL)
	}

	return nil
}

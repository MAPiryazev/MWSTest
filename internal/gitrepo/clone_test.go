package gitrepo

import (
	"os"
	"testing"
)

func TestValidateURLHelper(t *testing.T) {
	//валидная ссылка
	err := validateURLHelper("https://github.com/user/repo.git")
	if err != nil {
		t.Errorf("validateURLHelper: неправильная валидация ссылки")
	}

	// Пустая строка
	err = validateURLHelper("")
	if err == nil {
		t.Errorf("validateURLHelper: ожидалась ошибка при пустом URL, но её нет")
	}

	err = validateURLHelper("github.com/user/repo.git")
	if err == nil {
		t.Errorf("validateURLHelper: ожидалась ошибка при URL без схемы, но её нет")
	}

	// Битый URL
	err = validateURLHelper("://jgf;gj")
	if err == nil {
		t.Errorf("validateURLHelper: ожидалась ошибка при битом URL, но её нет")
	}
}

func TestClone_InvalidURL(t *testing.T) {
	_, err := Clone("not-a-valid-url")
	if err == nil {
		t.Errorf("Clone: ожидалась ошибка при передаче невалидного URL")
	}
}

func TestClone_NonExistingRepo(t *testing.T) {
	tmpDir, err := Clone("https://github.com/some/-08-99880--repo.git")
	if err == nil {
		t.Errorf("Clone: ожидалась ошибка при клонировании несуществующего репозитория")
	}

	// Проверяем, что временная папка удалена
	if tmpDir != "" {
		if _, statErr := os.Stat(tmpDir); !os.IsNotExist(statErr) {
			t.Errorf("Clone: временная папка не была удалена при ошибке, осталась: %s", tmpDir)
		}
	}
}

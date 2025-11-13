package gomod

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGoModsHelper(t *testing.T) {
	tmpDir := t.TempDir()

	// создаём несколько тестовых файлов
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module example.com/test"), 0644)
	os.Mkdir(filepath.Join(tmpDir, "subdir"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "subdir", "go.mod"), []byte("module example.com/sub"), 0644)

	files, err := findGoModsHelper(tmpDir)
	if err != nil {
		t.Fatalf("findGoModsHelper вернул ошибку: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("ожидалось 2 go.mod, а найдено %d", len(files))
	}
}

func TestParseGoModHelper(t *testing.T) {
	tmpDir := t.TempDir()
	modPath := filepath.Join(tmpDir, "go.mod")

	content := []byte(`
module example.com/test
go 1.25
require github.com/stretchr/testify v1.8.4
`)

	err := os.WriteFile(modPath, content, 0644)
	if err != nil {
		t.Fatalf("ошибка при создании go.mod: %v", err)
	}

	mod, err := parseGoModHelper(modPath)
	if err != nil {
		t.Fatalf("parseGoModHelper вернул ошибку: %v", err)
	}

	if mod.Name != "example.com/test" {
		t.Errorf("ожидалось имя модуля example.com/test, получено %s", mod.Name)
	}

	if mod.GoVersion != "1.25" {
		t.Errorf("ожидалась версия Go 1.25, получено %s", mod.GoVersion)
	}

	if len(mod.Deps) != 1 || mod.Deps[0].Path != "github.com/stretchr/testify" {
		t.Errorf("неправильно разобраны зависимости: %+v", mod.Deps)
	}
}

func TestParseAllModules(t *testing.T) {
	tmpDir := t.TempDir()

	// создаём два модуля
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module example.com/main\ngo 1.22"), 0644)
	os.Mkdir(filepath.Join(tmpDir, "sub"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "sub", "go.mod"), []byte("module example.com/sub\ngo 1.22"), 0644)

	mods, err := ParseAllModules(tmpDir)
	if err != nil {
		t.Fatalf("ParseAllModules вернул ошибку: %v", err)
	}

	if len(mods) != 2 {
		t.Errorf("ожидалось 2 модуля, а получено %d", len(mods))
	}
}

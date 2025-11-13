package gomod

import (
	"fmt"
	"os"
	"path/filepath"

	"upgit/internal/models"

	"golang.org/x/mod/modfile"
)

// ParseAllModules ищет все go.mod в репозитории и парсит их
func ParseAllModules(repoPath string) ([]*models.ModuleInfo, error) {
	modFiles, err := findGoModsHelper(repoPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось нйти go.mod: %w", err)
	}
	if len(modFiles) == 0 {
		return nil, fmt.Errorf("в репозитории не найдено go.mod файлов")
	}

	var modules []*models.ModuleInfo
	for _, modPath := range modFiles {
		mod, err := parseGoModHelper(modPath)
		if err != nil {
			return nil, fmt.Errorf("ошибка при парсинге %s: %w", modPath, err)
		}
		modules = append(modules, mod)
	}

	return modules, nil
}

// findGoModsHelper ищет все go.mod файлы начиная с root
func findGoModsHelper(root string) ([]string, error) {
	var mods []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "go.mod" {
			mods = append(mods, path)
		}
		return nil
	})
	return mods, err
}

// parseGoModHelper парсит go.mod и возвращает ModuleInfo
func parseGoModHelper(path string) (*models.ModuleInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла %s: %w", path, err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("файл %s пустой", path)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при парсинге %s: %w", path, err)
	}

	if f.Module == nil {
		return nil, fmt.Errorf("в go.mod %s не найден модуль (module)", path)
	}

	mod := &models.ModuleInfo{
		Name:      f.Module.Mod.Path,
		GoVersion: f.Go.Version,
	}

	if f.Require != nil {
		for _, r := range f.Require {
			mod.Deps = append(mod.Deps, &models.Dependency{
				Path:    r.Mod.Path,
				Version: r.Mod.Version,
			})
		}
	}

	return mod, nil
}

package models

// ModuleInfo информация о модуле и его зависимостях
type ModuleInfo struct {
	Name      string
	GoVersion string
	Deps      []*Dependency
}

// Dependency - зависимости из go.mod
type Dependency struct {
	Path    string
	Version string
}

// DependencyUpdate информация о обновлениях
type DependencyUpdate struct {
	Path    string
	Version string
	Update  *struct {
		Version string // Новая версия если есть
	}
}

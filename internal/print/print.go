package output

import (
	"encoding/json"
	"fmt"
	"strings"

	"upgit/internal/gomod"
)

// PrintModulesPlain выводит информацию о модулях и зависимостях в обычном формате
func PrintModulesPlain(modules []*gomod.ModuleInfo) {
	for _, mod := range modules {
		fmt.Printf("Модуль: %s\n", mod.Name)
		fmt.Printf("Go версия: %s\n", mod.GoVersion)
		if len(mod.Deps) == 0 {
			fmt.Println("Зависимости отсутствуют")
		} else {
			fmt.Println("Зависимости:")
			for _, dep := range mod.Deps {
				fmt.Printf("  - %s %s\n", dep.Path, dep.Version)
			}
		}
		fmt.Println(strings.Repeat("-", 40))
	}
}

// PrintUpdatesPlain выводит список зависимостей с доступными обновлениями
func PrintUpdatesPlain(updates []*gomod.DependencyUpdate) {
	if len(updates) == 0 {
		fmt.Println("Все зависимости актуальны")
		return
	}

	fmt.Println("Доступные обновления зависимостей:")
	for _, u := range updates {
		fmt.Printf("  - %s %s → %s\n", u.Path, u.Version, u.Update.Version)
	}
}

// PrintModulesJSON выводит модули и зависимости в JSON
func PrintModulesJSON(modules []*gomod.ModuleInfo, updates []*gomod.DependencyUpdate) {
	data := map[string]interface{}{
		"modules": modules,
		"updates": updates,
	}

	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка при формировании JSON: %v\n", err)
		return
	}

	fmt.Println(string(out))
}

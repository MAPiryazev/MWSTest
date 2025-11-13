package main

import (
	"flag"
	"fmt"
	"os"

	"upgit/internal/gitrepo"
	"upgit/internal/gomod"
	"upgit/internal/models"
	output "upgit/internal/print"
)

func main() {

	flags := parseFlags()
	if flags.RepoURL == "" {
		fmt.Println("Ошибка: не указан URL репозитория")
		flag.Usage()
		os.Exit(1)
	}

	tmpDir, err := gitrepo.Clone(flags.RepoURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при клонировании репозитория: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	modules, err := gomod.ParseAllModules(tmpDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при парсинге модулей: %v\n", err)
		os.Exit(1)
	}

	updates, err := gomod.CheckUpdates(tmpDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при проверке обновлений: %v\n", err)
		os.Exit(1)
	}

	if flags.JSONOutput {
		output.PrintModulesJSON(modules, updates)
	} else {
		output.PrintModulesPlain(modules)
		output.PrintUpdatesPlain(updates)
	}
}

// parseFlags флаги и возвращает структуру Flags
func parseFlags() *models.Flags {
	repoURL := flag.String("repo", "", "URL git репозитория")
	jsonOutput := flag.Bool("json", false, "Вывод в формате JSON")
	flag.Parse()

	return &models.Flags{
		RepoURL:    *repoURL,
		JSONOutput: *jsonOutput,
	}
}

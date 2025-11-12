package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"upgit/internal/gitrepo"
	"upgit/internal/gomod"
	output "upgit/internal/print"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	repoURL := flag.String("repo", "", "URL git репозитория")
	jsonOutput := flag.Bool("json", false, "Вывод в формате JSON")
	flag.Parse()

	if *repoURL == "" {
		fmt.Println("Ошибка: не указан URL репозитория")
		flag.Usage()
		os.Exit(1)
	}

	tmpDir, err := gitrepo.Clone(*repoURL)
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

	if *jsonOutput {
		output.PrintModulesJSON(modules, updates)
	} else {
		output.PrintModulesPlain(modules)
		output.PrintUpdatesPlain(updates)
	}
}

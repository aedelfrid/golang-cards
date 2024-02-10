package main

import (
	"github.com/manifoldco/promptui"
	"os"
)

func startPrompt() {
	prompt := promptui.Select{
		Label: "Would you like to play?",
		Items: []string{"Yes", "No"},
	}

	_, result, _ := prompt.Run()

	if result == "No" {
		os.Exit(0)
	}

	return
}

func gameSelect() playset{
	prompt := promptui.Select{
		Label: "What would you like to play?",
		Items: []string{"Texas Hold'em"},
	}

	_, result, _ := prompt.Run()

	return gameImports(result)
}
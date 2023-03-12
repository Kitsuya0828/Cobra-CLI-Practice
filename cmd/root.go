/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type commitType struct {
	Emoji string
	Name     string
	Description  string
}

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gommit",
	Short: "Provides commit messages for the git commit command in an easy-to-understand, uniform format",
	Run: func(cmd *cobra.Command, args []string) {
		// Opens an already existing repository.
		r, err := git.PlainOpen(".")
		CheckIfError(err)

		w, err := r.Worktree()
		CheckIfError(err)

		commitTypes := []commitType{
			{Emoji: "✨", Name: "feat", Description: "A new feature"},
			{Emoji: "🐛", Name: "fix", Description: "A bug fix"},
			{Emoji: "🎨", Name: "style", Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
			{Emoji: "♻", Name: "refactor", Description: "A code change that neither fixes a bug nor adds a feature"},
			{Emoji: "🐎", Name: "perf", Description: "A code change that improves performance"},
			{Emoji: "📚", Name: "docs", Description: "Documentation only changes"},
			{Emoji: "🚨", Name: "test", Description: "Adding missing tests or correcting existing tests"},
			{Emoji: "⏪", Name: "revert", Description: "Reverts a previous commit"},
		}
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "{{ .Emoji }} {{ .Name | cyan }}",
			Inactive: "  {{ .Name | cyan }}",
			Selected: "Type: {{ .Name | red | cyan }}",
			Details: `
--------- Type ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}`,
		}
		searcher := func(input string, index int) bool {
			commitType := commitTypes[index]
			name := strings.Replace(strings.ToLower(commitType.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
	
			return strings.Contains(name, input)
		}
		promptType := promptui.Select{
			Label: "Select your commit type",
			Items: commitTypes,
			Templates: templates,
			Searcher: searcher,
		}
		i, _, err := promptType.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		promptDescription := promptui.Prompt{
			Label:    "Enter a description",
		}
		resultDescription, err := promptDescription.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		commitMessage := fmt.Sprintf("%s: %s", commitTypes[i].Name, resultDescription)
		Info(fmt.Sprintf("git commit -m \"%s\"", commitMessage))
		commit, err := w.Commit(commitMessage, &git.CommitOptions{})
		CheckIfError(err)
		obj, err := r.CommitObject(commit)
		CheckIfError(err)

		fmt.Println(obj)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

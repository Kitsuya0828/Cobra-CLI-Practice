/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "io/ioutil"
	"os"
	// "path/filepath"
	"github.com/manifoldco/promptui"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

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

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commit called")
		// CheckArgs("<directory>")
		// directory := os.Args[1]
		directory := "."

		// Opens an already existing repository.
		r, err := git.PlainOpen(directory)
		CheckIfError(err)

		w, err := r.Worktree()
		CheckIfError(err)

		// ... we need a file to commit so let's create a new file inside of the
		// worktree of the project using the go standard library.
		// Info("echo \"hello world!\" > example-git-file")
		// filename := filepath.Join(directory, "example-git-file")
		// err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
		// CheckIfError(err)

		// Adds the new files to the staging area.
		// Info("git add .")
		// _, err = w.Add("example-git-file")
		// CheckIfError(err)

		prompt := promptui.Select{
			Label: "Select a type",
			Items: []string{
				"build",
				"chore",
				"ci",
				"docs",
				"feat",
				"fix",
				"perf",
				"refactor",
				"revert",
				"style",
				"test",
			},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		// Commits the current staging area to the repository, with the new file
		// just created. We should provide the object.Signature of Author of the
		// commit Since version 5.0.1, we can omit the Author signature, being read
		// from the git config files.
		// Info("git commit -m \"add commit from code\"")
		commit, err := w.Commit(result + ": commit from code", &git.CommitOptions{
			Author: &object.Signature{
				Name:  "Kitsuya0828",
				Email: "kitsuyaazuma@gmail.com",
				When:  time.Now(),
			},
		})

		CheckIfError(err)
		// Prints the current HEAD to verify that all worked well.
		Info("git show -s")
		obj, err := r.CommitObject(commit)
		CheckIfError(err)

		fmt.Println(obj)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

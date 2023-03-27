# gommit

A CLI tool that provides commit messages for the git commit command in an easy-to-understand, uniform format.

https://user-images.githubusercontent.com/60843722/227822377-6fe091c6-d9de-4c98-8994-3557078934e1.mov


```zsh
# Install
go install github.com/Kitsuya0828/gommit

# Usage
git add .
gommit
git push

# Uninstall
go clean -i -n github.com/Kitsuya0828/gommit
```

* [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
* [manifoldco/promptui: Interactive prompt for command\-line applications](https://github.com/manifoldco/promptui)
* [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
* [gitmoji \| An emoji guide for your commit messages](https://gitmoji.dev/)

# Git-Secrets

Vega has been pre-configured with [git-secrets](https://github.com/awslabs/git-secrets) which prevents
developers from committing password and other sensitive information to a git repository.

Currently vega has following rules pre-configured:-

* AWS Secrets scanners: Looks for AWS Secrets in your git repos.
* Drupal database settings scanner: Looks for Drupal database settings written in plain text in your repo based on a simple regex.

Following git-hooks are also pre-configured:-

  1. ```pre-commit```: Used to check if any of the files changed in the commit
       use prohibited patterns.
  2. ```commit-msg```: Used to determine if a commit message contains a
       prohibited patterns.
  3. ```prepare-commit-msg```: Used to determine if a merge commit will
       introduce a history that contains a prohibited pattern at any point.
       Please note that this hook is only invoked for non fast-forward merges.

# Quick Start

```
curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install.sh | bash
```
Please follow the output of the command carefully.
```
vega init
```
Goto your project repositories and execute following
```
vega hooks install
```
Note: ```vega hooks install``` overrides any current git hooks if you have added any. In case you would like to have multiple
hooks please refer: https://gist.github.com/carlos-jenkins/89da9dcf9e0d528ac978311938aade43

## Migration from older releases
In case you are running older release of vega(<1.0.6), please perform following steps

```
git config -l
```
In case the output has
```
core.hookspath=/Users/viz/.git/hooks
```
```
vim ~/.gitconfig
```
and delete the ```core.hookspath=/Users/viz/.git/hooks``` line.

## FAQs

**Question**: I use Docksal, Lando or Drupal VM in my project, can I continue using it?

**Answer**: You can use your current development stack for your project in case you want to use the secret management capability of vega.
Unless you have git commit hooks set up in your project there will not be any impact on your local development stack.

**Question**: I get ```git-secrets``` command not found when I try to commit code.

**Answer**: Have you exported the PATH in ```~/.bashrc``` or equivalent file(```.zshrc, .bash_profile etc```). 
Check if all the binaries are there in ```/home/<user-name>/.local/bin```. In case you have just installed vega and opened a new terminal,
the session is not retained so type following in your terminal ```source ~/.bashrc``` or equivalent file(```.zshrc, .bash_profile etc```).
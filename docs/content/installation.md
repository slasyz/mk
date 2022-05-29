# Installation

## Leap of faith

```shell
curl -sf -L https://raw.githubusercontent.com/slasyz/mk/master/install.sh 2>/dev/null | sudo sh
```

## Building from sources

```shell
$ go build -o mk .
$ chmod a+x mk
$ sudo cp mk /usr/local/bin/mk
```

## Shell autocompletion

```shell
# Add this to your .zshrc if you use zsh:
autoload -Uz +X compinit bashcompinit && compinit && bashcompinit
complete -o nosort -C mk mk

# Add this to your .bash_profile if you use bash:
complete -o nosort -C mk mk
```

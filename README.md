# mk

**mk** is a more user-friendly `make` alternative.


## Description

You can create `mk.yml` file just like you create `Makefile`, and define all targets using modern trendy YAML syntax. After that, you can run `mk` to execute these targets.

It's just proof of concept yet, so `mk.yml` schema in unstable, and everything is going to be rewritten multiple times. If you have enough courage to use it, I'd be glad to hear any feedback.

It has nothing to do with Plan 9 `mk`, I just took the same name because it seems nice.


## Installation

### Leap of faith
```shell
curl -sf -L https://raw.githubusercontent.com/slasyz/mk/master/install.sh 2>/dev/null | sudo sh
```

### Building from sources

```shell
$ go build -o mk .
$ chmod a+x mk
$ sudo cp mk /usr/local/bin/mk
```


## Usage

```shell
$ mk help
```

## Shell autocompletion

```shell
# Add this to your .zshrc if you use zsh:
autoload -Uz +X compinit bashcompinit && compinit && bashcompinit
complete -o nosort -C mk mk

# Add this to your .bash_profile if you use bash:
complete -o nosort -C mk mk
```

## To Do

* flags in addition to positional arguments
* using other commands to set variables
* fancy colorful logs with emoji for my young friends
* more examples (correct and incorrect)
* a lot of testing and bugfixes
* website with documentation (use mk.syrovats.ky address to promote personal brand)
* ~~make it :rocket:blazing fast:rocket: by rewriting in Rust~~ no


## Contributing

Pull requests are welcome.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

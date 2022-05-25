# mk

**mk** is a more user-friendly `make` alternative.


## Description

You can create `mk.yml` file just like you create `Makefile`, and define all targets using modern trendy YAML syntax. After that, you can run `mk` to execute these targets.

It's just proof of concept yet, so `mk.yml` schema in unstable, and everything is going to be rewritten multiple times. If you have enough courage to use it, I'd be glad to hear any feedback.

It has nothing to do with Plan 9 `mk`, I just took the same name because it seems nice.


## Installation

```shell
$ go build -o mk .
$ chmod a+x mk
$ sudo cp mk /usr/local/bin/mk
```


## Usage

```shell
$ mk help
```


## To Do

* flags in addition to positional arguments
* shell autocompletion
* using other commands to set variables
* fancy colorful logs with emoji for my young friends
* more examples (correct and incorrect)
* a lot of testing and bugfixes
* website with documentation (use mk.syrovats.ky address to promote personal brand)
* single line instruction to install the tool (`curl ... | bash`) to anyone who doesn't really care about security


## Contributing

Pull requests are welcome.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

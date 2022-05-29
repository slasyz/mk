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
# BadgeIT

A tool for quickly downloading [shields.io](https://shields.io) badges in PNG format.
I use these in design work where the behaviour of SVG and text and be odd sometimes.

## Usage

```zsh
$ go run main.go <URL> <OUTPUT FILE NAME>
```

example to download a docker badge as a png

```zsh
$ go run main.go https://img.shields.io/badge/docker-%230db7ed.svg\?style\=for-the-badge\&logo\=docker\&logoColor\=white docker
```

### Installing

run

```zsh
$ mkdir -p dist && go build -o dist/badgeit
```

or with [Just](https://just.systems/)

```zsh
$ just build
```

then copy the outputted binary `dist/badgeit` to your path

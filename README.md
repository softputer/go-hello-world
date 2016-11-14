# go-hello-world

This is a simple Hello-world repo written in Golang. By being  dockerized based on alpine, we can get an image only about 13M.

## Usage

You should build the repo first. By default, if using `net`  package, a build will likely produce a binary with some dynamic linking, e.g, to libc. So when you use this binary in alpine, it will be wrong. To solve this problem, there are two ways.

- Disable CGO, via CGO_ENABLED=0
- go build -tags netgo -a -v



## Build

`docker build .`

## Run

`docker run -d -p <port>:80  <image_name> `


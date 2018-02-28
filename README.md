
# Life.Church IT Downloads and Documentation


## Getting Started

`go get -u github.com/jessebarton/it-download`

IT-Download has been tested on Go v1.9

## Usage
```
$ cd it-download
$ docker build -t jessebarton/it-download:VERSION .
$ docker run --rm -it -p 8080:8080 jessebarton/it-download:VERSION
```
## Upgrade
```
$ docker service update --image jessebarton/it-download:VERSION
```
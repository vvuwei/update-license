[![Build](https://github.com/vvuwei/update-license/actions/workflows/build.yml/badge.svg)](https://github.com/vvuwei/update-license/actions/workflows/build.yml)
[![made_with golang](https://img.shields.io/badge/made_with-golang-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Coverage](https://github.com/vvuwei/update-license/wiki/coverage.svg)](https://raw.githack.com/wiki/vvuwei/update-license/coverage.html)

# update-license

This is a small tool that updates the license header in Golang files.

## Installation

```
go install github.com/vvuwei/update-license@v0.0.1
```

## Usage

Run command in the root of your application directory

```
update-license -path=./ -license=./LICENSE [-dry]
```

## Further Work

* Support licenses by name (MIT, Apache 2.0, etc.), url (http GET)

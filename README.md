[![Go Report Card](https://goreportcard.com/badge/github.com/postfinance/single)](https://goreportcard.com/report/github.com/postfinance/single)
[![Coverage Status](https://coveralls.io/repos/github/postfinance/single/badge.svg?branch=master)](https://coveralls.io/github/postfinance/single?branch=master)
[![Build Status](https://github.com/postfinance/single/workflows/build/badge.svg)](https://github.com/postfinance/single/actions)
[![GoDoc](https://godoc.org/github.com/postfinance/single?status.svg)](https://godoc.org/github.com/postfinance/single)

# Single

Single ensures that only one instance of your program is running

# Usage

Use a persistent `*Single` object for the duration of your program, or Go's garbage collector will automatically close the lockfile. 

//go:build tools

package main

// This file imports dev tools to lock its version.

import _ "github.com/Yamashou/gqlgenc"

// Run following command to generate code:
//   go generate tools.go

//go:generate go run github.com/Yamashou/gqlgenc

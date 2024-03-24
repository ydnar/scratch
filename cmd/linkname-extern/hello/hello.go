package hello

import _ "unsafe"

// World is implemented in package main
func World() string

type String string

// String is implemented in package main
func (s String) String() string

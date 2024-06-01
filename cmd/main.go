package main

import (
	"flag"

	"github.com/xnacly/verdunkel"
)

func main() {
	c := &verdunkel.Config{}

	// var Default = &Config{
	// 	Function:   true,
	// 	Variables:  true,
	// 	Structs:    true,
	// 	Interfaces: true,
	// 	OutputDir:  "./out",
	// }

	// TODO: add all flags
	flag.BoolVar(&c.Function, "funcs", true, "function name obfuscation")
	flag.BoolVar(&c.Variables, "vars", true, "variable name obfuscation")
	flag.BoolVar(&c.Structs, "structs", true, "struct name and field obfuscation")
	flag.Parse()
}

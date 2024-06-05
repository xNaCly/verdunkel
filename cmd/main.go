package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"regexp"

	"github.com/xnacly/verdunkel"
	"github.com/xnacly/verdunkel/log"
)

func main() {
	c := &verdunkel.Config{}

	slog.SetDefault(slog.New(log.New(os.Stdout, slog.LevelDebug)))

	flag.BoolVar(&c.Function, "funcs", true, "function name obfuscation")
	flag.BoolVar(&c.Variables, "vars", true, "variable name obfuscation")
	flag.BoolVar(&c.Structs, "structs", true, "struct name and field obfuscation")
	flag.BoolVar(&c.Interfaces, "interfaces", true, "interface name and method name obfuscation")
	flag.BoolVar(&c.ConstsValue, "consts", false, "const value obfuscation")
	flag.BoolVar(&c.Logs, "logs", false, "stripping of all fmt.Print, fmt.Println, print, println, slog and log calls")
	flag.BoolVar(&c.Packages, "packages", false, "package name and path obfuscation")
	flag.BoolVar(&c.FileNames, "files", false, "file name obfuscation")
	flag.StringVar(&c.OutputDir, "out", "./out", "output directory")
	exclude := flag.String("exclude", "", "matched agains file path, if matched, skips the file")
	flag.Parse()
	slog.Debug("startup...")
	if *exclude != "" {
		e, err := regexp.Compile(*exclude)
		if err != nil {
			slog.Error("failed to compile exclude regexp", "exp", *exclude, "err", err)
			return
		}
		c.Exclude = e
	}

	slog.Info("configuration parsing done")
	fmt.Println(c.String())
}

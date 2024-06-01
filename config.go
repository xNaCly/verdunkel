package verdunkel

import "regexp"

// Config configures the actions verdunkel performs on the input
type Config struct {
	Function    bool           // function name obfuscation, default: true
	Variables   bool           // variable name obfuscation, default: true
	Structs     bool           // struct name and field obfuscation, default: true
	Interfaces  bool           // interface name and method name obfuscation, default: true
	ConstsValue bool           // const obfuscation (string with byte array, numbers with binary, in the future encryption), default: false
	Logs        bool           // stripping of all fmt.Print, fmt.Println, print, println, slog and log calls, default: false
	Packages    bool           // package name and path obfuscation, default: false
	Exclude     *regexp.Regexp // matched agains file path, if matched, skips the file, default: nil
	OutputDir   string         // path to write the processed go code to, default: "./out"
}

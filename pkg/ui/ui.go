package ui

import (
	"fmt"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
)

func DisplayLogo(version string) {
	logo := `
` + Cyan + Bold + `
   █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ███████║███████╗█████╔╝ ███████║██████╔╝
  ██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
  ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝
` + Reset + `
` + Magenta + `              A Laravel-inspired Go Framework` + Reset + `
` + Yellow + `              Version ` + version + ` ✨` + Reset + `
`
	fmt.Print(logo)
}

func PrintStep(message string) {
	fmt.Printf("%s▸%s %s\n", Green+Bold, Reset, message)
}

func PrintError(message string) {
	fmt.Printf("\n%s✗ Error:%s %s\n\n", Red+Bold, Reset, message)
}

func PrintSuccess(message string) {
	fmt.Printf("\n%s%s✓ %s%s\n\n", Green, Bold, message, Reset)
}

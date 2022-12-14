package startupPrinter

import "github.com/TwiN/go-color"

func printAsciiArt() {
	println("  _____             __   _")
	println(" |  __ \\           / _| | |")
	println(" | |__) |  _   _  | |_  | |_")
	println(" |  _  /  | | | | |  _| | __|")
	println(" | | \\ \\  | |_| | | |   | |_")
	println(" |_|  \\_\\  \\__, | |_|    \\__|")
	println("            __/ |")
	println("           |___/")
	println(color.GreenBackground + color.Black + "   - Ryft Framework For Go -   " + color.Reset)
	println("")
}

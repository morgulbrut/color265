// Source: https://github.com/esimov/diagram/blob/master/color/color.go

package color256

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mattn/go-isatty"
)

// Color ansi colors (own type for type checking)
type Color int

// Format  text format (own type for type checking)
type Format int

// Colors
const (
	ColBlack Color = iota
	ColRed
	ColGreen
	ColYellow
	ColBlue
	ColMagenta
	ColCyan
	ColWhite
	ColHiBlack
	ColHiRed
	ColHiGreen
	ColHiYellow
	ColHiBlue
	ColHiMagenta
	ColHiCyan
	ColHiWhite

	ColOrange Color = 208
)

// Formats
const (
	Reset Format = iota
	Bold
	Faint
	Italic
	Underlined
	BlinkSlow
	BlinkRapid
	Reversed
	Concealed
	CrossedOut
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

// String returns a color escape string.
func String(c Color, str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", c, str)
}

// StringBoth fg and bg colors.
func StringBoth(fg, bg Color, str string) string {
	return fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%dm%s\x1b[0m", bg, fg, str)
}

// StringFormat returns a color escape string with extra options.
func StringFormat(c Color, str string, f Format) string {
	return fmt.Sprintf("\x1b[38;5;%d;%dm%s\x1b[0m", c, f, str)
}

// StringFormatBoth fg and bg colors.
func StringFormatBoth(fg, bg Color, str string, f Format) string {
	return fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%d;%dm%s\x1b[0m", bg, fg, f, str)
}

//Random  is a convenient helper function to print with random foreground.
// A newline is appended to format by default.
func Random(format string, a ...interface{}) { fmt.Println(StringRandom(format, a...)) }

// StringRandom  is a convenient helper function to print with black foreground.
// A newline is appended to format by default.
func StringRandom(format string, a ...interface{}) string {
	return String(Color(RandomColor(180, 231)), fmt.Sprintf(format, a...))
}

// RandomColor returns a random  color number.
func RandomColor(min, max int) int {
	return rand.Intn(max-min) + min
}

// Init random number generator (Quick and dirty Windows hack)
func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// List all the colors with its color code
func List() {
	for i := 0; i < 256; i++ {
		fmt.Println(String(Color(i), fmt.Sprintf("%d", i)))
	}
}

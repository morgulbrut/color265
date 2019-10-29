// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import "fmt"

//---------------------------------------------------------------------
// Bold
//---------------------------------------------------------------------

// PrintBold is a convenient helper function to print a bold text.
// A newline is appended to format by default.
func PrintBold(format string, a ...interface{}) { fmt.Println(Bold(format,a...)) }

//Bold is a convenient to get a bold string.
func Bold(format string, a ...interface{}) string { return Format(FmtBold, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Faint
//---------------------------------------------------------------------

// PrintFaint is a convenient helper function to print a faint text.
// A newline is appended to format by default.
func PrintFaint(format string, a ...interface{}) { fmt.Println(Faint(format,a...)) }

//Faint is a convenient to get a faint string.
func Faint(format string, a ...interface{}) string { return Format(FmtFaint, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Italic
//---------------------------------------------------------------------

// PrintItalic is a convenient helper function to print a italic text.
// A newline is appended to format by default.
func PrintItalic(format string, a ...interface{}) { fmt.Println(Italic(format,a...)) }

//Italic is a convenient to get a italic string.
func Italic(format string, a ...interface{}) string { return Format(FmtItalic, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Underlined
//---------------------------------------------------------------------

// PrintUnderlined is a convenient helper function to print a underlined text.
// A newline is appended to format by default.
func PrintUnderlined(format string, a ...interface{}) { fmt.Println(Underlined(format,a...)) }

//Underlined is a convenient to get a underlined string.
func Underlined(format string, a ...interface{}) string { return Format(FmtUnderlined, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// BlinkSlow
//---------------------------------------------------------------------

// PrintBlinkSlow is a convenient helper function to print a blinkslow text.
// A newline is appended to format by default.
func PrintBlinkSlow(format string, a ...interface{}) { fmt.Println(BlinkSlow(format,a...)) }

//BlinkSlow is a convenient to get a blinkslow string.
func BlinkSlow(format string, a ...interface{}) string { return Format(FmtBlinkSlow, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// BlinkRapid
//---------------------------------------------------------------------

// PrintBlinkRapid is a convenient helper function to print a blinkrapid text.
// A newline is appended to format by default.
func PrintBlinkRapid(format string, a ...interface{}) { fmt.Println(BlinkRapid(format,a...)) }

//BlinkRapid is a convenient to get a blinkrapid string.
func BlinkRapid(format string, a ...interface{}) string { return Format(FmtBlinkRapid, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Reversed
//---------------------------------------------------------------------

// PrintReversed is a convenient helper function to print a reversed text.
// A newline is appended to format by default.
func PrintReversed(format string, a ...interface{}) { fmt.Println(Reversed(format,a...)) }

//Reversed is a convenient to get a reversed string.
func Reversed(format string, a ...interface{}) string { return Format(FmtReversed, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Concealed
//---------------------------------------------------------------------

// PrintConcealed is a convenient helper function to print a concealed text.
// A newline is appended to format by default.
func PrintConcealed(format string, a ...interface{}) { fmt.Println(Concealed(format,a...)) }

//Concealed is a convenient to get a concealed string.
func Concealed(format string, a ...interface{}) string { return Format(FmtConcealed, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// CrossedOut
//---------------------------------------------------------------------

// PrintCrossedOut is a convenient helper function to print a crossedout text.
// A newline is appended to format by default.
func PrintCrossedOut(format string, a ...interface{}) { fmt.Println(CrossedOut(format,a...)) }

//CrossedOut is a convenient to get a crossedout string.
func CrossedOut(format string, a ...interface{}) string { return Format(FmtCrossedOut, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Fraktur
//---------------------------------------------------------------------

// PrintFraktur is a convenient helper function to print a fraktur text.
// A newline is appended to format by default.
func PrintFraktur(format string, a ...interface{}) { fmt.Println(Fraktur(format,a...)) }

//Fraktur is a convenient to get a fraktur string.
func Fraktur(format string, a ...interface{}) string { return Format(FmtFraktur, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// DoublyUnderlined
//---------------------------------------------------------------------

// PrintDoublyUnderlined is a convenient helper function to print a doublyunderlined text.
// A newline is appended to format by default.
func PrintDoublyUnderlined(format string, a ...interface{}) { fmt.Println(DoublyUnderlined(format,a...)) }

//DoublyUnderlined is a convenient to get a doublyunderlined string.
func DoublyUnderlined(format string, a ...interface{}) string { return Format(FmtDoublyUnderlined, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Framed
//---------------------------------------------------------------------

// PrintFramed is a convenient helper function to print a framed text.
// A newline is appended to format by default.
func PrintFramed(format string, a ...interface{}) { fmt.Println(Framed(format,a...)) }

//Framed is a convenient to get a framed string.
func Framed(format string, a ...interface{}) string { return Format(FmtFramed, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Encircled
//---------------------------------------------------------------------

// PrintEncircled is a convenient helper function to print a encircled text.
// A newline is appended to format by default.
func PrintEncircled(format string, a ...interface{}) { fmt.Println(Encircled(format,a...)) }

//Encircled is a convenient to get a encircled string.
func Encircled(format string, a ...interface{}) string { return Format(FmtEncircled, fmt.Sprintf(format,a...))}

//---------------------------------------------------------------------
// Overlined
//---------------------------------------------------------------------

// PrintOverlined is a convenient helper function to print a overlined text.
// A newline is appended to format by default.
func PrintOverlined(format string, a ...interface{}) { fmt.Println(Overlined(format,a...)) }

//Overlined is a convenient to get a overlined string.
func Overlined(format string, a ...interface{}) string { return Format(FmtOverlined, fmt.Sprintf(format,a...))}


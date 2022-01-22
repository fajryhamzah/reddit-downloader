package log

import (
	"fmt"
)

const RED string = "\033[31m"
const RESET string = "\033[0m"
const GREEN string = "\033[32m"
const BLUE string = "\033[34m"

func Error(message ...interface{}) {
	fmt.Print(string(RED), "[X] ")
	fmt.Println(message...)

	reset()
}

func Errorf(format string, message ...interface{}) {
	fmt.Print(string(RED), "[X] ")
	printfln(format, message...)

	reset()
}

func Success(message ...interface{}) {
	fmt.Print(string(GREEN), "[+] ")
	fmt.Println(message...)

	reset()
}

func Successf(format string, message ...interface{}) {
	fmt.Print(string(GREEN), "[+] ")
	printfln(format, message...)

	reset()
}

func Logf(format string, message ...interface{}) {
	fmt.Print(string(BLUE), "[-] ")
	printfln(format, message...)

	reset()
}

func printfln(format string, message ...interface{}) {
	fmt.Printf(format+"\n", message...)
}

func reset() {
	fmt.Print(string(RESET))
}

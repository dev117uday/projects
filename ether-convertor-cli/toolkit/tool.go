package toolkit

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

// ParseFlags :  parse and return input flags
func ParseFlags() (string, string, string, string) {

	var to, from, value, help string

	fmt.Println("Welcome to ether convertor")

	commandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	commandLine.StringVar(&from, "from", "wei", "Unit of value to convert from")
	commandLine.StringVar(&to, "to", "wei", "Unit of value to convert to")
	commandLine.StringVar(&value, "value", "1", "Enter the value in this field")
	commandLine.StringVar(&help, "help", "hell", "Print all the options present")

	err := commandLine.Parse(os.Args[1:])
	if err != nil {
		panic("Error in parsing command line arguments")
	}

	return to, from, value, help
}

// PrintHelp : print all the help guide to the user
func PrintHelp() {
	fmt.Println("Printing help")
	fmt.Println("Welcome to ether convertor CLI\n\n### Flags :\n- `help` : To print help related to this package\n- `from` : unit to convert *from*\n- `to` :  unit to convert *to*\n- `value` : amount of ether \n\n### Usage :\n```\n./main -from milliether -to ether -value 12\n./main.go -from microether -to ether -value 12\n```")
}

// CheckInputs : check all input for validity
func CheckInputs(to, from, value *string) bool {
	if !valueCheck(*to) {
		fmt.Println("Incorrect `to` value")
		return false
	}
	if !valueCheck(*from) {
		fmt.Println("Incorrect `from` value")
		return false
	}
	_, err := strconv.ParseFloat(*value, 64)
	if err != nil {
		fmt.Println("Problem parsing `value`")
		return false
	}
	return true
}

func valueCheck(val string) bool {
	if val == "wei" || val == "kwei" || val == "mwei" || val == "gwei" || val == "microether" || val == "milliether" || val == "ether" {
		return true
	} else {
		return false
	}
}

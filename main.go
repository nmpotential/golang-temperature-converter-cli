package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	/* Validate arguments
	• Check if the length of args is different than 2.
	• If so, invoke the printError() function, passing
	errInvalidArguments as the argument.
	 */
	if len(os.Args) != 1 {
		printError(errInvalidArguments)
	}
	/* Read origin unit
	• Invoke the strings.ToUpper() function passing os.Args[1] as the argument.
	This ensures consistency when reading command line arguments provided by the user.
	• Assign the result to the previously defined originUnit variable.
	 */
	originUnit := strings.ToUpper(os.Args[1])

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		/* Read current temperature
		• Invoke the fmt.Scanln() function, passing &originValue as the argument.
		• Assign the two return values to the variables _ and err respectively.
		 */
		_, err := fmt.Scanln(&originValue)
		/*
		• Create an if statement checking if err != nil, and if that's true,
		invoke the printError() function, passing errReadingInput as its argument.
		 */
		if err != nil {
			printError(errReadingInput)
		}
		/* Convert the temperature
		• Create an if statement to check if originUnit is equal to "C"
		(notice the capital casing).
		• If so, invoke the convertToFahrenheit()
		function passing originValue as the argument.
		• Otherwise, create an else block and invoke the convertToCelsius() function,
		passing originValue as the argument.
		 */
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}
		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		/* Prompt to convert again
		• Invoke the fmt.Scanln() function passing &shouldConvertAgain as its argument.
		• Assign the two return values to the previously defined variables _ and err respectively.
		• On the following line, create an if statement checking if err != nil.
		• If that condition is true, invoke printError() passing errReadingInput as its argument.
		 */
		_, err = fmt.Scanln(&shouldConvertAgain)
		/* Parse prompt answer
		• Currently, we can't always be sure that the value assigned to shouldConvertAgain is
		going to be in the casing we expect.
		• Let's fix this by making a small change to the condition on the last if statement
		inside the for loop.
		• The condition should be a combination of calling the strings.ToUpper() function,
		passing it the result of calling the strings.TrimSpace() function with shouldConvertAgain
		as its argument.
		•If the result of all of that is NOT equal to "Y", then the existing if block containing
		fmt.Println("Good bye!") followed by break should be run.
		 */
		if err != nil {
			printError(errReadingInput)
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}

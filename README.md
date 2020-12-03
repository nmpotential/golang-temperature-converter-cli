# Project Objectives

In this project, we built a command line application which has the functionality to parse 
CLI arguments and reads keyboard inputs from the terminal in Go.

## How to: Build a CLI app that parses CLI arguments and reads keyboard inputs from the terminal in Go

### 1. Declare variables needed for expected input and expected/unexpected output
- [ ] Outline how the app supposed to work and put in README.md(this will be updated as 
      functionality and implementation are refined):

- [ ] What are expected inputs:

- [ ] What are expected outputs:

- [ ] Finally, declare variables needed for expected/unexpected input and expected/unexpected 
output (Hint: Think about what we need in between input -> output):

### Example: Temperature Converter CLI

- [x] Outline how the apps supposed to work:
        - *User enters value and unit of temperature and then app prints out converted temperature 
between Fahrenheit and Celsius.*

- [x] What are the expected inputs:
        - Value and unit of temperature
        
- [x] What are the expected outputs of the app:
        - Value of temperature in converted unit 



- [x] Finally, declare variables needed for expected/unexpected input and expected/unexpected 
output (Hint: Think about what we need in between input -> output):

```
package main

var originUnit string	// Expected Input
var originValue float64	// Input

var shouldConvertAgain string //Output

var err error //Error output

var errInvalidArguments = errors.New("Invalid arguments")	// Error output for invalid arguments
var errReadingInput = errors.New("Error reading input")	// Error output for issue with reading input
```

### 2. Create methods for parsing CLI arguments and reading keyboard inputs from the terminal 
- Take methods and find out what its required to create these methods
- What are the steps?
- Work on methods func by func

###   2.1. Create method for parsing CLI arguments 
- In the main func (or otherwise) validate arguments by check if the length of args 
is different than 2. If so, invoke the printError() function, passing errInvalidArguments
as the argument.
        
#   2.2. Create method for reading keyboard input
- In the main func (or otherwise) validate arguments by check if the length of args 
is different than 2. If so, invoke the printError() function, passing errInvalidArguments
as the argument.


# 2. Write tests for each method
- //TODO breakdown how to write tests for each method i.e. test for checking args method, 
test for reading input

## Usage

Compile with `go build -o temp`.

Then, invoke the binary passing as argument the unit of temperature we want to convert **from**.
For example:

`./temp C` to convert from Celsius to Fahrenheit or `./temp F` to convert from Fahrenheit to Celsius.

To verify your work locally for this module, run the following command in a terminal window: 

`go test -v -run M2`

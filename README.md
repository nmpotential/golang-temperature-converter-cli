
### 1. Declare variables needed for expected input and expected/unexpected output



### 2. Create methods for parsing CLI arguments and reading keyboard inputs from the terminal 
- Take methods and find out what its required to create these methods
- What are the steps?
- Work on methods func by func

###   2.1. Create method for parsing CLI arguments 
- In the main func (or otherwise) validate arguments by check if the length of args 
is different from10.  2. If so, invoke the printError() function, passing errInvalidArguments
as the argument.
        
#   2.2. Create method for reading keyboard input
- In the main func (or otherwise) validate arguments by check if the length of args 
is different from 2. If so, invoke the printError() function, passing errInvalidArguments
as the argument.


# 2. Write tests for each method
- //TODO breakdown how to write tests for each method i.e. test for checking args method, 
test for reading input

## Usage

Compile with `go build -o temp`.

Then, invoke the binary passing as argument the unit of temperatdure we want to convert **from**.
For example:

`./temp C` to convert from Celsius to Fahrenheit or `./temp F` to convert from Fahrenheit to Celsius.

To verify your work locally for this module, run the following command in a terminal window: 

`go test -v -run M2`

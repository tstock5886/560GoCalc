# 560GoCalc

## Testing
You'll need to have go installed, and perhaps the coverage plugin, which can be installed with

`go get golang.org/x/tools/cmd/cover`

Test the code by running . . .

`go test -coverprofile -run ''`

## Usage

The *calculator* package has it's own object, *MyNum*, and supports 3 creation functions, and 6 operations as follows.

- GetNumFromBin(string), takes a binary string and returns a MyNum.
- GetNumFromHex(string), takes a hex string and returns a MyNum.
- GetNumFromOct(string), takes an oct string and returns a MyNum.
- GetNumFromInt(int), takes an integer and returns a MyNum.
- MyNum.Add(MyNum*) - add a MyNum to the receiver MyNum (returns the receiver)
- MyNum.Subtract(MyNum*) - subtract a MyNum from the receiver MyNum (returns the receiver)
- MyNum.Multiply(MyNum*) - multiply a MyNum by the receiver MyNum (returns the receiver)
- MyNum.Divide(MyNum*) - divide a MyNum into the receiver MyNum (returns the receiver)
- MyNum.SqRoot() - calculate the square root of the receiver MyNum (returns the receiver)
- MyNum.Exponent(MyNum*) - raise the receiver to a MyNum power (returns the receiver)
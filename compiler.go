package main

import "fmt"

type op int

const (
   imm op = iota
   arg
   plus
   min
   mul
   div
)

type AST struct {
   Op op
   A  *AST
   B  *AST
   N  int
}



func main() {
    //a := &AST{Op: imm, N: 5}
    //b := &AST{Op: plus, A: a, B: &AST{Op: arg, N: 0}}
    input := "[ a b ] a*a + b*b"
    value := []rune(input)
    for index, char := range value {
        //make a stack and start pusing the [] to identify the start of a function and its end
        //also check on a *-+/ for starting new operations with the last value and the next
        //or can be a ( which pushes last value to a stack and picks up a new "first value" for this operation and ) indicating that order of operation is finish and
        //you should pull again the last value that you pushed

        //found [ start registering args to a map
        // variable inside [] add to map of variables
        // found ] stop registering new variables for the map
        // found variable put to stack
        // found inmmediate number put to stack
        // found operation activate operation mode and upon next variable or inmediate add to the structure
        // found ( can push another variable to stack and increment the indent counter
        // found ) if no operation mode (that would be an error) then add the operations to the structure and decrement the indnet counter
    }
}

package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

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
	Op     op
	Left   *AST
	Right  *AST
	Value  int
	Parent *AST
}

func main() {
	//TODO(josuer08): Change this for a argv reader and all of the printing can
	// be moved to writing to a file or just use standalone with redirection not
	//quite sure about it.
	input := "[ a b ] ((a*b) + (5*5))-3"

	variables, program := extractVariables(input)
	//fmt.Println(variables, program)
	Tree := AST{}
	firstPass(variables, program, &Tree)
	//fmt.Println(Tree)
	secondPass(&Tree)
	slices.Sort(variables)
	thirdPass(&Tree, variables)
	//printer(&Tree)

}

// printer si a function that prints in Reverse Pollish Notation the AST
func printer(tree *AST) {
	switch {
	case tree.Op == imm:
		fmt.Print(tree.Value)
	case tree.Op == arg:
		fmt.Printf("%c", tree.Value)
	default:
		fmt.Print("(")
		switch tree.Op {
		case min:
			fmt.Print("-")
		case plus:
			fmt.Print("+")
		case div:
			fmt.Print("/")
		case mul:
			fmt.Print("*")
		}
		fmt.Print(",")
		printer(tree.Left)
		fmt.Print(",")
		printer(tree.Right)
		fmt.Print(")")

	}
}

// firstPass Is a function that makes the first pass of the compiler,
// it converts the variable and program into an AST
func firstPass(variables, program []rune, node *AST) {
	pass := node
	switch program[0] {
	case '-':
		node.Op = min
	case '+':
		node.Op = plus
	case '*':
		node.Op = mul
	case '/':
		node.Op = div
	case '(':
		if node.Left == nil {
			node.Left = &AST{}
			node.Left.Parent = node
			pass = node.Left
		} else {
			node.Right = &AST{}
			node.Right.Parent = node
			pass = node.Right
		}
	case ')':
		pass = node.Parent

	default:
		if program[0] > 47 && program[0] < 58 {
			var zeroOp op
			if node.Op == zeroOp {
				node.Left = &AST{Op: imm, Value: int(program[0]) - 48}
			} else {
				node.Right = &AST{Op: imm, Value: int(program[0]) - 48}
			}
		} else if slices.Contains(variables, program[0]) {
			if node.Op != plus && node.Op != min && node.Op != mul && node.Op != div {
				node.Left = &AST{Op: arg, Value: int(program[0])}
			} else {
				node.Right = &AST{Op: arg, Value: int(program[0])}
			}
		}
	}
	if len(program) > 1 {
		firstPass(variables, program[1:], pass)
	}
}

// secondPass takes an AST and reduces the operations that only include imm
// values so the program results in a more compact one with precalculated imms
func secondPass(node *AST) {
	if node.Op == arg {
		return
	}
	if node.Op == imm {
		return
	}
	if node.Right.Op == imm && node.Left.Op == imm {
		switch node.Op {
		case min:
			node.Value = node.Left.Value - node.Right.Value
		case plus:
			node.Value = node.Left.Value + node.Right.Value
		case mul:
			node.Value = node.Left.Value * node.Right.Value
		case div:
			node.Value = node.Left.Value / node.Right.Value
		}
		node.Op = imm
		node.Left = nil
		node.Right = nil
		return
	}
	if node.Left.Op != arg && node.Left.Op != imm {
		secondPass(node.Left)
	}
	if node.Right.Op != arg && node.Right.Op != imm {
		secondPass(node.Right)
	}
}

func thirdPass(node *AST, variables []rune) {
	switch node.Op {
	case arg:
		number, found := slices.BinarySearch(variables, rune(node.Value))
		if found {
			fmt.Printf("AR %d\n", number)
		}
	case imm:
		fmt.Printf("IM %d\n", node.Value)
	default:
		switch node.Left.Op {
		case arg:
			number, valid := slices.BinarySearch(variables, rune(node.Left.Value))
			if valid {
				fmt.Printf("AR %d\n", number)
			}
		case imm:
			fmt.Printf("IM %d\n", node.Left.Value)
		default:
			thirdPass(node.Left, variables)
		}
		switch node.Right.Op {
		case arg:
			fmt.Println("SW")
			number, valid := slices.BinarySearch(variables, rune(node.Right.Value))
			if valid {
				fmt.Printf("AR %d\n", number)
			}
			fmt.Println("SW")
		case imm:
			fmt.Println("SW")
			fmt.Printf("IM %d\n", node.Right.Value)
			fmt.Println("SW")
		default:
			fmt.Println("PU")
			thirdPass(node.Right, variables)
			fmt.Println("SW")
			fmt.Println("PO")
		}
		switch node.Op {
		case mul:
			fmt.Println("MU")
		case div:
			fmt.Println("DI")
		case min:
			fmt.Println("SU")
		case plus:
			fmt.Println("AD")

		}

	}

}

// extractVariables receives the original program string and converts it in
// two rune slices, the first containing the variables and a second containing
// the trimmed program
func extractVariables(input string) ([]rune, []rune) {
	variables := strings.Split(input, "]")
	// Cleaning out the variables that are gettting extracted
	variables[0] = strings.Split(variables[0], "[")[1]
	variables[0] = strings.Trim(variables[0], " ")
	cleanVariables := []rune(variables[0])
	var resultVariables []rune
	for _, v := range cleanVariables {
		if v != ' ' {
			resultVariables = append(resultVariables, v)
		}
	}
	//Cleaning out the program that is getting extracted
	variables[1] = strings.Trim(variables[1], " ")
	cleanProgram := []rune(variables[1])
	var resultProgram []rune
	for _, v := range cleanProgram {
		if v != ' ' {
			resultProgram = append(resultProgram, v)
		}
	}
	return resultVariables, resultProgram
}

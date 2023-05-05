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
	//a := &AST{Op: imm, N: 5
	//b := &AST{Op: plus, A: a, B: &AST{Op: arg, N: 0}}
	input := "[ a b ] (a*a) + (5*b)"
	//value := []rune(input)

	variables, program := extractVariables(input)

	fmt.Println(variables, program)
	Tree := AST{}
	firstPass(variables, program, &Tree)
	fmt.Println(Tree)
	printer(&Tree)
	//si es una letra y el stack esta sin setear pon en el A del stack un AST arg
	//si es una operacion setea la op en el stack
	//si es un abrir parentesis apunta al lado que este disponible del AST
	//si es una letra y ya esta seteada la op mete un AST arg a la otra letra
	//si es un cerrar parentesis coge para el pai.

	//los numeros se portan justo como las letras.

}

func printer(tree *AST) {
	switch {
	case tree.Op == imm:
		fmt.Print(tree.Value)
		//fmt.Print(tree.Value - 48)
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
				//a := &AST{Op: imm, N: 5
			} else {
				node.Right = &AST{Op: imm, Value: int(program[0]) - 48}
			}
		} else if slices.Contains(variables, program[0]) {
			//var zeroOp op
			if node.Op != 2 && node.Op != 3 && node.Op != 4 && node.Op != 5 {
				node.Left = &AST{Op: arg, Value: int(program[0])}
				//a := &AST{Op: imm, N: 5
			} else {
				node.Right = &AST{Op: arg, Value: int(program[0])}
			}

		}

	}
	if len(program) > 1 {
		firstPass(variables, program[1:], pass)

	}
	return

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

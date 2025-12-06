package d06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nozarashi20/aoc-2025/helpers"
)

type Operation struct {
	operator string
	operands []string
}

func (op *Operation) Result() int64 {
	total, _ := strconv.ParseInt(op.operands[0], 10, 64)
	for i := 1; i < len(op.operands); i++ {
		operand, _ := strconv.ParseInt(op.operands[i], 10, 64)
		total = Eval(total, operand, op.operator)
	}
	return total
}

func Eval(a, b int64, op string) int64 {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	default:
		panic(fmt.Sprintf("unknown operator %s", op))
	}
}

func getOperators(lines []string) []string {
	var operators []string
	for _, c := range lines[len(lines)-1] {
		if isOperator(c) {
			operators = append(operators, string(c))
		}
	}
	return operators
}

func ParseOperations(lines []string) []Operation {
	operators := getOperators(lines)
	operations := make([]Operation, len(operators))
	for i, op := range operators {
		operations[i] = Operation{operator: op}
	}

	for _, line := range lines[0 : len(lines)-1] {
		nums := strings.Fields(line)
		for i, num := range nums {
			operations[i].operands = append(operations[i].operands, num)
		}
	}

	return operations
}

func ParseOperations2(lines []string) []Operation {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	operators := getOperators(lines)
	operations := make([]Operation, len(operators))
	for i, op := range operators {
		operations[i] = Operation{operator: op}
	}

	grid = helpers.TransposeRagged(grid)
	cursor := len(operators) - 1
	for i := len(grid) - 1; i >= 0; i-- {
		row := grid[i]
		if strings.TrimSpace(strings.Join(row, "")) == "" {
			cursor--
			continue
		}
		num := strings.TrimSpace(strings.Join(row[0:len(row)-1], ""))
		operations[cursor].operands = append(operations[cursor].operands, num)
	}

	return operations
}

func PartOne(lines []string) int64 {
	operations := ParseOperations(lines)
	var total int64
	for _, op := range operations {
		result := op.Result()
		total += result
	}
	return total
}

type Range struct {
	start int
	end   int
}

func PartTwo(lines []string) int64 {
	operations := ParseOperations2(lines)
	var total int64
	for _, op := range operations {
		result := op.Result()
		total += result
	}
	return total
}

func isOperator(c rune) bool {
	return c == '+' || c == '*'
}

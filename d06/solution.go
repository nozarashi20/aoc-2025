package d06

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	operator    string
	operands    []int64
	operandsStr []string
}

func (op *Operation) Result() int64 {
	total := op.operands[0]
	for i := 1; i < len(op.operands); i++ {
		total = Eval(total, op.operands[i], op.operator)
	}
	return total
}

func (op *Operation) Result2() int64 {
	colCount := len(op.operandsStr[0])
	newOperands := make([]int64, colCount)
	for i := 0; i < colCount; i++ {
		var num string
		for j := 0; j < len(op.operandsStr); j++ {
			digit := op.operandsStr[j][i]
			if digit == 'X' {
				continue
			}
			num += string(digit)
		}
		operand, _ := strconv.ParseInt(num, 10, 64)
		newOperands[i] = operand
	}

	total := newOperands[0]
	for i := 1; i < len(newOperands); i++ {
		total = Eval(total, newOperands[i], op.operator)
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

func ParseOperations(lines []string) []Operation {
	var operations []Operation
	for lineNum, line := range lines {
		row := strings.Fields(line)
		for i, val := range row {
			if i >= len(operations) {
				operations = append(operations, Operation{operands: []int64{}})
			}
			if lineNum == len(lines)-1 {
				operations[i].operator = val
			} else {
				operand, _ := strconv.ParseInt(val, 10, 64)
				operations[i].operands = append(operations[i].operands, operand)
			}
		}
	}
	return operations
}

func ParseOperations2(lines []string) []Operation {
	var operations []Operation
	sizes := operationsSizes(lines[len(lines)-1])
	maxDigitPerOp := make([]int, len(sizes)+1)
	for lineNum, line := range lines {
		split := splitByOperationSizes(line, sizes)
		for opNum, chunk := range split {
			if lineNum == 0 {
				operations = append(operations, Operation{operands: []int64{}})
			}

			if lineNum == len(lines)-1 {
				operations[opNum].operator = strings.TrimSpace(chunk)
			} else {
				chunk = strings.ReplaceAll(chunk, " ", "X")
				maxDigitPerOp[opNum] = max(maxDigitPerOp[opNum], len(chunk))
				operations[opNum].operandsStr = append(operations[opNum].operandsStr, chunk)
			}
		}

		for k, op := range operations {
			op.operands = make([]int64, len(op.operandsStr))
			maxDigit := maxLen(op.operandsStr)
			for i := 0; i < len(op.operandsStr); i++ {
				op.operandsStr[i] = op.operandsStr[i] + strings.Repeat("X", maxDigit-len(op.operandsStr[i]))
				operand, _ := strconv.ParseInt(op.operandsStr[i], 10, 64)
				op.operands[i] = operand
			}
			operations[k] = op
		}
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

func operationsSizes(line string) []Range {
	var sizes []Range
	var start int
	for i := 1; i < len(line); i++ {
		if isOperator(rune(line[i])) {
			sizes = append(sizes, Range{start, i - 1})
			start = i
		}
	}
	return sizes
}

func splitByOperationSizes(line string, sizes []Range) []string {
	var chunks []string
	for _, size := range sizes {
		chunks = append(chunks, line[size.start:size.end])
	}
	lastStart := sizes[len(sizes)-1].end + 1
	if lastStart < len(line) {
		chunks = append(chunks, line[lastStart:])
	}
	return chunks
}

func PartTwo(lines []string) int64 {
	operations := ParseOperations2(lines)
	var total int64
	for _, op := range operations {
		result := op.Result2()
		total += result
	}
	return total
}

func maxLen(strs []string) int {
	m := 0
	for _, str := range strs {
		if len(str) > m {
			m = len(str)
		}
	}
	return m
}

func isOperator(c rune) bool {
	return c == '+' || c == '*'
}

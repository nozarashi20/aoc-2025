package d03

import (
	"math"
	"strconv"
)

func ParseBank(line string) []int {
	var bank []int
	for _, c := range line {
		n, _ := strconv.Atoi(string(c))
		bank = append(bank, n)
	}
	return bank
}

func PartOne(lines []string) int {
	res := 0
	for _, line := range lines {
		res += LargestJoltage(line)
	}
	return res
}

func PartTwo(lines []string) int64 {
	res := int64(0)
	for _, line := range lines {
		res += LargestJoltage2(line)
	}
	return res
}

func LargestJoltage(line string) int {
	bank := ParseBank(line)
	res := 0
	firstBattery := bank[0]

	for i := 1; i < len(bank); i++ {
		battery := bank[i]
		potentialJoltage := 10*firstBattery + battery
		if potentialJoltage > res {
			res = potentialJoltage
		}

		if battery > firstBattery {
			firstBattery = battery
		}
	}

	return res
}

func LargestJoltage2(line string) int64 {
	bank := ParseBank(line)
	idx := int64(0)
	res := int64(0)
	for i := 11; i >= 0; i-- {
		if len(bank) == 0 {
			break
		}
		idx = MaxNumber(bank, i)
		res += int64(math.Pow10(i)) * int64(bank[idx])
		bank = bank[idx+1:]
	}
	return res
}

func MaxNumber(numbers []int, power10 int) int64 {
	if len(numbers) == 0 || len(numbers) < power10+1 {
		return 0
	}
	res := int64(0)
	bestNum := int64(numbers[0])
	bestNumIndex := int64(0)
	for i := 1; i < len(numbers)-power10; i++ {
		num := int64(numbers[i])
		potential := num * int64(math.Pow10(power10))
		if potential > res {
			res = num
		}
		if num > bestNum {
			bestNum = num
			bestNumIndex = int64(i)
		}
	}
	return bestNumIndex
}

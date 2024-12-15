package day13

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return parts(lines, false)
}

func Part2(lines []string) int {
	return parts(lines, true)
}

func parts(lines []string, isPart2 bool) int {
	sum := 0

	for i := 0; i < len(lines); i += 4 {
		// Button A
		btnA := strings.Split(strings.Split(lines[i], ": ")[1], ", ")
		x1, _ := strconv.Atoi(btnA[0][1:])
		y1, _ := strconv.Atoi(btnA[1][1:])

		// Button B
		btnB := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")
		x2, _ := strconv.Atoi(btnB[0][1:])
		y2, _ := strconv.Atoi(btnB[1][1:])

		// Prize
		prizes := strings.Split(strings.Split(lines[i+2], ": ")[1], ", ")
		xE, _ := strconv.Atoi(prizes[0][2:])
		yE, _ := strconv.Atoi(prizes[1][2:])

		if isPart2 {
			xE += 10000000000000
			yE += 10000000000000
		}

		a, b := solveLE(Equation{x1, x2, xE}, Equation{y1, y2, yE})
		if a != -1 && b != -1 {
			sum += a*3 + b
		}
	}

	return sum
}

type Equation struct {
	factorA, factorB, equals int
}

// Gaussian elimination
func solveLE(eq1, eq2 Equation) (int, int) {
	// printGauss(eq1, eq2)

	// 1. eliminate a in eq2
	ggTA := ggT(eq1.factorA, eq2.factorA)
	quot1 := eq1.factorA / ggTA
	quot2 := eq2.factorA / ggTA

	// fmt.Printf("ggT: %d quot1: %d quot2: %d\n", ggTA, quot1, quot2)

	eq2.factorA = 0
	eq2.factorB = eq2.factorB*quot1 - eq1.factorB*quot2
	eq2.equals = eq2.equals*quot1 - eq1.equals*quot2

	// printGauss(eq1, eq2)

	// 2. reduce b to 1 in eq2
	if (eq2.equals % eq2.factorB) != 0 {
		return -1, -1
	}
	eq2.equals = eq2.equals / eq2.factorB
	eq2.factorB = 1

	// printGauss(eq1, eq2)

	// 3. eliminate b in eq1
	eq1.equals = eq1.equals - eq1.factorB*eq2.equals
	eq1.factorB = 0

	// printGauss(eq1, eq2)

	// 4. reduce a to 1 in eq1
	if (eq1.equals % eq1.factorA) != 0 {
		return -1, -1
	}
	eq1.equals = eq1.equals / eq1.factorA
	eq1.factorA = 1

	// printGauss(eq1, eq2)

	// a, b
	return eq1.equals, eq2.equals
}

// func printGauss(eq1, eq2 Equation) {
// 	fmt.Printf("%4d %4d | %4d\n", eq1.factorA, eq1.factorB, eq1.equals)
// 	fmt.Printf("%4d %4d | %4d\n", eq2.factorA, eq2.factorB, eq2.equals)
// 	fmt.Println()
// }

func ggT(a, b int) int {
	var h int
	for b != 0 {
		h = a % b
		a = b
		b = h
	}

	return a
}

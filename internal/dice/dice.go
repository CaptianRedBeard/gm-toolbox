package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

// DefaultDiceSides sets the default die type for a roll with no expression.
const DefaultDiceSides = 20

type Result struct {
	Total int
	Rolls []int
	Expr  string
}

func Roll(expression string) Result {
	if expression == "" {
		total, rolls := rollNDice(1, DefaultDiceSides)
		return Result{
			Total: total,
			Rolls: rolls,
			Expr:  fmt.Sprintf("d%d", DefaultDiceSides),
		}
	}

	re := regexp.MustCompile(`^^(\d*)d(\d+)$`)
	matches := re.FindStringSubmatch(expression)
	if matches != nil {
		count, _ := strconv.Atoi(matches[1])
		if count == 0 {
			count = 1 // Handles "d6"
		}
		sides, _ := strconv.Atoi(matches[2])

		total, rolls := rollNDice(count, sides)

		// The parser is responsible for creating the final, clean result
		return Result{
			Total: total,
			Rolls: rolls,
			Expr:  fmt.Sprintf("%dd%d", count, sides),
		}
	}

	// Fallback for invalid expressions
	return Result{}
}

// rollNDice performs the core logic of rolling N dice with Y sides.
// It returns the total and the slice of individual rolls.
func rollNDice(count, sides int) (int, []int) {
	rolls := make([]int, count)
	total := 0
	for i := 0; i < count; i++ {
		roll := rand.Intn(sides) + 1
		rolls[i] = roll
		total += roll
	}
	return total, rolls
}

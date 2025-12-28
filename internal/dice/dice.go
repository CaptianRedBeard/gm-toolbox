package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// DefaultDiceSides sets the default die type for a roll with no expression.
const DefaultDiceSides = 20

type Result struct {
	Total int
	Rolls []int
	Expr  string
}

// Roll handles the dice rolling logic.
func Roll(expression string) Result {
	if expression == "" {
		// Use the constant for the default expression string too
		return singleDieRoll(DefaultDiceSides)
	}

	if strings.HasPrefix(expression, "d") {
		sidesStr := expression[1:]
		diceSides, err := strconv.Atoi(sidesStr)

		if err != nil {
			// If parsing fails (e.g., "dABC"), fall back to a d20 roll
			// but show the original, bad expression in the result.
			result := singleDieRoll(DefaultDiceSides)
			result.Expr = expression
			return result
		}

		return singleDieRoll(diceSides)
	}

	// Placeholder for all other cases
	return Result{}
}

// singleDieRoll is a helper for the common "roll one die" logic.
func singleDieRoll(sides int) Result {
	singleRoll := rand.Intn(sides) + 1
	return Result{
		Total: singleRoll,
		Rolls: []int{singleRoll},
		Expr:  fmt.Sprintf("d%d", sides), // Create the canonical expression string
	}
}

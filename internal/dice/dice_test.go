package dice

import (
	"testing"
)

// TestRollMultipleDice verifies rolling multiple dice, like "2d4".
func TestRollMultipleDice(t *testing.T) {
	// Arrange
	const expectedCount = 2
	const expectedSides = 4
	expression := "2d4"

	// Act
	result := Roll(expression)

	// Assert
	if result.Expr != expression {
		t.Errorf("Expected Expr '%s', got '%s'", expression, result.Expr)
	}

	if len(result.Rolls) != expectedCount {
		t.Errorf("Expected %d rolls, got %d", expectedCount, len(result.Rolls))
	}

	// Check that each roll is within the valid range for a d4
	for _, roll := range result.Rolls {
		if roll < 1 || roll > expectedSides {
			t.Errorf("Expected roll between 1 and %d, got %d", expectedSides, roll)
		}
	}

	// Check that the total is the sum of all rolls
	expectedTotal := 0
	for _, roll := range result.Rolls {
		expectedTotal += roll
	}
	if result.Total != expectedTotal {
		t.Errorf("Expected Total to be %d, got %d", expectedTotal, result.Total)
	}
}

// TestRollSingleDie verifies that rolling a simple expression like "d6"
func TestRollSingleDie(t *testing.T) {
	// Arrange
	const expectedSides = 6
	const expectedRollsCount = 1
	expression := "1d6"

	// Act
	result := Roll(expression)

	// Assert
	if result.Expr != expression {
		t.Errorf("Expected Expr '%s', got '%s'", expression, result.Expr)
	}

	if len(result.Rolls) != expectedRollsCount {
		t.Errorf("Expected %d roll, got %d", expectedRollsCount, len(result.Rolls))
	}

	// Check that the single roll is within the valid range for a d6
	roll := result.Rolls[0]
	if roll < 1 || roll > expectedSides {
		t.Errorf("Expected roll between 1 and %d, got %d", expectedSides, roll)
	}

	// The total must equal the single roll
	if result.Total != roll {
		t.Errorf("Expected Total to be %d, got %d", roll, result.Total)
	}
}

// TestRollDefaultToD20 verifies that rolling with an empty string
func TestRollDefaultToD20(t *testing.T) {
	// Arrange
	const expectedSides = 20
	const expectedRollsCount = 1

	// Act
	result := Roll("")

	// Assert
	if result.Expr != "d20" {
		t.Errorf("Expected Expr 'd20', got '%s'", result.Expr)
	}

	if len(result.Rolls) != expectedRollsCount {
		t.Errorf("Expected %d roll, got %d", expectedRollsCount, len(result.Rolls))
	}

	// Check that the single roll is within the valid range for a d20
	roll := result.Rolls[0]
	if roll < 1 || roll > expectedSides {
		t.Errorf("Expected roll between 1 and %d, got %d", expectedSides, roll)
	}

	// The total must equal the single roll
	if result.Total != roll {
		t.Errorf("Expected Total to be %d, got %d", roll, result.Total)
	}
}

package new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	// Test for value 1000
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax, "Tax for 1000 should be 10")

	// Test for value 0
	tax = CalculateTax(0.0)
	assert.Equal(t, 0.0, tax, "Tax for 0 should be 0")

	// Test for negative values
	tax = CalculateTax(-100.0)
	assert.Equal(t, 0.0, tax, "Tax for 100 should be 0")

	// Test for value 500
	tax = CalculateTax(500)
	assert.Equal(t, 5.0, tax, "Tax for 500 should be 5")
}

package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	assert.Equal(t, answer, int32(3))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	assert.Equal(t, answer, int32(1))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	assert.Equal(t, answer, int32(2))
}

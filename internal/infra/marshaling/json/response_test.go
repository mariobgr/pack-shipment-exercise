package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromCalculatedShipment(t *testing.T) {
	input := map[int]int{250: 2, 500: 1, 1000: 2, 2000: 3}

	expected := CalculatedShipmentResponse{
		NumPacks: []string{
			"2 x 250",
			"1 x 500",
			"2 x 1000",
			"3 x 2000",
		},
	}

	result := NewFromCalculatedShipment(input)
	assert.ElementsMatch(t, expected.NumPacks, result.NumPacks)
}

func TestNewFromDomainPacks(t *testing.T) {
	input := []int{250, 500, 1000}

	expected := PackListResponse{
		Packs: []string{
			"250 items",
			"500 items",
			"1000 items",
		},
	}

	result := NewFromDomainPacks(input)
	assert.Equal(t, expected, result)
}

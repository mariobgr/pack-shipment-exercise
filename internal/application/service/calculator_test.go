package service

import (
	"fmt"
	"testing"

	"github.com/mariobgr/pack-shipment-exercise/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestPackCalculatorService_Calculate(t *testing.T) {

	tests := []struct {
		input    int
		expected map[int]int
	}{
		{
			input:    250,
			expected: map[int]int{250: 1},
		},
		{
			input:    256,
			expected: map[int]int{500: 1},
		},
		{
			input:    1,
			expected: map[int]int{250: 1},
		},
		{
			input:    501,
			expected: map[int]int{500: 1, 250: 1},
		},
		{
			input:    12001,
			expected: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			input:    12251,
			expected: map[int]int{5000: 2, 2000: 1, 500: 1},
		},
	}

	service := NewCalculatorService(packGetter{})

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("test with input %d", tt.input), func(t *testing.T) {
			result := service.Calculate(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

type packGetter struct{}

func (pg packGetter) GetPacks() domain.Packs {
	return []int{250, 500, 1000, 2000, 5000}
}

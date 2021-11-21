package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToString(t *testing.T) {
	str := IntToString(10)
	assert.Equal(t, str, "10")

	i := StringToInt("10")
	assert.Equal(t, i, 10)

	assert.Equal(t, IntToInt64(100), int64(100))
}

func TestStringToFloat(t *testing.T) {
	assert.Equal(t, StringToFloat("100"), float64(100))
	assert.Equal(t, FloatToString(100.01), "100.01")
}

func TestFloat64ToInt(t *testing.T) {
	assert.Equal(t, Float64ToInt(100.9), 100)
	assert.Equal(t, IntToFloat64(100), 100.0)
}

package algebraic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/madshov/data-structures/algebraic"
)

func TestNewVector(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return a new vector with dimension 3 and set values", func(t *testing.T) {
		v := algebraic.NewVector(3, 1, 2, 3)

		assert.Len(v, 3)
		assert.Equal(1.0, v[0])
		assert.Equal(2.0, v[1])
		assert.Equal(3.0, v[2])
	})

	t.Run("should return a new vector with dimension 4 and two set values", func(t *testing.T) {
		v := algebraic.NewVector(4, 1, 2)

		assert.Len(v, 4)
		assert.Equal(1.0, v[0])
		assert.Equal(2.0, v[1])
		assert.Equal(0.0, v[2])
		assert.Equal(0.0, v[3])
	})

	t.Run("should return a new vector with dimension 2 and only two set values", func(t *testing.T) {
		v := algebraic.NewVector(2, 1, 2, 3)

		assert.Len(v, 2)
		assert.Equal(1.0, v[0])
		assert.Equal(2.0, v[1])
	})
}

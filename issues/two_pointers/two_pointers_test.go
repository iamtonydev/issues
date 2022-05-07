package two_pointers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		data := []int{2, 5, 6, 7, 11, 12, 56, 70, 123, 142}
		target := 129
		expected := []int{6, 123}
		res, err := TwoSum(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("sum not find", func(t *testing.T) {
		data := []int{2, 5, 6, 7, 11, 12, 56, 70, 123, 142}
		target := 200
		expected := []int{0, 0}
		res, err := TwoSum(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raise error", func(t *testing.T) {
		_, err := TwoSum(nil, 0)
		require.Error(t, err)
		require.Equal(t, errors.New("input data invalid"), err)
	})
}

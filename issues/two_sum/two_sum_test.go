package two_sum

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		data := []int{11, 142, 56, 70, 6, 12, 123, 5, 9, 2}
		target := 129
		expected := []int{6, 123}
		res, err := TwoSum(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("base case", func(t *testing.T) {
		data := []int{11, 142, 56, 70, 6, 12, 123, 5, 9, 2}
		target := 13
		expected := []int{11, 2}
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

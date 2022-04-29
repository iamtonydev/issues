package most_frequent

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostFrequent(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		data := []int{1, 1, 1, 2, 2, 5, 5, 5, 5, 5, 8, 8, 8, 9, 9, 9, 9, 9, 9}
		expected := []int{9, 5}
		res, err := MostFrequent(data, 2)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("data is empty", func(t *testing.T) {
		data := make([]int, 0, 0)
		expected := make([]int, 0, 0)
		res, err := MostFrequent(data, 2)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("n less than repeats count", func(t *testing.T) {
		data := []int{1, 1, 1, 1, 2, 2, 5, 5, 5, 5, 5, 8, 8, 8, 9, 9, 9, 9, 9, 9}
		expected := []int{9, 5, 1, 8, 2}
		res, err := MostFrequent(data, 10)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("n is zero", func(t *testing.T) {
		data := []int{1, 1, 1, 2, 2, 5, 5, 5, 5, 5, 8, 8, 8, 9, 9, 9, 9, 9, 9}
		expected := make([]int, 0, 0)
		res, err := MostFrequent(data, 0)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raise error", func(t *testing.T) {
		_, err := MostFrequent(nil, 2)
		require.Error(t, err)
		require.Equal(t, errors.New("input data invalid"), err)
	})
}

package binary_search

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchElement(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		target := 4
		data := []int{1, 3, 4, 5, 7, 9, 11, 15}
		expected := 2
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target on right edge", func(t *testing.T) {
		target := 15
		data := []int{1, 3, 4, 5, 7, 9, 11, 15}
		expected := 7
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target on left edge", func(t *testing.T) {
		target := 1
		data := []int{1, 3, 4, 5, 7, 9, 11, 15}
		expected := 0
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target not in data", func(t *testing.T) {
		target := 16
		data := []int{1, 3, 4, 5, 7, 9, 11, 15}
		expected := -1
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target not in data 2", func(t *testing.T) {
		target := -5
		data := []int{1, 3, 4, 5, 7, 9, 11, 15}
		expected := -1
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raise error", func(t *testing.T) {
		target := 1
		_, err := SearchElement(nil, target)
		require.Error(t, err)
		require.Equal(t, errors.New("input data invalid"), err)
	})
}

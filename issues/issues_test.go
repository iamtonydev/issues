package issues

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchElement(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		var target int64 = 4
		data := []int64{1, 3, 4, 5, 7, 9, 11, 15}
		var expected int64 = 2
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target on right edge", func(t *testing.T) {
		var target int64 = 15
		data := []int64{1, 3, 4, 5, 7, 9, 11, 15}
		var expected int64 = 7
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target on left edge", func(t *testing.T) {
		var target int64 = 1
		data := []int64{1, 3, 4, 5, 7, 9, 11, 15}
		var expected int64 = 0
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target not in data", func(t *testing.T) {
		var target int64 = 16
		data := []int64{1, 3, 4, 5, 7, 9, 11, 15}
		var expected int64 = -1
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("target not in data 2", func(t *testing.T) {
		var target int64 = -5
		data := []int64{1, 3, 4, 5, 7, 9, 11, 15}
		var expected int64 = -1
		res, err := SearchElement(data, target)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raise error", func(t *testing.T) {
		var target int64 = 1
		_, err := SearchElement(nil, target)
		require.Error(t, err)
		require.Equal(t, errors.New("input data invalid"), err)
	})
}

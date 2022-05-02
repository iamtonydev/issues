package most_frequent_words

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostFrequentWords(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		text := "qwe qwe, - xcv rty fgh vbn\nuio jkl! bnm fgh qwe, fgh rty fgh qwe, fgh qwe, rty qwe, vbn, uio"
		expected := []string{"fgh", "qwe,", "rty", "uio", "-", "bnm", "jkl!", "qwe", "vbn", "vbn,"}
		res, err := MostFrequentWords(text)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("lexicographic order", func(t *testing.T) {
		text := "x a B A D C c d j g"
		expected := []string{"A", "B", "C", "D", "a", "c", "d", "g", "j", "x"}
		res, err := MostFrequentWords(text)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("words count less than 10", func(t *testing.T) {
		text := "qwe, abc def - rty def def def"
		expected := []string{"def", "-", "abc", "qwe,", "rty"}
		res, err := MostFrequentWords(text)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty string", func(t *testing.T) {
		expected := make([]string, 0)
		res, err := MostFrequentWords("")
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})
}

package numberOfPairsWithDiffK

import "testing"

func TestNameOfPairsWithDiffK(t *testing.T) {
	n := numberOfPairsWithDiffK([]int{
		1, 7, 5, 9, 2, 12, 3,
		// 1, 2, 3, 5, 7, 9, 12, // sorted
	}, 2)

	if n != 4 {
		t.Error("The result is", n, "but should be 4")
	}

	t.Log("The result is", n)
}

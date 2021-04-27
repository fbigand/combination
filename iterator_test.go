package combination

import (
	"fmt"
	"testing"
)

func TestCombinationIterator(t *testing.T) {
	listA := []string{"A1", "A2", "A3"}
	listB := []string{"B1", "B2"}
	listC := []string{"C1"}
	combinationIterator := NewCombinationIterator([]int{
		len(listA),
		len(listB),
		len(listC),
	})

	var wanted Combination
	var hasNext bool

	wanted = Combination{0, 0, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, true)
	wanted = Combination{1, 0, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, true)
	wanted = Combination{0, 1, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, true)
	wanted = Combination{1, 1, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, true)
	wanted = Combination{2, 0, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, true)
	wanted = Combination{2, 1, 0}
	checkCorrectCombination(t, combinationIterator, wanted)

	hasNext = combinationIterator.Next()
	checkHasNext(t, combinationIterator, hasNext, false)
}

func checkHasNext(t *testing.T, iterator CombinationIterator, currentHasNext bool, wantedHasNext bool) {
	if currentHasNext != wantedHasNext {
		t.Errorf("Has next should be %v. Outputted combination: %v", wantedHasNext, iterator.GetCombination())
	}
}

func checkCorrectCombination(t *testing.T, iterator CombinationIterator, wanted Combination) {
	currentCombination := iterator.GetCombination()
	fmt.Println(currentCombination)
	if len(currentCombination) != len(wanted) {
		t.Errorf("current combination not the same length as wanted. current = %v, wanted = %v", currentCombination, wanted)
	}

	for i := range currentCombination {
		if currentCombination[i] != wanted[i] {
			t.Errorf("Combination should be %v, instead is %v", wanted, currentCombination)
			break
		}
	}
}

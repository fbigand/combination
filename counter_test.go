package combination

import (
	"testing"

	"github.com/fbigand/combination/util"
)

func TestCounterState(t *testing.T) {
	counter := newCounter(3, []uint{0, 0, 3})

	checkState(t, counter, []uint{3, 0})
	checkCombinations(t, counter, []Combination{{0, 0, 0}})

	counter.next()
	checkState(t, counter, []uint{2, 1})
	checkCombinations(t, counter, []Combination{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})

	counter.next()
	checkState(t, counter, []uint{1, 2})
	checkCombinations(t, counter, []Combination{{1, 1, 0}, {1, 0, 1}, {0, 1, 1}})

	counter.next()
	checkState(t, counter, []uint{0, 3})
	checkCombinations(t, counter, []Combination{{1, 1, 1}})

	counter.next()
	checkState(t, counter, []uint{2, 0, 1})
	checkCombinations(t, counter, []Combination{{2, 0, 0}, {0, 2, 0}, {0, 0, 2}})

	counter.next()
	checkState(t, counter, []uint{1, 1, 1})
	checkCombinations(t, counter, []Combination{{2, 1, 0}, {2, 0, 1}, {1, 2, 0}, {0, 2, 1}, {1, 0, 2}, {0, 1, 2}})

	counter.next()
	checkState(t, counter, []uint{0, 2, 1})
	checkCombinations(t, counter, []Combination{{2, 1, 1}, {1, 2, 1}, {1, 1, 2}})

	counter.next()
	checkState(t, counter, []uint{1, 0, 2})
	checkCombinations(t, counter, []Combination{{2, 2, 0}, {2, 0, 2}, {0, 2, 2}})

	counter.next()
	checkState(t, counter, []uint{0, 1, 2})
	checkCombinations(t, counter, []Combination{{2, 2, 1}, {2, 1, 2}, {1, 2, 2}})

	counter.next()
	checkState(t, counter, []uint{0, 0, 3})
	checkCombinations(t, counter, []Combination{{2, 2, 2}})

	hasNext := counter.next()
	if hasNext {
		t.Error("Counter should not have next at this point")
	}
}

func checkState(t *testing.T, counter *counter, wantedState []uint) {
	if !util.EqualsSliceUint(counter.state, wantedState) {
		t.Errorf("Wrong counter state. Current: %v, wanted: %v", counter.state, wantedState)
	}
}

func checkCombinations(t *testing.T, counter *counter, wantedCombinations []Combination) {
	counterCombs := counter.getCombinations()
	for combIndex := range wantedCombinations {
		for valueIndex := range wantedCombinations[combIndex] {
			if counterCombs[combIndex][valueIndex] != wantedCombinations[combIndex][valueIndex] {
				t.Errorf("Wrong counter combinations. Current state: %v, current combinations: %v, wanted combinations: %v", counter.state, counterCombs, wantedCombinations)
			}
		}
	}
}

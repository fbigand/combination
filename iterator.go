package combination

import (
	"github.com/fbigand/combination/util"
)

// CombinationIterator responds to the following rules.
//
// Consider n (initialValue) lists ordered by importance of their elements.
// This iterator will give in order the combinations with most important
// element to combinations with less important element. It will provide first
// the combination with the top element of each list, then the combinations
// with five top elements and one second top element etc.
type CombinationIterator struct {
	counter           *counter
	combinationsQueue []Combination
	combinationIndex  uint
	lastCombination   Combination
}

type Combination []int

// NewCombinationIterator creates and returns a new CombinationIterator.
// See CombinationIterator documentation for more information.
func NewCombinationIterator(lastCombination Combination) CombinationIterator {
	initialValue := uint(len(lastCombination))
	endStateValue := util.MaxInt(lastCombination...)
	endState := make([]uint, initialValue)
	util.InitSliceUint(endState, uint(endStateValue))

	counter := newCounter(initialValue, endState)
	firstCombinations := counter.getCombinations()

	combinationIterator := CombinationIterator{
		counter:           counter,
		combinationsQueue: firstCombinations,
		combinationIndex:  0,
		lastCombination:   lastCombination,
	}
	return combinationIterator
}

// Next iterates one step on the iterator. Returns false if the iterator
// reached the end.
func (ci *CombinationIterator) Next() bool {
	hasNext := true
	if int(ci.combinationIndex) == len(ci.combinationsQueue) {
		hasNext = ci.counter.next()
		if hasNext {
			ci.combinationsQueue = ci.counter.getCombinations()
			ci.filterCombinations()
			ci.combinationIndex = 0
		}
	} else {
		ci.combinationIndex++
	}
	return hasNext
}

// GetCombination returns the current combination of the iterator.
func (ci *CombinationIterator) GetCombination() Combination {
	return ci.combinationsQueue[ci.combinationIndex]
}

// filterCombinations filter the combinations to test. Basically, this
// function removes all combinations containing out of range index
func (ci *CombinationIterator) filterCombinations() {
	for combIndex, combination := range ci.combinationsQueue {
		for listIndex := range combination {
			if combination[listIndex] > ci.lastCombination[listIndex] {
				// remove combination from combinations queue
				ci.combinationsQueue = append(ci.combinationsQueue[:combIndex], ci.combinationsQueue[combIndex+1:]...)
				break
			}
		}
	}
}

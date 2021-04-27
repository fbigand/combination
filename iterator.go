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
	lastCombination   Combination
}

type Combination []int

// NewCombinationIterator creates and returns a new CombinationIterator.
// See CombinationIterator documentation for more information.
func NewCombinationIterator(listsLength []int) CombinationIterator {
	initialValue := uint(len(listsLength))
	endStateValue := util.MaxInt(listsLength...)
	endState := make([]uint, initialValue)
	endState[initialValue-1] = uint(endStateValue)

	counter := newCounter(initialValue, endState)
	firstCombinations := counter.getCombinations()

	lastCombination := make(Combination, len(listsLength))
	for i := range lastCombination {
		lastCombination[i] = listsLength[i] - 1
	}

	combinationIterator := CombinationIterator{
		counter:           counter,
		combinationsQueue: firstCombinations,
		lastCombination:   lastCombination,
	}
	return combinationIterator
}

// Next iterates one step on the iterator. Returns false if the iterator
// reached the end.
func (ci *CombinationIterator) Next() bool {
	hasNext := true
	if len(ci.combinationsQueue) == 1 {
		hasNext = ci.requeueCombinations()
	} else {
		ci.combinationsQueue = ci.combinationsQueue[1:]
	}
	return hasNext
}

func (ci *CombinationIterator) requeueCombinations() bool {
	hasNext := ci.counter.next()
	if hasNext {
		ci.combinationsQueue = ci.counter.getCombinations()
		ci.removeOutOfRangeCombinations()
		if len(ci.combinationsQueue) == 0 {
			hasNext = ci.requeueCombinations()
		}
	}
	return hasNext
}

// GetCombination returns the current combination of the iterator.
func (ci *CombinationIterator) GetCombination() Combination {
	return ci.combinationsQueue[0]
}

func (ci *CombinationIterator) removeOutOfRangeCombinations() {
	for combIndex := len(ci.combinationsQueue) - 1; combIndex >= 0; combIndex-- {
		// for combIndex, combination := range ci.combinationsQueue {
		combination := ci.combinationsQueue[combIndex]
		for listIndex := range combination {
			if combination[listIndex] > ci.lastCombination[listIndex] {
				ci.combinationsQueue = removeFromCombinationSlice(ci.combinationsQueue, combIndex)
				break
			}
		}
	}
}

func removeFromCombinationSlice(slice []Combination, index int) []Combination {
	return append(slice[:index], slice[index+1:]...)
}

package combination

import (
	"strconv"

	"github.com/fbigand/combination/util"
)

// counter is used internally by CombinationIterator.
// At each step, it will give how many elements from which
// position in the list the combination will have.
//
// Example: first values starting with initialValue = 6
//    1: 6
//    2: 5 1
//    3: 4 2
//    4: 3 3
//    5: 2 4
//    6: 1 5
//    7: 0 6
//    8: 1 0 5
type counter struct {
	initialValue          uint
	state                 []uint
	endState              []uint
	combinationUnsetIndex int
}

// counterCombsCache record counter combinations that has already been computed
// by the method GetCombinations
var counterCombinationsCache map[string][]Combination = make(map[string][]Combination)

// newCounter creates and returns a new counter.
// See CombinationIterator documentation for more information.
func newCounter(initialValue uint, endState []uint) *counter {
	c := counter{}

	c.initialValue = initialValue
	c.combinationUnsetIndex = int(initialValue) + 1

	c.state = make([]uint, 2)
	c.state[0] = initialValue

	if util.MaxUint(endState...) > initialValue {
		panic("newCounter: endState cannot be reached with these parameters")
	}
	if util.SumUint(endState...) != initialValue {
		panic("newCounter: endState sum must be equal to initialValue")
	}
	c.endState = endState

	return &c
}

func (c *counter) getStateToString() string {
	strState := ""
	for _, v := range c.state {
		strState += strconv.FormatUint(uint64(v), 10)
	}
	return strState
}

// next set the counter to the next position. Undoable.
//
// Returns true if the operation was successful. Returns false
// if reached the end state
func (c *counter) next() bool {
	hasNext := true
	if util.EqualsSliceUint(c.state, c.endState) {
		hasNext = false
	} else {
		c.incrementState()
	}

	return hasNext
}

func (c *counter) incrementState() {
	if c.state[0] == 0 {
		iReset := 1
		for c.state[iReset] == 0 {
			iReset++
		}
		c.state[iReset] = 0

		// prevent out of range
		if iReset+1 == len(c.state) {
			c.state = append(c.state, 0)
			// c.state = c.state[:len(c.state)+1]
		}
		c.state[iReset+1]++

		c.state[0] = c.initialValue - util.SumUint(c.state...)
	} else {
		c.state[0]--
		c.state[1]++
	}
}

func (c *counter) getCombinations() []Combination {
	var combinations []Combination

	strState := c.getStateToString()
	cachedCombs, combsInCache := counterCombinationsCache[strState]
	if combsInCache {
		combinations = cachedCombs
	} else {
		combinations = c.computeCounterCombinations(len(c.state) - 1)
		counterCombinationsCache[strState] = combinations
	}
	return combinations
}

func (c *counter) computeCounterCombinations(index int) []Combination {
	combinations := make([]Combination, 0)

	if index != -1 {
		lowerIndexCombs := c.computeCounterCombinations(index - 1)
		if len(lowerIndexCombs) == 0 {
			combinations = c.startCombinationsCreation(index)
		} else {
			combinations = c.updateLowerIndexCombinations(index, lowerIndexCombs)
		}
	}

	return combinations
}

func (c *counter) startCombinationsCreation(index int) []Combination {
	combinations := make([]Combination, 0)
	binomialCombs := getBinomialCombinations(c.state[index], c.initialValue)
	for _, binomialComb := range binomialCombs {
		combination := make(Combination, c.initialValue)
		util.InitSliceInt(combination, c.combinationUnsetIndex)

		for _, v := range binomialComb {
			combination[v] = index
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func (c *counter) updateLowerIndexCombinations(index int, lowerIndexCombs []Combination) []Combination {
	var combinations []Combination
	binomialCombs := getBinomialCombinations(c.state[index], c.initialValue)
	if len(binomialCombs) == 0 {
		combinations = lowerIndexCombs
	} else {
		combinations = make([]Combination, 0)
		for _, binomialComb := range binomialCombs {
			for _, lowerIndexComb := range lowerIndexCombs {
				// check if values in lowerIndexComb not already taken
				anySpotTaken := false
				for _, v := range binomialComb {
					if lowerIndexComb[v] != c.combinationUnsetIndex {
						anySpotTaken = true
						break
					}
				}
				if !anySpotTaken {
					// combine lower index combination with binomial combination
					// eg: [7, 7, 7, 0, 0, 7] + [1, 2, 5] -> [7, 1, 1, 0, 0, 1]
					combination := make(Combination, c.initialValue)
					copy(combination, lowerIndexComb)
					for _, v := range binomialComb {
						combination[v] = index
					}
					combinations = append(combinations, combination)
				}
			}
		}
	}

	return combinations
}

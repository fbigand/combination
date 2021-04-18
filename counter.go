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
	initialValue uint
	state        []uint
	endState     []uint
}

// counterCombsCache record counter combinations that has already been computed
// by the method GetCombinations
var counterCombinationsCache map[string][]Combination = make(map[string][]Combination)

// newCounter creates and returns a new counter.
// See CombinationIterator documentation for more information.
func newCounter(initialValue uint, endState []uint) *counter {
	c := counter{}

	c.initialValue = initialValue

	c.state = make([]uint, 2)
	c.state[0] = initialValue

	if util.MaxUint(endState...) > initialValue {
		panic("newCounter: endstate cannot be reached with these parameters")
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
	if util.EqualsSliceUint(c.state, c.endState) {
		return false
	}

	if c.state[0] == 0 {
		iReset := 1
		for c.state[iReset] == 0 {
			iReset++
		}
		c.state[iReset] = 0

		// prevent out of range
		if iReset+1 == len(c.state) {
			c.state = c.state[:len(c.state)+1]
		}
		c.state[iReset+1]++

		c.state[0] = c.initialValue - util.SumUint(c.state...)
	} else {
		c.state[0]--
		c.state[1]++
	}

	return true
}

func (c *counter) getCombinations() []Combination {
	var combinations []Combination

	strState := c.getStateToString()
	cachedCombs, combsInCache := counterCombinationsCache[strState]
	if combsInCache {
		combinations = cachedCombs
	} else {
		combinations = getCounterCombinations(c, len(c.state)-1)
		counterCombinationsCache[strState] = combinations
	}
	return combinations
}

func getCounterCombinations(c *counter, index int) []Combination {

	binomialCombs := getBinomialCombinations(c.state[index], c.initialValue)
	combinations := make([]Combination, 0)

	if index == 0 {
		// recursive stop case
		for _, binomialComb := range binomialCombs {
			combination := make(Combination, c.initialValue)
			// set default value (outside of c.state values)
			util.InitSliceInt(combination, int(c.initialValue)+1)

			for _, v := range binomialComb {
				// set actual values
				combination[v] = index
			}
			combinations = append(combinations, combination)
		}
	} else {
		lowerIndexCombs := getCounterCombinations(c, index-1)
		if c.state[index] == 0 {
			// pass to upper index
			combinations = lowerIndexCombs
		} else {
			for _, binomialComb := range binomialCombs {
				for _, lowerIndexComb := range lowerIndexCombs {
					// check if values in lowerIndexComb not already taken
					anySpotTaken := false
					for _, v := range binomialComb {
						if lowerIndexComb[v] != int(c.initialValue+1) {
							anySpotTaken = true
							break
						}
					}
					if !anySpotTaken {
						// combine lower inder combination with binomial combination
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
	}
	return combinations
}

package combination

// binomialCoefficient represents a binomial coefficient (K and N)
type binomialCoefficient struct {
	k uint
	n uint
}

// binomialCombsCache is the list of combinations of k elements from n.
// It serves as cache and record combinations that has already been computed
// by getBinomialCombinations
//
// Example:
//    binomialCombsCache[{2,4}] = [[0 1] [0 2] [0 3] [1 2] [1 3] [2 3]]
var binomialCombsCache map[binomialCoefficient][][]uint = make(map[binomialCoefficient][][]uint)

// getBinomialCombinations returns all unique and unordered combinations
// of k element in the list [0..n-1]
func getBinomialCombinations(k, n uint) [][]uint {
	var binomialCombinations [][]uint

	binomialCoefficient := binomialCoefficient{k: k, n: n}
	cachedCombs, combsInCache := binomialCombsCache[binomialCoefficient]
	if combsInCache {
		binomialCombinations = cachedCombs
	} else {
		binomialCombinations = make([][]uint, 0)
		if n >= k && k != 0 && n != 0 {
			path := make([]uint, 0)
			dfs(k, n, 1, path, &binomialCombinations)
		}
		binomialCombsCache[binomialCoefficient] = binomialCombinations
	}

	return binomialCombinations
}

// Depth-First Search algorithm
func dfs(k, n, begin uint, path []uint, res *[][]uint) {
	if uint(len(path)) == k {
		var copiedPath []uint = make([]uint, len(path))
		copy(copiedPath, path)
		*res = append(*res, copiedPath)
	} else {
		for i := begin; i <= n; i++ {
			path = append(path, i-1)
			dfs(k, n, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

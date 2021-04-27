# Combination

## Installation

`go get -u github.com/fbigand/combination`

## Purpose

I created this package for a personal project where I needed to combine things in a smart order.

Let's say you have multiple 3 lists of items ordered by pertinence, and you want to have combinations of these three list of items. But you want to test first the items with the most pertinence, and decrease gradually the pertinence of the items in the combination.

If you have these three lists
| ListA | ListB | ListC |
|-------|-------|-------|
|  A1   |  B1   |  C1   |
|  A2   |  B2   |  C2   |
|  A3   |  B3   |
|       |  B4   |

You want to test the items in this order:
- A1-B1-C1
- A2-B1-C1
- A2-B2-C1
- A2-B1-C2
- A1-B2-C2
- A3-B1-C1
- etc.

This package provides in order the indexes of the elements to combine through an iterator. The only parameter you need to give when creating the iterator, is the length of your lists.

## How to use it

```golang
package main

import (
    "fmt"

    "github.com/fbigand/combination"
)

func main() {
    // /!\ Lists must not be empty
    list1 := []string{"one", "two", "three"}
    list2 := []uint{10, 32}
    list3 := []int{5}

    listsLength := make([]int, 3)
    listsLength[0] = len(list1)
    listsLength[1] = len(list2)
    listsLength[2] = len(list3)

    combIt := combination.NewCombinationIterator(listsLength)
    
    var comb combination.Combination // alias for []int
    hasNext := true
    for hasNext {
        comb = combIt.GetCombination()
        fmt.Printf("[%v %v %v]\n", list1[comb[0]], list2[comb[1]], list3[comb[2]])

        hasNext = combIt.Next()
    }

    // Print output
    //
    // ["one" 10 5]
    // ["two" 10 5]
    // ["one" 32 5]
    // ["two" 32 5]
    // ["three" 10 5]
    // ["three" 32 5]
}
```
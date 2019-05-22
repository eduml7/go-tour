/*
1. Implement the Walk function.
2. Test the Walk function.
The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
4. Test the Same function.
Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
The documentation for Tree can be found here. 
*/

package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walki(t *tree.Tree, ch chan int) {
	if t != nil {

		Walki(t.Left, ch)
		ch <- t.Value
		Walki(t.Right, ch)
	}
}
func Walk(t *tree.Tree, ch chan int) {

	Walki(t, ch)
	close(ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	bool := true
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {

		if v != <-ch2 {

			bool = false
		}
	}
	return bool
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

/*
1
2
3
4
5
6
7
8
9
10
true
false

*/

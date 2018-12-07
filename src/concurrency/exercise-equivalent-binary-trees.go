package main

import (
	//"fmt"
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	inorderTraverseEntry(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		if ok1 && ok2 {
			if val1 != val2 {
				return false
			}
		} else if ok1 && !ok2 || !ok1 && ok2 {
			return false
		} else {
			return true
		}
	}
}

//func checkBinaryTreesEquivalence(root1 *ds.TreeNode, root2 *ds.TreeNode) bool {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go inorderTraverseEntry(root1, ch1)
//	go inorderTraverseEntry(root2, ch2)
//	for {
//		val1, ok1 := <-ch1
//		val2, ok2 := <-ch2
//		if ok1 && ok2 {
//			if val1 != val2 {
//				return false
//			}
//		} else if ok1 && !ok2 || !ok1 && ok2 {
//			return false
//		} else {
//			return true
//		}
//	}
//}

func inorderTraverseEntry(root *tree.Tree, ch chan int) {
	inorderTraverse(root, ch)
	close(ch)
}

func inorderTraverse(root *tree.Tree, ch chan int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraverse(root.Left, ch)
	}
	ch <- root.Value
	if root.Right != nil {
		inorderTraverse(root.Right, ch)
	}
}

//func testCheckBinaryTreesEquivalence(root1 *ds.TreeNode, root2 *ds.TreeNode) {
//	res := "Trees are "
//	if same := checkBinaryTreesEquivalence(root1, root2); same {
//		res += "the same!"
//	} else {
//		res += "not the same!"
//	}
//	fmt.Printf("%s\n", res)
//}

func main() {
	//r1 := &ds.TreeNode{Val: 3}
	//r11 := &ds.TreeNode{Val: 1}
	//r111 := &ds.TreeNode{Val: 1}
	//r18 := &ds.TreeNode{Val: 8}
	//r12 := &ds.TreeNode{Val: 2}
	//r15 := &ds.TreeNode{Val: 5}
	//r112 := &ds.TreeNode{Val: 12}
	//r113 := &ds.TreeNode{Val: 13}
	//
	//r1.Left = r11
	//r1.Right = r18
	//r11.Left = r111
	//r11.Right = r12
	//r18.Left = r15
	//r18.Right = r113
	//r15.Left = r112
	//
	//r2 := &ds.TreeNode{Val: 8}
	//r21 := &ds.TreeNode{Val: 1}
	//r211 := &ds.TreeNode{Val: 1}
	//r23 := &ds.TreeNode{Val: 3}
	//r22 := &ds.TreeNode{Val: 2}
	//r25 := &ds.TreeNode{Val: 5}
	//r213 := &ds.TreeNode{Val: 13}
	//r212 := &ds.TreeNode{Val: 12}
	//
	//r2.Left = r23
	//r2.Right = r213
	//r23.Left = r21
	//r23.Right = r25
	//r21.Left = r211
	//r21.Right = r22
	//r25.Left = r212
	//
	//testCheckBinaryTreesEquivalence(r1, r2)

	fmt.Printf("%t\n", Same(tree.New(3), tree.New(3)))
	fmt.Printf("%t\n", Same(tree.New(1), tree.New(2)))
}

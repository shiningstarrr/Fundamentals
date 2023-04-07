/*
Binary Search Tree implementation.
Fast insertion, removal, and lookup of items while offering an efficient way to iterate them in sorted order.
*/
package main

import "fmt"

// Define TreeNode struct with Key, Left and Right
type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Description - Insert a key to a TreeNode
 * @Input() - key -> integer
 * TC - O(logn) | SC - O(1)
 */
func (node *TreeNode) Insert(key int) {
	if node.Key < key {
		if node.Right == nil {
			node.Right = &TreeNode{Key: key}
		} else {
			node.Right.Insert(key)
		}
	} else if node.Key > key {
		if node.Left == nil {
			node.Left = &TreeNode{Key: key}
		} else {
			node.Left.Insert(key)
		}
	}
}

/*
 * Description - Search a key in TreeNode
 * @Input() - key -> Integer
 * @Output() -> Boolean
 * TC - O(logn) | SC - O(1)
 */
func (node *TreeNode) Search(key int) bool {
	if node == nil {
		return false
	}
	if node.Key < key {
		return node.Right.Search(key)
	} else if node.Key > key {
		return node.Left.Search(key)
	}
	return true
}

func main() {
	tree := &TreeNode{Key: 100}
	tree.Insert(50)
	tree.Insert(300)
	tree.Insert(30)
	tree.Insert(88)
	tree.Insert(231)
	fmt.Println("Tree:", *tree)  //{100 0xc00010a000 0xc00010a018}
	fmt.Println(tree.Search(77)) //false
	fmt.Println(tree.Search(88)) //true
}

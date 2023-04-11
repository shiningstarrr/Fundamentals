/* Tree implementation
 */
package fundamentals

import "fmt"

//Define a TreeNode struct with value, left and right
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

/*
 * Description - preOrderTraversal() - to add to root first,
 * then left and right nodes
 * @Input - node -> *TreeNode
 * TC - O(n) | SC - O(1)
 */
func preOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.value, " ")
	preOrderTraversal(node.left)
	preOrderTraversal(node.right)
}

/*
 * Description - postOrderTraversal() - to add to left and right nodes first, then root node
 * @Input - node -> *TreeNode
 * TC - O(n) | SC - O(1)
 */
func postOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	preOrderTraversal(node.left)
	preOrderTraversal(node.right)
	fmt.Println(node.value, " ")
}

/*
 * Description - postOrderTraversal() - to add left node, then root node and at last right node
 * @Input - node -> *TreeNode
 * TC - O(n) | SC - O(1)
 */
func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	preOrderTraversal(node.left)
	fmt.Println(node.value, " ")
	preOrderTraversal(node.right)
}

// func main() {
// 	root := &TreeNode{value: 1}

// 	// Add child nodes to the root node
// 	root.left = &TreeNode{value: 2}
// 	root.right = &TreeNode{value: 3}

// 	// Add child nodes to the left child node
// 	root.left.left = &TreeNode{value: 4}
// 	root.left.right = &TreeNode{value: 5}

// 	// Add child nodes to the right child node
// 	root.right.left = &TreeNode{value: 6}
// 	root.right.right = &TreeNode{value: 7}

// 	// Traverse the binary tree
// 	fmt.Println("Pre-Order Traversal:")
// 	preOrderTraversal(root)
// 	// Pre-Order Traversal: 1 2 4 5 3 6 7

// 	fmt.Println("\nIn-Order Traversal:")
// 	inOrderTraversal(root)
// 	// In-Order Traversal: 2 4 5 1 3 6 7

// 	fmt.Println("\nPost-Order Traversal:")
// 	postOrderTraversal(root)
// 	// Post-Order Traversal: 2 4 5 3 6 7 1
// }

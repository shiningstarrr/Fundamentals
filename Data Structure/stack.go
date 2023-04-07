/*
Stack implementation
*/

package fundamentals

type Stack struct {
	items []int
}

/*
 * Description - Push() function to push values to the stack.
 */
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

/*
 * Description - Pop() function to pop/remove values from stack.
 */
func (s *Stack) pop() int {
	last := len(s.items) - 1
	removeAt := s.items[last]
	s.items = s.items[:last]
	return removeAt
}

/*
*Description - isEmpty() function determines whether the stack is empty or not.
 */
func isEmpty(s Stack) bool {
	return len(s.items) == 0
}

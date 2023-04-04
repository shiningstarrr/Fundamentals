/*
Arrays&Slice implementation
*/
package fundamentals

import "fmt"

/*
 * Description: Initializing or array creation
 * @Output - array of integers
 * TC - O(1) | SC - O(1)
 */
func createArray() []int {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Array created: %d\n", arr)
	return arr
}

/*
 * Description: Reading elements in array
 * @Input - array of integers
 * TC - O(n) | SC - O(1)
 */
func readArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Reading elements from array: %d\n", arr[i])
	}
}

/*
 * Description: Updating elements in array
 * @Input - array of integers
 * TC - O(1) | SC - O(1)
 * Note: If we are appending on range/for loop,
 * then space complexity becomes O(N)
 */
func updateArray(arr []int, index, updateNum int) {
	arr[index] = updateNum
}

/*
 * Description: Deleting an array
 * @Input - array of integers
 * @Output - array of integers
 * TC - O(1) | SC - O(1)
 */
func deleteArray(arr []int) []int {
	arr = []int{}
	fmt.Printf("Array is deleted from the memory: %d\n", arr)
	return arr
}

/*
 * Description: Deleting an elements by index in an array
   * Basically, we are reslicing after removing an element by index
   * To achieve the correct form of elements in an array, we need to
   * reslice the elements before the index and with the appending slice
   * we need to return all the resliced elements after the index in an array
   * for all this we need to append to the same array, thus this will do this
   * all by callbyreference to update the original array
 * @Input - slice -> array of integers, index -> integer
 * @Output - array of integers
 * TC - O(1) | SC - O(1)
*/
func removeElemByIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

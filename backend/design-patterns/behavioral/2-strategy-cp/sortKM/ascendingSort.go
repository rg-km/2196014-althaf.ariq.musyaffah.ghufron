package sortKM

import "fmt"

// TODO: answer here
type AscSort interface {
	Sort([]int)
}

//concrete strategy implementation
type AscendingSort struct{}

func (as *AscendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here

	//merge sort ascending
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)

}

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func UnThreadedMergeSorted(left []int, right []int) []int{
	var result []int
	for (len(left) > 0) && (len(right) > 0){
		if(left[0] > right[0]){
			result = append(result, right[0])
			right = right[1:len(right)]
		}else{
			result = append(result,left[0])
			left = left[1:len(left)]
		}
	}

	if(len(left) > 0){
		result = append(result, left...)
	}
	if(len(right) > 0){
		result = append(result, right...)
	}

	return result
}

func UnThreadedMergeSort(unsorted []int) []int{
	var result []int
	if(len(unsorted) <= 1){
		return unsorted
	}else{
		left := UnThreadedMergeSort(unsorted[:len(unsorted)/2])
		right := UnThreadedMergeSort(unsorted[len(unsorted)/2:])

		result = UnThreadedMergeSorted(left, right)
	}
	return result
}

func MergeSorted(left []int, right []int, ch chan []int){
	var result []int
	for (len(left) > 0) && (len(right) > 0){
		if(left[0] > right[0]){
			result = append(result, right[0])
			right = right[1:len(right)]
		}else{
			result = append(result,left[0])
			left = left[1:len(left)]
		}
	}

	if(len(left) > 0){
		result = append(result, left...)
	}
	if(len(right) > 0){
		result = append(result, right...)
	}

	ch <- result
}

func MergeSort(unsorted []int,ch chan []int) []int{
	if(len(unsorted) <= 1){
		return unsorted
	}else{
		left := MergeSort(unsorted[:len(unsorted)/2], ch)
		right := MergeSort(unsorted[len(unsorted)/2:], ch)
		
		go MergeSorted(left, right, ch)
	}
	return <- ch
}


func main() {

	var A []int
	for len(A) < 1000000{
		A = append(A, rand.Intn(10000000))
	}

	timer := time.Now()
	ch := make(chan []int)

	B := A

	fmt.Println("LET THE RACE BEGIN\n","+-+-+-+-+-+-+-+-+-+")

	MergeSort(A,ch)

	meTime := time.Since(timer)
	fmt.Println("Threaded Time to Sort: ",meTime)

	timer = time.Now()

	B = UnThreadedMergeSort(B)
	stupidGoLangTime := time.Since(timer)
	fmt.Println("Unthreaded Time to Sort: ",stupidGoLangTime)

	if meTime < stupidGoLangTime {
		fmt.Println("Finally I've won!")
	}else{
		fmt.Println("Whatever, I don't care. Nobody won.")
	}

}
package main

func main() {

	arr := []int{6, 7, 1, 0, 9, 8, 12}
	temparr := [len(arr)]int{}

	MergeSort(arr, 0, arr[len(arr)-1], temparr[:])

}

func MergeSort(arr []int, start int, end int, temparr []int) {
	leftEnd := len(arr) / 2
	rightStart := leftEnd + 1
	MergeSort(arr, arr[0], arr[leftEnd], temparr)
	MergeSort(arr, arr[rightStart], arr[len(arr)-1], temparr)
	Merge(arr[0:leftEnd], arr[rightStart:len(arr)-1])

}

func Merge(leftArr []int, rightArr []int) {

	for i:=0; i<= 
	for 
}

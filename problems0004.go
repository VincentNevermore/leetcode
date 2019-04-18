package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := sortedArray(nums1, nums2)
	return median(res)
}

func sortedArray(array1, array2 []int) []int {
	len1, i := len(array1), 0
	len2, j := len(array2), 0

	res := make([]int, len1+len2)
	for k := 0; k <=len1+len2-1; k++ {
		if i == len1 ||
			(i < len1 && j < len2 && array1[i] > array2[j]) {
			res[k] = array2[j]
			j++
			continue
		}
		if j == len2 ||
			(i < len1 && j < len2 && array1[i] <= array2[j]) {
			res[k] = array1[i]
			i++
		}

	}
	return res
}

func median(res []int) float64 {
	length := len(res)
	if length == 0 {
		panic("Error Input! Please restart and check")
	} else if length/2 == 0 {
		return float64(res[length/2] + res[length/2+1]) /2.0
	} else {
		return float64(res[length/2])
	}

}

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3})
}

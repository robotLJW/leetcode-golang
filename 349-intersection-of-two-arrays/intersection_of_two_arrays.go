package _49_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	hashMap := make(map[int]bool)
	hashMapTwo := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		if _, ok := hashMap[nums1[i]]; ok {
		} else {
			hashMap[nums1[i]] = true
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := hashMap[nums2[i]]; ok {
			if _, isExist := hashMapTwo[nums2[i]]; isExist {
			} else {
				ans = append(ans, nums2[i])
				hashMapTwo[nums2[i]] = true
			}
		}
	}
	return ans
}

package _41_valid_mountain_array

func validMountainArray(A []int) bool {
	if len(A) < 2 {
		return false
	}
	flag, upFlag := true, false
	for i := 0; i < len(A); i++ {
		if i+1 >= len(A) {
			break
		}
		if A[i+1] == A[i] {
			return false
		}
		if A[i+1] < A[i] {
			if flag == true {
				flag = false
			}
			continue
		}
		if A[i+1] > A[i] {
			if flag == false {
				return false
			}
			if upFlag == false {
				upFlag = true
			}
		}
	}
	if upFlag == true && flag == false {
		return true
	}
	return false
}

func validMountainArrayTwo(A []int) bool {
	var i int
	for ; i+1 < len(A) && A[i+1] > A[i]; i++ {
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for ; i+1 < len(A) && A[i+1] < A[i]; i++ {
	}
	return i == len(A)-1
}

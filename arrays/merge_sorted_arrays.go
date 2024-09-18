package arrays

func merge(nums1 []int, m int, nums2 []int, n int)  {
    combined_array := make([]int, 0, m+n)
    i, j := 0, 0
    for i < m && j < n {
        if nums1[i] <= nums2[j] {
            combined_array = append(combined_array, nums1[i])
            i++
        } else {
            combined_array = append(combined_array, nums2[j])
            j++
        }
    }

    for i < m {
        combined_array = append(combined_array, nums1[i])
        i++
    }

    for j < n {
        combined_array = append(combined_array, nums2[j])
        j++
    }
    copy(nums1, combined_array)
}

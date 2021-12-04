package helper

// MapKeys - get keys of map data as a Array
// in php,the keys you want always is string or number
// here,let it be string
func MapKeys(data map[string]interface{}) []string {
	if len(data) < 1 {
		return []string{}
	}
	var resData []string
	for index := range data {
		resData = append(resData, index)
	}
	return resData
}

// ListIntersect — Computes the intersection of lists
func ListIntersect(nums1, nums2 []int) []int {

	res := make([]int, 0, len(nums1))
	nc := make(map[int]int)

	for _, n := range nums1 {
		nc[n]++
	}

	for _, n := range nums2 {
		if nc[n] > 0 {
			res = append(res, n)
			nc[n]--
		}
	}

	return res
}

// ListMerge — Merge one or more lists
func ListMerge(arr ...[]interface{}) []interface{} {
	s := make([]interface{}, 0)
	for _, v := range arr {
		s = append(s, v...)
	}

	return s
}

// ListReverse - Return a list with elements in reverse order
func ListReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
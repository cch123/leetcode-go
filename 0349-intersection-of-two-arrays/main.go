package main

func intersection(nums1 []int, nums2 []int) []int {
    res := []int{}
    var valMap = map[int]bool{}
    for i:=0;i<len(nums1);i++{
        valMap[nums1[i]]= true
    }

    for i:=0;i<len(nums2);i++{
        if valMap[nums2[i]] == true {
            res = append(res, nums2[i])
            valMap[nums2[i]] = false
        }
    }
    return res
}

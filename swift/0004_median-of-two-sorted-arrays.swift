class Solution {
  func findMedianSortedArrays(_ nums1: [Int], _ nums2: [Int]) -> Double {
    let n = nums1.count, m = nums2.count
    if n > m {
      return findMedianSortedArrays(nums2, nums1)
    }
    var l = 0, r = n, k = (n + m + 1) / 2
    var x1, x2, y1, y2: Int
    while l < r {
      var mid = (l + r) / 2
      x1 = get(nums1, index: mid - 1)
      x2 = get(nums1, index: mid)
      y1 = get(nums2, index: k - mid - 1)
      y2 = get(nums2, index: k - mid)
      if x1 <= y2 && y1 <= x2 {
        l = mid
        break
      }
      if x1 > y2 {
        r = mid
      } else {
        l = mid + 1
      }
    }
    x1 = get(nums1, index: l - 1)
    x2 = get(nums1, index: l)
    y1 = get(nums2, index: k - l - 1)
    y2 = get(nums2, index: k - l)
    if (n + m) % 2 == 1 {
      return Double(max(x1, y1))
    }
    return Double(max(x1, y1) + min(x2, y2)) / 2.0
  }

  func get(_ arr: [Int], index i: Int) -> Int {
    let n = arr.count
    if i < 0 {
      return Int.min
    }
    return i < n ? arr[i] : Int.max
  }
}

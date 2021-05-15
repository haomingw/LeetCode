class Solution {
  func lengthOfLongestSubstring(_ s: String) -> Int {
    var m = [Character: Int]()
    var res = 0, j = 0
    for (i, c) in s.enumerated() {
      if let prev = m[c] {
        j = max(j, prev + 1)
      }
      res = max(res, i - j + 1)
      m[c] = i
    }
    return res
  }
}

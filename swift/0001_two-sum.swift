class Solution {
  func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
    var m = [Int: Int]()
    for (i, v) in nums.enumerated() {
      if let idx = m[target - v] {
        return [idx, i]
      }
      m[v] = i
    }
    return [-1, -1]
  }
}

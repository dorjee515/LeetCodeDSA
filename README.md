# LeetCode DSA (Go)

A collection of LeetCode and algorithm practice solutions implemented in Go.

**Overview**
- This folder contains small Go programs and algorithm implementations used for practice and learning.
- Each file implements one or more algorithmic problems or helpers; `main.go` runs example cases.

**Files**
- `main.go`: Example runner that calls the implemented solutions. See [main.go](LeetCode DSA/main.go).
- `TwoSum.go`: Implementation of the Two Sum problem (`twoSum`). See [TwoSum.go](LeetCode DSA/TwoSum.go). https://leetcode.com/problems/two-sum/description/
- `BestTimetoBuyandSellStock.go`: Single-transaction max profit function (`maxProfit`). See [BestTimetoBuyandSellStock.go](LeetCode DSA/BestTimetoBuyandSellStock.go). https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/4868897/most-optimized-kadanes-algorithm-java-c-2yt85/
- `KadaneAlgo.go`: Kadane's algorithm implementation (`KadaneMaxSum`). See [KadaneAlgo.go](LeetCode DSA/KadaneAlgo.go). https://leetcode.com/problems/maximum-subarray/
- `ContainsDuplicate.go`: Contains Duplicate problem implementation (`containsDuplicate`). See [ContainsDuplicate.go](LeetCode DSA/ContainsDuplicate.go). https://leetcode.com/problems/contains-duplicate/
- `ProductofArrayExceptSelf.go`: Product of Array Except Self implementation (`productExceptSelf`). See [ProductofArrayExceptSelf.go](LeetCode DSA/ProductofArrayExceptSelf.go). https://leetcode.com/problems/product-of-array-except-self/
- `go.mod`: Module definition for the package.

**How to run**
Ensure Go is installed and `GOPATH`/`GOMOD` are set up. From the repository root run:

```powershell
cd D:\Golang\LeetCode DSA
go run .
```

Or run all Go files explicitly:

```powershell
go run *.go
```

Notes:
- Running `go run main.go` may fail with `undefined` symbol errors because it only builds that single file. Use `go run .` or `go run *.go` to include all files.
- These are small practice examples — feel free to refactor into packages, add tests, or improve algorithms.

**Next steps / Suggestions**
- Add unit tests (`*_test.go`) for each algorithm.
- Add a short description for each algorithm and link to the corresponding LeetCode problem.

---
**Recent Updates (2026-05-22):**
- Added `ContainsDuplicate.go` - Detects duplicate elements in an array
- Added `ProductofArrayExceptSelf.go` - Calculates product of array elements except self
- Linked LeetCode problems for all algorithm implementations

Generated on: 2026-05-22

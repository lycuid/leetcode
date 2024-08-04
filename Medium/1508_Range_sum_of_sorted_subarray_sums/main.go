// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
package main

type Stream struct {
	items      []int
	index, sum int
}

func (this Stream) Done() bool     { return this.index >= len(this.items) }
func (this *Stream) GetValue() int { return this.sum + this.items[this.index] }
func (this *Stream) Advance()      { this.sum, this.index = this.sum+this.items[this.index], this.index+1 }

func rangeSum(nums []int, n int, left int, right int) (ret int) {
	const MOD = 1e9 + 7
	streams := make([]*Stream, n)
	for i := range streams {
		streams[i] = &Stream{items: nums, index: i}
	}
	for i := 1; i <= right; i++ {
		min_stream := streams[0]
		for _, stream := range streams {
			if min_stream.Done() || (!stream.Done() && stream.GetValue() < min_stream.GetValue()) {
				min_stream = stream
			}
		}
		if i >= left {
			ret += min_stream.GetValue()
			ret %= MOD
		}
		min_stream.Advance()
	}
	return ret % MOD
}

func main() {}

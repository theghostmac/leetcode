package dayone

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 5},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, testData := range tests {
		got := twoSum(testData.nums, testData.target)
		if !reflect.DeepEqual(got, testData.want) {
			t.Errorf("TwoSum(%v, %d) got %v, want %v", testData.nums, testData.target, got, testData.want)
		}
	}
}

func TestTwoSumWithHashMap(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 5},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, testData := range tests {
		got := twoSumWithHashmap(testData.nums, testData.target)
		if !reflect.DeepEqual(got, testData.want) {
			t.Errorf("TwoSum(%v, %d) got %v, want %v", testData.nums, testData.target, got, testData.want)
		}
	}
}

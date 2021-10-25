package main

import (
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"A", "A"}, []string{"A"}},
		{"2", args{"A", "B"}, []string{""}},
		{"3", args{"BCD", "ACBAD"}, []string{"BD", "CD"}},
		{"4", args{"ABCD", "ACBAD"}, []string{"ABD", "ACD"}},
		{"5", args{"ABCDEFG", "CBADFEG"}, []string{"BDEG", "BDFG", "ADEG", "ADFG", "CDEG", "CDFG"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.s1, tt.args.s2); !equal(got, tt.want) {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	m := map[string]int{}
	for _, v := range arr1 {
		m[v]++
	}
	for _, v := range arr2 {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

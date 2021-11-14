package _77_map_sum_pairs

import "strings"

type MapSum struct {
	KV map[string]int
}

func Constructor() MapSum {
	return MapSum{
		make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.KV[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	ans := 0
	for key, value := range this.KV {
		if strings.HasPrefix(key, prefix) {
			ans = ans + value
		}
	}
	return ans
}

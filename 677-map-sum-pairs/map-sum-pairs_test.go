package _77_map_sum_pairs

import "testing"

func TestMapSumPairs(t *testing.T) {
	t.Run("input: [\"MapSum\", \"insert\", \"sum\", \"insert\", \"sum\"]\n[[], [\"apple\", 3], [\"ap\"], [\"app\", 2], [\"ap\"]]\n,"+
		"output: [null, null, 3, null, 5]", func(t *testing.T) {
		mapSum := Constructor()
		mapSum.Insert("apple", 3)
		ans := mapSum.Sum("app")
		if ans != 3 {
			t.Error("wrong answer, should return 3")
		}
		mapSum.Insert("app", 2)
		ans = mapSum.Sum("app")
		if ans != 5 {
			t.Error("wrong answer, should return 5")
		}
	})
}

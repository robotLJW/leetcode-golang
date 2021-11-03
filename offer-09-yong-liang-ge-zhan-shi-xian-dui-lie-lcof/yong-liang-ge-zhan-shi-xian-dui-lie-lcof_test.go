package offer_09_yong_liang_ge_zhan_shi_xian_dui_lie_lcof

import "testing"

func TestQueue(t *testing.T) {
	t.Run("input: [\"CQueue\",\"appendTail\",\"deleteHead\",\"deleteHead\"]\n[[],[3],[],[]],output: [null,null,3,-1]", func(t *testing.T) {
		queue := Constructor()
		queue.AppendTail(3)
		value := queue.DeleteHead()
		if value != 3 {
			t.Error("answer is wrong, should return 3")
		}
		value = queue.DeleteHead()
		if value != -1 {
			t.Error("answer is wrong, should return -1")
		}
	})

	t.Run("input: [\"CQueue\",\"deleteHead\",\"appendTail\",\"appendTail\",\"deleteHead\",\"deleteHead\"]\n[[],[],[5],[2],[],[]],output: [null,-1,null,null,5,2]", func(t *testing.T) {
		queue := Constructor()
		value := queue.DeleteHead()
		if value != -1 {
			t.Error("answer is wrong, should return -1")
		}
		queue.AppendTail(5)
		queue.AppendTail(2)
		value = queue.DeleteHead()
		if value != 5 {
			t.Error("answer is wrong, should return 5")
		}
		value = queue.DeleteHead()
		if value != 2 {
			t.Error("answer is wrong, should return 2")
		}
	})

}

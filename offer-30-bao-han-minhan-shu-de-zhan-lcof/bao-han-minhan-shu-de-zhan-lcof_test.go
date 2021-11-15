package offer_30_bao_han_minhan_shu_de_zhan_lcof

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("MinStack minStack = new MinStack();\n"+
		"minStack.push(-2);\n"+
		"minStack.push(0);\n"+
		"minStack.push(-3);\n"+
		"minStack.min();   --> 返回 -3.\n"+
		"minStack.pop();\n"+
		"minStack.top();      --> 返回 0.\n"+
		"minStack.min();   --> 返回 -2.", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		if minStack.Min() != -3 {
			t.Error("wrong answer, should return -3")
		}
		minStack.Pop()
		if minStack.Top() != 0 {
			t.Error("wrong answer, should return 0")
		}
		if minStack.Min() != -2 {
			t.Error("wrong answer, should return -2")
		}
	})
}

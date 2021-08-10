package leetcode

type Stack struct {
	items []Item
}
type Item struct {
	value int
	pos   int
}

// Push
func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

// Pop
func (s *Stack) Pop() Item {
	lenStack := len(s.items)
	if lenStack == 0 {
		return Item{
			pos: -1,
		}
	}
	toRemove := s.items[lenStack-1]
	s.items = s.items[:lenStack-1]
	return toRemove
}

func nextGreaterElements(nums []int) []int {

	myStack := Stack{}
	tempStack := Stack{}
	lenNums := len(nums)

	nums = append(nums, nums...)
	for i := 0; i < len(nums); i++ {
		for len(myStack.items) > 0 {
			item := myStack.Pop()
			if item.value < nums[i] {
				nums[item.pos] = nums[i]
			} else {
				tempStack.Push(item)
				break
			}
		}
		for len(tempStack.items) > 0 {
			item := tempStack.Pop()
			myStack.Push(item)
		}
		myStack.Push(Item{
			value: nums[i],
			pos:   i,
		})
		nums[i] = -1
	}

	return nums[:lenNums]
}

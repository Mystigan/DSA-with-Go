package leetcode

func mergeLists(l1, l2 *Node) *Node {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var prev, base, slave, temp *Node

	if l1.Val < l2.Val {
		slave = l1
		prev = l2
		base = l2.Next
	} else {
		slave = l2
		prev = l1
		base = l1.Next
	}

	head := prev
	if base == nil {
		prev.Next = slave
		return head
	}

	for base != nil && slave != nil {
		if slave.Val <= base.Val {
			temp = slave.Next
			prev.Next = slave
			slave.Next = base
			prev = slave
			slave = temp
		} else {
			prev = prev.Next
			base = base.Next
		}
	}

	if slave != nil {
		prev.Next = slave
	}
	return head
}

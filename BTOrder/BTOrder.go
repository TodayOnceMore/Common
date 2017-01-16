package BTNode

type BTNode struct {
	value  int
	lchild *BTNode
	rchild *BTNode
}

// 先序遍历
func PreOrderBT(root *BTNode) {
	if root == nil {
		return
	}

	p := root
	stack := make([]*BTNode, 0, 128)
	for (len(stack) > 0) || (p != nil) {
		for p != nil {
			fmt.Printf("%d ", p.value)
			stack = append(stack, p)
			p = p.lchild
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:(len(stack) - 1)]
			p = p.rchild
		}
	}
	fmt.Println()
}

// 中序遍历
func InOrderBT(root *BTNode) {
	if root == nil {
		return
	}

	p := root
	stack := make([]*BTNode, 0, 128)
	for (len(stack) > 0) || (p != nil) {
		for p != nil {
			stack = append(stack, p)
			p = p.lchild
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]
			fmt.Printf("%d ", p.value)
			stack = stack[:(len(stack) - 1)]
			p = p.rchild
		}
	}
	fmt.Println()
}

// 后序遍历
func PostOrderBT(root *BTNode) {
	if root == nil {
		return
	}

	p, pVisited := root, new(BTNode)
	stack := make([]*BTNode, 0, 128)
	for p != nil {
		stack = append(stack, p)
		p = p.lchild
	}

	for len(stack) > 0 {
		p = stack[len(stack)-1]
		stack = stack[:(len(stack) - 1)]
		if (p.rchild == nil) || (p.rchild == pVisited) {
			fmt.Printf("%d ", p.value)
			pVisited = p
		} else {
			stack = append(stack, p)
			p = p.rchild

			for p != nil {
				stack = append(stack, p)
				p = p.lchild
			}
		}
	}
	fmt.Println()
}

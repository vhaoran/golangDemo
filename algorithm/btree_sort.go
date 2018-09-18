package main

import (
	"fmt"
)

type Node struct {
	L    *Node
	R    *Node
	Data int
}

type ComparerBSort func(int, int) int

func (node *Node) Add(n *Node, c ComparerBSort) {
	if c == nil {
		panic("comparerBSort must be implement!")
	}
	i := c(n.Data, node.Data)
	if i < 0 { //
		//place left Tree
		if node.L != nil {
			node.L.Add(n, c)
		} else {
			node.L = n
		}
	} else {
		if node.R != nil {
			node.R.Add(n, c)
		} else {
			node.R = n
		}
	}
}

func print_all_asc(n *Node) {
	//中序遍历,依次从小到从打印L-root-L
	//中序遍历,依次从大到小打印R-root-L
	if n != nil {
		print_all_asc(n.L)
		fmt.Println("#### R ###", n.Data)
		print_all_asc(n.R)
	}
}

func print_all_desc(n *Node) {
	//中序遍历,依次从小到从打印L-root-L
	//中序遍历,依次从大到小打印R-root-L
	if n != nil {
		print_all_desc(n.R)
		fmt.Println("#### R ###", n.Data)
		print_all_asc(n.L)
	}
}

func print_shape(prefix string, n *Node) {
	//先序遍历,打印出来各个子树
	if n == nil {
		return
	}
	fmt.Println(prefix, "  ", n.Data)
	//
	print_shape(prefix+"*", n.L)
	print_shape(prefix+"*", n.R)
}

func R_Rotate(p *Node) *Node {
	if p.L != nil {
		L := p.L  //
		p.L = L.R //
		L.R = p   //
		return L
	}
	return nil
}

func deepth(i int, node *Node) int {
	if node == nil {
		return i
	}

	l := deepth(i+1, node.R)
	r := deepth(i+1, node.L)
	if l > r {
		return l
	}
	return r
}

func comp(a int, b int) int {
	return (a - b)
}

func main() {
	fmt.Println("--------------")
	lst := []int{100, 5, 200, 7, 20, 38, 4, 1, 5, 3, 4, 1, 8, 6}

	var first *Node
	for _, v := range lst {
		if first == nil {
			first = &Node{}
			first.Data = v
			continue
		}
		n := &Node{}
		n.Data = v
		first.Add(n, comp)
	}
	fmt.Println("deept L:", deepth(1, first.L), " R:", deepth(1, first.R))
	fmt.Println("init-------------")
	print_shape("", first)

	//---------------------------
	first = R_Rotate(first)
	fmt.Println(first.Data)
	fmt.Println("first rotate:-------------")
	print_shape("", first)

	fmt.Println("deept L:", deepth(1, first.L))
	fmt.Println("deept R:", deepth(1, first.R))

	fmt.Println("aaa")
	first = R_Rotate(first)
	fmt.Println("--------------------------")
	print_shape("", first)
	fmt.Println("deept L:", deepth(1, first.L))
	fmt.Println("deept R:", deepth(1, first.R))

	fmt.Println("-------------------2-------------")
	print_all_desc(first)
}

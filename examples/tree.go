package examples

import "github.com/badgerodon/goreify/generics"

type tree struct {
	root *treeNode
}

type treeNode struct {
	elem     generics.T1
	children []*treeNode
}

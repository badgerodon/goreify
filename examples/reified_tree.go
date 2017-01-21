package examples

type treeInt struct {
	root *treeNodeInt
}

type treeNodeInt struct {
	elem     int
	children []*treeNodeInt
}

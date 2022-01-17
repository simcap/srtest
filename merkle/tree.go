package merkle

type tree struct {
	height int
	root   *node
}

type node struct {
	hash        string
	left, right *node
}

func (t *tree) Height() int { return t.height }

func (t *tree) Root() string {
	if t != nil && t.root != nil {
		return t.root.hash
	}
	return ""
}

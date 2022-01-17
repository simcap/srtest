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
	if t.root != nil {
		return t.root.hash
	}
	return ""
}

func (t *tree) Level(index int) (hashes []string) {
	if index < 0 || index > t.height || t.root == nil {
		return
	}

	var level int
	nodes := []*node{t.root}

	for level < index {
		var tmp []*node
		for _, n := range nodes {
			if n.left != nil {
				tmp = append(tmp, n.left)
			}
			if n.right != nil {
				tmp = append(tmp, n.right)
			}
		}
		nodes = tmp
		level++
	}

	for _, n := range nodes {
		hashes = append(hashes, n.hash)
	}

	return
}

package merkle

import (
	"crypto/sha256"
	"fmt"
)

func BuildTree(data ...string) *tree {
	if len(data) == 0 {
		return new(tree)
	}

	var nodes []*node
	for _, d := range data {
		nodes = append(nodes, &node{hash: shasum(d)})
	}

	height := 1
	for len(nodes) != 1 {
		var reduced []*node
		for i := 0; i < len(nodes); i = i + 2 {
			var n *node
			if len(nodes) > i+1 {
				n = buildParent(nodes[i], nodes[i+1])
			} else {
				n = buildParent(nodes[i], nil)
			}
			reduced = append(reduced, n)
		}
		nodes = reduced
		height++
	}

	return &tree{
		height: height,
		root:   nodes[0],
	}
}

func buildParent(nodeL, nodeR *node) *node {
	if nodeR == nil {
		return &node{
			left: nodeL,
			hash: nodeL.hash,
		}
	}

	return &node{
		left:  nodeL,
		right: nodeR,
		hash:  shasum(fmt.Sprintf("%s%s", nodeL.hash, nodeR.hash)),
	}
}

func shasum(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

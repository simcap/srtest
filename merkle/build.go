package merkle

import (
	"crypto/sha256"
	"fmt"
)

func BuildTree(data ...string) *tree {
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
				n = buildParent(nodes[i].hash, nodes[i+1].hash)
			} else {
				n = buildParent(nodes[i].hash, "")
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

func buildParent(hashL, hashR string) *node {
	if hashR == "" {
		return &node{
			left: &node{hash: hashL},
			hash: hashL,
		}
	}

	return &node{
		left:  &node{hash: hashL},
		right: &node{hash: hashR},
		hash:  shasum(fmt.Sprintf("%s%s", hashL, hashR)),
	}
}

func shasum(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

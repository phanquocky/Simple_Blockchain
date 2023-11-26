package main

import (
	"bytes"
	"crypto/sha256"
)

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Data  []byte
	Left  *MerkleNode
	Right *MerkleNode
}

func NewMerkleNode(data []byte, left *MerkleNode, right *MerkleNode) *MerkleNode {
	return &MerkleNode{Data: data, Left: left, Right: right}
}

func NewMerkleTree(transactions []*Transaction) *MerkleTree {

	nodes := make([]*MerkleNode, 0)

	for _, transaction := range transactions {
		nodes = append(nodes, NewMerkleNode(transaction.Data, nil, nil))
	}

	for len(nodes) > 1 {
		var newLevel []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]

			var right *MerkleNode
			// If there are an odd number of nodes,
			// the node without a partner is hashed with a copy of itself.
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right = nodes[i]
			}

			hash := sha256.Sum256(append(left.Data, right.Data...))
			parent := NewMerkleNode(hash[:], left, right)
			newLevel = append(newLevel, parent)
		}
		nodes = newLevel
	}
	return &MerkleTree{Root: nodes[0]}
}

func traverse(node *MerkleNode, targetValue []byte) bool {
	if node == nil {
		return false
	}
	if bytes.Equal(node.Data, targetValue) {
		return true
	}
	return traverse(node.Left, targetValue) || traverse(node.Right, targetValue)
}

func (tree *MerkleTree) MerkleProof(targetValue []byte) bool {
	return traverse(tree.Root, targetValue)
}

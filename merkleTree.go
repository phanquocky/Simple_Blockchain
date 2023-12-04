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

	if len(nodes) == 1 {
		// If only have 1 transaction, duplicate itself for make safer root
		nodes = append(nodes, nodes[0])
	}

	for len(nodes) > 1 {
		var newLevel []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]

			var right *MerkleNode
			// If there are an odd number of nodes/leafs,
			// the nodes/leafs without a partner is hashed with a copy of itself.
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

func traverse(node *MerkleNode, targetValue []byte) (bool, []byte) {
	if node == nil {
		return false, make([]byte, 32)
	}

	if bytes.Equal(node.Data, targetValue) {
		return true, node.Data
	}
	left, leftHash := traverse(node.Left, targetValue)
	right, rightHash := traverse(node.Right, targetValue)

	if left == true || right == true {
		temp := sha256.Sum256(append(leftHash, rightHash...))
		return true, temp[:]
	}

	return false, node.Data

}

func MerkleVerify(tree *MerkleTree, data []byte) bool {
	found, root := traverse(tree.Root, data)
	return found == true && bytes.Equal(root, tree.Root.Data)
}

package mpt

import (
	"bytes"
	"errors"

	"github.com/joeqian10/neo-gogogo/helper"
)

//Trie mpt tree
type Trie struct {
	db   *trieDb
	root node
}

//NewTrie new a trie instance
func NewTrie(root []byte, db IKVReadOnlyDb) (*Trie, error) {
	if db == nil {
		return nil, errors.New("failed initialize Trie, invalid db")
	}
	t := &Trie{
		db: newTrieDb(db),
	}
	r, err := t.resolve(hashNode(root))
	if err != nil {
		return nil, err
	}
	t.root = r
	return t, nil
}

func (t *Trie) resolve(hash hashNode) (node, error) {
	return t.db.node(hash)
}

//Get try get value
func (t *Trie) Get(path []byte) ([]byte, error) {
	path = helper.ToNibbles(path)
	vn, err := t.get(t.root, path)
	v, ok := vn.(valueNode)
	if !ok {
		return nil, err
	}
	return ([]byte)(v), err
}

func (t *Trie) get(n node, path []byte) (node, error) {
	switch n.(type) {
	case valueNode:
		if len(path) == 0 {
			return n, nil
		}
		return n, errors.New("trie cant find the path")
	case fullNode:
		f := n.(fullNode)
		if len(path) == 0 {
			return t.get(f.children[16], path)
		}
		return t.get(f.children[path[0]], path[1:])
	case shortNode:
		s := n.(shortNode)
		if !bytes.HasPrefix(path, s.key) {
			return nil, errors.New("trie cant find the path")
		}
		return t.get(s.next, bytes.TrimPrefix(path, s.key))
	case hashNode:
		nn, err := t.resolve(n.(hashNode))
		if err != nil {
			return nil, err
		}
		return t.get(nn, path)
	}
	return nil, errors.New("trie cant find the path")
}

//VerifyProof directly verify proof
func VerifyProof(root, path []byte, proof [][]byte) ([]byte, error) {
	proofdb := NewProofDb(proof)
	trie, err := NewTrie(root, proofdb)
	if err != nil {
		return nil, err
	}
	value, err := trie.Get(path)
	return value, err
}

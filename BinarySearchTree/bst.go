package main

import "fmt"

type Node struct {
        left *Node
        right *Node
        value int
}

type BST struct {
        root *Node;
}

func NewBST () * BST {
        return new(BST)
}

func (root *Node) insert (new_node *Node) {
        if new_node.value > root.value {
                if root.right == nil {
                        root.right = new_node
                } else {
                        root.right.insert(new_node)
                }
        } else if new_node.value < root.value{
                if root.left == nil {
                        root.left = new_node
                } else {
                        root.left.insert(new_node)
                }
        }
}

func (tree *BST) Insert(value int) {
        if tree.root == nil {
                tree.root = &Node{nil, nil, value}
        }
        tree.root.insert(&Node{nil, nil, value})
}

func minValue(n *Node) int {
        if n.left == nil {
                return n.value
        }
        return minValue(n.right)
}


func link(parent *Node, n *Node) {
        if parent.left == n {
                if n.left != nil {
                        parent.left = n.left
                } else {
                        parent.left = n.right
                }
        } else if parent.right == n {
                if n.left != nil {
                        parent.right = n.left
                } else {
                        parent.right = n.right
                }
        }
}

func del(n *Node, parent *Node, v int) bool {
        switch {
        case n.value == v:
                if n.left != nil && n.right != nil {
                        n.value = minValue(n.right)
                        return del(n.right, n, n.value)
                }
                link(parent, n); return true
        case n.value > v:
                if n.left == nil {
                        return false
                }
                return del(n.left, n, v)
        case n.value < v:
                if n.right == nil {
                        return false
                }
                return del(n.right, n, v)
        }
        return false
}

func (bst *BST) Delete(v int) bool {
        if bst.root == nil {
                return false
        }

        if bst.root.value == v {
                tempRoot := &Node{nil, nil, 0}
                tempRoot.left = bst.root
                r := del(bst.root, tempRoot, v)
                bst.root = tempRoot.left
                return r
        }
        return del(bst.root.left, bst.root, v) || del(bst.root.right, bst.root, v)
}

func showInOrder(root *Node) {
        if (root != nil) {
                showInOrder(root.left)
                fmt.Println(root.value)
                showInOrder(root.right)
        }
}

func showPreOrder(root *Node) {
        if (root != nil) {
                fmt.Println(root.value)
                showInOrder(root.left)
                showInOrder(root.right)
        }
}

func showPostOrder(root *Node) {
        if (root != nil) {
                fmt.Println(root.value)
                showInOrder(root.left)
                showInOrder(root.right)
        }
}

package main

import "fmt"

func main() {
        test := NewBST()
        test.Insert(4)
        test.Insert(2)
        test.Insert(5)
        test.Insert(15)
        test.Insert(25)
        test.Insert(1)
        test.Insert(3)

        fmt.Println("Inorder Traversal")
        showInOrder(test.root)
        fmt.Println("Preorder Traversal")
        showPreOrder(test.root)
        fmt.Println("Postorder Traversal")
        showPostOrder(test.root)

        fmt.Println("Removing value 15")
        test.Delete(15)
        showInOrder(test.root)
}

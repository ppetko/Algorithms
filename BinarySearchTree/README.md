# Data Structure - Binary Search Tree Implementation

## Sample main 

```
  1 package main
  2 
  3 func main() {
  4         test := NewBST()
  5         test.Insert(4)
  6         test.Insert(2)
  7         test.Insert(5)
  8         test.Insert(15)
  9         test.Insert(25)
 10         test.Insert(1)
 11         test.Insert(3)
 12 
 13         fmt.Println("Inorder Traversal")
 14         showInOrder(test.root)
 15         fmt.Println("Preorder Traversal")
 16         showPreOrder(test.root)
 17         fmt.Println("Postorder Traversal")
 18         showPostOrder(test.root)
 19 
 20         fmt.Println("Removing value 15")
 21         test.Delete(15)
 22         showInOrder(test.root)
 23 }
```
## Run the code

```
go vet main.go bst.go 
go build main.go bst.go 
./main 
Inorder Traversal
1
2
3
4
5
15
25
Preorder Traversal
4
1
2
3
5
15
25
Postorder Traversal
4
1
2
3
5
15
25
Removing value 15
1
2
3
4
5
25

```

## Running time in Big(O)

```
			Access	         Search		Insertion	Deletion	Access	Search	Insertion Deletion Worst	
Binary Search Tree	Θ(log(n))	Θ(log(n))	Θ(log(n))	Θ(log(n))	O(n)	O(n)	O(n)	  O(n)	    O(n)
```

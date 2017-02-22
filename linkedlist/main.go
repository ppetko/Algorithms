package main

func main(){
        list := &List{}
        list.add(5)
        list.add(10)
        list.add(15)
        list.add(20)
        list.add(25)
        list.add(30)
        list.print()
        list.delete(5)
        list.delete(30)
        list.print()
        list.search(10)
}

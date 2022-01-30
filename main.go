package main

import "fmt"

type LRUCache struct {
	capacity int
	cacheMap map[int]*cacheNode
	head     *cacheNode
	tail     *cacheNode
}

type cacheNode struct {
	next  *cacheNode
	prev  *cacheNode
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.capacity = capacity
	l.cacheMap = make(map[int]*cacheNode)
	return l
}

func (this *LRUCache) Get(key int) int {
	// first check if we don't have the value
	if this.cacheMap[key] == nil {
		return -1
	}
	// if the value already exists and we only have one node just return the value without doing any swapping
	if this.cacheMap[key] != nil && len(this.cacheMap) == 1 {
		return this.cacheMap[key].value
	}
	// if it is equal to the tail we make certain to adjust the tail before moving the node
	// if we do not have a single space for capacity
	if this.capacity != 1 && this.tail == this.cacheMap[key] {
		prevNode := this.tail.prev
		moveToHead(this.head, this.cacheMap[key])
		this.tail = prevNode
		this.head = this.cacheMap[key]
		return this.cacheMap[key].value
	}
	// only move if it isn't the head
	if this.head != this.cacheMap[key] {
		moveToHead(this.head, this.cacheMap[key])
		this.head = this.cacheMap[key]
	}
	return this.cacheMap[key].value
}

// Put - implemented a doubly linked list. Head is the beginning of queue. Tail being the end
func (this *LRUCache) Put(key int, value int) {
	// if our linked list is empty
	if this.head == nil {
		cn := cacheNode{
			key:   key,
			value: value,
		}
		this.head = &cn
		this.tail = &cn
		this.cacheMap[key] = &cn
		return
	}
	// if this value already exists in our linked list we update the value and put it at the beginning
	if this.cacheMap[key] != nil {
		// update in place if the value already exists
		if len(this.cacheMap) == 1 {
			this.cacheMap[key].value = value
			return
		}
		// don't move the head if it is already the HEAD
		if this.cacheMap[key] == this.head {
			this.cacheMap[key].value = value
			return
		}
		// otherwise update and move
		if this.capacity != 1 && this.tail == this.cacheMap[key] {
			prevNode := this.tail.prev
			prevNode.next = nil
			this.tail = prevNode
		}
		this.cacheMap[key].value = value
		moveToHead(this.head, this.cacheMap[key])
		this.head = this.cacheMap[key]
		return
	}
	// eviction (LRU)
	if len(this.cacheMap) == this.capacity {
		// the code in the else block will not work with a capacity of 1
		// we only need to remove the value from the cache map
		if this.capacity == 1 {
			delete(this.cacheMap, this.head.key)
		} else {
			prevNode := this.tail.prev
			// remove the previous node's connect to our current tail
			prevNode.next = nil
			// remove it from our map
			delete(this.cacheMap, this.tail.key)
			// remove from the tail connection
			this.tail = prevNode
		}
		// fall through to the next block of code to add to our head
	}
	cn := cacheNode{
		next:  this.head,
		key:   key,
		value: value,
	}
	this.cacheMap[key] = &cn
	node := this.head
	node.prev = &cn
	this.head = &cn
}

func (this LRUCache) display() {
	node := this.head
	for {
		if node.next == nil {
			fmt.Printf("%+v\n", node)
			break
		}
		fmt.Printf("%+v -> ", node)
		node = node.next
	}
}

func moveToHead(head *cacheNode, mover *cacheNode) {
	// actions to replace the moved node in the list
	// replace the previous nodes next connection with the node following the one we are moving
	// replace the next nodes previous connection with the node prior to the one we are moving
	if mover.prev != nil {
		mover.prev.next = mover.next
	}
	if mover.next != nil {
		mover.next.prev = mover.prev
	}
	// since this is now the head there is no previous node
	mover.prev = nil
	// since it is at the head we need to now set this to be the old head
	mover.next = head
	// head node is now the second in the list set its previous to the moved node
	head.prev = mover
}

func main() {
	// lRUCache := Constructor(2)
	// lRUCache.Put(1, 1) // cache is {1=1}
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(1)) // return 1
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(3)) // return 3
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(4)) // return 4
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)

	// lRUCache := Constructor(1)
	// lRUCache.Put(2, 1)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(2))
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(3, 2)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(2))
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(3))
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", *lRUCache.head, *lRUCache.tail, lRUCache.cacheMap)

	// lRUCache := Constructor(2)
	// lRUCache.Put(2, 1)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(2, 2)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(2))
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(1, 1)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.Put(4, 1)
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// fmt.Println(lRUCache.Get(2))
	// fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)

	lRUCache := Constructor(10)
	lRUCache.Put(10, 13)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 17)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(6, 11)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(10, 5)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 10)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(13))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(2, 19)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(2))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(3))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(5, 25)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(8))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 22)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(5, 5)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(1, 30)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(11))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 12)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(7))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(5))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(8))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(9))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(4, 30)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 3)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(9))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(10))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(10))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(6, 14)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 1)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(3))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(10, 11)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(8))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(2, 14)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(1))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(5))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(4))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(11, 4)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(12, 24)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(5, 18)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(13))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(7, 23)
	fmt.Println(lRUCache.Get(8))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(12))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 27)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(2, 12)
	fmt.Println(lRUCache.Get(5))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(2, 9)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(13, 4)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(8, 18)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(1, 7)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(6))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 29)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(8, 21)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(5))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(6, 30)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(1, 12)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(10))
	lRUCache.Put(4, 15)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(7, 22)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(11, 26)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(8, 17)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(9, 29)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(5))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 4)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(11, 30)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(12))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(4, 29)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(3))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(9))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(6))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 4)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(1))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(10))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 29)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(10, 28)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(1, 20)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(11, 13)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	fmt.Println(lRUCache.Get(3))
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 12)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	// lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 8)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(10, 9)
	fmt.Printf("head: %+v tail: %+v cacheMap: %+v\n", lRUCache.head, lRUCache.tail, lRUCache.cacheMap)
	lRUCache.display()
	fmt.Println()
	lRUCache.Put(3, 26)
	fmt.Println(lRUCache.Get(8))
	fmt.Println(lRUCache.Get(7))
	fmt.Println(lRUCache.Get(5))
	lRUCache.Put(13, 17)
	lRUCache.Put(2, 27)
	lRUCache.Put(11, 15)
	fmt.Println(lRUCache.Get(12))
	lRUCache.Put(9, 19)
	lRUCache.Put(2, 15)
	lRUCache.Put(3, 16)
	fmt.Println(lRUCache.Get(1))
	lRUCache.Put(12, 17)
	lRUCache.Put(9, 1)
	lRUCache.Put(6, 19)
	fmt.Println(lRUCache.Get(4))
	fmt.Println(lRUCache.Get(5))
	fmt.Println(lRUCache.Get(5))
	lRUCache.Put(8, 1)
	lRUCache.Put(11, 7)
	lRUCache.Put(5, 2)
	lRUCache.Put(9, 28)
	fmt.Println(lRUCache.Get(1))
	lRUCache.Put(2, 2)
	lRUCache.Put(7, 4)
	lRUCache.Put(4, 22)
	lRUCache.Put(7, 24)
	lRUCache.Put(9, 26)
	lRUCache.Put(13, 28)
	lRUCache.Put(11, 26)
}

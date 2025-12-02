package main

import "fmt"

const ArraySize = 7 //js, java, python, go, jsx, tsx, cpp
// hash table structure
type HashTable struct {
	array [ArraySize]*bucket //pointers
}

// insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].del(key)
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// insert
// insert will take in a key, create a node with the key and then insert it into the buckert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

// search
// search will take in a key and return true of bucket contains
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil { //very interesting for loop acting as a while loop
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
// delete takes in a key adn deletes it from a bucket
func (b *bucket) del(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		} else {
			previousNode = previousNode.next
		}
	}
}

// bucket node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// init
// init creates a bucket in each slot of the hash table
func create() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}

// hashing function
func hash(key string) int {
	index := 0
	for _, char := range key {
		index += int(char)

	}
	return (index % ArraySize)
}
func main() {
	testTable := create()
	list := []string{
		".jsx",
		".js",
		".java",
		".go",
		".py",
		".cpp",
		".tsx",
	}

	for _, v := range list {
		testTable.Insert(v)
	}

	fmt.Println(testTable)
}

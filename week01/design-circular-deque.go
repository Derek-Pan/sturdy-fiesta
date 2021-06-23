package main

// 双端队列特点: 1. 队头队尾都可进行insert, 时间复杂的为O(1); 2. 队头队尾都可进行delete, 时间复杂度为O(1);

// 思路: 使用map维护存储的数据,minIndex记录对头最小值, maxIndex记录队尾最大值, total记录可存储的容量, storeNum记录已使用的空间

type MyCircularDeque struct {
	data     map[int]int
	minIndex int
	maxIndex int
	storeNum int
	total    int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:     make(map[int]int),
		minIndex: -9999,
		maxIndex: -9999,
		storeNum: k,
		total:    k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.storeNum == 0 {
		return false
	}
	if this.minIndex == -9999 {
		this.data[0] = value
		this.minIndex = 0
		this.maxIndex = 0
		this.storeNum--
		return true
	}
	this.data[this.minIndex-1] = value
	this.storeNum--
	this.minIndex--
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.storeNum == 0 {
		return false
	}
	if this.maxIndex == -9999 {
		this.data[1] = value
		this.maxIndex = 1
		this.minIndex = 1
		this.storeNum--
		return true
	}
	this.data[this.maxIndex+1] = value
	this.storeNum--
	this.maxIndex++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.storeNum == this.total {
		return false
	}
	delete(this.data, this.minIndex)
	this.minIndex++
	this.storeNum++
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.storeNum == this.total {
		return false
	}
	delete(this.data, this.maxIndex)
	this.maxIndex--
	this.storeNum++
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.storeNum == this.total {
		return -1
	}
	return this.data[this.minIndex]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.storeNum == this.total {
		return -1
	}
	return this.data[this.maxIndex]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.storeNum == this.total
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.storeNum == 0
}

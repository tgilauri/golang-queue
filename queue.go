package queue

const DEFAULT_TRESHOLD = 1024

type SQueue[T any] struct {
	queue       []T
	size        int
	length      int
	startIdx    int
	expandable  bool
	defaultSize int
}

type Queue[T any] interface {
	PushLeft(item T) bool
	PushRight(item T) bool
	PopRight() T
	PopLeft() T
	GetLast() T
	GetFirst() T
	IsEmpty() bool
	Clear()
	GetLength() int
	GetSize() int
}

func NewQueue[T any](size int, expandable bool) Queue[T] {
	queue := new(SQueue[T])

	queue.queue = make([]T, size)
	queue.size = size
	queue.length = 0
	queue.startIdx = 0
	queue.expandable = expandable
	queue.defaultSize = DEFAULT_TRESHOLD

	if size > DEFAULT_TRESHOLD {
		queue.defaultSize = size
	}

	return queue
}

func FromArray[T any](array []T, expandable bool) Queue[T] {
	queue := new(SQueue[T])

	queue.queue = make([]T, len(array))
	queue.size = len(array)
	queue.length = len(array)
	queue.startIdx = 0
	queue.expandable = expandable

	copy(queue.queue, array)

	return queue
}

// PUBLIC METHODS

func (this *SQueue[T]) GetLength() int {
	return this.length
}

func (this *SQueue[T]) GetSize() int {
	return this.size
}

func (this *SQueue[T]) PushLeft(item T) bool {
	if this.expandable {
		this.extend(1)
	}
	if this.length < this.size {
		newIdx := this.size - 1
		if this.startIdx > 0 {
			newIdx = this.startIdx - 1
		}
		this.queue[newIdx] = item
		this.startIdx = newIdx
		this.length += 1
		return true
	} else {
		return false
	}
}

func (this *SQueue[T]) PushRight(item T) bool {
	if this.expandable {
		this.extend(1)
	}
	if this.length < this.size {
		newIdx := this.getNewItemIdx()
		this.queue[newIdx] = item
		this.length += 1
		return true
	} else {
		return false
	}
}

func (this *SQueue[T]) PopLeft() T {
	if this.IsEmpty() {
		panic("No items in queue.")
	}
	newIdx := 0
	if this.startIdx < this.size-1 {
		newIdx = this.startIdx + 1
	}

	item := this.queue[this.startIdx]
	this.length -= 1
	this.startIdx = 0
	if this.length > 0 {
		this.startIdx = newIdx
	}
	return item
}

func (this *SQueue[T]) PopRight() T {
	if this.IsEmpty() {
		panic("No items in queue.")
	}

	lastItemIdx := this.getLastItemIdx()
	item := this.queue[lastItemIdx]
	this.length -= 1
	if this.length == 0 {
		this.startIdx = 0
	}
	return item
}

func (this *SQueue[T]) GetFirst() T {
	return this.queue[this.startIdx]
}

func (this *SQueue[T]) GetLast() T {
	return this.queue[this.getLastItemIdx()]
}

func (this *SQueue[T]) IsEmpty() bool {
	return this.length == 0
}

func (this *SQueue[T]) Clear() {
	this.length = 0
	this.startIdx = 0
	this.queue = make([]T, this.size)
}

// PRIVATE METHODS

func (this *SQueue[T]) getNewItemIdx() int {
	return (this.startIdx + this.length) % this.size
}

func (this *SQueue[T]) getLastItemIdx() int {
	if this.startIdx == 0 && this.length > 0 {
		return this.startIdx + this.length - 1
	} else {
		return this.getNewItemIdx() - 1
	}
}

/* func (this *SQueue[T]) extend(newSize int) {
	if this.length == this.size {
		newQueue := make([]T, this.size*2)
		if this.startIdx != 0 {
			tmpQueue := append(this.queue[this.startIdx:], this.queue[0:this.startIdx]...)
			newQueue = append(tmpQueue, make([]T, this.size)...)
			this.startIdx = 0
		}
		this.queue = newQueue
		this.size *= 2
	}
} */

func (this *SQueue[T]) extend(num int) {
	newSize := this.size + num
	newLen := this.length + num
	doubleSize := this.size * 2

	if newLen <= this.size {
		return
	}

	if newSize < doubleSize {
		if this.length < this.defaultSize {
			newSize = doubleSize
		} else {
			newSize = this.size + int(this.defaultSize/4)
		}
	}

	newQueue := make([]T, newSize)
	if this.startIdx != 0 {
		tmpQueue := append(this.queue[this.startIdx:], this.queue[0:this.startIdx]...)
		newQueue = append(tmpQueue, make([]T, this.size)...)
		this.startIdx = 0
	}
	this.queue = newQueue
	this.size = newSize
}

# golang-queue

Queue implementation in golang

## Create queue

### Create empty queue of known size and type

`NewQueue` method creates new queue of given type and size. `Expandable` parameter indicates wheter queue should expand in case there is no room for new elements

To create an empty queue run:

```golang
// size - the size of the queue
// expandable - indicates whether queue will expand if there is no room for new elements
// type - type of the elements that will be stored in queue
queue := NewQueue[type](size, expanable)
```

### Create queue from array

`FromArray` method will create a queue contains all the elements from array and will have size and length of the source array.

To create a queue from given array run:

```golang
queue := FromArray[type](array)
```

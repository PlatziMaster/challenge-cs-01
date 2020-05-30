package main

import "fmt"

func main() {
	// Dynamic array usage example
	myList := CreateList()

	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	myList.Append(4)
	myList.Append(55)
	myList.Append(6)
	myList.Append(7)
	myList.RemoveOne(4)
}

// Dynamic array data structure implementation.
type DynamicArray struct {
	data 		[]interface{}
	// Logical space.
	length 		int
	// Total space.
	capacity 	int
}

// Return a new dynamic array with no data and length of 0.
func CreateList() DynamicArray {
	return DynamicArray{}
}

// Adds a new element at the beginning of the array.
func (d *DynamicArray) Add(value interface{}) {
	if d.length == d.capacity {
		d.resizeListSpace()
	}
	oldData := make([]interface{}, d.length)
	copy(oldData, d.data)
	newData := make([]interface{}, 1)
	newData[0] =  value
	data := append(newData, oldData...)
	d.data = data
	d.length++
}

// Deletes the first element of the array.
func (d *DynamicArray) Delete() {
	if d.length <= 0 {
		fmt.Errorf("List is empty. Cannot perform delete.")
	}
	newData := make([]interface{}, d.capacity)
	copy(newData, d.data[1:])
	d.data = newData
	d.length--
}

// Appends a new element.
func (d *DynamicArray) Append(value interface{}) {
	if d.length == d.capacity {
		d.resizeListSpace()
	}
	d.data[d.length] = value
	d.length++
}

// Pop deletes the last element of the array.
func (d *DynamicArray) Pop() {
	if d.length <= 0 {
		fmt.Errorf("List is empty. Cannot perform pop.")
	}
	d.data[d.length - 1] = nil
	d.length--
}

// Inserts a new element at the desired position.
func (d *DynamicArray) Insert(value interface{}, index int) {
	if d.length == d.capacity {
		d.resizeListSpace()
	}
	if index > d.length - 1 {
		fmt.Errorf("Index %d is out of range %d", index, d.length)
	} else {
		leftSide := make([]interface{}, index)
		copy(leftSide, d.data[:index])
		newData := make([]interface{}, 1)
		newData[0] =  value
		leftData := append(leftSide, newData...)
		data := append(leftData, d.data[index:]...)
		d.data = data
		d.length++
	}
}

// Removes an element at the desired position.
func (d *DynamicArray) RemoveOne(index int) {
	if d.length <= 0 {
		fmt.Errorf("List is empty. Cannot perform delete.")
	}
	leftSide := make([]interface{}, index)
	copy(leftSide, d.data[:index])
	data := append(leftSide, d.data[index+1:]...)
	d.data = data
	d.length--
}

// Length returns the length of the array. 
func (d *DynamicArray) Length() int {
	return int(d.length)
}

// Resize the array to save memory for new future elements.
func (d *DynamicArray) resizeListSpace() {
	// Prevent resize if there is free space yet.
	if d.capacity < d.capacity {
		return
	} else if d.capacity == 0 || d.capacity == 1 {
		d.capacity = 2
	} else {
		d.capacity = d.capacity * 2
	}
	newData := make([]interface{}, d.capacity)
	copy(newData, d.data)
	d.data = newData
}


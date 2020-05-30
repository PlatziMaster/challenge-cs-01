package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	arrayUnxOpt = cmp.AllowUnexported(DynamicArray{})
)

func compareArrays(t *testing.T, expected, got *DynamicArray, producer string) {
	diff := cmp.Diff(expected, got, arrayUnxOpt);
	if diff != "" {
		t.Fatalf("%s produced unwanted array: %d\nwant %d\ndiff want -> got\n%s",
			producer, got, expected, diff)
	}
}

func TestArray_CreateList(t *testing.T) {
	expected := DynamicArray{data: nil, length: 0, capacity: 0}
	got := CreateList()
	compareArrays(t, &expected, &got, "CreateList")
}

func TestArray_Append(t *testing.T) {
	tests := []struct {
		name 		  string
		init, want 	  DynamicArray
		insertedData  []int
	} {
		{
			name: "Append element to new array",
			init: DynamicArray{data: []interface{}{}, length: 0, capacity: 0},
			want: DynamicArray{
				data: 	  []interface{}{1, nil},
				length:   1,
				capacity: 2,
			},
			insertedData: []int{1},
		},
		{
			name: "Append a new element before resize",
			init: DynamicArray{data: []interface{}{}, length: 0, capacity: 0},
			want: DynamicArray{
				data: 	  []interface{}{1, 2},
				length:   2,
				capacity: 2,
			},
			insertedData: []int{1, 2},
		},
		{
			name: "Resize array",
			init: DynamicArray{data: []interface{}{}, length: 0, capacity: 0},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, 3, nil},
				length:   3,
				capacity: 4,
			},
			insertedData: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, i := range test.insertedData {
				test.init.Append(i)
			}
		})
		compareArrays(t, &test.want, &test.init, "Append")
	}
}

func TestArray_Add(t *testing.T) {
	tests := []struct {
		name 		  string
		init, want 	  DynamicArray
		insertedData  []int
	} {
		{
			name: "Add element to new array",
			init: DynamicArray{data: []interface{}{}, length: 0, capacity: 0},
			want: DynamicArray{
				data: 	  []interface{}{1},
				length:   1,
				capacity: 2,
			},
			insertedData: []int{1},
		},
		{
			name: "Append a new element before resize",
			init: DynamicArray{data: []interface{}{}, length: 0, capacity: 0},
			want: DynamicArray{
				data: 	  []interface{}{2, 1},
				length:   2,
				capacity: 2,
			},
			insertedData: []int{1, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, i := range test.insertedData {
				test.init.Add(i)
			}
		})
		compareArrays(t, &test.want, &test.init, "Add")
	}
}

func TestArray_Pop(t *testing.T) {
	tests := []struct {
		name 	    string
		init, want  DynamicArray
	} {
		{
			name: "Pops the last element in a 1 length array",
			init: DynamicArray{data: []interface{}{1, nil}, length: 1, capacity: 2},
			want: DynamicArray{
				data: 	  []interface{}{nil, nil},
				length:   0,
				capacity: 2,
			},
		},
		{
			name: "Pops one time",
			init: DynamicArray{data: []interface{}{1, 2}, length: 2, capacity: 2},
			want: DynamicArray{
				data: 	  []interface{}{1, nil},
				length:   1,
				capacity: 2,
			},
		},
		{
			name: "Pops after resize",
			init: DynamicArray{data: []interface{}{1, 2, 3, nil}, length: 3, capacity: 4},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, nil, nil},
				length:   2,
				capacity: 4,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.Pop()
		})
		compareArrays(t, &test.want, &test.init, "Pop")
	}
}

func TestArray_Delete(t *testing.T) {
	tests := []struct {
		name 	    string
		init, want  DynamicArray
	} {
		{
			name: "Deletes the first element of a 2 elements array",
			init: DynamicArray{data: []interface{}{1, 2}, length: 2, capacity: 2},
			want: DynamicArray{
				data: 	  []interface{}{2, nil},
				length:   1,
				capacity: 2,
			},
		},
		{
			name: "Deletes after resize",
			init: DynamicArray{data: []interface{}{1, 2, 3, nil}, length: 3, capacity: 4},
			want: DynamicArray{
				data: 	  []interface{}{2, 3, nil, nil},
				length:   2,
				capacity: 4,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.Delete()
		})
		compareArrays(t, &test.want, &test.init, "Delete")
	}
}

func TestArray_Insert(t *testing.T) {
	tests := []struct {
		name 	    string
		init, want  DynamicArray
		value, idx  int
	} {
		{
			name: "Add an element before resize",
			init: DynamicArray{data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8}, length: 8, capacity: 8},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, 3, 4, 55, 5, 6, 7, 8, nil, nil, nil, nil, nil, nil, nil, nil},
				length:   9,
				capacity: 16,
			},
			value: 55,
			idx: 4,
		},
		{
			name: "Add an element without resize",
			init: DynamicArray{data: []interface{}{1, 2, 3, 4, 5, 6, 7, nil}, length: 7, capacity: 8},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, 3, 4, 5, 55, 6, 7, nil},
				length:   8,
				capacity: 8,
			},
			value: 55,
			idx: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.Insert(test.value, test.idx)
		})
		compareArrays(t, &test.want, &test.init, "Insert")
	}
}

func TestArray_RemoveOne(t *testing.T) {
	tests := []struct {
		name 	    string
		init, want  DynamicArray
		idx			int
	} {
		{
			name: "Delete one element without resize",
			init: DynamicArray{data: []interface{}{1, 2, 3, 4, 55, 5, 6, 7, 8, nil, nil, nil, nil, nil, nil, nil, nil}, length: 9, capacity: 16},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, 3, 4, 5, 6, 7, 8, nil, nil, nil, nil, nil, nil, nil, nil},
				length:   8,
				capacity: 16,
			},
			idx: 4,
		},
		{
			name: "Delete element when array has enough capacity",
			init: DynamicArray{data: []interface{}{1, 2, 3, 4, 55, 6, 7, nil}, length: 8, capacity: 8},
			want: DynamicArray{
				data: 	  []interface{}{1, 2, 3, 4, 6, 7, nil},
				length:   7,
				capacity: 8,
			},
			idx: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.RemoveOne(test.idx)
		})
		compareArrays(t, &test.want, &test.init, "Remove one")
	}
}

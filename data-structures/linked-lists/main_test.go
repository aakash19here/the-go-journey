package main

import (
	"reflect"
	"testing"
)

func buildList(values []int) (*LinkedList, map[int]*Node) {
	list := &LinkedList{}
	nodes := make(map[int]*Node, len(values))

	for i := len(values) - 1; i >= 0; i-- {
		node := &Node{Data: values[i]}
		nodes[values[i]] = node
		list.InsertAtStart(node)
	}

	return list, nodes
}

func listValues(list *LinkedList) []int {
	var values []int

	for current := list.Head; current != nil; current = current.Next {
		values = append(values, current.Data)
	}

	return values
}

func TestInsertAtStart(t *testing.T) {
	list := &LinkedList{}

	list.InsertAtStart(&Node{Data: 1})
	list.InsertAtStart(&Node{Data: 2})
	list.InsertAtStart(&Node{Data: 3})

	if got, want := listValues(list), []int{3, 2, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected list order: got %v, want %v", got, want)
	}

	if got, want := list.Length, 3; got != want {
		t.Fatalf("unexpected length: got %d, want %d", got, want)
	}
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name           string
		initial        []int
		deleteValue    int
		useMissingNode bool
		useNilNode     bool
		wantValues     []int
		wantLength     int
	}{
		{
			name:        "delete from empty list",
			initial:     nil,
			deleteValue: 1,
			wantValues:  nil,
			wantLength:  0,
		},
		{
			name:        "delete head node",
			initial:     []int{1, 2, 3},
			deleteValue: 1,
			wantValues:  []int{2, 3},
			wantLength:  2,
		},
		{
			name:        "delete middle node",
			initial:     []int{1, 2, 3, 4},
			deleteValue: 3,
			wantValues:  []int{1, 2, 4},
			wantLength:  3,
		},
		{
			name:        "delete tail node",
			initial:     []int{1, 2, 3, 4},
			deleteValue: 4,
			wantValues:  []int{1, 2, 3},
			wantLength:  3,
		},
		{
			name:           "delete missing node",
			initial:        []int{1, 2, 3},
			deleteValue:    99,
			useMissingNode: true,
			wantValues:     []int{1, 2, 3},
			wantLength:     3,
		},
		{
			name:       "delete nil node",
			initial:    []int{1, 2, 3},
			useNilNode: true,
			wantValues: []int{1, 2, 3},
			wantLength: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, nodes := buildList(tt.initial)

			switch {
			case tt.useNilNode:
				list.DeleteNode(nil)
			case tt.useMissingNode:
				list.DeleteNode(&Node{Data: tt.deleteValue})
			default:
				list.DeleteNode(nodes[tt.deleteValue])
			}

			if got := listValues(list); !reflect.DeepEqual(got, tt.wantValues) {
				t.Fatalf("unexpected list values: got %v, want %v", got, tt.wantValues)
			}

			if got := list.Length; got != tt.wantLength {
				t.Fatalf("unexpected length: got %d, want %d", got, tt.wantLength)
			}
		})
	}
}

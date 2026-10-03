/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap struct {
    array []int
}

func (heap *MinHeap) Insert(element int) {
    heap.array = append(heap.array, element)
    l := len(heap.array)-1
    heap.heapifyUp(l)
}

func (heap *MinHeap) Extract() int {
    min := heap.array[0]
    l := len(heap.array)-1
    heap.array[0] = heap.array[l]
    heap.array = heap.array[:l]
    heap.heapifyDown(0)
    return min
}

func (heap *MinHeap) Swap(i, j int) {
    heap.array[i], heap.array[j] = heap.array[j], heap.array[i]
}

func parent(i int) int {
    return (i-1)/2
}

func left(i int) int {
    return 2*i+1
}

func right(i int) int {
    return 2*i+2
}

func (heap *MinHeap) heapifyUp(i int) {
    index := i
    for heap.array[index] < heap.array[parent(index)] {
        heap.Swap(index, parent(index))
        index = parent(index)
    }
}

func (heap *MinHeap) heapifyDown(i int) {
    l := len(heap.array)-1
    index := i
    for left(index) <= l {
        le, ri := left(index), right(index)
        cmp := index

        if le > l {
            return
        }

        if le == l {
            cmp = le
        } else if heap.array[le] < heap.array[ri] {
            cmp = le
        } else {
            cmp = ri
        }

        if heap.array[cmp] >= heap.array[index] {
            return
        }

        heap.Swap(index, cmp)
        index = cmp
    }
}

func mergeKLists(lists []*ListNode) *ListNode {
    heap := MinHeap{}

    remaining := true
    for remaining {
        remaining = false
        for i := range lists {
            if lists[i] != nil {
                remaining = true
                heap.Insert(lists[i].Val)
                lists[i] = lists[i].Next
            }
        }
    }

    merged := ListNode {}
    head := &merged

    for len(heap.array) > 0 {
        head.Next = &ListNode { Val: heap.Extract() }
        head = head.Next
    }

    return merged.Next
}

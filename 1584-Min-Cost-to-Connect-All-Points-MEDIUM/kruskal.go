import "container/heap"

type Edge struct {
	i1 int
	x1 int
	y1 int
	i2 int
	x2 int
	y2 int
}

type MinHeap []Edge
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return dist(h[i]) < dist(h[j]) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) { *h = append(*h, x.(Edge)) }
func (h *MinHeap) Pop() any {
	n := h.Len()
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type DisjointSet map[int]int
func (d *DisjointSet) Add(x int) {
	if _, exists := (*d)[x]; !exists {
		(*d)[x] = -1
	}
}
func (d *DisjointSet) Find(x int) int {
	if (*d)[x] == -1 {
		return x
	}
	(*d)[x] = d.Find((*d)[x])
	return (*d)[x]
}
func (d *DisjointSet) Union(x, y int) {
	x = d.Find(x)
	y = d.Find(y)
	if x != y {
		(*d)[y] = x
	}
}

func minCostConnectPoints(points [][]int) int {
	h := &MinHeap{}
	d := DisjointSet{}
	totalCost := 0

	heap.Init(h)

	for i1, p1 := range points {
		x1, y1 := p1[0], p1[1]

		for i2 := i1 + 1; i2 < len(points); i2++ {
			p2 := points[i2]
			x2, y2 := p2[0], p2[1]

			heap.Push(h, Edge{
				i1, x1, y1,
				i2, x2, y2,
			})
		}
	}

	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)

		d.Add(edge.i1)
		d.Add(edge.i2)

		if d.Find(edge.i1) != d.Find(edge.i2) {
			totalCost += dist(edge)
			d.Union(edge.i1, edge.i2)
		}
	}

	return totalCost
}

func dist(e Edge) int {
	return abs(e.x2-e.x1) + abs(e.y2-e.y1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

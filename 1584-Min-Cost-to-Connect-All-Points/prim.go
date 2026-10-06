import "container/heap"

type Point struct {
	i int
	w int
}

type MinHeap []Point

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].w < h[j].w
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() any {
	n := h.Len()
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func minCostConnectPoints(points [][]int) int {
	h := &MinHeap{}
	v := make(map[int]struct{})
	totalCost := 0

	heap.Init(h)
	heap.Push(h, Point{0, 0})

	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		if _, visited := v[p.i]; visited {
			continue
		}
		totalCost += p.w
		v[p.i] = struct{}{}
		x1, y1 := points[p.i][0], points[p.i][1]

		for i2, p2 := range points {
			if _, visited := v[i2]; visited {
				continue
			}
			x2, y2 := p2[0], p2[1]
			heap.Push(h, Point{i2, dist(x1, y1, x2, y2)})
		}
	}

	return totalCost
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

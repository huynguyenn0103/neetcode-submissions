import "slices"
func abs(x, y int)int{
	return int(math.Abs(float64(x-y)))
}
func distance(p1, p2 []int)int{
	x1,y1,x2,y2 := p1[0],p1[1],p2[0],p2[1]
	return abs(x1,x2)+abs(y1,y2)
}

type Point [2]int

type UnionFind struct {
	parents map[Point]Point
	ranks map[Point]int
}

func NewUnion(n int) *UnionFind{
	return &UnionFind{
		parents: make(map[Point]Point),
		ranks: make(map[Point]int),
	}
}

func (u *UnionFind) Find(v Point)Point{
	if parent, existed := u.parents[v]; !existed{
		u.parents[v] = v
		return v
	} else {
		if parent == v{
			return parent
		} else {
			u.parents[v]=u.Find(u.parents[v])
			return u.parents[v]
		}
	}
}

func (u *UnionFind) Union(v1,v2 Point)bool{
	root1 := u.Find(v1)
	root2 := u.Find(v2)
	if root1 == root2{
		return false
	}
	rank1 := u.ranks[root1]
	rank2 := u.ranks[root2]
	if rank1 < rank2{
		u.parents[root1]=root2
	} else if rank1 > rank2{
		u.parents[root2]=root1
	} else {
		u.ranks[root2]++
		u.parents[root1]=root2
	}
	return true
}

func (u *UnionFind)Connected(v1,v2 Point)bool{
	return u.Find(v1) == u.Find(v2)
}

type Edge struct{
	nodeA Point
	nodeB Point
	weight int
}

func minCostConnectPoints(points [][]int) int {
    n := len(points)
	edges := make([]Edge,0)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++{
			edges = append(edges,Edge{nodeA: Point(points[i]), nodeB: Point(points[j]), weight: distance(points[i],points[j]) })
		}
	}
	mUnion := NewUnion(n)
	result := 0
	slices.SortFunc(edges, func(e1,e2 Edge)int{
		return e1.weight-e2.weight
	})
	for _, edge := range edges{
		nodeA, nodeB, weight := edge.nodeA, edge.nodeB, edge.weight
		if !mUnion.Connected(nodeA,nodeB){
			result += weight
			mUnion.Union(nodeA,nodeB)
		}
	}
	return result
}

import "slices"

type UnionFind struct{
	parents []int
	ranks []int
}

func NewUnion(n int) *UnionFind{
	parents := make([]int,n)
	for node := range parents{
		parents[node]=node
	}
	return &UnionFind{
		parents: parents,
		ranks: make([]int,n),
	}
}

func (u *UnionFind) Find(v int)int{
	if u.parents[v] == v{
		return v
	} else {
		u.parents[v] = u.Find(u.parents[v])
		return u.parents[v]
	}
}

func (u *UnionFind) Union(v1,v2 int)bool{
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
	} else{
		u.ranks[root2]++
		u.parents[root2]=root1
	}
	return true
}

func (u *UnionFind) Connected(v1, v2 int)bool{
	return u.Find(v1) == u.Find(v2)
}


func abs(x int)int{
	if x >= 0 {return x}
	return -x
}

func findDist(p1, p2 []int)int{
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func minCostConnectPoints(points [][]int) int {
    n := len(points)
	edges := make([][]int,0)
	for i := 0; i < n; i++{
		for j := i + 1; j < n; j++{
			edges = append(edges, []int{i,j,findDist(points[i],points[j])})
		}
	}
	slices.SortFunc(edges, func(e1,e2 []int)int{return e1[2]-e2[2]})
	u := NewUnion(n)
	result := 0
	for _, edge := range edges{
		node1, node2, dist := edge[0],edge[1],edge[2]
		if u.Connected(node1,node2){continue}
		u.Union(node1,node2)
		result = result + dist
	}
	return result
}

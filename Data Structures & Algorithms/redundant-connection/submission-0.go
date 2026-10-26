type UnionFind struct{
	parents []int
	ranks []int
}

func NewUnion(n int) *UnionFind{
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i]=i
	}
	return &UnionFind{
		parents: parents,
		ranks: make([]int, n),
	}
}

func (u *UnionFind)Find(v int)int{
	if u.parents[v] == v{
		return v
	} else {
		u.parents[v] = u.Find(u.parents[v])
		return u.parents[v]
	}
}

func (u *UnionFind)Union(v1, v2 int)bool{
	root1 := u.Find(v1)
	root2 := u.Find(v2)
	if root1 == root2{
		return false
	}
	rank1 := u.ranks[root1]
	rank2 := u.ranks[root2]
	if rank1 > rank2{
		u.parents[root2]=root1
	} else if rank1 < rank2{
		u.parents[root1]=root2
	} else {
		u.ranks[root2]++
		u.parents[root1]=root2
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
    fUnion := NewUnion(len(edges))
	for _, edge := range edges {
		isSuccess := fUnion.Union(edge[0]-1, edge[1]-1)
		if !isSuccess{
			return edge
		}
	}
	return []int{}
}

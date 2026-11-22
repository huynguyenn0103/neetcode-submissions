type UnionFind struct {
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

func (u *UnionFind) Union(v1, v2 int)bool{
	root1 := u.Find(v1)
	root2 := u.Find(v2)
	if root1 == root2 {return false}
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

func (u *UnionFind) NumOfConnected() int{
	visited := make(map[int]bool)
	for node := range u.parents{
		visited[u.Find(node)]=true
	}
	return len(visited)
}

func validTree(n int, edges [][]int) bool {
    u := NewUnion(n)
	for _, edge := range edges{
		node1, node2 := edge[0], edge[1]
		if !u.Union(node1, node2){
			return false
		}
	}
	return u.NumOfConnected() == 1
}

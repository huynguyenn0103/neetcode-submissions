
type DSU struct {
    parent []int
    rank   []int
    count  int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    rank := make([]int, n)

    for i := 0; i < n; i++ {
        parent[i] = i
    }

    return &DSU{
        parent: parent,
        rank:   rank,
        count:  n,
    }
}

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
    rootX := d.Find(x)
    rootY := d.Find(y)

    if rootX == rootY {
        return false
    }

    if d.rank[rootX] < d.rank[rootY] {
        rootX, rootY = rootY, rootX
    }

    d.parent[rootY] = rootX

    if d.rank[rootX] == d.rank[rootY] {
        d.rank[rootX]++
    }

    d.count--
    return true
}

type Edge struct {
    u, v, weight int
    index        int
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    // Preserve original edge indices.
    sortedEdges := make([]Edge, len(edges))

    for i, e := range edges {
        sortedEdges[i] = Edge{
            u:      e[0],
            v:      e[1],
            weight: e[2],
            index:  i,
        }
    }

    // Sort by weight.
    sort.Slice(sortedEdges, func(i, j int) bool {
        return sortedEdges[i].weight < sortedEdges[j].weight
    })

    // Calculate the original MST weight.
    mstWeight := kruskal(n, sortedEdges, -1, -1)

    critical := []int{}
    pseudoCritical := []int{}

    for _, edge := range sortedEdges {
        // 1. Skip this edge and calculate MST again.
        weightWithout := kruskal(n, sortedEdges, edge.index, -1)

        if weightWithout > mstWeight {
            critical = append(critical, edge.index)
            continue
        }

        // 2. Force this edge into the MST.
        weightWith := kruskal(n, sortedEdges, -1, edge.index)

        if weightWith == mstWeight {
            pseudoCritical = append(pseudoCritical, edge.index)
        }
    }

    return [][]int{critical, pseudoCritical}
}

func kruskal(n int, edges []Edge, skipIndex, forceIndex int) int {
    dsu := NewDSU(n)
    totalWeight := 0
    usedEdges := 0

    // Force the selected edge first.
    if forceIndex != -1 {
        for _, edge := range edges {
            if edge.index == forceIndex {
                if dsu.Union(edge.u, edge.v) {
                    totalWeight += edge.weight
                    usedEdges++
                }
                break
            }
        }
    }

    // Run normal Kruskal.
    for _, edge := range edges {
        if edge.index == skipIndex || edge.index == forceIndex {
            continue
        }

        if dsu.Union(edge.u, edge.v) {
            totalWeight += edge.weight
            usedEdges++

            if usedEdges == n-1 {
                break
            }
        }
    }

    // Cannot construct a spanning tree.
    if usedEdges != n-1 {
        return int(^uint(0) >> 1)
    }

    return totalWeight
}
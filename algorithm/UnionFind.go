package algorithm

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	size := make([]int, n)
	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) Find(x int) int {
	// 查找根节点
	root := x
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	// 路径压缩，将root节点和x之间的所有节点和root直接连接
	for x != root {
		uf.parent[x], x = root, uf.parent[x]
		// origParent := uf.parent[x]
		// uf.parent[x] = root
		// x = origParent
	}
	return root
}

func (uf *UnionFind) Mmerge(x, y int) {
	xRoot, yRoot := uf.Find(x), uf.Find(y)
	if xRoot != yRoot {
		uf.parent[xRoot] = yRoot
	}
}

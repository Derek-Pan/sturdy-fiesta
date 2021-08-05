package week07

/*
题目: 冗余连接(并查集解法)
思路: 遍历每一条边, 判断边的两个顶点是否属于同一个集合,当两个顶点属于同一个集合的时,说明已经形成了环,
*/

type DisJoinSet struct {
    fa []int
}

func (djs *DisJoinSet) find(x int) int {
    if x == djs.fa[x] {
        return x
    }
    djs.fa[x] = djs.find(djs.fa[x])
    return djs.fa[x]
}

func (djs *DisJoinSet) UnionSet(x, y int) bool {
    x = djs.find(x)
    y = djs.find(y)
    if x != y {
        djs.fa[x] = y
        return false
    }
    return true
}

func findRedundantConnection(edges [][]int) []int {
    var fa []int = make([]int, len(edges) + 1, len(edges) + 1)
    for index, _ := range fa {
        fa[index] = index
    }

    var djs *DisJoinSet = &DisJoinSet{fa: fa}

    for _, edge := range edges {
		// 判断两个顶点是否属于同一个父亲
        if djs.UnionSet(edge[0], edge[1]) {
            return edge
        }
    }
    return nil
}

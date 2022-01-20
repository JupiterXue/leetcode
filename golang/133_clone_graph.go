package main

type Node struct {
	Val int
	Neighbors []*Node
}

// 1. 题解，AC！深搜
func cloneGraph(node *Node) *Node {
	// 记录已访问的边（节点与节点的关系）
	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	// 匿名函数，拷贝图
	cg = func(node *Node) *Node {
		if node == nil {  // 如果图为空返回
			return nil
		}

		// 如果节点已经被访问，直接从哈希表中取出对应的克隆节点返回
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &Node{node.Val, []*Node{}}
		// 哈希表存储
		visited[node] = cloneNode

		// 遍历该节点的邻居，并更新克隆节点的邻居列表
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(node)
}
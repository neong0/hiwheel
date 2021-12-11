package wheel

import (
	"strings"
)

type routeNode struct {
	pattern string
	path    string
	childs  []*routeNode
	isAll   bool
}

func (rn *routeNode) matchFirstChild(path string) *routeNode {
	for _, child := range rn.childs {
		if child.path == path || child.isAll {

			return child
		}
	}

	return nil
}

func (rn *routeNode) matchChildren(path string) []*routeNode {
	res := make([]*routeNode, 0)
	for _, child := range rn.childs {
		if child.path == path || child.isAll {
			res = append(res, child)
		}
	}

	return res
}

func (rn *routeNode) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		rn.pattern = pattern

		return
	}
	part := parts[height]
	child := rn.matchFirstChild(part)
	if child == nil {
		child = &routeNode{
			path:  part,
			isAll: part[0] == ':' || part[0] == '*',
		}
		rn.childs = append(rn.childs, child)

	}
	child.insert(pattern, parts, height+1)
}

func (rn *routeNode) search(parts []string, height int) *routeNode {
	if len(parts) == height || strings.HasPrefix(rn.path, "*") {
		if rn.pattern == "" {

			return nil
		}

		return rn
	}
	part := parts[height]
	childs := rn.matchChildren(part)
	for _, child := range childs {
		result := child.search(parts, height+1)
		if result != nil {

			return result
		}
	}

	return nil
}

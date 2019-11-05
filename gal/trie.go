package gal

import "strings"

type node struct {
	pattern  string
	self     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if part == child.self || child.isWild == true {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) (children []*node) {
	children = make([]*node, 0)
	for _, child := range n.children {
		if part == child.self || child.isWild == true {
			children = append(children, child)
		}
	}
	return children
}

func (n *node) insert(pattern string, parts []string, pos int) {
	if pos == len(parts) {
		n.pattern = pattern
		return
	}

	part := parts[pos]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			self:     part,
			children: make([]*node, 0),
			isWild:   part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, pos+1)
}

func (n *node) search(parts []string, pos int) *node {
	// The reason of not using n.self[0] is n.self could be ""(an empty string)
	if pos == len(parts) || strings.HasPrefix(n.self, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[pos]
	children := n.matchChildren(part)

	for _, child := range children {
		if ans := child.search(parts, pos+1); ans != nil {
			return ans
		}
	}
	return nil
}

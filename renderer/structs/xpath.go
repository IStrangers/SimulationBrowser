package structs

import "fmt"

func getXPath(node *NodeDOM) string {
	if node == nil {
		return ""
	}

	xPath := getXPath(node.Parent)

	if node.Parent == nil {
		return xPath + "/" + node.NodeName
	}

	if len(node.Parent.Children) > 0 {
		selfIdx := getSelfIdx(node)
		if selfIdx == -1 {
			xPath += "/" + node.NodeName
		} else {
			xPath += "/" + node.NodeName + "[" + fmt.Sprint(selfIdx) + "]"
		}
	} else {
		xPath += "/" + node.NodeName
	}

	return xPath
}

func getSelfIdx(node *NodeDOM) int {
	var sameTypeChildren []*NodeDOM

	parent := node.Parent
	for _, child := range parent.Children {
		if child.NodeName == node.NodeName {
			sameTypeChildren = append(sameTypeChildren, child)
		}
	}

	for idx, sameTypeChild := range sameTypeChildren {
		if sameTypeChild == node {
			return idx + 1
		}
	}

	return -1
}

package structs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	NodeType_Root = iota
	NodeType_Element
	NodeType_Text
	NodeType_Common
)

type NodeDOM struct {
	Document *Document
	//父节点
	Parent *NodeDOM
	//节点类型
	NodeType int
	//标签名
	NodeName string
	//是否是自闭合标签
	IsSelfClosing bool
	//文本节点和注释节点内容
	TextContent string
	//属性
	Attributes []*Attribute
	//是否需要回流
	NeedsReflow bool
	//是否需要重绘
	NeedsRepaint bool
	//样式表
	Style *CSSStyleSheet
	//渲染盒子
	RenderBox *RenderBox
	//子节点
	Children []*NodeDOM
	//节点解析信息
	Location *Selection
}

func (node *NodeDOM) Print(d int) {
	spacing := strings.Repeat("-", d)
	fmt.Printf("|%s> %s [%s]\n", spacing, node.NodeName, node.TextContent)

	for _, child := range node.Children {
		child.Print(d + 1)
	}
}

func (node *NodeDOM) JSON() string {
	res, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(res)
}

func (node *NodeDOM) FindByXPath(xPath string) (*NodeDOM, error) {
	if node.GetXPath() == xPath {
		return node, nil
	}

	for _, child := range node.Children {
		foundChild, err := child.FindByXPath(xPath)
		if err != nil {
			var noChild NoSuchElementError
			if errors.As(err, &noChild) {
				continue
			}

			// Some other error
			return nil, err
		}

		return foundChild, nil
	}

	return nil, NoSuchElementError(xPath)
}

func (node *NodeDOM) GetXPath() string {
	return getXPath(node)
}

func (node *NodeDOM) FindChildByName(childName string) (*NodeDOM, error) {
	if node.NodeName == childName {
		return node, nil
	}

	for _, child := range node.Children {
		foundChild, err := child.FindChildByName(childName)
		if err != nil {
			var noChild NoSuchElementError
			if errors.As(err, &noChild) {
				continue
			}

			return nil, err
		}

		return foundChild, nil
	}

	return nil, NoSuchElementError(childName)
}

func (node *NodeDOM) Attr(attrName string) string {
	for _, attribute := range node.Attributes {
		if attribute.Name == attrName {
			return attribute.Value
		}
	}

	return ""
}

func (node *NodeDOM) CalcPointIntersection(x, y float64) *NodeDOM {
	var intersectedNode *NodeDOM
	if x > float64(node.RenderBox.Left) &&
		x < float64(node.RenderBox.Left+node.RenderBox.Width) &&
		y > float64(node.RenderBox.Top) &&
		y < float64(node.RenderBox.Top+node.RenderBox.Height) {
		intersectedNode = node
	}

	for i := 0; i < len(node.Children); i++ {
		tempNode := node.Children[i].CalcPointIntersection(x, y)
		if tempNode != nil {
			intersectedNode = tempNode
		}
	}

	return intersectedNode
}

func (node NodeDOM) RequestRepaint() {
	node.NeedsRepaint = true

	for _, childNode := range node.Children {
		childNode.RequestRepaint()
	}
}

func (node NodeDOM) RequestReflow() {
	node.NeedsReflow = true

	for _, childNode := range node.Children {
		childNode.RequestReflow()
	}
}

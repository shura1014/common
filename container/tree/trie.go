package tree

import (
	"strings"
)

const FullMatchKey = "FullMatchKey"

// Keys Context
// 方便记录上下文参数
// 对于trie /user/:id这类请求应该记录参数
// 例如 /user/:id/:name /user/1001/wendell
// Params ["id":1001},{"name":"wendell"}]
type Keys map[string]any

// Node 前缀树
// / -> /user,/order
// /user -> /addUser,/delete
type Node struct {
	Name       string  //节点名称  比如 /user/get/:id 可能是 user get :id
	Children   []*Node // 子节点
	RouterName string  // /user/get/:id
	IsEnd      bool    // true

	// 是否是一个参数路径，参数名叫什么
	paramName string
	isParam   bool

	// 是否是/**路径，这个路径作为全匹配路径
	isFullMatch bool
}

// put path: /user/get/:id

func (t *Node) Put(path string) {
	root := t
	paths := strings.Split(path, "/")
	routerName := ""
	for index, name := range paths {
		if index == 0 {
			continue
		}
		children := t.Children
		isMatch := false
		for _, childNode := range children {
			// 如果user匹配到了，下一次就开始判断get
			if childNode.Name == name {
				isMatch = true
				routerName += "/" + childNode.Name
				childNode.RouterName = routerName
				t = childNode
				break
			}
		}
		if !isMatch {
			isEnd := false
			if index == len(paths)-1 {
				isEnd = true
			}
			// 没有匹配到，那么这是一个新的路径，创建一个节点对象
			childNode := &Node{Name: name, Children: make([]*Node, 0), IsEnd: isEnd}
			// 如果是一个参数路径
			if strings.HasPrefix(name, ":") {
				childNode.paramName = name[1:]
			}

			if name == "**" {
				childNode.isFullMatch = true
				childNode.paramName = FullMatchKey
			}

			if childNode.paramName != "" {
				childNode.isParam = true
			}

			routerName += "/" + childNode.Name
			childNode.RouterName = routerName
			children = append(children, childNode)
			t.Children = children
			t = childNode
		}
	}
	t = root
}

//get path: /user/get/1

func (t *Node) Get(path string, keys Keys) *Node {
	paths := strings.Split(path, "/")

	//	剩余的路径
	//	每匹配到一次就删减一点,主要作为 /**Name 这样路径的参数
	//	如 /user/get/1 -> /get/1 -> /1
	residuePath := path
	for index, name := range paths {
		if index == 0 {
			continue
		}
		//isMatch := false
		children := t.Children
		for _, childNode := range children {
			//strings.Contains(childNode.Name, ":")
			if childNode.isFullMatch {
				if childNode.isParam && nil != keys {
					keys[childNode.paramName] = name
				}
				return childNode
			}
			if childNode.Name == name || childNode.Name == "*" || childNode.isParam {
				if childNode.isParam && nil != keys {
					keys[childNode.paramName] = name
				}
				residuePath = strings.TrimPrefix(residuePath, "/"+name)
				//isMatch = true

				t = childNode
				if index == len(paths)-1 {
					return childNode
				}
				break
			}
		}
		//// 如果说没有匹配到，需要检查一下是否有**的匹配
		//if !isMatch {
		//	for _, childNode := range Children {
		//		// /user/**
		//		// /user/get/userInfo
		//		// /user/aa/bb
		//		if childNode.Name == "**" {
		//			return childNode
		//		}
		//	}
		//}
	}
	return nil
}

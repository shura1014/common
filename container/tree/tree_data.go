package tree

import "strings"

// Element 表示一个数据节点
type Element[V any] struct {
	Name     string        // 名字
	FullName string        // 全名，加上了父节点的名字
	Children []*Element[V] // 子节点
	IsEnd    bool          // 是否是尾节点
	Data     *V            // 该节点所存放的数据
}

var DefaultSplit = "/"

func NewTreeData[V any](name string) *Element[V] {
	return &Element[V]{
		Name:     name,
		Children: make([]*Element[V], 0),
	}
}

func (t *Element[V]) Put(path string, data *V) {
	root := t
	paths := strings.Split(path, DefaultSplit)
	fullName := ""
	prefix := strings.HasPrefix(path, DefaultSplit)
	for index, name := range paths {
		if prefix && index == 0 {
			continue
		}
		children := t.Children
		isMatch := false
		for _, childElement := range children {
			// 如果user匹配到了，下一次就开始判断get
			if childElement.Name == name {
				isMatch = true
				fullName += DefaultSplit + childElement.Name
				childElement.FullName = fullName
				t = childElement
				break
			}
		}
		if !isMatch {
			isEnd := false
			if index == len(paths)-1 {
				isEnd = true
			}
			// 没有匹配到，那么这是一个新的路径，创建一个节点对象
			childElement := &Element[V]{Name: name, Children: make([]*Element[V], 0), IsEnd: isEnd}
			if isEnd {
				childElement.Data = data
			}
			fullName += DefaultSplit + childElement.Name
			childElement.FullName = fullName
			children = append(children, childElement)
			t.Children = children
			t = childElement
		}
	}
	t = root
}

func (t *Element[V]) Get(path string) *Element[V] {
	paths := strings.Split(path, DefaultSplit)
	prefix := strings.HasPrefix(path, DefaultSplit)
	residuePath := path
	for index, name := range paths {
		if prefix && index == 0 {
			continue
		}
		children := t.Children
		for _, childElement := range children {
			if childElement.Name == name {

				residuePath = strings.TrimPrefix(residuePath, DefaultSplit+name)

				t = childElement
				if index == len(paths)-1 {
					return childElement
				}
				break
			}
		}
	}
	return nil
}

func (t *Element[V]) delete(path string) {
	paths := strings.Split(path, DefaultSplit)
	prefix := strings.HasPrefix(path, DefaultSplit)

	residuePath := path
	for index, name := range paths {
		if prefix && index == 0 {
			continue
		}
		children := t.Children
		for _, childElement := range children {
			if childElement.Name == name {

				residuePath = strings.TrimPrefix(residuePath, DefaultSplit+name)

				t = childElement
				if index == len(paths)-1 {
					childElement = nil
				}
				break
			}
		}
	}
}

func (t *Element[V]) Iterator(fun func(name string, fullName string, data *V)) {
	Iterator[V](t, fun)
}

func Iterator[V any](Element *Element[V], fun func(name string, fullName string, data *V)) {
	for _, children := range Element.Children {
		if children.IsEnd {
			fun(children.Name, children.FullName, children.Data)
		} else {
			Iterator(children, fun)
		}
	}
}

package glib

import (
	"unicode/utf8"
)

/* ================================================================================
 * 关键字过滤
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	Keyword struct {
		Root *keywordNode
	}

	keywordNode struct {
		Next  map[rune]*keywordNode
		IsEnd bool
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 新关键词
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewKeyword() *Keyword {
	keyword := new(Keyword)
	keyword.Root = NewkeywordNode()
	return keyword
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 新节点
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewkeywordNode() *keywordNode {
	node := new(keywordNode)
	node.Next = make(map[rune]*keywordNode)
	return node
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 增加关键字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Keyword) AddKeywords(keywords ...string) {
	if len(keywords) == 0 {
		return
	}

	for _, v := range keywords {
		s.AddKeyword(v)
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 增加关键字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Keyword) AddKeyword(keyword string) {
	if len(keyword) == 0 {
		return
	}

	currentNode := s.Root
	keywords := []rune(keyword)
	for i := 0; i < len(keywords); i++ {
		if _, isExists := currentNode.Next[keywords[i]]; !isExists {
			currentNode.Next[keywords[i]] = NewkeywordNode()
		}
		currentNode = currentNode.Next[keywords[i]]
	}

	currentNode.IsEnd = true
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤关键字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Keyword) Filter(content string, args ...string) string {
	if len(content) == 0 {
		return content
	}

	var results []rune = nil

	replaceContent := "*"
	if len(args) > 0 {
		replaceContent = args[0]
	}

	currentNode := s.Root
	contents := []rune(content)
	count := len(contents)

	for i := 0; i < count; i++ {
		if _, isExists := currentNode.Next[contents[i]]; isExists {
			currentNode = currentNode.Next[contents[i]]

			for j := i + 1; j < count; j++ {
				if _, isExists := currentNode.Next[contents[j]]; isExists {
					currentNode = currentNode.Next[contents[j]]

					if currentNode.IsEnd {
						if results == nil {
							results = contents
						}

						for offset := i; offset <= j; offset++ {
							replaceContentRune, _ := utf8.DecodeRuneInString(replaceContent)
							contents[offset] = replaceContentRune
						}

						i = j
						currentNode = s.Root
						break
					}
				}
			}
			currentNode = s.Root
		}
	}

	if results == nil {
		return content
	}

	return string(results)
}

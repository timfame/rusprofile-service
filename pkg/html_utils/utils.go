package html_utils

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	IDAttrKey    = "id"
	ClassAttrKey = "class"
	HrefAttrKey  = "href"
)

// Returns html.Node with corresponding tag and attribute key and value, and bool exists it or not.
// If there are several possible result nodes, returns the first one
func findTagByAttribute(node *html.Node, tagAtom atom.Atom, attrKey, attrValue string) (*html.Node, bool) {
	if node.DataAtom == tagAtom {
		if value, ok := GetAttributeValueByKey(node, attrKey); ok && value == attrValue {
			return node, true
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, ok := findTagByAttribute(c, tagAtom, attrKey, attrValue); ok {
			return result, true
		}
	}
	return nil, false
}

// Returns html.Node which is nth in crawl order of tree
// with corresponding attribute key and value, and int how many corresponding nodes was found
func findNthByAttribute(node *html.Node, nth int, attrKey, attrValue string) (*html.Node, int) {
	found := 0
	if value, ok := GetAttributeValueByKey(node, attrKey); ok && value == attrValue {
		if nth == 1 {
			return node, 1
		}
		found += 1
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result, cnt := findNthByAttribute(c, nth-found, attrKey, attrValue)
		found += cnt
		if found >= nth {
			return result, found
		}
	}
	return nil, found
}

// Returns html.Node with tag <div> and corresponding attribute key and value, and bool exists it or not
// If there are several possible result nodes, returns the first one
func FindDivByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, bool) {
	return findTagByAttribute(node, atom.Div, attrKey, attrValue)
}

// Returns html.Node with tag <span> and corresponding attribute key and value, and bool exists it or not.
// If there are several possible result nodes, returns the first one
func FindSpanByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, bool) {
	return findTagByAttribute(node, atom.Span, attrKey, attrValue)
}

// Returns html.Node which is nth in crawl order of tree
// and class key with corresponding value, and bool if it exists or not
func FindNthByClass(node *html.Node, nth int, classValue string) (*html.Node, bool) {
	if result, found := findNthByAttribute(node, nth, ClassAttrKey, classValue); found >= nth {
		return result, true
	}
	return nil, false
}

// Returns attribute value with corresponding key, and bool exists it or not
func GetAttributeValueByKey(node *html.Node, attrKey string) (string, bool) {
	for i := range node.Attr {
		if node.Attr[i].Key == attrKey {
			return node.Attr[i].Val, true
		}
	}
	return "", false
}

func IsA(node *html.Node) bool {
	return node.DataAtom == atom.A
}

// Returns href value from node, and bool if it exists or not.
// Also it checks that node is with tag <a>
func GetHref(node *html.Node) (string, bool) {
	if !IsA(node) {
		return "", false
	}
	return GetAttributeValueByKey(node, HrefAttrKey)
}

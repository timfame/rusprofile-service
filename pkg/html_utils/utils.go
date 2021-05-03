package html_utils

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	IDAttrKey    = "id"
	ClassAttrKey = "class"
	HrefAttrKey  = "href"
)

// Returns all text data from node
func GetText(node *html.Node) string {
	var ts []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			ts = append(ts, strings.TrimSpace(c.Data))
		}
	}
	return strings.Join(ts, " ")
}

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

// Returns html.Node with corresponding text and attribute key and value, and bool if it exists of not
func findByAttributeAndText(node *html.Node, attrKey, attrValue, text string) (*html.Node, bool) {
	if value, ok := GetAttributeValueByKey(node, attrKey); ok && value == attrValue && GetText(node) == text {
		return node, true
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, ok := findByAttributeAndText(c, attrKey, attrValue, text); ok {
			return result, true
		}
	}
	return nil, false
}

// Returns html.Node with corresponding tag among children of node, and bool if it exists or not.
// If there are several possible nodes, returns the first one
func FindTagAmongChildren(node *html.Node, tag atom.Atom) (*html.Node, bool) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.DataAtom == tag {
			return c, true
		}
	}
	return nil, false
}

// Returns html.Node among next siblings with corresponding attribute key and value, and bool if it exists or not
func FindAmongNextSiblingsByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, bool) {
	for s := node.NextSibling; s != nil; s = s.NextSibling {
		if value, ok := GetAttributeValueByKey(s, attrKey); ok && value == attrValue {
			return s, true
		}
	}
	return nil, false
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

// Returns html.Node with corresponding text and class attribute value, and bool if it exists of not
func FindSpanByClassAndText(node *html.Node, classValue, text string) (*html.Node, bool) {
	return findByAttributeAndText(node, ClassAttrKey, classValue, text)
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

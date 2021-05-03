package html_utils

import (
	"fmt"
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

// Returns html.Node with corresponding tag and attribute key and value.
// If there are several possible result nodes, returns the first one
func findTagByAttribute(node *html.Node, tagAtom atom.Atom, attrKey, attrValue string) (*html.Node, error) {
	if node.DataAtom == tagAtom {
		if value, err := GetAttributeValueByKey(node, attrKey); err == nil && value == attrValue {
			return node, nil
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, err := findTagByAttribute(c, tagAtom, attrKey, attrValue); err == nil {
			return result, nil
		}
	}
	return nil, fmt.Errorf("cannot find node with tag <%s> with attribute key \"%s\" and value \"%s\"",
		tagAtom.String(),
		attrKey,
		attrValue)
}

// Returns html.Node with corresponding text and attribute key and value
func findByAttributeAndText(node *html.Node, attrKey, attrValue, text string) (*html.Node, error) {
	if value, err := GetAttributeValueByKey(node, attrKey); err == nil && value == attrValue && GetText(node) == text {
		return node, nil
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, err := findByAttributeAndText(c, attrKey, attrValue, text); err == nil {
			return result, nil
		}
	}
	return nil, fmt.Errorf("cannot find node with with text \"%s\" attribute key \"%s\" and value \"%s\"",
		text,
		attrKey,
		attrValue)
}

// Returns html.Node with corresponding tag among children of node.
// If there are several possible nodes, returns the first one
func FindTagAmongChildren(node *html.Node, tag atom.Atom) (*html.Node, error) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.DataAtom == tag {
			return c, nil
		}
	}
	return nil, fmt.Errorf("cannot find node with tag <%s> among children", tag.String())
}

// Returns html.Node among next siblings with corresponding attribute key and value
func FindAmongNextSiblingsByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, error) {
	for s := node.NextSibling; s != nil; s = s.NextSibling {
		if value, err := GetAttributeValueByKey(s, attrKey); err == nil && value == attrValue {
			return s, nil
		}
	}
	return nil, fmt.Errorf("cannot find node with attribute key \"%s\" and value \"%s\" among next siblings",
		attrKey,
		attrValue)
}

// Returns html.Node with tag <div> and corresponding attribute key and value.
// If there are several possible result nodes, returns the first one
func FindDivByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, error) {
	return findTagByAttribute(node, atom.Div, attrKey, attrValue)
}

// Returns html.Node with tag <span> and corresponding attribute key and value.
// If there are several possible result nodes, returns the first one
func FindSpanByAttribute(node *html.Node, attrKey, attrValue string) (*html.Node, error) {
	return findTagByAttribute(node, atom.Span, attrKey, attrValue)
}

// Returns html.Node with corresponding text and class attribute value.
func FindSpanByClassAndText(node *html.Node, classValue, text string) (*html.Node, error) {
	return findByAttributeAndText(node, ClassAttrKey, classValue, text)
}

// Returns attribute value with corresponding key
func GetAttributeValueByKey(node *html.Node, attrKey string) (string, error) {
	for i := range node.Attr {
		if node.Attr[i].Key == attrKey {
			return node.Attr[i].Val, nil
		}
	}
	return "", fmt.Errorf("cannot find attribute with key \"%s\"", attrKey)
}

package helper

import (
	"errors"
	"strings"

	"github.com/jbowtie/gokogiri/xml"
)

// SearchSubtree searches the subtree of the given root node.
func SearchSubtree(root xml.Node, xpath string) (nodes []xml.Node, err error) {
	// http://stackoverflow.com/a/25387687/358804
	if !strings.HasPrefix(xpath, ".") {
		err = errors.New("XPath must have leading period (`.`) to only search the subtree")
		return
	}

	return root.Search(xpath)
}

// SearchOne searches the subtree of the given root node and returns the first result.
func SearchOne(root xml.Node, xpath string) (xml.Node, error) {
	results, err := SearchSubtree(root, xpath)
	if err != nil {
		return nil, err
	}
	return results[0], nil
}

// SearchLast searches the subtree of the given root node and returns the last result.
func SearchLast(root xml.Node, xpath string) (xml.Node, error) {
	results, err := SearchSubtree(root, xpath)
	if err != nil {
		return nil, err
	}
	lengthOfResults := len(results)
	return results[lengthOfResults-1], nil
}

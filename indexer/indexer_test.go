package indexer

import "testing"

func TestIndexDocument(t *testing.T) {
	document := NewDocument(map[string]interface{}{
		"id":   4,
		"name": "joão",
	})

	IndexDocument(document)
}

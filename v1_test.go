package server

import (
	"testing"

	"github.com/flyaways/golang-lru/simplelru"
	"github.com/gin-gonic/gin"
)

func TestMeta(t *testing.T) {
	type args struct {
		v     *gin.RouterGroup
		cache simplelru.LRUCache
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Meta(tt.args.v, tt.args.cache)
		})
	}
}

func TestObject(t *testing.T) {
	type args struct {
		v     *gin.RouterGroup
		cache simplelru.LRUCache
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Object(tt.args.v, tt.args.cache)
		})
	}
}

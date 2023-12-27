package goutils

import (
	"fmt"
	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"
	"testing"
)

type URL struct {
	Url string `json:"url"`
}

type TITLE struct {
	Title string `json:"title"`
}

func TestFromJson(t *testing.T) {

	t.Run("compond structure: case 1", func(t *testing.T) {
		type compound struct {
			URL
			TITLE
		}
		var composite compound
		FromJson([]byte(`{  "url": "www.google.com",  "title": "Google"}`), &composite)
		fmt.Println(composite)
		fmt.Println(string(ToJson(composite)))
		assert.Panics(t, func() {
			FromJson(nil, &composite)
		}, "should panic from nil json bytes")
	})

	t.Run("compound structure", func(t *testing.T) {
		type TITLE struct {
			Title string
		}
		type compound struct {
			URL
			TITLE
		}
		var composite = compound{URL{"www.google.com"}, TITLE{"Google"}}

		fmt.Println(composite)
		fmt.Println(string(ToJson(composite)))
		assert.Panics(t, func() {
			FromJson(nil, &composite)
		}, "should panic from nil json bytes")
	})
	t.Run("json array", func(t *testing.T) {
		var arr []string
		FromJson([]byte(`["whoami"]`), &arr)
		assert.Equal(t, `[whoami]`, fmt.Sprint(arr))
	})
}

func TestToJson(t *testing.T) {
	t.Run("url array", func(t *testing.T) {
		var urls = []URL{{"www.google.com"}, {"facebook"}}
		var jsonBytes = ToJson(urls)
		fmt.Println(string(jsonBytes))
	})

	t.Run("map2json", func(t *testing.T) {
		var _map = map[string]string{
			"abc": "2",
		}

		assert.Equal(t, `{"abc":"2"}`, string(ToJson(_map)))
	})

	t.Run("string array", func(t *testing.T) {
		var stringArray = []string{
			"whoami",
		}
		var jsonBytes = ToJson(stringArray)
		fmt.Println(string(jsonBytes))
	})
	t.Run("without structure", func(t *testing.T) {
		var obj = map[string]interface{}{
			"a": 1,
			"b": "buffer",
		}
		var jsonBytes = ToJson(obj)
		fmt.Println(string(jsonBytes))
	})
	t.Run("auto loading", func(t *testing.T) {
		var _map = map[string]string{
			"a": Base64Encode([]byte("b1234")),
			"c": Base64Encode([]byte("d1234")),
		}
		var byteMap = map[string][]byte{}
		var _json = ToJson(_map)
		FromJson(_json, &byteMap)
		utter.Dump(byteMap)
		for _, value := range byteMap {
			println(string(value))
		}
	})
	t.Run("handler pointer", func(t *testing.T) {
		type Version struct {
			BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
			TxNum    uint64 `protobuf:"varint,2,opt,name=tx_num,json=txNum,proto3" json:"tx_num,omitempty"`
		}
		type KVRead struct {
			Key     string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
			Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
		}
		type KVRWSet struct {
			Reads []*KVRead `protobuf:"bytes,1,rep,name=reads,proto3" json:"reads,omitempty"`
		}
		var KVRead_1 KVRead
		KVRead_1.Key = "hub"
		KVRead_1.Version = &Version{TxNum: 123, BlockNum: 2}
		var rwset = KVRWSet{Reads: []*KVRead{
			&KVRead_1,
		}}
		fmt.Printf("%s", ToJson(rwset))
	})

}

// Code generated by go generate via internal/cmd/meta; DO NOT EDIT.
package fs

import (
	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager

// Type is the type for fs
const Type = "fs"

var allowedStoragePairs = map[string]map[string]struct{}{
	"init": {
		"work_dir": struct{}{},
	},
	"list_dir": {
		"dir_func":  struct{}{},
		"file_func": struct{}{},
	},
	"read": {
		"offset": struct{}{},
		"size":   struct{}{},
	},
	"write": {
		"size": struct{}{},
	},
}

var allowedServicePairs = map[string]map[string]struct{}{}

type pairStorageInit struct {
	HasWorkDir bool
	WorkDir    string
}

func parseStoragePairInit(opts ...*types.Pair) (*pairStorageInit, error) {
	result := &pairStorageInit{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["init"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["init"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.WorkDir]
	if !ok {
		return nil, types.NewErrPairRequired(pairs.WorkDir)
	}
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}
	return result, nil
}

type pairStorageListDir struct {
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
}

func parseStoragePairListDir(opts ...*types.Pair) (*pairStorageListDir, error) {
	result := &pairStorageListDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["list_dir"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["list_dir"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.DirFunc]
	if ok {
		result.HasDirFunc = true
		result.DirFunc = v.(types.ObjectFunc)
	}
	v, ok = values[pairs.FileFunc]
	if ok {
		result.HasFileFunc = true
		result.FileFunc = v.(types.ObjectFunc)
	}
	return result, nil
}

type pairStorageRead struct {
	HasOffset bool
	Offset    int64
	HasSize   bool
	Size      int64
}

func parseStoragePairRead(opts ...*types.Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["read"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["read"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Offset]
	if ok {
		result.HasOffset = true
		result.Offset = v.(int64)
	}
	v, ok = values[pairs.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}
	return result, nil
}

type pairStorageWrite struct {
	HasSize bool
	Size    int64
}

func parseStoragePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["write"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["write"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}
	return result, nil
}

package gcs

import "go.beyondstorage.io/v5/types"

type objectPageStatus struct {
	delimiter string
	prefix    string
	offset    string
}

func (i *objectPageStatus) ContinuationToken() string {
	return i.prefix
}

func offsetData(objects []*types.Object) ([]*types.Object, string) {
	lastObject := objects[len(objects)-1]
	return objects[:len(objects)-1], lastObject.ID
}

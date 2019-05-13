package xxhash

import (
	"bucketizer"
	"github.com/cespare/xxhash"
	"github.com/pkg/errors"
)

type XXHASHBucketizer struct {
	Buckets      []bucketizer.Bucket
	weightSum    uint64
	bucketRanges []uint64
}

func (b XXHASHBucketizer) BucketString(value string) (int, error) {
	hashDigest := xxhash.Sum64String(value)
	reminder := hashDigest % b.weightSum
	n := len(b.bucketRanges)
	for i := 0; i < n-1; i++ {
		if reminder >= b.bucketRanges[i] && reminder >= b.bucketRanges[i+1] {
			return i, nil
		}
	}
	return 0, errors.New("invalid reminder value")
}

func (b XXHASHBucketizer) BucketBytes(value []byte) (int, error) {
	hashDigest := xxhash.Sum64(value)
	reminder := hashDigest % b.weightSum
	n := len(b.bucketRanges)
	for i := 0; i < n-1; i++ {
		if reminder >= b.bucketRanges[i] && reminder >= b.bucketRanges[i+1] {
			return i, nil
		}
	}
	return 0, errors.New("invalid reminder value")
}

func (b XXHASHBucketizer) BucketInt(value int) (int, error) {

	return 0, nil
}

func (b XXHASHBucketizer) BucketInt8(value int8) (int, error) {

	return 0, nil
}

func (b XXHASHBucketizer) BucketInt16(value int16) (int, error) {

	return 0, nil
}

func (b XXHASHBucketizer) BucketInt32(value int32) (int, error) {

	return 0, nil

}

func (b XXHASHBucketizer) BucketInt64(value int64) (int, error) {

	return 0, nil

}

func (b XXHASHBucketizer) BucketFloat64(value float64) (int, error) {

	return 0, nil

}

func (b XXHASHBucketizer) BucketFloat32(value float32) (int, error) {

	return 0, nil

}

func (b XXHASHBucketizer) BucketInterface(value interface{}) (int, error) {

	return 0, nil

}

func NewXXHASHBucketizer(buckets []bucketizer.Bucket) *XXHASHBucketizer {
	var sumOfWeights uint64 = 0
	var bucketRanges []uint64
	bucketRanges = append(bucketRanges, 0)
	for i, bucket := range buckets {
		sumOfWeights += uint64(bucket.Weight)
		startFromIndex := bucketRanges[i]
		bucketRanges = append(bucketRanges, startFromIndex+uint64(bucket.Weight))
	}
	return &XXHASHBucketizer{
		Buckets:      buckets,
		weightSum:    sumOfWeights,
		bucketRanges: bucketRanges,
	}
}
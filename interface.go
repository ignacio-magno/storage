package storageS3

type IStorageS3 interface {
	GetBucket() string
	GetKey() string
}

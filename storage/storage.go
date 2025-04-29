package storage

type StorageProvider interface {
	CreateBucket(bucketName string) error
}

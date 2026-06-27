package storage

type DiskStorage struct {
	Id int
	root string
}


type BucketFile struct {
	Name string
	Size int64
	Folder string
}

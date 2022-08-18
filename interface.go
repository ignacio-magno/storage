package storage

type IBase interface {
	getBucket() string
	getKey() string
}

type IDownload interface {
	IBase
}

type IUpload interface {
	IBase
	getData() []byte
}

// ===================================================================

type base struct {
	bucket string
	key    string
}

func (d *base) constructorBase(bucket, key string) {
	d.bucket = bucket
	d.key = key
}

func (d *base) getBucket() string {
	return d.bucket
}

func (d *base) getKey() string {
	return d.key
}

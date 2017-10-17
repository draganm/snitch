package data

import "github.com/elazarl/go-bindata-assetfs"

//go:generate go-bindata-assetfs -pkg data -ignore .*\.go  ./

func AssetFS() *assetfs.AssetFS {
	afs := assetFS()
	afs.Prefix = "/"
	return afs
}

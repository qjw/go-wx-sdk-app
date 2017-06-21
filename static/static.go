package static

// todo 用go generate命令不行
//go:generate go-bindata-assetfs -pkg static -o static_gen.go -prefix "" html/...

import (
	"github.com/elazarl/go-bindata-assetfs"
	"fmt"
)

func AssetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		fmt.Print(k)
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: ""}
	}
	panic("unreachable")
}
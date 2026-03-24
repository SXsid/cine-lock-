package root

import (
	"embed"
	"io/fs"
)

//go:embed assests/*

var emdeedAssests embed.FS

var StaticAssests fs.FS

func init() {
	var err error
	StaticAssests, err = fs.Sub(emdeedAssests, "assests")
	if err != nil {
		panic(err)
	}
}

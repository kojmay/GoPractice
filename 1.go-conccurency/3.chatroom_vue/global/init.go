package global

import (
	"os"
	"path/filepath"
	"sync"
)

// RootDir : root dir
var RootDir string

var once = new(sync.Once)

// Init : init
func Init() {
	once.Do(func() {
		inferRootDir()
		initConfig()
	})
}

// inferRootDir 递归找出路径
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + "/template") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

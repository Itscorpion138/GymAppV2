package main

import "os"

func ensureUploadsDir(path string) {
	_ = os.MkdirAll(path, os.ModePerm)
}

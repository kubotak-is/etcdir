package action

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return []byte(nil), err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

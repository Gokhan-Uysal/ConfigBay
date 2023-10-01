package mapper

import (
	"os"
	"path/filepath"
)

func FilesToPaths(path string) (map[string]string, error) {
	var (
		dirEntries []os.DirEntry
		paths      = make(map[string]string)
		err        error
	)
	dirEntries, err = os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filename = new(string)
	var filePath = new(string)

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		*filename = dirEntry.Name()
		*filePath = filepath.Join(path, *filename)

		paths[*filename] = *filePath
	}

	return paths, nil
}

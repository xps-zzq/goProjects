package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
//读取zip
func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//读取zip文件
	r, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self *ZipEntry) String() string { return self.absPath }

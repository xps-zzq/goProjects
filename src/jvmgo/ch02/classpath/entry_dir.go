package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

//相当于构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
//读取class文件
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
//返回绝对路径
func (self *DirEntry) String() string {
	return self.absDir
}

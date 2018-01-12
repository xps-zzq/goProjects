package classfile

/**这2个标记长度为0 在Class文件中也仅仅作为标记使用*/
/*@Deprecated*/
type DeprecatedAttribute struct{ MarkerAttribute }

/*@Synthetic*/
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

package classfile

type ConstantInvokeDynamicInfo struct {
	tag                      uint8
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.tag = reader.readUint8()
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

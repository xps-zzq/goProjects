package classfile

type ConstantMethodTypeInfo struct {
	tag             uint8
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.tag = reader.readUint8()
	self.descriptorIndex = reader.readUint16()
}

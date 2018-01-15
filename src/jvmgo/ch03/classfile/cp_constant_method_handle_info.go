package classfile

type ConstantMethodHandleInfo struct {
	tag            uint8
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.tag = reader.readUint8()
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

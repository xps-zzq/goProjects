package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}
//函数读取字段或方法数据
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}
//函数读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp), //  见 3.4
	}
}
func (self *MemberInfo) AccessFlags() uint16 { ... } // getter
func (self *MemberInfo) Name() string        {
	return self.cp.getUtf8(self.nameIndex)
	}
//从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string  {
	return self.cp.getUtf8(self.descriptorIndex)
}

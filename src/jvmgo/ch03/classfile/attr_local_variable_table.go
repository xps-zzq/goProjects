package classfile

type LocalVariableTable struct {
	lineNumberTable []*LocalVariableTable
}

type LocalVariableTableEntry struct {
	startPc        uint16
	length         uint32
	nameIndex      uint16
	name           uint16
	signatureIndex uint16
}

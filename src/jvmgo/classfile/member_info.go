package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
    accessFlags     uint16
    nameIndex       uint16
    descriptorIndex uint16
    AttributeTable
}

type FieldInfo struct {
    MemberInfo
}

type MethodInfo struct {
    MemberInfo
}

func (self *MemberInfo) read(reader *ClassReader, cp *ConstantPool) {
    self.accessFlags = reader.readUint16()
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
    self.attributes = readAttributes(reader, cp)
}

func (self *MemberInfo) AccessFlags() (uint16) {
    return self.accessFlags
}
func (self *MemberInfo) GetName(cp *ConstantPool) (string) {
    return cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) GetDescriptor(cp *ConstantPool) (string) {
    return cp.getUtf8(self.descriptorIndex)
}
package bankv2alpha1

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
)

var (
	md_Balance         protoreflect.MessageDescriptor
	fd_Balance_address protoreflect.FieldDescriptor
	fd_Balance_denom   protoreflect.FieldDescriptor
	fd_Balance_amount  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_bank_v2alpha1_state_proto_init()
	md_Balance = File_cosmos_bank_v2alpha1_state_proto.Messages().ByName("Balance")
	fd_Balance_address = md_Balance.Fields().ByName("address")
	fd_Balance_denom = md_Balance.Fields().ByName("denom")
	fd_Balance_amount = md_Balance.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_Balance)(nil)

type fastReflection_Balance Balance

func (x *Balance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Balance)(x)
}

func (x *Balance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v2alpha1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Balance_messageType fastReflection_Balance_messageType
var _ protoreflect.MessageType = fastReflection_Balance_messageType{}

type fastReflection_Balance_messageType struct{}

func (x fastReflection_Balance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Balance)(nil)
}
func (x fastReflection_Balance_messageType) New() protoreflect.Message {
	return new(fastReflection_Balance)
}
func (x fastReflection_Balance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Balance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Balance) Descriptor() protoreflect.MessageDescriptor {
	return md_Balance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Balance) Type() protoreflect.MessageType {
	return _fastReflection_Balance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Balance) New() protoreflect.Message {
	return new(fastReflection_Balance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Balance) Interface() protoreflect.ProtoMessage {
	return (*Balance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Balance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Address) != 0 {
		value := protoreflect.ValueOfBytes(x.Address)
		if !f(fd_Balance_address, value) {
			return
		}
	}
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_Balance_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_Balance_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Balance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		return len(x.Address) != 0
	case "cosmos.bank.v2alpha1.Balance.denom":
		return x.Denom != ""
	case "cosmos.bank.v2alpha1.Balance.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Balance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		x.Address = nil
	case "cosmos.bank.v2alpha1.Balance.denom":
		x.Denom = ""
	case "cosmos.bank.v2alpha1.Balance.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Balance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		value := x.Address
		return protoreflect.ValueOfBytes(value)
	case "cosmos.bank.v2alpha1.Balance.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "cosmos.bank.v2alpha1.Balance.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Balance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		x.Address = value.Bytes()
	case "cosmos.bank.v2alpha1.Balance.denom":
		x.Denom = value.Interface().(string)
	case "cosmos.bank.v2alpha1.Balance.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Balance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		panic(fmt.Errorf("field address of message cosmos.bank.v2alpha1.Balance is not mutable"))
	case "cosmos.bank.v2alpha1.Balance.denom":
		panic(fmt.Errorf("field denom of message cosmos.bank.v2alpha1.Balance is not mutable"))
	case "cosmos.bank.v2alpha1.Balance.amount":
		panic(fmt.Errorf("field amount of message cosmos.bank.v2alpha1.Balance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Balance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Balance.address":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.bank.v2alpha1.Balance.denom":
		return protoreflect.ValueOfString("")
	case "cosmos.bank.v2alpha1.Balance.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Balance"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Balance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Balance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v2alpha1.Balance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Balance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Balance) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Balance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Balance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Balance)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Balance)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Balance)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Balance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = append(x.Address[:0], dAtA[iNdEx:postIndex]...)
				if x.Address == nil {
					x.Address = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Supply        protoreflect.MessageDescriptor
	fd_Supply_denom  protoreflect.FieldDescriptor
	fd_Supply_supply protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_bank_v2alpha1_state_proto_init()
	md_Supply = File_cosmos_bank_v2alpha1_state_proto.Messages().ByName("Supply")
	fd_Supply_denom = md_Supply.Fields().ByName("denom")
	fd_Supply_supply = md_Supply.Fields().ByName("supply")
}

var _ protoreflect.Message = (*fastReflection_Supply)(nil)

type fastReflection_Supply Supply

func (x *Supply) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Supply)(x)
}

func (x *Supply) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v2alpha1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Supply_messageType fastReflection_Supply_messageType
var _ protoreflect.MessageType = fastReflection_Supply_messageType{}

type fastReflection_Supply_messageType struct{}

func (x fastReflection_Supply_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Supply)(nil)
}
func (x fastReflection_Supply_messageType) New() protoreflect.Message {
	return new(fastReflection_Supply)
}
func (x fastReflection_Supply_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Supply
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Supply) Descriptor() protoreflect.MessageDescriptor {
	return md_Supply
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Supply) Type() protoreflect.MessageType {
	return _fastReflection_Supply_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Supply) New() protoreflect.Message {
	return new(fastReflection_Supply)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Supply) Interface() protoreflect.ProtoMessage {
	return (*Supply)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Supply) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_Supply_denom, value) {
			return
		}
	}
	if x.Supply != "" {
		value := protoreflect.ValueOfString(x.Supply)
		if !f(fd_Supply_supply, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Supply) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		return x.Denom != ""
	case "cosmos.bank.v2alpha1.Supply.supply":
		return x.Supply != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supply) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		x.Denom = ""
	case "cosmos.bank.v2alpha1.Supply.supply":
		x.Supply = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Supply) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "cosmos.bank.v2alpha1.Supply.supply":
		value := x.Supply
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supply) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		x.Denom = value.Interface().(string)
	case "cosmos.bank.v2alpha1.Supply.supply":
		x.Supply = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supply) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		panic(fmt.Errorf("field denom of message cosmos.bank.v2alpha1.Supply is not mutable"))
	case "cosmos.bank.v2alpha1.Supply.supply":
		panic(fmt.Errorf("field supply of message cosmos.bank.v2alpha1.Supply is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Supply) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v2alpha1.Supply.denom":
		return protoreflect.ValueOfString("")
	case "cosmos.bank.v2alpha1.Supply.supply":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v2alpha1.Supply"))
		}
		panic(fmt.Errorf("message cosmos.bank.v2alpha1.Supply does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Supply) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v2alpha1.Supply", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Supply) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supply) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Supply) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Supply) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Supply)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Supply)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Supply)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Supply) > 0 {
			i -= len(x.Supply)
			copy(dAtA[i:], x.Supply)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Supply)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Supply)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Supply: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Supply = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/bank/v2alpha1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount  string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v2alpha1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v2alpha1_state_proto_rawDescGZIP(), []int{0}
}

func (x *Balance) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Balance) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Balance) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type Supply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Supply string `protobuf:"bytes,2,opt,name=supply,proto3" json:"supply,omitempty"`
}

func (x *Supply) Reset() {
	*x = Supply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v2alpha1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supply) ProtoMessage() {}

// Deprecated: Use Supply.ProtoReflect.Descriptor instead.
func (*Supply) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v2alpha1_state_proto_rawDescGZIP(), []int{1}
}

func (x *Supply) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Supply) GetSupply() string {
	if x != nil {
		return x.Supply
	}
	return ""
}

var File_cosmos_bank_v2alpha1_state_proto protoreflect.FileDescriptor

var file_cosmos_bank_v2alpha1_state_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x24, 0xf2, 0x9e, 0xd3, 0x8e,
	0x03, 0x1e, 0x0a, 0x0f, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2c, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01, 0x18, 0x01,
	0x22, 0x49, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x3a, 0x11, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x0b,
	0x0a, 0x07, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x42, 0xdc, 0x01, 0x0a, 0x18,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x62, 0x61,
	0x6e, 0x6b, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58,
	0xaa, 0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x56,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x20, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x6e, 0x6b,
	0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cosmos_bank_v2alpha1_state_proto_rawDescOnce sync.Once
	file_cosmos_bank_v2alpha1_state_proto_rawDescData = file_cosmos_bank_v2alpha1_state_proto_rawDesc
)

func file_cosmos_bank_v2alpha1_state_proto_rawDescGZIP() []byte {
	file_cosmos_bank_v2alpha1_state_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_v2alpha1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_v2alpha1_state_proto_rawDescData)
	})
	return file_cosmos_bank_v2alpha1_state_proto_rawDescData
}

var file_cosmos_bank_v2alpha1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_bank_v2alpha1_state_proto_goTypes = []interface{}{
	(*Balance)(nil), // 0: cosmos.bank.v2alpha1.Balance
	(*Supply)(nil),  // 1: cosmos.bank.v2alpha1.Supply
}
var file_cosmos_bank_v2alpha1_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_bank_v2alpha1_state_proto_init() }
func file_cosmos_bank_v2alpha1_state_proto_init() {
	if File_cosmos_bank_v2alpha1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_v2alpha1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_bank_v2alpha1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_bank_v2alpha1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_bank_v2alpha1_state_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_v2alpha1_state_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_v2alpha1_state_proto_msgTypes,
	}.Build()
	File_cosmos_bank_v2alpha1_state_proto = out.File
	file_cosmos_bank_v2alpha1_state_proto_rawDesc = nil
	file_cosmos_bank_v2alpha1_state_proto_goTypes = nil
	file_cosmos_bank_v2alpha1_state_proto_depIdxs = nil
}

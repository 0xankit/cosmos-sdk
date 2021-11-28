package ormtable

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	io "io"
	"math"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/backend/kv"
	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
	"github.com/cosmos/cosmos-sdk/orm/model/ormindex"
)

type TableImpl struct {
	*ormindex.PrimaryKey
	msgType               protoreflect.MessageType
	indexers              []ormindex.Indexer
	indexes               []ormindex.Index
	indexesByFields       map[ormkv.Fields]ormindex.Index
	uniqueIndexesByFields map[ormkv.Fields]ormindex.UniqueIndex
	indexesById           map[uint32]ormindex.Index
	tablePrefix           []byte
	typeResolver          TypeResolver
	customValidator       func(proto.Message) error
}

func (t TableImpl) MessageType() protoreflect.MessageType {
	return t.msgType
}

type TypeResolver interface {
	protoregistry.MessageTypeResolver
	protoregistry.ExtensionTypeResolver
}

func (t TableImpl) Save(store kv.IndexCommitmentStore, message proto.Message, mode SaveMode) error {
	mref := message.ProtoReflect()
	pkValues, pk, err := t.EncodeFromMessage(mref)
	if err != nil {
		return err
	}

	existing := mref.New().Interface()
	haveExisting, err := t.GetByKeyBytes(store, pk, pkValues, existing)
	if err != nil {
		return err
	}

	if haveExisting {
		if mode == SAVE_MODE_CREATE {
			return ormerrors.PrimaryKeyConstraintViolation.Wrapf("%q", mref.Descriptor().FullName())
		}
	} else {
		if mode == SAVE_MODE_UPDATE {
			return ormerrors.NotFoundOnUpdate.Wrapf("%q", mref.Descriptor().FullName())
		}
	}

	// temporarily clear primary key
	t.ClearValues(mref)

	// store object
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(message)
	err = store.CommitmentStore().Set(pk, bz)
	if err != nil {
		return err
	}

	// set primary key again

	t.SetValues(mref, pkValues)

	// set indexes
	indexStore := store.IndexStore()
	if !haveExisting {
		for _, idx := range t.indexers {
			err = idx.OnCreate(indexStore, mref)
			if err != nil {
				return err
			}

		}
	} else {
		existingMref := existing.ProtoReflect()
		for _, idx := range t.indexers {
			err = idx.OnUpdate(indexStore, mref, existingMref)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t TableImpl) Delete(store kv.IndexCommitmentStore, primaryKey []protoreflect.Value) error {
	pk, err := t.Encode(primaryKey)
	if err != nil {
		return err
	}

	msg := t.Type.New().Interface()
	found, err := t.GetByKeyBytes(store, pk, primaryKey, msg)
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	// delete object
	err = store.CommitmentStore().Delete(pk)
	if err != nil {
		return err
	}

	// clear indexes
	mref := msg.ProtoReflect()
	indexStore := store.IndexStore()
	for _, idx := range t.indexers {
		err := idx.OnDelete(indexStore, mref)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t TableImpl) GetIndex(fields ormkv.Fields) ormindex.Index {
	return t.indexesByFields[fields]
}

func (t TableImpl) GetUniqueIndex(fields ormkv.Fields) ormindex.UniqueIndex {
	return t.uniqueIndexesByFields[fields]
}

func (t TableImpl) Indexes() []ormindex.Index {
	return t.indexes
}

func (t TableImpl) DefaultJSON() json.RawMessage {
	return json.RawMessage("[]")
}

func (t TableImpl) decodeJson(reader io.Reader, onMsg func(message proto.Message) error) error {
	decoder := json.NewDecoder(reader)
	token, err := decoder.Token()
	if err != nil {
		return err
	}

	if token != json.Delim('[') {
		return ormerrors.JSONImportError.Wrapf("expected [ got %s", token)
	}

	for decoder.More() {
		var rawJson json.RawMessage
		err = decoder.Decode(&rawJson)
		if err != nil {
			return ormerrors.JSONImportError.Wrapf("%s", err)
		}

		msg := t.msgType.New().Interface()
		err = protojson.UnmarshalOptions{Resolver: t.typeResolver}.Unmarshal(rawJson, msg)
		if err != nil {
			return err
		}

		err = onMsg(msg)
		if err != nil {
			return err
		}
	}

	if token != json.Delim(']') {
		return ormerrors.JSONImportError.Wrapf("expected ] got %s", token)
	}

	return nil
}

func DefaultValidator(message proto.Message) error {
	if v, ok := message.(interface{ ValidateBasic() error }); ok {
		err := v.ValidateBasic()
		if err != nil {
			return err
		}
	}

	if v, ok := message.(interface{ Validate() error }); ok {
		err := v.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t TableImpl) ValidateJSON(reader io.Reader) error {
	return t.decodeJson(reader, func(message proto.Message) error {
		if t.customValidator != nil {
			return t.customValidator(message)
		} else {
			return DefaultValidator(message)
		}
	})
}

func (t TableImpl) ImportJSON(store kv.IndexCommitmentStore, reader io.Reader) error {
	return t.decodeJson(reader, func(message proto.Message) error {
		return t.Save(store, message, SAVE_MODE_DEFAULT)
	})
}

func (t TableImpl) ExportJSON(store kv.IndexCommitmentReadStore, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(json.Delim('['))
	if err != nil {
		return err
	}

	it, _ := t.PrefixIterator(store, nil, ormindex.IteratorOptions{})
	for {
		found, err := it.Next()
		if err != nil {
			return err
		}

		if found {
			err := encoder.Encode(json.Delim(']'))
			if err != nil {
				return err
			}
		}

		msg := t.msgType.New().Interface()
		err = it.GetMessage(msg)
		if err != nil {
			return err
		}

		bz, err := protojson.Marshal(msg)

		err = encoder.Encode(json.RawMessage(bz))
		if err != nil {
			return err
		}
	}
}

func (t TableImpl) DecodeKV(k, v []byte) (ormkv.Entry, error) {
	r := bytes.NewReader(k)
	if bytes.HasPrefix(k, t.tablePrefix) {
		err := ormkv.SkipPrefix(r, t.tablePrefix)
		if err != nil {
			return nil, err
		}

		id, err := binary.ReadUvarint(r)
		if err != nil {
			return nil, err
		}

		if id == 0 {
			return t.PrimaryKey.DecodeKV(k, v)
		}

		if id > math.MaxUint32 {
			return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("uint32 varint id out of range %d", id)
		}

		idx, ok := t.indexesById[uint32(id)]
		if !ok {
			return nil, ormerrors.UnexpectedDecodePrefix.Wrapf("can't find field with id %d", id)
		}

		return idx.DecodeKV(k, v)
	} else {
		return nil, ormerrors.UnexpectedDecodePrefix
	}
}

func (t TableImpl) EncodeKV(entry ormkv.Entry) (k, v []byte, err error) {
	switch entry := entry.(type) {
	case ormkv.PrimaryKeyEntry:
		return t.PrimaryKey.EncodeKV(entry)
	case ormkv.IndexKeyEntry:
		idx, ok := t.indexesByFields[entry.Fields]
		if !ok {
			return nil, nil, ormerrors.BadDecodeEntry.Wrapf("can't find index with fields %s", entry.Fields)
		}

		return idx.EncodeKV(entry)
	default:
		return nil, nil, ormerrors.BadDecodeEntry.Wrapf("%s", entry)
	}
}

var _ Table = &TableImpl{}

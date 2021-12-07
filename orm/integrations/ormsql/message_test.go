package ormsql

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/internal/testpb"
)

func TestMessageCodec(t *testing.T) {
	gormDb, err := gorm.Open(sqlite.Open("file:test.sqlite"), &gorm.Config{})
	assert.NilError(t, err)
	sch := db{
		gormDb: gormDb,
		schema: &schema{
			jsonMarshalOptions:   protojson.MarshalOptions{},
			jsonUnmarshalOptions: protojson.UnmarshalOptions{},
			resolver:             protoregistry.GlobalTypes,
			messageCodecs:        map[protoreflect.FullName]*messageCodec{},
		},
		migratedMsgCodecs: map[protoreflect.FullName]*messageCodec{},
	}

	x := &testpb.A{U32: 7, I32: 4,
		Map:      map[string]uint32{"abc": 4},
		Msg:      &testpb.B{X: "foo"},
		Repeated: []uint32{1, 2, 4, 7, 9},
		Sum:      &testpb.A_Oneof{Oneof: 3},
	}
	y := &testpb.A{U32: 12, I32: 8,
		Map:      map[string]uint32{"wah": 37},
		Msg:      &testpb.B{X: "bar"},
		Repeated: []uint32{2, 4, 7, 9, 1},
		Sum:      &testpb.A_Oneof{Oneof: 12},
	}
	origJson, err := protojson.Marshal(x)
	t.Logf("orig:%s", origJson)

	sch.Save(x)
	assert.NilError(t, sch.Error())
	sch.Save(y)
	assert.NilError(t, sch.Error())

	var x2 testpb.A
	sch.First(&x2)
	assert.NilError(t, sch.Error())
	decodedJson, err := protojson.Marshal(&x2)
	assert.NilError(t, err)
	t.Logf("decoded:%s", decodedJson)
	var xs []*testpb.A
	sch.Where("i32 = ?", 8).Find(&xs)
	assert.NilError(t, sch.Error())
	for i := 0; i < len(xs); i++ {
		t.Logf("xs[%d]:%+v", i, xs[i])
	}
}

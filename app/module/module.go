package module

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/types/descriptorpb"

	appv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/app/v1alpha1"
	"github.com/cosmos/cosmos-sdk/app/module/internal"
)

func Register(configType proto.Message, options ...Option) {
	desc := configType.ProtoReflect().Descriptor()
	modDesc := proto.GetExtension(desc.Options(), appv1alpha1.E_IsModule).(*appv1alpha1.ModuleDescriptor)
	if modDesc == nil {
		panic(fmt.Errorf("protobuf type %s must have the %s extension to be used as an app module config type",
			desc.FullName(), appv1alpha1.E_IsModule.TypeDescriptor().FullName()))
	}

	if modDesc.GoImport == "" {
		fileOpts := desc.ParentFile().Options().(*descriptorpb.FileOptions)
		if fileOpts.GoPackage == nil || *fileOpts.GoPackage == "" {
			panic(fmt.Errorf(
				"either go_package must be set on the file containing %s or the ModuleDescriptor.go_import field must be set",
				desc.FullName(),
			))
		}
	}

	config := &internal.ModuleConfig{
		Params: map[reflect.Type]internal.ModuleParamType{},
	}

	for _, opt := range options {
		err := opt.apply(config)
		if err != nil {
			panic(err)
		}
	}

	internal.ModuleRegistry[desc.FullName()] = config
}

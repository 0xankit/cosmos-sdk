package example

import (
	"context"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	types "github.com/cosmos/cosmos-sdk/store/v2alpha1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/capability"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/core/appmodule"
)

func init() {
	// register the module with the app-wiring dependency injection framework
	appmodule.Register(&Module{},
		appmodule.Provide(provideKeeper,
			func(k keeper) *appmodule.Handler {
				h := &appmodule.Handler{}
				RegisterMsgServer(h, k)
				RegisterQueryServer(h, k)
				return h
			}),
	)
}

func provideKeeper(storeKey storetypes.KVStoreKey) keeper {
	return keeper{
		contextFactory: sdk.NewModuleContextFactory[ExampleContext](),
		kvStoreKey:     storeKey,
	}
}

// the module's dependency injection inputs
type keeper struct {
	contextFactory sdk.ModuleContextFactory[ExampleContext]
	kvStoreKey     types.KVStoreKey
}

const (
	nameInfoPrefix byte = iota
)

func nameInfoKey(name string) []byte {
	return append([]byte{nameInfoPrefix}, name...)
}

type ExampleContext interface {
	sdk.ModuleContext
	capability.BlockInfoService
	capability.KVStoreService
	capability.EventService
}

// implement MsgServer
func (s keeper) RegisterName(ctx context.Context, msg *MsgRegisterName) (*MsgRegisterNameResponse, error) {
	moduleContext := s.contextFactory.Make(ctx)

	kvStore := moduleContext.KVStoreService(s.kvStoreKey).Open(ctx)
	key := nameInfoKey(msg.Name)
	if kvStore.Has(key) {
		return nil, status.Error(codes.AlreadyExists, "name already registered")
	}

	blockInfoService := moduleContext.BlockInfoService()
	height := blockInfoService.GetBlockInfo(ctx).Height()
	bz, err := proto.Marshal(&NameInfo{
		Owner:            msg.Sender,
		RegisteredHeight: height,
	})
	if err != nil {
		return nil, err
	}

	kvStore.Set(key, bz)
	err = moduleContext.EventService().GetManager(ctx).Emit(&EventRegisterName{
		Name:  msg.Name,
		Owner: msg.Sender,
	})
	return &MsgRegisterNameResponse{}, err
}

// implement QueryServer
func (s keeper) Name(ctx context.Context, request *QueryNameRequest) (*QueryNameResponse, error) {
	moduleContext := s.contextFactory.Make(ctx)

	kvStore := moduleContext.KVStoreService(s.kvStoreKey).Open(ctx)
	key := nameInfoKey(request.Name)
	bz := kvStore.Get(key)
	if bz == nil {
		return nil, status.Error(codes.NotFound, "name not found")
	}

	var info NameInfo
	err := proto.Unmarshal(bz, &info)
	if err != nil {
		return nil, err
	}

	return &QueryNameResponse{Info: &info}, nil
}

func (s keeper) mustEmbedUnimplementedMsgServer()   {}
func (s keeper) mustEmbedUnimplementedQueryServer() {}

var _ MsgServer = keeper{}
var _ QueryServer = keeper{}

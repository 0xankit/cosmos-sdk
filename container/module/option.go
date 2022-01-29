package module

import (
	"github.com/cosmos/cosmos-sdk/container"
	"github.com/cosmos/cosmos-sdk/container/module/internal"
)

type Option interface {
	apply(*internal.ModuleConfig) error
}

func Provide(providers ...interface{}) Option {
	return option(func(config *internal.ModuleConfig) error {
		for _, provider := range providers {
			desc, err := container.ExtractProviderDescriptor(provider)
			if err != nil {
				return err
			}

			config.Providers = append(config.Providers, desc)
		}
		return nil
	})
}

func Supply(values ...interface{}) Option {
	return option(func(config *internal.ModuleConfig) error {
		panic("TODO")
	})
}

type ParamType = internal.ModuleParamType

func DefineParam(params ...ParamType) Option {
	return option(func(config *internal.ModuleConfig) error {
		panic("TODO")
	})
}

type option func(*internal.ModuleConfig) error

func (o option) apply(opts *internal.ModuleConfig) error {
	return o(opts)
}

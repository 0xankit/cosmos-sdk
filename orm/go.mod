module github.com/cosmos/cosmos-sdk/orm

go 1.17

require (
	github.com/cosmos/cosmos-proto v1.0.0-alpha4
	github.com/cosmos/cosmos-sdk/api v0.1.0-alpha1
	github.com/cosmos/cosmos-sdk/errors v0.0.0
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/protobuf v1.27.1
	gotest.tools/v3 v3.0.3
	pgregory.net/rapid v0.4.7
)

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/confio/ics23/go v0.6.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/graphql-go/graphql v0.8.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sasha-s/go-deadlock v0.2.1-0.20190427202633-1595213edefa // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tendermint/tendermint v0.34.14 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/datatypes v1.0.4 // indirect
	gorm.io/driver/mysql v1.2.1 // indirect
	gorm.io/gorm v1.22.4 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk/api => ../api

replace github.com/cosmos/cosmos-sdk/errors => ../errors

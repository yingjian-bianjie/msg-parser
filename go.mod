module github.com/kaifei-bianjie/msg-parser

go 1.15

require (
	github.com/bianjieai/iritamod v1.2.1-0.20220809083453-a38b88e9aac1
	github.com/bianjieai/spartan-cosmos v1.0.1
	github.com/bianjieai/tibc-go v0.2.0-alpha
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/irismod v1.6.0
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/stretchr/testify v1.7.5
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20210908054213-781a5fed16d6
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.6.1-0.20220808105048-55bdb6a8f178
)

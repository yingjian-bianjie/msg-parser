package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/tendermint/tendermint/types"
)

var (
	appModules []module.AppModuleBasic
	encodecfg  params.EncodingConfig
)

// 初始化账户地址前缀
func MakeEncodingConfig() {
	encodingConfig := params.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics := module.NewBasicManager(appModules...)
	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	encodecfg = encodingConfig
}

func GetTxDecoder() sdk.TxDecoder {
	return encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Codec {
	return encodecfg.Marshaler
}

func GetSigningTx(txBytes types.Tx) (signing.Tx, error) {
	Tx, err := GetTxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	signingTx := Tx.(signing.Tx)
	return signingTx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	appModules = append(appModules, module...)
}

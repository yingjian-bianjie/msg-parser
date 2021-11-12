package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestTibc() {
	cases := []SubTest{
		{
			"NftTransfer",
			NftTransfer,
		}, {
			"RecvPacket",
			RecvPacket,
		}, {
			"AcknowledgePacket",
			AcknowledgePacket,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func NftTransfer(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0aaa010aa7010a292f746962632e617070732e6e66745f7472616e736665722e76312e4d73674e66745472616e73666572127a0a046c616e671204693030311a2a69616131636871306e637434353066336e3974767a6a3234706563796335337333776664773334307735222a69616131347361786e6c7664667463647037737539366335773932386e66767570367638667673646c6a2a09746962632d746573743209746962632d7465737412640a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102585b0763171ff17f3c1d0ccb906fac6591169df6ec2bf17ef7cd136001f29eea12040a020801181312100a0a0a057374616b6512013210c09a0c1a406f0a3cebbf23cd32480f0dd4becd7ffa28d1c68c92e851c3db32582c7b5dd34d3ccd037da0fee0edf3320436b019996c016d6e78467dfda59460338a9f3e6d46")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Tibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func RecvPacket(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0ade050adb050a222f746962632e636f72652e7061636b65742e76312e4d7367526563765061636b657412b4050a8901080412034e46541a0a746962632d74657374312209746962632d7465737432690a0574657374311206733232337931222a69616131777577726a77746c76797736366e703530717367766867797a36777566327339337432766b382a2a6961613134613436396b70713476797978786d6877676c6574773432767930633476783068356c787364300112f4030ac1020abe020a2c636f6d6d69746d656e74732f746962632d74657374312f746962632d746573742f73657175656e6365732f341220e2de2e3b50d049d2c401678b700c1679556d5079e6e51c95fe4386cb5e4f7bcd1a0c0801180120012a0400029419222a080112260204941920a41a78fe238aae077475e154e90b285d7a46ccbdeaef56d4d6c810efded21c5620222c0801120504069419201a2120acea3fef9fd556593a2fb5374bae5913c006afe2cc53cbc6fdf47ff3ba764c98222c08011205060a9419201a2120faa8ba33a7f0f3bbcf3989a7c3284a716279566642d2cef571b3e35b49e4251a222a0801122608169419200c6f4c6553b0d76bb52a3412a4dea09a453a1929be011d60fbae3a884631500720222a080112260a2e941920c9245187ebdb5a0f80d918ef257b7560466247f192eae4db2a3e54db1e0ace41200aad010aaa010a047469626312203362fa7095860f2e060fd1a443a6ab541c03bc71ba8c8b6d5c9cd4cc2e70b2e21a090801180120012a0100222708011201011a20c9c8849ed125cc7681329c4d27b83b1fc8acf7a865c9d1d1df575cca56f48dbe22250801122101b1fa483555c512177f8a63e199748a3666b25c8960325cf411fd0108ff65c012222508011221013c80938e278137394006b8240ff425818c7d33cf9b240015df0927fc20a918e01a0310cb0c222a69616131636871306e637434353066336e3974767a6a323470656379633533733377666477333430773512660a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102585b0763171ff17f3c1d0ccb906fac6591169df6ec2bf17ef7cd136001f29eea12040a020801182112120a0c0a057374616b65120331303010c09a0c1a40021bcd3a6d5e4ea217e9d8e63128961f1eacc9a71d2eb81aa8931e2ebb9f52791872b7b2e26d5e2f86cc20477d9c56a7dc45d533bc1979ed2769707f8fdf3471")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Tibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func AcknowledgePacket(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0aa7070aa4070a272f746962632e636f72652e7061636b65742e76312e4d736741636b6e6f776c656467656d656e7412f8060ae201080412034e46541a0f697269736875622d6d61696e6e6574220e62736e6875622d6d61696e6e657432b7010a2d6e66742f62736e6875622d6d61696e6e65742f697269736875622d6d61696e6e65742f6e667464656e6f6d696412076e6674696432311a25687474703a2f2f7777772e62656a736f6e2e636f6d2f6a736f6e7669657765726e65772f39222a696161316571766b667468747272393367347039717370703534773664746a74726e32376172377270772a2a696161316c667a74666333347972646a30666b667a616a3932797667676d75616c6e776c3965613739651204aa0101011ad8040afc020af9020a2f61636b732f697269736875622d6d61696e6e65742f62736e6875622d6d61696e6e65742f73657175656e6365732f341220e2e240ed1d7b1ee6be77e9101b573c90800cf8d61d6eff892f9d7d987ccc33831a0d0801180120012a050002c4e202222b080112270204c4e202203aa49ff8b2c7f3836e2b38eb277bea2ce378ab3b13c49aba56dd6f61dcc73de620222b080112270408c4e202203d3903763c4f06043a75888876da2129d7cc4450010219e7076caf94666ce05f20222d08011206060ec4e202201a212063724e184df772142ddf3fedcbafeb86d3c507ea33daab7c9de7a6f6ef074b6f222d08011206081cc4e202201a2120acb683c2b152745a89a85292e76f09bb68af0fbca10903cca2fde26f1b68b99f222d080112060c50c4e202201a21205b3529034ab3dc2f992c2be3f69ca96be1a1d383222322a3cc31be49b74abbd0222e080112070e8401c4e202201a212058f962aeda2b68373719416eb0e35f8b0e4ade909070802680f1dc2dccd2278e0ad6010ad3010a0474696263122037e83555328ed529ae2f234e3e61373673190cd821233d6c31bd289c20d1d2e41a090801180120012a01002225080112210111662a7e092533cb81a2cca4887dacec1e502135a7387fe5fdebcf724fce9516222708011201011a206966c029f03f1f7f11e8dc63e23324f859070e7d0ab68d440ddf02c59a6c5a22222708011201011a20d0fe1707f3cf6b7d6a350c8530518ab638ebe716eb08641c49d7e070b3a984662225080112210186f60a189580aa3fc05823a115b33c32f2f114c4b4077d91b44490456047e64b220410a3b1012a2a69616131333373706e3072717273363678716c687668356a6c6675773233387239757139363677756e7a12680a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103e8c4cc19bd867361b3e4d5a89e6fc3a2fe84acbb350456f02802cbdb4fe489c712040a020801184212140a0e0a05756e79616e1205313030303010c09a0c1a4065b11fc6fbc9c908d152f650a30caa86b264341e40171c74c5605f8991ab1fe42f123f152bd4f956fa1a85135b13c28ce2627731c1b4fd82dae321bcc427a44b")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Tibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

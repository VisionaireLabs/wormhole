package sdk

import "github.com/wormhole-foundation/wormhole/sdk/vaa"

// KnownTestnetEmitters is a list of known emitters on the various L1 testnets.
var KnownTestnetEmitters = buildKnownEmitters(knownTestnetTokenbridgeEmitters, knownTestnetNFTBridgeEmitters)

// KnownTestnetTokenbridgeEmitters is a map of known tokenbridge emitters on the various L1 testnets.
var KnownTestnetTokenbridgeEmitters = buildEmitterMap(knownTestnetTokenbridgeEmitters)
var knownTestnetTokenbridgeEmitters = map[vaa.ChainID]string{
	vaa.ChainIDSolana:          "3b26409f8aaded3f5ddca184695aa6a0fa829b0c85caf84856324896d214ca98",
	vaa.ChainIDEthereum:        "000000000000000000000000f890982f9310df57d00f659cf4fd87e65aded8d7",
	vaa.ChainIDTerra:           "0000000000000000000000000c32d68d8f22613f6b9511872dad35a59bfdf7f0",
	vaa.ChainIDTerra2:          "c3d4c6c2bcba163de1defb7e8f505cdb40619eee4fa618678955e8790ae1448d",
	vaa.ChainIDBSC:             "0000000000000000000000009dcf9d205c9de35334d646bee44b2d2859712a09",
	vaa.ChainIDPolygon:         "000000000000000000000000377D55a7928c046E18eEbb61977e714d2a76472a",
	vaa.ChainIDAvalanche:       "00000000000000000000000061e44e506ca5659e6c0bba9b678586fa2d729756",
	vaa.ChainIDOasis:           "00000000000000000000000088d8004a9bdbfd9d28090a02010c19897a29605c",
	vaa.ChainIDAlgorand:        "6241ffdc032b693bfb8544858f0403dec86f2e1720af9f34f8d65fe574b6238c",
	vaa.ChainIDAptos:           "0000000000000000000000000000000000000000000000000000000000000001",
	vaa.ChainIDAurora:          "000000000000000000000000d05ed3ad637b890d68a854d607eeaf11af456fba",
	vaa.ChainIDFantom:          "000000000000000000000000599cea2204b4faecd584ab1f2b6aca137a0afbe8",
	vaa.ChainIDKarura:          "000000000000000000000000e157115ef34c93145Fec2FE53706846853B07F42",
	vaa.ChainIDAcala:           "000000000000000000000000e157115ef34c93145Fec2FE53706846853B07F42",
	vaa.ChainIDKlaytn:          "000000000000000000000000c7a13be098720840dea132d860fdfa030884b09a",
	vaa.ChainIDCelo:            "00000000000000000000000005ca6037ec51f8b712ed2e6fa72219feae74e153",
	vaa.ChainIDNear:            "c2c0b6ecbbe9ecf91b2b7999f0264018ba68126c2e83bf413f59f712f3a1df55",
	vaa.ChainIDMoonbeam:        "000000000000000000000000bc976d4b9d57e57c3ca52e1fd136c45ff7955a96",
	vaa.ChainIDArbitrum:        "00000000000000000000000023908A62110e21C04F3A4e011d24F901F911744A",
	vaa.ChainIDOptimism:        "000000000000000000000000C7A204bDBFe983FCD8d8E61D02b475D4073fF97e",
	vaa.ChainIDXpla:            "b66da121bd3621c8d2604c08c82965640fe682d606af26a302ee09094f5e62cf",
	vaa.ChainIDInjective:       "00000000000000000000000003f3e7b2e363f51cf6e57ef85f43a2b91dbce501",
	vaa.ChainIDSui:             "40440411a170b4842ae7dee4f4a7b7a58bc0a98566e998850a7bb87bf5dc05b9",
	vaa.ChainIDBase:            "000000000000000000000000A31aa3FDb7aF7Db93d18DDA4e19F811342EDF780",
	vaa.ChainIDSei:             "9328673cb5de3fd99974cefbbd90fea033f4c59a572abfd7e1a4eebcc5d18157",
	vaa.ChainIDScroll:          "00000000000000000000000022427d90B7dA3fA4642F7025A854c7254E4e45BF",
	vaa.ChainIDMantle:          "00000000000000000000000075Bfa155a9D7A3714b0861c8a8aF0C4633c45b5D",
	vaa.ChainIDBlast:           "000000000000000000000000430855B4D43b8AEB9D2B9869B74d58dda79C0dB2",
	vaa.ChainIDXLayer:          "000000000000000000000000dA91a06299BBF302091B053c6B9EF86Eff0f930D",
	vaa.ChainIDLinea:           "000000000000000000000000C7A204bDBFe983FCD8d8E61D02b475D4073fF97e",
	vaa.ChainIDBerachain:       "000000000000000000000000a10f2eF61dE1f19f586ab8B6F2EbA89bACE63F7a",
	vaa.ChainIDSnaxchain:       "000000000000000000000000a10f2eF61dE1f19f586ab8B6F2EbA89bACE63F7a",
	vaa.ChainIDUnichain:        "000000000000000000000000a10f2eF61dE1f19f586ab8B6F2EbA89bACE63F7a",
	vaa.ChainIDWorldchain:      "000000000000000000000000430855B4D43b8AEB9D2B9869B74d58dda79C0dB2",
	vaa.ChainIDInk:             "000000000000000000000000376428e7f26D5867e69201b275553C45B09EE090",
	vaa.ChainIDHyperEVM:        "0000000000000000000000004a8bc80Ed5a4067f1CCf107057b8270E0cC11A78",
	vaa.ChainIDSepolia:         "000000000000000000000000DB5492265f6038831E89f495670FF909aDe94bd9",
	vaa.ChainIDHolesky:         "00000000000000000000000076d093BbaE4529a342080546cAFEec4AcbA59EC6",
	vaa.ChainIDArbitrumSepolia: "000000000000000000000000C7A204bDBFe983FCD8d8E61D02b475D4073fF97e",
	vaa.ChainIDBaseSepolia:     "00000000000000000000000086F55A04690fd7815A3D802bD587e83eA888B239",
	vaa.ChainIDOptimismSepolia: "00000000000000000000000099737Ec4B815d816c49A385943baf0380e75c0Ac",
	vaa.ChainIDPolygonSepolia:  "000000000000000000000000C7A204bDBFe983FCD8d8E61D02b475D4073fF97e",
	vaa.ChainIDWormchain:       "ef5251ea1e99ae48732800ccc7b83b57881232a73eb796b63b1d86ed2ea44e27",
}

// KnownTestnetNFTBridgeEmitters is a map  of known NFT emitters on the various L1 testnets.
var KnownTestnetNFTBridgeEmitters = buildEmitterMap(knownTestnetNFTBridgeEmitters)
var knownTestnetNFTBridgeEmitters = map[vaa.ChainID]string{
	vaa.ChainIDSolana:          "752a49814e40b96b097207e4b53fdd330544e1e661653fbad4bc159cc28a839e",
	vaa.ChainIDEthereum:        "000000000000000000000000d8e4c2dbdd2e2bd8f1336ea691dbff6952b1a6eb",
	vaa.ChainIDBSC:             "000000000000000000000000cd16e5613ef35599dc82b24cb45b5a93d779f1ee",
	vaa.ChainIDPolygon:         "00000000000000000000000051a02d0dcb5e52f5b92bdaa38fa013c91c7309a9",
	vaa.ChainIDAvalanche:       "000000000000000000000000d601baf2eee3c028344471684f6b27e789d9075d",
	vaa.ChainIDOasis:           "000000000000000000000000c5c25b41ab0b797571620f5204afa116a44c0eba",
	vaa.ChainIDAurora:          "0000000000000000000000008f399607e9ba2405d87f5f3e1b78d950b44b2e24",
	vaa.ChainIDFantom:          "00000000000000000000000063ed9318628d26bdcb15df58b53bb27231d1b227",
	vaa.ChainIDKlaytn:          "00000000000000000000000094c994fc51c13101062958b567e743f1a04432de",
	vaa.ChainIDCelo:            "000000000000000000000000acd8190f647a31e56a656748bc30f69259f245db",
	vaa.ChainIDMoonbeam:        "00000000000000000000000098a0f4b96972b32fcb3bd03caeb66a44a6ab9edb",
	vaa.ChainIDArbitrum:        "000000000000000000000000Ee3dB83916Ccdc3593b734F7F2d16D630F39F1D0",
	vaa.ChainIDOptimism:        "00000000000000000000000023908A62110e21C04F3A4e011d24F901F911744A",
	vaa.ChainIDBase:            "000000000000000000000000F681d1cc5F25a3694E348e7975d7564Aa581db59",
	vaa.ChainIDSepolia:         "0000000000000000000000006a0B52ac198e4870e5F3797d5B403838a5bbFD99",
	vaa.ChainIDHolesky:         "000000000000000000000000c8941d483c45eF8FB72E4d1F9dDE089C95fF8171",
	vaa.ChainIDArbitrumSepolia: "00000000000000000000000023908A62110e21C04F3A4e011d24F901F911744A",
	vaa.ChainIDBaseSepolia:     "000000000000000000000000268557122Ffd64c85750d630b716471118F323c8",
	vaa.ChainIDOptimismSepolia: "00000000000000000000000027812285fbe85BA1DF242929B906B31EE3dd1b9f",
	vaa.ChainIDPolygonSepolia:  "00000000000000000000000023908A62110e21C04F3A4e011d24F901F911744A",
}

// KnownTestnetAutomaticRelayerEmitters is a list of well-known testnet emitters for the Automatic Relayers.
// It is based on this: https://github.com/wormhole-foundation/wormhole/blob/2c9703670eadc48a7dc8967e81ed2823affcc679/sdk/js/src/relayer/consts.ts#L14
// Note that the format of this is different from the other maps because we don't want to limit it to one per chain.
var KnownTestnetAutomaticRelayerEmitters = []struct {
	ChainId vaa.ChainID
	Addr    string
}{
	{ChainId: vaa.ChainIDEthereum, Addr: "00000000000000000000000028D8F1Be96f97C1387e94A53e00eCcFb4E75175a"},
	{ChainId: vaa.ChainIDBSC, Addr: "00000000000000000000000080aC94316391752A193C1c47E27D382b507c93F3"},
	{ChainId: vaa.ChainIDPolygon, Addr: "0000000000000000000000000591C25ebd0580E0d4F27A82Fc2e24E7489CB5e0"},
	{ChainId: vaa.ChainIDAvalanche, Addr: "000000000000000000000000A3cF45939bD6260bcFe3D66bc73d60f19e49a8BB"},
	{ChainId: vaa.ChainIDCelo, Addr: "000000000000000000000000306B68267Deb7c5DfCDa3619E22E9Ca39C374f84"},
	{ChainId: vaa.ChainIDMoonbeam, Addr: "0000000000000000000000000591C25ebd0580E0d4F27A82Fc2e24E7489CB5e0"},
	{ChainId: vaa.ChainIDArbitrum, Addr: "000000000000000000000000Ad753479354283eEE1b86c9470c84D42f229FF43"},
	{ChainId: vaa.ChainIDOptimism, Addr: "00000000000000000000000001A957A525a5b7A72808bA9D10c389674E459891"},
	{ChainId: vaa.ChainIDBase, Addr: "000000000000000000000000ea8029CD7FCAEFFcD1F53686430Db0Fc8ed384E1"},
	{ChainId: vaa.ChainIDSepolia, Addr: "0000000000000000000000007B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470"},
	{ChainId: vaa.ChainIDArbitrumSepolia, Addr: "0000000000000000000000007B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470"},
	{ChainId: vaa.ChainIDOptimismSepolia, Addr: "00000000000000000000000093BAD53DDfB6132b0aC8E37f6029163E63372cEE"},
	{ChainId: vaa.ChainIDBaseSepolia, Addr: "00000000000000000000000093BAD53DDfB6132b0aC8E37f6029163E63372cEE"},
}

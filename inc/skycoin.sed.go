package main

const skycoinSedTemplate = `
# This script is used to modify !skycoin.go!

# change BlockchainPubkeyStr
s/0328c576d3f420e7682058a981173a4b374c7cc5ff55bf394d3cf57059bbe6456a/$BlockchainPubkeyStr/

# change GenesisAddressStr
s/2jBbGxZRGoQG1mqhPBnXnLTxK6oxsTf8os6/$GenesisAddressStr/

# change GenesisSignatureStr
s/eb10468d10054d15f2b6f8946cd46797779aa20a7617ceb4be884189f219bc9a164e56a5b9f7bec392a804ff3740210348d73db77a37adb542a8e08d429ac92700/$GenesisSignatureStr/

# change GenesisTimestamp
s/1426562704/$GenesisTimestamp/

# change GenesisCoinVolume
s/100e12/$GenesisCoinVolume/

# change default peers 
s/"118.178.135.93:6000"/"$nodewithport001"/
s/"47.88.33.156:6000"/"$nodewithport002"/

# delete extra default peers, only leaving two 
/"121.41.103.148/d
/"120.77.69.188/d
/"104.237.142.206/d
/"176.58.126.224/d
/"172.104.85.6/d
/"139.162.7.132/d

# add more peers 
# to be added

# change WebInterfacePort
s/6420/$WebInterfacePort/

# change Port 
s/6000/$Port/

# change RPCInterfacePort
s/6430/$RPCInterfacePort/

# change PeerListURL
# s#https://downloads\.skycoin\.net/blockchain/peers.txt#$PeerListURL#

# change data folder 
s#\.skycoin#\.$coinname#

# change distribution address

#=> change GenesisUxid
# s/043836eb6f29aaeb8b9bfce847e07c159c72b25ae17d291f32125e7f1912e2a0/$GenesisUxid/
`

func skycoinSed(fn string, c CoinConfigT) {

	cText := injectValues(skycoinSedTemplate, c)
	write2File(fn, cText)
}

package main

// const skycoinSedTemplate = `
// # This script is used to modify !skycoin.go!

// # change BlockchainPubkeyStr
// s/0328c576d3f420e7682058a981173a4b374c7cc5ff55bf394d3cf57059bbe6456a/$BlockchainPubkeyStr/

// # change GenesisAddressStr
// s/2jBbGxZRGoQG1mqhPBnXnLTxK6oxsTf8os6/$GenesisAddressStr/

// # change GenesisSignatureStr
// s/eb10468d10054d15f2b6f8946cd46797779aa20a7617ceb4be884189f219bc9a164e56a5b9f7bec392a804ff3740210348d73db77a37adb542a8e08d429ac92700/$GenesisSignatureStr/

// # change GenesisTimestamp
// s/1426562704/$GenesisTimestamp/

// # change GenesisCoinVolume
// s/100e12/$GenesisCoinVolume/

// # change default peers
// s/"118.178.135.93:6000"/"$nodewithport001"/
// s/"47.88.33.156:6000"/"$nodewithport002"/

// # delete extra default peers, only leaving two
// /"121.41.103.148/d
// /"120.77.69.188/d
// /"104.237.142.206/d
// /"176.58.126.224/d
// /"172.104.85.6/d
// /"139.162.7.132/d

// # add more peers
// # to be added

// # change WebInterfacePort
// s/6420/$WebInterfacePort/

// # change Port
// s/6000/$Port/

// # change RPCInterfacePort
// s/6430/$RPCInterfacePort/

// # change PeerListURL
// # s#https://downloads\.skycoin\.net/blockchain/peers.txt#$PeerListURL#

// # change data folder
// s#\.skycoin#\.$coinname#

// # change distribution address

// #=> change GenesisUxid
// # s/043836eb6f29aaeb8b9bfce847e07c159c72b25ae17d291f32125e7f1912e2a0/$GenesisUxid/
// `

const skycoinSedTemplate = `
# Change skycoin.go so we can issue a complete new coin 
# for v0.24.0 Note this is skycoin version
#
#
# the following parameters are to be changed 
# 1. GenesisSignatureStr
# 2. GenesisAddressStr
# 3. BlockchainPubkeyStr
# 4. GenesisTimestamp
# 5. GenesisCoinVolume
# 6. DefaultConnections
# 7. Port
# 8. WebInterfacePort
# 9. DataDirectory
# 10. ProfileCPUFile
#

# 01 change GenesisSignatureStr
s/eb10468d10054d15f2b6f8946cd46797779aa20a7617ceb4be884189f219bc9a164e56a5b9f7bec392a804ff3740210348d73db77a37adb542a8e08d429ac92700/$GenesisSignatureStr/

# 02 change GenesisAddressStr
s/2jBbGxZRGoQG1mqhPBnXnLTxK6oxsTf8os6/$GenesisAddressStr/

# 03 change BlockchainPubkeyStr
s/0328c576d3f420e7682058a981173a4b374c7cc5ff55bf394d3cf57059bbe6456a/$BlockchainPubkeyStr/

# 04 change GenesisTimestamp
s/1426562704/$GenesisTimestamp/

# 05 change GenesisCoinVolume
s/100000000000000/$GenesisCoinVolume/

# 06 change default peers 
s/"118.178.135.93:6000"/"$nodewithport001"/
s/"47.88.33.156:6000"/"$nodewithport002"/

# delete extra default peers, only leaving two 
/"121.41.103.148/d
/"120.77.69.188/d
/"104.237.142.206/d
/"176.58.126.224/d
/"172.104.85.6/d
/"139.162.7.132/d

# 07 change Port 
s/6000/$Port/
 
# 08 change WebInterfacePort
s/6420/$WebInterfacePort/

# 09 change data folder 
s#\.skycoin#\.$coinname#

# 10 change profile CPU file
s#skycoin\.prof#$coinname\.prof#

# change PeerListURL
# s/https:\/\/downloads.skycoin.net/blockchain/peers.txt/??????/

# change distribution address`

func skycoinSed(fn string, c CoinConfigT) {

	cText := injectValues(skycoinSedTemplate, c)
	write2File(fn, cText)
}

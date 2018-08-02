package main

const configurationMdTemplate = `

# Configurations

## Core configurations

| Item                  | Value           |
|-----------------------|-----------------|
|              Coin name| $CoinName
|            Coin symbol| $CoinSymbol
|    BlockchainPubkeyStr| $BlockchainPubkeyStr
|    BlockchainSeckeyStr| $BlockchainSeckeyStr
|      GenesisAddressStr| $GenesisAddressStr
|             PrivateKey| $gasPrivateKey
|                       | $gasSeed
|    GenesisSignatureStr| $GenesisSignatureStr
|       GenesisTimestamp| $GenesisTimestamp
|                       | $TimestampReadable
|      GenesisCoinVolume| $GenesisCoinVolume
|            PeerListURL| $PeerListURL
|                   Port| $Port
|       WebInterfacePort| $WebInterfacePort
|       RPCInterfacePort| $RPCInterfacePort Not used anymore from v0.24.0
|            GenesisUxID| $GenesisUxID
|                 Node 1| $node001
|                 Node 2| $node002
|                 Node 3| $node003
|                 Node 4| $node004

# Deploy nodes

scp -P XXXX $coinname-0.24.1-bin-linux-x64.tar.gz root@$node001:/root
scp -P XXXX $coinname-0.24.1-bin-linux-x64.tar.gz root@$node002:/root
scp -P XXXX $coinname-0.24.1-bin-linux-x64.tar.gz root@$node003:/root
scp -P XXXX $coinname-0.24.1-bin-linux-x64.tar.gz root@$node004:/root

# Start master node

./$coinname -master=true -master-secret-key=$BlockchainSeckeyStr  -launch-browser=false -download-peerlist=false -enable-wallet-api

# Start normal node

./$coinname -launch-browser=false -download-peerlist=false

# Distribution addresses

`

func configurationFile(fn string, c CoinConfigT) {
	configString := injectValues(configurationMdTemplate, c)
	configString = insertDistributionAddrs(configString, c)
	write2File(fn, configString)
}

func insertDistributionAddrs(text string, c CoinConfigT) string {
	for _, addr := range c.DistributionAddresses {
		text = text + addr
	}
	return text
}

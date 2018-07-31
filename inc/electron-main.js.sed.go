package main

// v0.23.0
// const electronMainSed = `
// # Usage
// # sed -i -f electron-main.js.sed <path-to-electron-main.js>

// # change port number
// s/6420/$WebInterfacePort/

// # change executable name
// s#esources/app/skycoin#esources/app/$coinname#

// # change log messages
// s/Starting skycoin/Starting $coinname/
// s/Skycoin already running/$Coinname already running/
// s/Failed to start skycoin/Failed to start $coinname/
// s/Cleared the caching of the skycoin wallet/Cleared the caching of the $coinname wallet/

// # download-peerlist=false
// s/-download-peerlist=true/-download-peerlist=false/

// # change menu
// s/About Skycoin/About $Coinname/

// # Note this should be the last one!
// s/"Skycoin"/"$Coinname"/

// `

//v0.24.0
const electronMainSed = `
# v0.24.0 Note this is Skycoin version

# Usage
# sed -i -f electron-main.js.sed <path-to-electron-main.js>

# change executable name
s#esources/app/skycoin#esources/app/$coinname#

# change log messages
s/Starting skycoin/Starting $coinname/
s/Skycoin already running/$Coinname already running/
s/Failed to start skycoin/Failed to start $coinname/
s/Skycoin closed/$Coinname closed/
s/Skycoin exited/$Coinname exited/
s/Cleared the caching of the skycoin wallet/Cleared the caching of the $coinname wallet/

# download-peerlist=false
s/-download-peerlist=true/-download-peerlist=false/

# change menu
s/Skycoin/$Coinname/

# Note this should be the last one!
s/'Skycoin'/'$Coinname'/
`

func eletronMainJs(fn string, c CoinConfigT) {
	cText := electronMainSed
	cText = injectValues(cText, c)
	write2File(fn, cText)
}

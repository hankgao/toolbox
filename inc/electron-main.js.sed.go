package main

const electronMainSed = `
# Usage
# sed -i -f electron-main.js.sed <path-to-electron-main.js>

# change port number
s/6420/$WebInterfacePort/

# change executable name
s/esources\/app\/skycoin/esources\/app\/$coinname/

# change log messages
s/Starting skycoin/Starting $Coinname/
s/Skycoin already running/$Coinname already running/
s/Failed to start skycoin/Failed to start $Coinname/
s/Cleared the caching of the skycoin wallet/Cleared the caching of the $Coinname wallet/

# download-peerlist=false
s/-download-peerlist=true/-download-peerlist=false/

# change menu
s/About Skycoin/About $Coinname/
s/"Skycoin"/"$Coinname"/

`

# Overview

bp - blockchain parser
bp analyzes blockchain database and retrieve information required

## Steps

- call `NewVisor`
- `visor.GetUnspentOutputs`
> GetUnspentOutputs returns ReadableOutputs
> Before we can call GetUnspentOutputs, we need to prepare visor for that.

## Packages to be imported

- "github.com/skycoin/skycoin/src/visor"
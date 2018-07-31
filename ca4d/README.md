# Overview

ca4d - Create 100 addresses for initial coin distribution.

This tool is used by project owner to protect the safety of their coins, as the seeds are not exposed to development team. 

## design

- a CLI tool
- create two csv files, one only containing 100 addresses, another containing seeds and addresses. The first file is sent to development team, the 2nd one is used by project owners to pick up coins using the seeds
- create a folder named coins
- create a subfolder under coins for the new coin
- create 2 csv files

### Another possible implementation

Used nodeJS

## usage

``` bash
ca4d 100 newcoinname hank.gao@gmail.com


```
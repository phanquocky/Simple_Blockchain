# Simple_Blockchain

## Build
Build project with command: 
```
go build
```
## Run
Now, you’ve created your executable, run it. 
On macOS or Linux, run the following command:
```
./blockchain_go <command> [options]
```
On Windows, run:
```
blockchain_go.exe <command> [options]
```

## Functional
Up to date, there are 4 main command: printchain, createchain, addblock, validtran.

Run with command for usage:
```
<program> help
```

## Example
With executable file on Window. Create a new blockchain:
```
$ blockchain_go.exe createchain
$ blockchain_go.exe printchain

Block index 0                                                                                                                   Timestamp:      2023-12-05 18:40:07 +0700 +07                                                                           Merkle Root:    e7a...0df                                        Prev Block:     000...000                                        Hash Value:     474...b04b
```
Now you can see your block chain has 1 block which is genesis block.
Let's add a new block to blockchain.
```
$ blockchain_go.exe addblock -data=xyz -data=abc
$ blockchain_go.exe printchain

Block index 0
        Timestamp:      2023-12-05 18:40:07 +0700 +07
        Merkle Root:    e7a...90df
        Prev Block:     0000...000
        Hash Value:     4747...04b
Block index 1
        Timestamp:      2023-12-05 18:40:26 +0700 +07
        Merkle Root:    591...cee
        Prev Block:     474...04b
        Hash Value:     6a3...da3
```
Validate an transaction is in a block (through block's hash value - you can get it with `printchain` command) by following command:
```
$ blockchain_go.exe validtran -data=xyz -data=123 -block=6a3...da3

Transaction data xyz is stored in the block                                                                             Transaction data 123 is NOT stored in the block
```


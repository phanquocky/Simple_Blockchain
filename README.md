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
blockchain_go.exe createchain
blockchain_go.exe printchain
```
Now you can see your block chain has 1 block which is genesis block.
Let's add a new block to blockchain.
```
blockchain_go.exe addblock -data xyz -data abc
blockchain_go.exe printchain
```
Validate an transaction is in a block (through block's hash value - you can get it with `printchain` command) by following command:
```
blockchain_go.exe validtran -data xyz -block <blockHashValue>
```


```bash
$ go build
$ ./blockchain_go printchain
$ ./blockchain_go addblock -data "Send 1 BTC to Ivan"


$ ./blockchain_go createwallet
$ ./blockchain_go createblockchain -address="1NknnPE4vyrYa6Ec8e2Xxh7hE7yGew7qUv"
$ ./blockchain_go getbalance -address="1NknnPE4vyrYa6Ec8e2Xxh7hE7yGew7qUv"
$ ./blockchain_go transfer -from="1NknnPE4vyrYa6Ec8e2Xxh7hE7yGew7qUv" -to="17TaGyudJ7NLvRTEGL3fFwGKEd94uM9SEe" -amount=3
$ ./blockchain_go getbalance -address="17TaGyudJ7NLvRTEGL3fFwGKEd94uM9SEe"

```

ref: https://github.com/Jeiwan/blockchain_go

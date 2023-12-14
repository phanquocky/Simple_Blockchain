# Installization

btcd: full bitcoin node (testnet)
ref: https://github.com/btcsuite/btcwallet

btcwallet: bitcoin wallet (testnet)
ref: https://github.com/btcsuite/btcd

# Run Command

```bash
# btcd --testnet -u <rpcuser> -P <rpcpassword>
$ btcd --testnet -u admin -P admin123

# btcwallet --testnet -u <rpcuser> -P <rpcpassword>
$ btcwallet --testnet -u admin -P admin123
```

# task 1

```bash
$ ./lab2 getnewaddress
2023/12/14 14:08:32 Your Address:  mh93UUSMtupq9dypmWaheUR48oV2n2z9kQ
2023/12/14 14:08:32 Your Private Key:  c1f3801bbaa60b1350bb7f0d04cce4a568184094550205cac9e56f95b3e1d642
```

send coin to this address using bitcoin testnet faucet
ref: https://coinfaucet.eu/en/btc-testnet/

prevhash = "4ebf66b5b2f0bea65ffa87e00868d34aa4090b4ada4b58c5c8a3d1c5bcb4f7e4"
outIdx = 1

```bash
$ ./lab2 spendfund -privkey="c1f3801bbaa60b1350bb7f0d04cce4a568184094550205cac9e56f95b3e1d642" -prevhash="4ebf66b5b2f0bea65ffa87e00868d34aa4090b4ada4b58c5c8a3d1c5bcb4f7e4" -outidx=1
privkey string 11cb6c960a9a8907ba535f4129d06cd364c1dc9c
2023/12/14 17:05:20 Send transaction success!, txhash:  d54caf9ac355d8a8cc319de237808b815d1205e4bf488658f6fcb655947e03b8
```

```bash
$ ./lab2 getnewmultisig
2023/12/14 17:53:43 Your multisig address:  2NCSFYjfPFpyazd5GetPjxbgrpL5Fv66FdY
2023/12/14 17:53:43 First Private Key:  4c75458e1a1ae0e3c205805eef0ad666f69dc46bc6581033472698da1ac904af
2023/12/14 17:53:43 Second Private Key:  c9e45684c6dfd93770cc76eeb49d9b2b9017f7e4ab632863c195f563caf01b3a
2023/12/14 17:53:43 Redeem script:  522103fd886e0bb6cfe9db2a6c28fc000bd66855f5cab89a43f8f1ae5150889dc9dad42103e943aee00b09df2398b5ba85b2b1dcf5c627d2f70cc6b55756f6b6ba21c74ba3ae
```

prevhash: 925c8b18e07c6f287952a782141820ad2bddfe47fdf7a0a412be89e3a122f568
outidx: 0

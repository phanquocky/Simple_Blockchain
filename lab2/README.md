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

### Get new p2pkh address

```bash
$ ./lab2 getp2pkhaddress
2023/12/14 14:08:32 Your Address:  mh93UUSMtupq9dypmWaheUR48oV2n2z9kQ
2023/12/14 14:08:32 Your Private Key:  c1f3801bbaa60b1350bb7f0d04cce4a568184094550205cac9e56f95b3e1d642
```

### Using bitcoin testnet faucet send coin to this address

send coin to this address using bitcoin testnet faucet

ref: https://coinfaucet.eu/en/btc-testnet/

Information transaction when send coin

```
prevhash = "4ebf66b5b2f0bea65ffa87e00868d34aa4090b4ada4b58c5c8a3d1c5bcb4f7e4"
outidx = 1
```

### Spend this utxo on p2pkh address

```bash
$ ./lab2 spendp2pkh -privkey="c1f3801bbaa60b1350bb7f0d04cce4a568184094550205cac9e56f95b3e1d642" -prevhash="4ebf66b5b2f0bea65ffa87e00868d34aa4090b4ada4b58c5c8a3d1c5bcb4f7e4" -outidx=1
privkey string 11cb6c960a9a8907ba535f4129d06cd364c1dc9c
2023/12/14 17:05:20 Send transaction success!, txhash:  d54caf9ac355d8a8cc319de237808b815d1205e4bf488658f6fcb655947e03b8
```

# Task 2

### Creaate new 2-2 multisig

```bash
$ ./lab2 getmultisigaddress
2023/12/14 23:00:58 Your multisig address:  2MsHA5Z5aRAbNTt5WhjsaGkPFYCHhPTtf4A
2023/12/14 23:00:58 First Private Key:  47ac054b280206ee77c09eeb1eaac2ecd7ea999e42b3fdc4225d44727c4e6674
2023/12/14 23:00:58 Second Private Key:  1f33e82515928bfd717991d1d2d0921d2a297364d5d5842a2b93c0dadc660fbf
2023/12/14 23:00:58 Redeem script:  522103a28100a3a6b247188bc9885a2fae6729642d976324b738a811b724b8470645b92103f488b4febc5e72e51d597535adadc6a1011d35c81007caf0120def2150e7a65252ae
```

### Using bitcoin testnet faucet send coin to this address

prevhash: 28fef2dfe59168e934f4af1d338d7fa615936e58dce1ba3c320d713aec0baee4

outidx: 0

### Spend this utxo on 2-2 multisig address

```bash
$ ./lab2 spendmultisig -privkey1="47ac054b280206ee77c09eeb1eaac2ecd7ea999e42b3fdc4225d44727c4e6674" -privkey2="1f33e82515928bfd717991d1d2d0921d2a297364d5
d5842a2b93c0dadc660fbf" -prevhash="28fef2dfe59168e934f4af1d338d7fa615936e58dce1ba3c320d713aec0baee4" -outidx=0 -redeem="522103a28100a3a6b247188bc9885a2fa
e6729642d976324b738a811b724b8470645b92103f488b4febc5e72e51d597535adadc6a1011d35c81007caf0120def2150e7a65252ae"
2023/12/14 23:07:39 Send transaction success!, txhash:  f66b71cbc3feed81c7b0ed7e05a12a233d3340620cd63cbbb10afc9359ca0f3f

```

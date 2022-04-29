# SBox VRF Random


This code will generate an array of winners like:

```js
winners = [105, 503, 203, ...]
```

We also have an array of sboxes, which will be open on the github and discord before we run the random code. Then we run the random code get token ID order for the sboxes.


## How to obtain 'contractRandomSeed'

Every random we will use our contract (verified on etherscan, check links below) and create a SBOX RANDOM SEED.

This transaction cannot be manipulated and it uses [ChainLink VRF](https://docs.chain.link/docs/chainlink-vrf/).

You can use Etherscan to obtain every random seed.

## Notes

We believe that, this system %100 transparent and can be noterized by everyone.



## Links

- [SBOXRandomSeedGenerator contract](https://etherscan.io/address/0x8a664be336304385b2886e706828e4ee1e549d44#readContract)

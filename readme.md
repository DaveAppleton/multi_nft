# Sale contract & multi sale contracts

the contracts in contracts/flat are the final ones because they will be verified

## Private Sales

private sales implement royalty splits

- between flexible numbers of recipients
- with initial sale and resale having different splits
- defined by `pool`
- where each specific token being assigned to a `pool`

So _token#1_ can have a different token split from _token#2_

---

## The NEW (18 June 2021) private sale contract & token

## Pool 1

initial sales
```
    creator1  = { beneficiary: creator, share: 500 }
    patron1   = { beneficiary: patron, share: 500 }
```

secondary sales
```
    creator1R = { beneficiary: creator, share: 750 }
    patron1R  = { beneficiary: patron, share: 250 }

```



first sale

```
buyer spends  0.5125     (1.025)
seller gains  0.45       (0.9)
patron  gets  0.025      (0.05)   500
creator gets  0.025      (0.05)   500
minty   gets  0.0125     (0.025)
```

second sale

```
buyer spend  0.5125      (1.025)
seller gains 0.45        (0.9)
patron  gets 0.0125      (0.025)  250
creator gets 0.0375      (0.075)  750
minty   gets 0.0125      (0.025)
```

## Pool 2

```
    creator2   = { beneficiary: creator, share: 500 }
    musician2  = { beneficiary: musician, share: 300 }
    patron2    = { beneficiary: patron, share: 200 }

    creator2R  = { beneficiary: creator, share: 700 }
    musician2R = { beneficiary: musician, share: 200 }
    patron2R   = { beneficiary: patron, share: 100 }
```

first sale 

```
buyer spend   1.025
seller gains  0.9
patron   gets 0.02    200
creator  gets 0.05    500
musician gets 0.03    300
minty    gets 0.025
```

secondary sale

```
buyer spend  1.025
seller gains  0.9
patron   gets 0.01  100
creator  gets 0.07  700
musician gets 0.02  200
minty    gets 0.025
```

## Token Deployment

Assuming
- locking1, locking2 - subscribing to either of these contracts lets you in
- "You need..." the revert error message if nothing locked for first sale
- `[creator1,patron1]`   the initial share (see above)
- `[creator1R,patron1R]` resale share
- `"C+P"` - name of this Pool

```
m1155 = await M1155.deploy(
        artistAddress,
        saleContractAddress,
        [locking1.address,locking2.address],
        "You need to be a MINTY patron or TEST PROJECT patron",
        100,
        [creator1,patron1],
        [creator1R,patron1R],
        "C+P"
)
```

Adding another pool

```
await m1155.connect(addr1).addPools([creator2,musician2,patron2],[creator2R,musician2R,patron2R],"C+P+M")
```

## Sale Contract
Putting a new item for sale

```
sale.offerNew(
    m1155.address,                    // token contract
    11,                               // tokenId
    "hash",                           // URI hash
    180,                              // quantity
    ethers.utils.parseEther("0.205"), // unit price
    1                                 // poolID
)
```

Resale does not need to know the poolID

```
sale.offerResale(
     m1155.address,                    // token contract
     11,                               // tokenId
     5,                                // quantity
     ethers.utils.parseEther("0.205")  // unit price
)
```


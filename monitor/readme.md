# Building this

creating the bindings

for each contract that you monitor you need to create a binding using ABIGEN (part of the GETH tools)

```bash
monitor $ cd pSale
pSale $ abigen --sol psale.sol --pkg pSale --out core.go
cd ..
```

You can import this into your code

```go
import(
	"art.minty.monitor/pSale"
)
```

instatiate a contract object 

```go
	sale, err := pSale.NewPMintysale(saleAddress, client)
```

inside a goroutine (thread) you check the filter object so pass that and anything else you need

```go
	go multiTransfers(&multitoken.PMintyMultiTokenFilterer, multiTokenAddress, &wg)
```


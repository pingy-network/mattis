# Erc20 Contract Binding

Generate the smart contract bindings using the [abigen] command line tool.

```
abigen --abi pkg/contract/erc20/abi.json --pkg erc20 --type Erc20Contract --out pkg/contract/erc20/contract.go
```

[abigen]: https://geth.ethereum.org/docs/tools/abigen

## Resources

[sdk.operatorframework.io](https://sdk.operatorframework.io/)
[golang operator](https://sdk.operatorframework.io/docs/building-operators/golang/)

## Pros & Cons

### Pros

### Cons


## Init

```shell
# get operator-sdk\
curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v1.13.1/operator-sdk_linux_amd64 > operator-sdk-1.13.1

# init the operator
operator-sdk-1.13.1 init --domain example.com --repo github.com/example/some-operator

# create the API
operator-sdk-1.13.1 create api --group mygroup --version v1alpha1 --kind MyKind --resource --controller
```



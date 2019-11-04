# Easy Chain Lennon (鏈上連儂牆)

![img](assets/images/preview.png)

### require:
* go1.12.9 above
* [gobuffalo](https://gobuffalo.io/en/]) - web framework
* postgresql database

### setup:
```
# enable go mod
export GO111MODULE=on
go mod vendor
buffalo db create
buffalo db migrate
cp env_example .env # this is environment setting file of go web backend
# get top 10000 transactions from trigger recevied address
curl http://localhost:3000/trigger_query/query/
```
### frontend

please refer `hellochain_dapp`. after compiled dist folder. please run `./copy_from_fd.sh`

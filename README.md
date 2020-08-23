# BlockIP

本项目旨在本地接收需要block 掉的ip规则

timeout 默认为300s

服务器默认开启: 127.0.0.1:9111

## Build

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o staticBuilds/blockip

## Production

    nohup ./blockip > nohup.log 2>&1 &

## Other

[ipset 参考](https://www.zybuluo.com/lniwn/note/899851)

```shell
TIMEOUT=300 # 300s
IPSET_HASH_NAME=ip_block
ipset create $IPSET_HASH_NAME hash:ip
ipset add -exist $IPSET_HASH_NAME ${ip=要封禁的ip} timeout $TIMEOUT
ipset del $IPSET_HASH_NAME x.x.x.x
ipset list $IPSET_HASH_NAME
ipset list
ipset flush $IPSET_HASH_NAME
ipset flush
ipset destroy $IPSET_HASH_NAME
ipset destroy
ipset save $IPSET_HASH_NAME
ipset save
ipset restore
```
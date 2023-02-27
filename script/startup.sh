#! /bin/bash

for i in `seq 1 4`
do
    mkdir -p /tmp/$i/object
    mkdir /tmp/$i/temp
done

cd dataservice
go run . -store /tmp/1/object/ -temp /tmp/1/temp/ -p 10100 -rpcn nagato.data -rpcp 10150 &
go run . -store /tmp/2/object/ -temp /tmp/2/temp/ -p 10101 -rpcn nagato.data -rpcp 10151 &
go run . -store /tmp/3/object/ -temp /tmp/3/temp/ -p 10102 -rpcn nagato.data -rpcp 10152 &
go run . -store /tmp/4/object/ -temp /tmp/4/temp/ -p 10103 -rpcn nagato.data -rpcp 10153 &

cd ../apiservice &
go run . -p 10000 &
go run . -p 10001 &
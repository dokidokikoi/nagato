#! /bin/bash

for i in `seq 1 4`
do
    mkdir -p /tmp/$i/object
    mkdir /tmp/$i/temp
done

cd dataservice
go run . -store /tmp/1/object/ -temp /tmp/1/temp/ -p 10100 &
go run . -store /tmp/2/object/ -temp /tmp/2/temp/ -p 10101 &
go run . -store /tmp/3/object/ -temp /tmp/3/temp/ -p 10102 &
go run . -store /tmp/4/object/ -temp /tmp/4/temp/ -p 10103 &

cd ../apiservice &
go run . -p 10000 &
go run . -p 10001 &
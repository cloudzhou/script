#!/bin/bash

usage=`top -b -n 1|head -n 3|grep ^Cpu|awk -F "%| " '{print $3}'`
usage=`printf '%.0f' $usage`

memory_used=`free -m|grep ^Mem|awk '{print $3}'`
memory_free=`free -m|grep ^Mem|awk '{print $4}'`

df_used=`df -h |grep \/$|awk '{print $3}'|replace 'G' ''`
df_used=`printf '%.0f' $df_used`
df_free=`df -h |grep \/$|awk '{print $4}'|replace 'G' ''`
df_free=`printf '%.0f' $df_free`

iostat_read=`iostat|grep ^xvda|awk '{print $3}'`
iostat_read=`printf '%.0f' $iostat_read`
iostat_write=`iostat|grep ^xvda|awk '{print $4}'`
iostat_write=`printf '%.0f' $iostat_write`

netstat_est=`netstat -an|grep ESTABLISHED|wc -l`
netstat_tiw=`netstat -an|grep TIME_WAIT|wc -l`

echo $usage, $memory_used, $memory_free, $df_used, $df_free, $iostat_read, $iostat_write, $netstat_est, $netstat_tiw

token='THERE IS YOUR DEVICE TOKEN'

curl -H "Authorization: token $token" -d "{\"datapoint\": {\"x\": $usage}}" http://iot.espressif.cn/v1/datastreams/cpu/datapoint/
curl -H "Authorization: token $token" -d "{\"datapoint\": {\"x\": $memory_used, \"y\": $memory_free}}" http://iot.espressif.cn/v1/datastreams/memory/datapoint/
curl -H "Authorization: token $token" -d "{\"datapoint\": {\"x\": $df_used, \"y\": $df_free}}" http://iot.espressif.cn/v1/datastreams/df/datapoint/
curl -H "Authorization: token $token" -d "{\"datapoint\": {\"x\": $iostat_read, \"y\": $iostat_write}}" http://iot.espressif.cn/v1/datastreams/iostat/datapoint/
curl -H "Authorization: token $token" -d "{\"datapoint\": {\"x\": $netstat_est, \"y\": $netstat_tiw}}" http://iot.espressif.cn/v1/datastreams/netstat/datapoint/

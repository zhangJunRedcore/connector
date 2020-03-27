#!/bin/bash

nohup /project/connector/bin/connectord -config_path=/project/connector/etc/conf.yaml >> connector_start.log 2>&1 &

echo "Ready to run connectord Server ..."
echo ""
sleep 1

#
#	check if connectord is runing
#

PROC_NAME=connectord
ProcNumber=`ps -aux | grep $PROC_NAME | grep /etc/conf | wc -l`
if [ $ProcNumber == 1 ];then
   echo "Staring connectord successful..."
else
   echo "Staring connectord fail..."
   echo ""
fi

exit

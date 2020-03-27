#!/bin/bash

PROC_NAME=connectord
killall -9 $PROC_NAME
sleep 1

ProcNumber=`ps -aux | grep $PROC_NAME | grep conf | wc -l`
if [ $ProcNumber == 0 ];then
	echo "Stop deeptund successful..."
	echo ""
else
	echo "Stop deeptund fail..."
	echo ""
fi
exit

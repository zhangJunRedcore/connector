#!/bin/bash

CURRENT_DIR=`cd \`dirname $0\`; pwd`

$CURRENT_DIR/stop_connector.sh
sleep 1
$CURRENT_DIR/run_connector.sh

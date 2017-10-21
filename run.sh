#!/usr/bin/env bash

tar -xzvf dev.action.360baige.com.tar.gz

killall dev.action.360baige.com

nohup ./dev.action.360baige.com  >/dev/null 2>error.log &

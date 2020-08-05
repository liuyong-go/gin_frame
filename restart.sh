#!/bin/bash
cd "/var/www/gin_frame"
git pull
kill -1 $(lsof -i:8080 |awk '{print $2}' | tail -n 1)
#!/bin/sh 
./goose -dir ./migrations $DB_DRIVER $DB_DNS up
./demo-app
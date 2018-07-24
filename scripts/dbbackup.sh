#!/bin/sh
DIR=`date +%m%d%y`
DEST=/data/db_backups/$DIR
mkdir $DEST
mongodump -o $DEST --db gene
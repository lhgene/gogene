#!/bin/sh
DIR=`date +%m%d%y`
DEST=/db_backups/$DIR
mkdir $DEST
mongodump -o $DEST --db gene
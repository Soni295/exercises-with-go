#!/bin/bash
file=$(date +"%H:%M-%D-backup.sql")

touch $file;
#mysqldump go_example -usion -p"123456" > example.sql
#mysqldump go_example -usion -p"123456" > $file

echo $file
#echo $(date +"%H:%M-%D-backup.sql")

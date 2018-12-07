#!/bin/bash
# Configuration information; modify it as needed
mysql_user="USER" #MySQL backup user
mysql_password="PASSWORD" # MySQL backup user's password
mysql_host="localhost"
mysql_port="3306"
mysql_charset="utf8" # MySQL encoding
backup_db_arr=("db1" "db2") # Name of the database to be backed up, separating multiple databases wih spaces ("DB1", "DB2" db3 ")
backup_location=/var/www/mysql # Backup data storage location; please do not end with a "/" and leave it at its default, for the program to automatically create a folder
expire_backup_delete="ON" # Whether to delete outdated backups or not
expire_days=3 # Set the expiration time of backups, in days (defaults to three days); this is only valid when the `expire_backup_delete` option is "ON"

# We do not need to modify the following initial settings below
backup_time=`date +%Y%m%d%H%M` # Define the backup time format
backup_Ymd=`date +%Y-%m-%d` # Define the backup directory date time
backup_3ago=`date-d '3 days ago '+%Y-%m-%d` # 3 days before the date
backup_dir=$backup_location/$backup_Ymd # Full path to the backup folder
welcome_msg="Welcome to use MySQL backup tools!" # Greeting

# Determine whether to MySQL is running; if not, then abort the backup
mysql_ps=`ps-ef | grep mysql | wc-l`
mysql_listen=`netstat-an | grep LISTEN | grep $mysql_port | wc-l`
if [[$mysql_ps==0]-o [$mysql_listen==0]]; then
  echo "ERROR: MySQL is not running! backup aborted!"
  exit
else
  echo $welcome_msg
fi

# Connect to the mysql database; if a connection cannot be made, abort the backup
mysql-h $mysql_host-P $mysql_port-u $mysql_user-p $mysql_password << end
use mysql;
select host, user from user where user='root' and host='localhost';
exit
end

flag=`echo $?`
if [$flag!="0"]; then
  echo "ERROR: Can't connect mysql server! backup aborted!"
  exit
else
  echo "MySQL connect ok! Please wait......"
   # Determine whether a backup database is defined or not. If so, begin the backup; if not, then abort
  if ["$backup_db_arr"!=""]; then
       # dbnames=$(cut-d ','-f1-5 $backup_database)
       # echo "arr is(${backup_db_arr [@]})"
      for dbname in ${backup_db_arr [@]}
      do
          echo "database $dbname backup start..."
          `mkdir -p $backup_dir`
          `mysqldump -h $mysql_host -P $mysql_port -u $mysql_user -p $mysql_password $dbname - default-character-set=$mysql_charset | gzip> $backup_dir/$dbname -$backup_time.sql.gz`
          flag=`echo $?`
          if [$flag=="0"]; then
              echo "database $dbname successfully backed up to $backup_dir/$dbname-$backup_time.sql.gz"
          else
              echo "database $dbname backup has failed!"
          fi

      done
  else
      echo "ERROR: No database to backup! backup aborted!"
      exit
  fi
   # If deleting expired backups is enabled, delete all expired backups
  if ["$expire_backup_delete"=="ON" -a "$backup_location"!=""]; then
      # `find $backup_location/-type d -o -type f -ctime + $expire_days-exec rm -rf {} \;`
      `find $backup_location/ -type d -mtime + $expire_days | xargs rm -rf`
      echo "Expired backup data delete complete!"
  fi
  echo "All databases have been successfully backed up! Thank you!"
  exit
fi
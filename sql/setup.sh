PATH_TO_SQL_FILE="/home/bjs/go-postgres/sql"
USERNAME="postgres"
HOSTNAME="localhost"
PORT=5432
DB="postgres"
sudo -u $USERNAME bash -c "psql -h $HOSTNAME -p $PORT < $PATH_TO_SQL_FILE/setup.sql"
sudo -U $USERNAME bash -c "psql -h $HOSTNAME -p $PORT -d $DB -f $PATH_TO_SQL_FILE/db.sql"

#postgresql restore from dump
#psql -U postgres -h localhost -p 5432 -f data_dump.sql

#restore postgres database from sql file
#psql -d database_name -f backup.sql

#!bin/sh

CREDS="creds-$MYSQL_ENV.json"
BUILD_SCRIPT="../brieefly/sql/build-tables.sql"
TEST_DATA="../brieefly/sql/test-data.sql"

echo $CREDS

DB_NAME="$(cat ../brieefly/secrets/$CREDS | jq -r .name)"
DB_USER="$(cat ../brieefly/secrets/$CREDS | jq -r .user)"
DB_PASSWORD="$(cat ../brieefly/secrets/$CREDS | jq -r .password)"

echo $DB_NAME $DB_USER $DB_PASSWORD

Q1="CREATE DATABASE IF NOT EXISTS $DB_NAME; "
Q2="CREATE USER IF NOT EXISTS '$DB_USER'@'localhost' IDENTIFIED BY '$DB_PASSWORD'; "
Q3="GRANT ALL ON $DB_NAME.* TO '$DB_USER'@'localhost'; "
Q4="FLUSH PRIVILEGES; "

SQL="${Q1}${Q2}${Q3}${Q4}"

echo $SQL

mysql -uroot -e "$SQL";
mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME < $BUILD_SCRIPT;
mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME < $TEST_DATA;
mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME;


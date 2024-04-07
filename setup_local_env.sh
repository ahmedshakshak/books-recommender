export $(cat .env | xargs)

while ! wget --spider --tries=1 localhost:$MYSQL_PORT; do echo "Don't panic we are just waiting for MySQL ..." && sleep 5; done
echo "MySQL is up"

docker exec -i books-recommender_database_1 mysql -u $MYSQL_USERNAME -p$MYSQL_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $MYSQL_DB_NAME"
docker exec -i books-recommender_database_1 sh -c "exec mysql -u"$MYSQL_USERNAME" -p"$MYSQL_PASSWORD" $MYSQL_DB_NAME" < ./testhelpers/dumps/mysql.sql

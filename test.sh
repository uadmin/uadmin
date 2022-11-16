# Set testing environment variables
export UADMIN_TEST_MYSQL_PORT=23306
export UADMIN_TEST_MYSQL_USERNAME=root
export UADMIN_TEST_MYSQL_PASSWORD=rootpassword
export UADMIN_TEST_MYSQL_HOST=127.0.0.1
export UADMIN_TEST_POSTGRES_PORT=25432
export UADMIN_TEST_POSTGRES_USERNAME=postgres
export UADMIN_TEST_POSTGRES_PASSWORD=rootpassword
export UADMIN_TEST_POSTGRES_HOST=127.0.0.1

# Clean up docker databases engines / if they exist
docker stop uadmin-test-mysql
docker rm uadmin-test-mysql
docker stop uadmin-test-postgres
docker rm uadmin-test-postgres

# Start docker database engines
docker run --name uadmin-test-mysql -e MYSQL_ROOT_PASSWORD=rootpassword -p 23306:3306 -d mysql:8.0
docker run --name uadmin-test-postgres -e POSTGRES_PASSWORD=rootpassword -p 25432:5432 -d postgres:15.1

# Wait MySQL to be live
echo "Waiting for MySQL Server ..."
while ! mysqladmin ping -h"$UADMIN_TEST_MYSQL_HOST" -P "$UADMIN_TEST_MYSQL_PORT" --silent; do
    sleep 1
done
echo "Waiting for Postgres Server ..."
while ! nc -z $UADMIN_TEST_POSTGRES_HOST $UADMIN_TEST_POSTGRES_PORT; do
    sleep 1;
done

# Run tests
go test -run '' -cover -v -coverprofile=coverage.out; go tool cover -html=coverage.out; rm coverage.out

# Clean up docker databases engines
docker stop uadmin-test-mysql
docker rm uadmin-test-mysql
docker stop uadmin-test-postgres
docker rm uadmin-test-postgres

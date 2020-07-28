# -*- mode: Python -*-

docker_compose('docker-compose.yml')
# php_version = str(local('source .env && echo $PHP_VERSION')).strip(). @TODO: This does not
# work on some systems, hence for workaround php_version is hardcoded.
# Change this in case you want to use someother version of php(supported: 7.2, 7.3 & .7.4).
# Please ensure to change PHP_VERSION in .env file as well.
php_version = "7.3"
docker_build('srijanlabs/php-fpm-apache:' + php_version + '-buster-local', '.',
  build_args={'PHP_VERSION': php_version},
  dockerfile='Dockerfile.dev',
  live_update = [
    sync('.', '/app'),
  ]
)

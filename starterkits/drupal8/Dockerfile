FROM srijanlabs/php-cli:7.2-1.x as build
ADD --chown=continua:continua . /app
RUN composer install --no-dev --prefer-dist --no-progress --no-suggest --no-interaction --optimize-autoloader

FROM srijanlabs/php-fpm:7.2-1.x
COPY --from=build /app /app

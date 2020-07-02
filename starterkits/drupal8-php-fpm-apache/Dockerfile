ARG PHP_VERSION=7.3
ARG BASE_VERSION=10
FROM srijanlabs/php-cli:${PHP_VERSION}-${BASE_VERSION} as builder
COPY composer.json composer.lock /app/
COPY patches ./patches
RUN composer install --no-dev --prefer-dist --no-progress --no-suggest --no-interaction --optimize-autoloader

FROM srijanlabs/php-fpm-apache:${PHP_VERSION}-${PHP_VERSION} as fpm
COPY --from=builder --chown=continua /app  /app
COPY --chown=continua . /app

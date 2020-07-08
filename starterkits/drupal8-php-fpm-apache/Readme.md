# Drupal 8 Starterkit
This starterkit is based on php-fpm and apache. The default php version is 7.3 which is configurable.
Refer [Configurations](#Configurations) section for further details.

## Configurations
All the configurations are managed via environment variables, some of mostly used environment variables
are documented in the table below.

| Parameter     | File         |   Variable   |   Default    |
|:-------------:|:-------------|:------------:|:------------:|
|   PHP version | .env         | PHP_VERSION  |     7.3      |
|  Project name | .env         | PROJECT_NAME |     drupal8  |

Note: Refer .env file for other configurations.

---
## Installing modules
```bash
docker-compose run --rm cli composer require drush/drush=^9.2
```

recommended alias:
```bash
alias dcomposer="docker-compose run --rm cli composer"
```
---
## Running drush commands
```bash
docker-compose run --rm php drush cr
```
recommended alias:
```bash
alias ddrush="docker-compose run --rm php drush"
```
---
## Profiling
All the starterkits in vega are shipped with blackfire for code profiling. This is also available on
the dev images and is not added to production images.

- Get your blackfire server id and token from https://blackfire.io/my/settings/credentials
- Configure the blackfire server id and token in .env file.

  |        Parameter       |   File  |        Variable        |
  |:----------------------:|:-------:|:----------------------:|
  |   Blackfire Server Id  | .env    | BLACKFIRE_SERVER_ID    |
  | Blackfire Server Token | .env    | BLACKFIRE_SERVER_TOKEN |

- Uncomment blackfire service in docker-compose.yml.
- Run vega up.
- Install blackfire agent extension in your browser.
- Click on the extension on your browser and profile your application.

---
## Debugging:

Follow the document to setup xdebug:
[Setup Xdebug](/starterkits/XDEBUG-SETUP.md)

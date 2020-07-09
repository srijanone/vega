# Drupal 8 Starterkit
This starterkit is based on php-fpm and apache. The default php version is 7.3 which is configurable.
Refer [Configurations](#Configurations) section for further details.

## Configurations
All the configurations are managed via environment variables, some of mostly used environment variables
are documented in the table below.

### Configuring project:
| Parameter     | File         |   Variable   |   Default    |
|:--------------|:-------------|:------------:|:------------:|
|   PHP version | .env         | PHP_VERSION  |     7.3      |
|  Project name | .env         | PROJECT_NAME |     drupal8  |

### Configuring Runtime:

##### These values can be configure using [.env](.env) file

|   Variable                  |   Default             | Parameter    |
|:----------------------------|:----------------------|:-------------|
| DOC_ROOT                    | /app/web              |  Drupal root (docroot) |
| PHP_HOST                    | localhost             |  Host/Service name where php-fpm is running|
| PHP_FPM_PORT                | 9000                  |  The port at which php-fpm is listing |
| PHP_FPM_MAX_CHILDREN        | 20                    |  The maximum number of child processes to be created |
| PHP_FPM_START_SERVERS       | 2                     |  The number of child processes created on startup |
| PHP_FPM_MIN_SPARE_SERVERS   | 2                     |  The desired minimum number of idle server processes |
| PHP_FPM_MAX_SPARE_SERVERS   | 10                    |  The desired maximum number of idle server processes |
| PHP_FPM_MAX_REQUESTS        | 500                   |  The number of requests each child process should execute before respawning | 
| PHP_MEMORY_LIMIT            | 128                   |  PHP memory limit per script  |
| XDEBUG_REMOTE_PORT          | 9001                  |  Port on IDE is listing |
| NEW_RELIC_ENABLED           | false                 |  Enable newrelic |
| NEW_RELIC_APP_NAME          | ''                    |  Application name from newrelic website |
| NEW_RELIC_LICENSE_KEY       | ''                    |  Newrelic license key  |

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

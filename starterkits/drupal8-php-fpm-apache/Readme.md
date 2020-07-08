# Drupal Starterkit

### Configuring project:
| Parameter     | File         |   Variable   |   Default    |
|:-------------:|:-------------|:------------:|:------------:|
|   PHP version | .env         | PHP_VERSION  |     7.3      |
|  Project name | .env         | PROJECT_NAME |     drupal8  |

### Configure Runtime:

##### These values can be configure using [.env](.env) file

|   Variable                  |   Default             | Parameter    |
|:----------------------------|:----------------------|:-------------|
| DOC_ROOT                    | /app/web              |  drupal root |
| PHP_HOST                    | localhost             |  host/service name of php-fpm container |
| PHP_FPM_PORT                | 9000                  |  The port at which php-fpm is listing |
| PHP_FPM_MAX_CHILDREN        | 20                    |  The maximum number of child processes to be created |
| PHP_FPM_START_SERVERS       | 2                     |  The number of child processes created on startup |
| PHP_FPM_MIN_SPARE_SERVERS   | 2                     |  The desired minimum number of idle server processes |
| PHP_FPM_MAX_SPARE_SERVERS   | 10                    |  The desired maximum number of idle server processes |
| PHP_FPM_MAX_REQUESTS        | 500                   |  The number of requests each child process should execute before respawning | 
| PHP_MEMORY_LIMIT            | 128                   |  php memory limit per script  |
| XDEBUG_REMOTE_PORT          | 9001                  |  Port on IDE is listing |
| NEW_RELIC_ENABLED           | false                 |  enable newrelic |
| NEW_RELIC_APP_NAME          | ''                    |  application name from newrelic website |
| NEW_RELIC_LICENSE_KEY       | ''                    |  newrelic license key  |

---
### Installing modules:
```bash
docker-compose run --rm cli composer require drush/drush=^9.2
```
recommended alias:
```bash
alias dcomposer="docker-compose run --rm cli composer"
```
---
### Running drush commands:exit
```bash
docker-compose run --rm php drush cr
```
recommended alias:
```bash
alias ddrush="docker-compose run --rm php drush"
```


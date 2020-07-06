# Drupal Starterkit

### Configuring PHP version
| Parameter     | File         |   Variable   |   Default    |
|:-------------:|:-------------|:------------:|:------------:|
|   PHP version | .env         | PHP_VERSION  |     7.3      |
|  Project name | .env         | PROJECT_NAME |     drupal8  |


Note: Most of the configurable values are available in .env file.

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
### Running drush commands:
```bash
docker-compose run --rm php drush cr
```
recommended alias:
```bash
alias ddrush="docker-compose run --rm php drush"
```


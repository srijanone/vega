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
---
### Profiling:
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
 Click on the extension on your browser and profile your application.

---
### Debugging:

Please follow the document to setup xdebug:
[Setup Xdebug](/starterkits/XDEBUG-SETUP.md)

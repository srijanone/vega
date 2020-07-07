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
### Configure Blackfire:
1. Login to https://blackfire.io/
2. Go to https://blackfire.io/my/settings/credentials
3. Find Server ID and Server Token under My Server Credentials Section.
4. Configure the blackfire server id and token in .env file under your vega project as given below.

|        Parameter       |   File  |        Variable        |
|:----------------------:|:-------:|:----------------------:|
|   Blackfire Server Id  | .env    | BLACKFIRE_SERVER_ID    |
| Blackfire Server Token | .env    | BLACKFIRE_SERVER_TOKEN |

5. Open docker-compose.yml under your vega project.
6. Uncomment blackfire service under blackfire section.
7. Run vega up.
8. Go to your browser and install Blackfire agent extension.
9. Open your site: http://localhost:8080/
10. Click on the extension on your browser and profile your application.

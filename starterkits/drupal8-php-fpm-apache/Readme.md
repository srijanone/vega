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

---
## NOTE:
Due to an issue as of now PHP_VERSION has to be updated in Tiltfile as well.
---

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
| PHP_UPLOAD_MAX_FILESIZE     | 2                     |  PHP max upload size  |
| PHP_MAX_EXECUTION_TIME      | 180                   |  PHP max execution time  |
| XDEBUG_REMOTE_PORT          | 9001                  |  Port on IDE is listing |
| NEW_RELIC_ENABLED           | false                 |  Enable newrelic |
| NEW_RELIC_APP_NAME          | ''                    |  Application name from newrelic website |
| NEW_RELIC_LICENSE_KEY       | ''                    |  Newrelic license key  |

---
## Setting up drupal site
On non linux platform performance is a big challenge with drupal sites.
To mitigate this issue, installation process of starterkits is divided into two steps given below.
```bash
docker-compose run --rm cli composer install
vega up
```
You can install profile via web browser but it is recommended to do it using drush.
Run below command to install profile.
```bash
docker-compose run --rm drupal drush site:install PROFILE_NAME --account-name USERNAME --account-mail EMAIL --account-pass PASSWORD --site-name SITE_NAME --site-mail SITE_EMAIL -y
```
Replace the PROFILE_NAME, USERNAME, EMAIL, PASSWORD, SITE_NAME & SITE_EMAIL with actuals in previous command.

@NOTE:
For performance improvement we have used live update feature of tilt. In result of which we got one or two extra tags of existing image created by tilt to perform live update of code changes in host.
In virtue of that, after doing vega down these extra tags needs to be cleaned up. For cleanup run below command in terminal
```bash
 docker image rm -f $(docker image ls -q --filter "label=builtby=tilt")
```
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
docker-compose run --rm drupal drush cr
```
recommended alias:
```bash
alias ddrush="docker-compose run --rm drupal drush"
```
---
## NOTE:
```
Problem:
  On Linux OS, during composer install cli service throws
  a permission denied error while installing dependencies.

Solution:
  Run the cli service with root user and run chown service after to change the ownership.
  This will prevent the error throws during the composer install.
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

## Mailhog

This starterkit ships with mailhog which can be accessed at localhost:8025. This can be a handy feature
to test email functionality on local systems without worrying about sending mails accidently to real users.
Mailhog is only available on dev images which are used for local development and not available on production
images.

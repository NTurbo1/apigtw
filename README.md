# Reverse proxy
This is an educational project to understand and learn about reverse proxy internals, load balancing,
rate limiting, caching, and so on. No frameworks and third-party/external libraries except the go's
standard library are used unless extremely necessary.
### Content
- [Configuration](#Configuration)
- [Flags](#Flags)
- [Environment Variables](#Environment-Variables)
- [License](#License)

## Configuration
- We use json files for configuration files just because we don't have to implement a parser from
scratch for that thanks to the go's stdlib.
- Default configuration filepaths are set in the [environment variables](#Environment-Variables). You should edit
the variables if you want to have custom configuration filepaths. You can check the default values [here](#Environment-Variables)
- You can specify the filepath for the main app config file using the flag `--app-config-file` which
updates the value of the env variable `REVERSE_PROXY_APP_CONFIG_FILE`, which has a default value 
`app.configs.json`. ALL CONFIGURATION FILE FORMATS ARE EXPECTED TO BE JSON! ANY CONFIGURATION VALUE
SPECIFIED IN THE MAIN APP CONFIGS FILE TAKES PRECEDENCE OR OVERWRITES ALL THE OTHER CORRESPONDING VALUES
SPECIFIED ELSEWHERE, INCLUDING THE ENVIRONMENT VARIABLES!
- `envFile` field in the app config json file contains the filepath of a json file that represents the 
application environment (like in NodeJS).
- You can enable debugging mode in the logs by including `--debug` flag in the command line. More info
about the flags are [here](#Flags)

## Flags
|Name|Description|Required|Usage|Default|
|:--:|:---------:|:------:|:---:|:-----:|
|`--app-config-file`|Edits the value of the env variable `REVERSE_PROXY_APP_CONFIG_FILE`|false|`--app-config-file=<filepath>`|Check [Environment Variables](#Environment-Variables)|
|`--debug`|Edits the value of the env variable `REVERSE_PROXY_LOG_LEVEL`|false|`--debug`|Check [Environment Variables](#Environment-Variables)|

## Environment Variables
|Name|Description|Required|Default|
|:--:|:---------:|:------:|:-----:|
|PORT|The port the server runs on|false|8080|
|REVERSE_PROXY_APP_CONFIG_FILE|Filepath to the main application configs file|false|app.configs.json|
|ROUTES_CONFIGS_MASTER_FILE|Filepath to the routes configs master files that includes filepaths<br>to the rest of the routes configs files|false|routes.master.json|
|REVERSE_PROXY_LOG_LEVEL|The log level throughout the application runtime|false|INFO|


## License
- [MIT License](LICENSE)

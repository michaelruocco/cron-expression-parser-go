# Cron Expression Parser

[![Build](https://github.com/michaelruocco/cron-expression-parser-go/workflows/pipeline/badge.svg)](https://github.com/michaelruocco/cron-expression-parser-go/actions)
[![BCH compliance](https://bettercodehub.com/edge/badge/michaelruocco/cron-expression-parser-go?branch=master)](https://bettercodehub.com/)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=michaelruocco_cron-expression-parser-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=michaelruocco_cron-expression-parser-go)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=michaelruocco_cron-expression-parser-go&metric=sqale_index)](https://sonarcloud.io/dashboard?id=michaelruocco_cron-expression-parser-go)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=michaelruocco_cron-expression-parser-go&metric=coverage)](https://sonarcloud.io/dashboard?id=michaelruocco_cron-expression-parser-go)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=michaelruocco_cron-expression-parser-go&metric=ncloc)](https://sonarcloud.io/dashboard?id=michaelruocco_cron-expression-parser-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This was my first attempt at using Go. I already had an existing implementation for this problem written
in [Java](https://github.com/michaelruocco/cron-expression-parser-java) so I decided to have a go at porting it across into Go in order to use it as an excuse to get started using it.

Parses Cron Expressions of the following format:

*   (Minute) (hour) (day of month) (month) (day of week) (command)
*   `*` means all possible time units
*   `-` a range of time units
*   `,` a comma separated list of individual time units
*   `/` intervals time units, the left value is the starting value and the right value is the max value

For example given the input argument:

`3,45/15 0 1,15 * MON-FRI /usr/bin/find`

The output should be:

```bash
minute        3 18 33 45 48
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   0 1 2 3 4
command       /usr/bin/find
```

## Installing modules

`go install`

## Running Unit Tests

`go test -v ./... -cover`

## Running Example

`go run main.go 3,45/15 0 1,15 * 0-4 /usr/bin/find`

or

`go run main.go -arguments 3,45/15 0 1,15 * MON-FRI /usr/bin/find`

If you are planning to pass any arguments with `*` notation you may
need to configure your terminal to disable globbing. I found that using zsh shell on my mac this was
and issue. If I try to run either of the commands above I get the following output:

`notation parser not found for value domain`

As you can see the * command has been expanded to include all the README.md file
from the current directory where the command is being run from. This can be fixed by
running the following command to disable globbing:

`set -o noglob`

## TODO

*   Refactor parseExpression method of cronExpressionParser?
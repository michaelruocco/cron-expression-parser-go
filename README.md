# Cron Parser

Parses Cron Expressions of the following format:

*   (Minute) (hour) (day of month) (month) (day of week) (command)
*   `*` means all possible time units
*   `-` a range of time units
*   `,` a comma separated list of individual time units
*   `/` intervals time units, the left value is the starting value and the right value is the max value

For example given the input argument:

`3,45/15 0 1,15 * 1-5 /usr/bin/find`

The output should be:

```bash
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
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

# TODO

*  Add github actions build pipeline
*  Rework time units to split into separate files for hours / minutes etc?
*  Refactor parseExpression method of cronExpressionParser?
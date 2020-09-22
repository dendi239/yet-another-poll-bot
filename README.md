# yet-another-poll-bot

![Tests](https://github.com/dendi239/yet-another-poll-bot/workflows/go%20test/badge.svg)
![Deploy](https://github.com/dendi239/yet-another-poll-bot/workflows/Heroku/badge.svg)

Telegram bot for polls with restrictions

## Design

### Supported methods

#### `/create_poll`

#### `/add_option`

#### `/add_restrictions`

restriction should contain some command from `logical operators` list.

#### `/set_description`

#### `/set_title`

### Grammar

#### Logical operators

- `^` - exactly one of arguments should be true
- `|` - at least one of the arguments should be true
- `&` - all arguments should be true
- `!` - negates the argument

Arguments are comma-separated list of `values`.

#### Values

1. Integer, representing poll option(1, 2, 3, ... len(options)).
2. ^(values...)
3. &(values...)
4. |(values...)
5. !(value)

Sample poll:





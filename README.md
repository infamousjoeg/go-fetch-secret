# CyberArk Conjur Secret Fetcher

A small footprint Go-based binary that fetches a secret from a DAP or Conjur service.

## Installation

Download the [latest release](https://github.com/infamousjoeg/go-secret-fetcher/releases).

Move or copy the `fetchsecret` binary to `/usr/bin/local` or another path designed in `$PATH`.

## Usage

`$ fetchsecret path/to/secret`

## Dependency

Uses the [conjur-api-go](https://github.com/cyberark/conjur-api-go) library.

## License

[MIT](LICENSE)
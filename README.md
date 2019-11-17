# CyberArk Conjur Secret Fetcher

A small footprint Go-based binary that fetches a secret from a DAP or Conjur service.

## Installation

Download the [latest release](https://github.com/infamousjoeg/go-secret-fetcher/releases).

Move or copy the `fetchsecret` binary to `/usr/bin/local` or another path designed in `$PATH`.

## Required Environment Variables

* `CONJUR_APPLIANCE_URL` this is the URL to access your DAP or Conjur service
* `CONJUR_ACCOUNT` this is the Org Account created during initial deployment
* `CONJUR_AUTHN_LOGIN` this is the username or host ID to use for authentication
* `CONJUR_AUTHN_API_KEY` this is the API key associated with `CONJUR_AUTHN_LOGIN`

_Note: It is recommended to use [Summon](https://cyberark.github.io/summon) to inject environment variables securely to `fetchsecret`._

## Usage

After all environment variables are set...

`$ fetchsecret path/to/secret`

## Dependency

Uses the [conjur-api-go](https://github.com/cyberark/conjur-api-go) library.

## License

[MIT](LICENSE)
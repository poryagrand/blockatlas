# Block Atlas by Trust Wallet

![Go Version](https://img.shields.io/github/go-mod/go-version/TrustWallet/blockatlas)
![CI](https://github.com/trustwallet/blockatlas/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/trustwallet/blockatlas/branch/master/graph/badge.svg)](https://codecov.io/gh/trustwallet/blockatlas)
[![Go Report Card](https://goreportcard.com/badge/trustwallet/blockatlas)](https://goreportcard.com/report/TrustWallet/blockatlas)

> BlockAtlas is a clean explorer API and transaction observer for cryptocurrencies.

BlockAtlas connects to nodes or explorer APIs of the supported coins and maps transaction data,
account transaction history into a generic, easy to work with JSON format.
It is in production use at the [Trust Wallet app](https://trustwallet.com/), 
the official cryptocurrency wallet of Binance. Also is in production at the [BUTTON Wallet](https://buttonwallet.com), Telegram based non-custodial wallet.
The observer API watches the chain for new transactions and generates notifications by guids.

#### Supported Coins

<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/binance/info/logo.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/nimiq/info/logo.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ripple/info/logo.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/stellar/info/logo.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/kin/info/logo.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tezos/info/logo.png" width="32" /></a>
<a href="https://aion.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/aion/info/logo.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/info/logo.png" width="32" /></a>
<a href="https://ethereumclassic.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/classic/info/logo.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/poa/info/logo.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/callisto/info/logo.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/gochain/info/logo.png" width="32" /></a>
<a href="https://wanchain.org" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/wanchain/info/logo.png" width="32" /></a>
<a href="https://thundercore.com" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/thundertoken/info/logo.png" width="32" /></a>
<a href="https://icon.foundation" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/icon/info/logo.png" width="32" /></a>
<a href="https://tron.network" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tron/info/logo.png" width="32" /></a>
<a href="https://vechain.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/vechain/info/logo.png" width="32" /></a>
<a href="https://www.thetatoken.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/theta/info/logo.png" width="32" /></a>
<a href="https://cosmos.network/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/info/logo.png" width="32" /></a>
<a href="https://bitcoin.org/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/bitcoin/info/logo.png" width="32" /></a>
<a href="https://harmony.one/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/harmony/info/logo.png" width="32" /></a>
<a href="https://elrond.com/" target="_blank"><img src="https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/elrond/info/logo.png" width="32" /></a>

## Architecture

#### NOTE
Currently Block Atlas is under active development and is not well documented. If you still want to run it on your own or help to contribute, **please** pay attention that currently integration, nemwan, functional tests are not working locally without all endpoints. We are fixing that issue and soon you will be able to test all the stuff locally

Blockatlas allows to:
- Get information about transactions, tokens, staking details, collectibles, crypto domains for supported coins.
- Subscribe for price notifications via Rabbit MQ

Platform API is independent service and can work with the specific blockchain only (like Bitcoin, Ethereum, etc)

Notifications:

(Observer Subscriber Producer) - Create new blockatlas.SubscriptionEvent [Not implemented at Atlas, write it on your own]

(Observer Subscriber) - Get subscriptions from queue, set them to the DB

(Observer Parser) - Parse the block, convert block to the transactions batch, send to queue

(Observer Notifier) - Check each transaction for having the same address as stored at DB, if so - send tx data and id to the next queue

(Observer Notifier Consumer) - Notify the user [Not implemented at Atlas, write it on your own]

```
New Subscriptions --(Rabbit MQ)--> Subscriber --> DB
                                                   |
                      Parser  --(Rabbit MQ)--> Notifier --(Rabbit MQ)--> Notifier Consumer --> User

```

The whole flow is not available at Atlas repo. We will have integration tests with it. Also there will be examples of all instances soon.

## Setup

### Prerequisite
 * [Go Toolchain](https://golang.org/doc/install) versions 1.14+
 
 Depends on what type of Blockatlas service you would like to run will also be needed.
 * [Postgres](https://www.postgresql.org/download) to store user subscriptions and latest parsed block number
 * [Rabbit MQ](https://www.rabbitmq.com/#getstarted) to pass subscriptions and send transaction notifications

### Quick Start

#### Get source code

Download source to `GOPATH`
```shell
go get -u github.com/trustwallet/blockatlas
cd $(go env GOPATH)/src/github.com/trustwallet/blockatlas
```

#### Build and run

Read [configuration](#configuration) info

```shell
# Start Platform API server at port 8420 with the path to the config.yml ./
go build -o platform-api-bin cmd/platform_api/main.go && ./platform-api-bin -p 8420

# Start observer_parser with the path to the config.yml ./ 
go build -o observer_parser-bin cmd/observer_parser/main.go && ./observer_parser-bin

# Start observer_notifier with the path to the config.yml ./ 
go build -o observer_notifier-bin cmd/observer_notifier/main.go && ./observer_notifier-bin

# Start observer_subscriber with the path to the config.yml ./ 
go build -o observer_subscriber-bin cmd/observer_subscriber/main.go && ./observer_subscriber-bin

# Startp Swagger API server at port 8422 with the path to the config.yml ./ 
go build -o swagger-api-bin cmd/swagger-api/main.go && ./swagger-api-bin -p 8423

# Start Platform API server with mocked config, at port 8437 ./ 
go build -o platform-api-bin cmd/platform_api/main.go && ./platform-api-bin -p 8437 -c configmock.yml
```

### make command

Build and start all services:
```shell
make go-build
make start
```

Build and start individual service:
```shell
make go-build-platform-api
make start
```

### Docker

Build and run all services:

```shell
docker-compose build
docker-compose up
```

Build and run individual service:
```shell
docker-compose build swagger_api
docker-compose start swagger_api
```

## Configuration
When any of Block Atlas services started they look up inside [default configuration](./config.yml).
Most coins offering public RPC/explorer APIs are enabled, thus Block Atlas can be started and used right away, no additional configuration needed.
By default starting any of the [services](#architecture) will enable all platforms

To run a specific service only by passing environmental variable, e.g: `platfrom_api` :
```shell
ATLAS_PLATFORM=ethereum go run cmd/platform_api/main.go

ATLAS_PLATFORM=ethereum binance bitcoin go run cmd/platform_api/main.go # for multiple platforms
```

or change in config file
```yaml
# Single
platform: [ethereum]
# Multiple 
platform: [ethereum, binance, bitcoin]
```

This way you can one platform per binary, for scalability and sustainability.

To enable use of private endpoint:
```yaml
nimiq:
  api: http://localhost:8648
```
It works the same for observer_worker - you can run all observer at 1 binary or 30 coins per 30 binaries

#### Environment

The rest gets loaded from environment variables.
Every config option is available under the `ATLAS_` prefix. Nested keys are joined via `_`.

Example:

```shell
ATLAS_NIMIQ_API=http://localhost:8648
```

## Tests

### Unit tests
```
make test
```
### Mocked tests

End-to-end tests with calls to external APIs has great value, but they are not suitable for regular CI verification, beacuse any external reason could break the tests.

Therefore mocked API-level tests are used, whereby external APIs are replaced by mocks.

* External mocks are implemented as a simple, own, golang `mockserver`.  It listens locally, and returns responses to specific API paths, taken from json data files.
* There is a file where API paths and corresponding data files are listed.
* Tests invoke into blockatlas through public APIs only, and are executed using *newman* (Postman cli -- `make newman-mocked`).
* Product code, and even test code should not be aware whether it runs with mocks or the real external endpoints.
* See Makefile for targets with 'mock'; platform can be started locally with mocks using `make start-platform-api-mock`.
* The newman tests can be executed with unmocked external APIs as well, but verifications may fail, because some APIs return variable responses.  Unmocked tests are not intended for regular CI execution, but as ad-hoc development tests.
* General steps for creating new mocked tests: replace endpoint to localhost:3347, observe incoming calls (visible in mockserver's output), obtain real response from external API (with exact same parameters), place response in a file, add path + file to data file list.  Restart mock, and verify that blockatlas provides correct output.  Also, add verifications of results to the tests.

## Docs

Swagger API docs provided at path `/swagger/index.html`

or you can install `go-swagger` and render it locally (macOS example)

Install:

```shell
brew tap go-swagger/go-swagger
brew install go-swagger
```

Render: 
```shell
swagger serve docs/swagger.yaml
```

#### Updating Docs

- After creating a new route, add comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).

- Run `$ make go-gen-docs` in root folder.

## Contributing

If you'd like to add support for a new blockchain, feel free to file a pull request.
Note that most tokens that run on top of other chains are already supported and
don't require code changes (e.g. ERC-20).

The best way to submit feedback and report bugs is to open a GitHub issue.
Please be sure to include your operating system, version number, and
[steps](https://gist.github.com/nrollr/eb24336b8fb8e7ba5630) to reproduce reported bugs.

More resources for developers are in [CONTRIBUTING.md](CONTRIBUTING.md).

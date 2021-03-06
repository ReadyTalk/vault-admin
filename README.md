# Vault Admin [![Build Status](https://travis-ci.org/PremiereGlobal/vault-admin.svg?branch=master)](https://travis-ci.org/PremiereGlobal/vault-admin)

This utility configures Vault audit devices, auth methods, policies and secrets engines by syncing with a set of standard JSON configuration files.

## Installation

This utility can be used via Docker or the CLI.

### CLI
Download and extract the latest binary for your OS on the [releases page](https://github.com/PremiereGlobal/vault-admin/releases)

Run `vadmin <flags>`.  See below for a description of the command line flags.

### Docker
The Docker container must be run in interactive mode with the `-it` parameter because it prompts for things like policy deletion, etc.

```
docker run \
  --rm \
	-it \
	-e VAULT_ADDR=https://vault.mysite.com:8200 \
	-e VAULT_TOKEN=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
	-v $(pwd)/config:/config
	premiereglobal/vault-admin:latest
```

Map wherever you have your Vault Admin configuration files to `/config` within the container.

## Options
All options can be set via environment variables or command line options

| Environment Variable               | Command Line Flags | Description                           |
| ----------------------- | ----------------------------------    | ---------------------------------------------------------- |
| `CONFIGURATION_PATH` | --configuration-path, -c | Path to the configuration files |
| `VAULT_ADDR` | --vault-addr, -a | Vault address (example: https://vault.mysite.com:8200) |
| `VAULT_TOKEN` | --vault-token, -t | Vault token to use |
| `VAULT_SKIP_VERIFY` | --vault-skip-verify, -K | Skip Vault TLS certificate verification |
| `VAULT_SECRET_BASE_PATH`  | --vault-secret-base-path, -s | Base secret path, in Vault, to pull secrets for substitution. Defaults to `secret/vault-admin` |
|   | --rotate-creds, -r | Perform key rotation on AWS secret engines |
| `DEBUG`  | --debug, -d | Turn on debug logging |
|   | --version, -v | Show version information |

## Configuration Files
The configuration files are what drive how Vault is configured.  See the [examples/](examples/) directory for more information on how to set up the configuration.

# secrets

Client Example for Managed Secrets - Relies on:

 * libs in https://github.com/scribd/vault-authenticator
 
 * vault configured by https://github.com/scribd/keymaster-cli

This repo is a reference implementation of a client for the public libraries listed above.

## Usage

Designed to be run from `dbt` https://github.com/nikogura/dbt, but can also be run as a bare binary.

## DBT Usage

For usage with `dbt` run the following to get subcommands and options: (assuming you have `dbt` set up in your organization of course)

    dbt secrets help
    
## Bare Usage

1. Fetch the version of your choice from wherever you keep the binanaries.

2. Run with:

    ./secrets help
    
As configured, this tool will attempt to authenticate you via LDAP if it cannot do so any other way. 

On a laptop this is cool, but when used in an automated fashion it's probably not what you want. Specify 'silent' mode by adding the -s flag thusly:

    [dbt] secrets <verb> -s
        
### Fetching Secrets in K8s

From within a pod in the appropriate Namespace in K8s, run:

    [dbt] secrets fetch -si <cluster> -r <team>-<role>
        
Fetching Secrets with TLS Auth
 
    secrets fetch -sr <team>-<role>-<environment>
    
TLS Authed  Hosts need to suffix the role name with the environment name. It's annoying, but that's the way the TLS auth backend was written by Hashicorp. Blame Mitch not me. Or if you think up a way around it, submit a PR. :D

### Fetch returning a single key from the Secret

    [dbt] secrets fetch -s -r <team>-<role> -k <key>

Fetch returning a single key from the Secret returning just the raw value associated with that key

    [dbt] secrets fetch -s -r <team>-<role> -k <key> -f raw

Exec With Secrets in K8s

    [dbt] secrets exec -si <cluster name> -r <team>-<role> <program to exec>
        
Your secrets will appear in ENV to be consumed by the program you have supplied as the final argument to the command.

Exec With Secrets with TLS Auth

    [dbt] secrets exec -sr <team>-<role>-<environment> <program to exec>
    
Your secrets will appear in ENV to be consumed by the program you have supplied as the final argument to the command.

Debugging Secret Access

If `[dbt] secrets fetch ... `or `[dbt] secrets exec ...` fail to delight you, try running again with the -v flag (verbose). 

It will dump a very large amount of information regarding what it's connecting to, and how it's trying to auth, and the like.  Happy Hacking!
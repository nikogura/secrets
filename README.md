# secrets

Client Example for Managed Secrets - Relies on:

 * libs in https://github.com/nikogura/vault-authenticator
 
 * vault configured by https://github.com/nikogura/keymaster-cli

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

On a laptop this is cool, but when used in an automated fashion it's probably not what you want. Specify 'silent' mode by adding the -s flag:

    [dbt] secrets <verb> -s
        
## Automated Usage

## Fetching Secrets with the IAM Realm

From an EC2 machine, run:

    [dbt] secrets fetch -r <team>-<role>

## Fetching Secrets with the Kubernetes Realm

From within a pod in the appropriate Namespace in K8s, run:

    [dbt] secrets fetch -i <cluster> -r <team>-<role>
    
## Fetching Secrets with the TLS Realm

From a machine possessing a certificate signed by your root CA, run:

    [dbt] secrets fetch -r <team>-<role>-<environment>

## Fetching Secrets using LDAP

From any machine that can connect to the LDAP server, run:

`[dbt] secrets fetch -r <team>-<role> -e <environment>`

## Fetch a single key-value pair from the Secret

   Append `-k <key>`
    
## Fetch a single value from the Secret

   Append `-k <key> -f raw`

## Fetch returning a single key from the Secret in a formatted presentation for stdout

   Append `-k <key>`
    
## Fetch returning a single key from the Secret returning the unformatted raw value

   Append `-k <key> -f raw`

## Executing A Command With Secrets

This example is for k8s. Other realms are similar to `fetch` syntax, above

    [dbt] secrets exec -i <cluster name> -r <team>-<role> '<program to exec>'
    
Your secrets will appear in ENV to be consumed by the program you have supplied as the final argument to the command.

## A note on "backup" LDAP auth

The `secrets` tool will attempt to authenticate you via your personal LDAP credentials if it cannot do so via k8s, tls, or iam.  In an interactive shell during development, this can be useful, but when used in an automated fashion, it will cause your script to hang (or exit), waiting for an LDAP password that will never arrive.  You can turn off the LDAP auth attempt in a script with the 'silent' switch:

    [dbt] secrets <verb> -s
    
which can be combined with the `-r` or `-i` switches for tls and k8s realm types, respectively:

    [dbt] secrets exec -sr <team>-<role>-<environment> '<program to exec>'
    
## Debugging Secret Access

If `[dbt] secrets fetch ... `or `[dbt] secrets exec ...` fail to delight you, try running again with the -v flag (verbose). 

It will dump a very large amount of information regarding what it's connecting to, and how it's trying to auth, and the like.  Happy Hacking!

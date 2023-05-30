# An Alfred 5 Workflow for Pass

## Original Author and Repository
- CGenie
- [Original repository](https://github.com/CGenie/alfred-pass)

## Rewritten for Alfred 5
The goal is to rewrite the original workflow to make it work with Alfred 5 and the minimun dependencies.

## Development

### Dependencies to build
- Golang 1.20.4 `brew install go`

### Dependencies to run
- pass 1.7.4 `brew install pass`
- pass-otp 1.2.0 `brew install pass-otp`
- alfred 5 `brew install alfred`
- gpg, and gpg-agent 2.4.1 `brew install gnupg`
- pinentry-mac 1.1.1 `brew install pinentry-mac` (this is GUI frontend for `gpg-agent`).
  - Get path for `pinentry-mac` with `which pinentry-mac`
  - Next configure `gpg-agent` to use `pinentry-mac` and not the bundled one, editing `~/.gnupg/gpg-agent.conf`:
    ```
    pinentry-program /path/to/pinentry-mac
    ```
### GPG tweaking

You can tweak some of the `gpg-agent` settings in `~/.gnupg/gpg-agent.conf`:

```
max-cache-ttl 7200
```

After 7200 seconds, GPG will forget your master password.


### Build for development
- `make dev` basically it will run the build for production and then it will open the workflow in Alfred

### Build for production
- `make`

## Usage

Basic Alfred commands:

## `pass <filter terms>`

This will search through your passwords using the filter terms you provided.

The password will be copied to clipboard and cleared after 45 seconds (this is the default
`pass -c` behavior).  You can change that time by modifying the env variable
`PASSWORD_STORE_CLIP_TIME`. Or in the `pass-show.sh` file you can change this line

```
pass show -c $QUERY
```

into this one

```
pass show $QUERY | awk 'BEGIN{ORS=""} {print; exit}' | pbcopy
```

to aviod auto-clearing of clipboard.

## `pg <id>`

Calls `pass generate` to add a new password with default length of 20 chars.

## `po <filter terms>`

This will search through your OTP passwords (requires `pass-otp`) using the filter terms you provided.

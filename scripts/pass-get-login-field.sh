#!/bin/bash

set -e

QUERY=$1
PATH=/opt/homebrew/bin:$HOME/.brew/bin:/usr/local/bin:$PATH
PASSWORD_STORE_ENABLE_EXTENSIONS=true
PASSWORD_STORE_EXTENSIONS_DIR=$HOME/.password-store/.extensions

build/pasawutil get-login "$QUERY"

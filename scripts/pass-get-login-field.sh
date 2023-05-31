#!/bin/bash

set -e

QUERY=$1
PATH=/opt/homebrew/bin:$HOME/.brew/bin:/usr/local/bin:$PATH

build/pasawutil get-login "$QUERY"
osascript -e 'display notification "Copied login field to clipboard" with title "Alfred x pass"'

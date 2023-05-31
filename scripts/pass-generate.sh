#!/bin/bash

set -e

QUERY=$1
PATH=/opt/homebrew/bin:$HOME/.brew/bin:/usr/local/bin:$PATH

pass generate "$QUERY" -n 20 -c 
osascript -e 'display notification "Copied generated password to clipboard" with title "Alfred x pass"'
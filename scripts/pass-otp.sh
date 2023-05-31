#!/bin/bash

set -e

QUERY=$1
PATH=/opt/homebrew/bin:$HOME/.brew/bin:/usr/local/bin:$PATH

pass otp -c "$QUERY"
osascript -e 'display notification "Copied OTP key to clipboard" with title "Alfred x pass"'

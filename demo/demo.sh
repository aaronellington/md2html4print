#!/usr/bin/env bash
set -Eeuox pipefail

cd demo/src
go run ../.. > index.html

open index.html

curl \
    --request POST https://demo.gotenberg.dev/forms/chromium/convert/html \
    --form files=@index.html \
    -o index.pdf

open index.pdf

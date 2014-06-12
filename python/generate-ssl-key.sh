#!/bin/bash

domain="$1"
openssl genrsa -des3 -out $domain.key 2048
openssl req -new -key $domain.key -out $domain.csr
cp $domain.key $domain.key.org
openssl rsa -in ${domain}.key.org -out ${domain}.key
openssl x509 -req -days 3650 -in ${domain}.csr -signkey $domain.key -out $domain.crt



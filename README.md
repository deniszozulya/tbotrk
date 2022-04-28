#!/bin/bash
#
# Generate self-signed certificates for Telegram Bot API
# (https://core.telegram.org/bots/self-signed)
# 
# created on : 2015.10.08.
# last update: 2015.10.11.
# 
# by meinside@gmail.com

DOMAIN="my.host.com"	# XXX - edit
EXPIRE_IN=3650	# XXX - edit

NUM_BITS=2048	# 2048 bits
C="US"
ST="New York"
L="Brooklyn"
O="Example Company"

PRIVATE_KEY="cert.key"
CERT_PEM="cert.pem"

openssl req -newkey rsa:$NUM_BITS -sha256 -nodes -keyout $PRIVATE_KEY -x509 -days $EXPIRE_IN -out $CERT_PEM -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$DOMAIN"

# finished
echo "> finished - certificate: $CERT_PEM / private key: $PRIVATE_KEY"
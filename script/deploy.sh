#!/usr/bin/env bash

scp -i ./deploy_key main.o $MACHINE_USERNAME@$MACHINE_ADDRESS:back/main.o
ssh -i ./deploy_key $MACHINE_USERNAME@$MACHINE_ADDRESS <<EOF
sudo systemctl restart tarantool-api.service
EOF

#!/usr/bin/env bash

echo "Coping..."
scp -i ./deploy_key main.o $MACHINE_USERNAME@$MACHINE_ADDRESS:back/main.o
echo "Restarting..."
ssh -i ./deploy_key $MACHINE_USERNAME@$MACHINE_ADDRESS <<EOF
sudo systemctl restart tarantool-api.service
EOF
echo "Finished!"

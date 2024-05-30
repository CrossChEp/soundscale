#!/bin/sh
SCRIPT_PATH=$(realpath "$(dirname "$0")")
if [ "$(pwd)" != "$SCRIPT_PATH" ]; then
    echo YOU MUST RUN THIS SCRIPT FROM SCRIPT DIRECTORY!
    exit
fi

echo Installing metallb...
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml \
|| (echo Failed to install metallb! && exit)
echo Done!


echo Waiting for metallb webhook-service to apply metallb configs...
kubectl wait --timeout 30s --for condition=available --namespace metallb-system deployment/controller
kubectl apply -f metallb-conf.yml
echo Metallb configs applied!
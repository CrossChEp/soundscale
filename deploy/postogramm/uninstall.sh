#!/bin/sh

echo This will delete metallb and istio!
printf "%s" "Do you want to continue? [y/N]: "
read -r answer
if [ "$(echo "$answer" | tr '[:upper:]' '[:lower:]')" != 'y' ]; then
    exit
fi

kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml 
# kubectl delete namespace istio-system
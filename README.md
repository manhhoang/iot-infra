# iot-infra

## How to run

```
export CLUSTER_NAME="kafka-events" #You can change this to whatever you want.
export AV_KEY="[your_AlphaVantage_API_key]"
chmod +x scripts/setup.sh
./scripts/setup.sh

```

## Clean resources

```
chmod +x scripts/cleanup.sh
./scripts/cleanup.sh
```

## Display logs of event-display

```
kubectl logs event-display-ftxzd-deployment-6856647cc9-njdh7 -n kafka-eventing -c user-container
```

## Decode data
From the logs, find the Data section

Data,
  "MTA5LjgyMDAwMDAw"
  
```
echo `echo MTA5LjgyMDAwMDAw | base64 --decode`
```
# iot-infra

## How to run

```
export CLUSTER_NAME="kafka-events" #You can change this to whatever you want.
export AV_KEY="[your_AlphaVantage_API_key]"
chmod +x scripts/setup.sh
./scripts/setup.sh

chmod +x scripts/cleanup.sh
./scripts/cleanup.sh
```
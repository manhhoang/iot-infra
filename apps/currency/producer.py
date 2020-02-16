import os
from avdemo.avdemo import AlphaKafka


CONFLUENT_HOST = os.environ.get("CONFLUENT_HOST")
CONFLUENT_KEY = os.environ.get("CONFLUENT_KEY")
CONFLUENT_SECRET = os.environ.get("CONFLUENT_SECRET")
ALPHAVANTAGE_KEY = os.environ.get("ALPHAVANTAGE_KEY")

origin = AlphaKafka(CONFLUENT_HOST, CONFLUENT_KEY, CONFLUENT_SECRET, ALPHAVANTAGE_KEY)

if __name__ == "__main__":
    msg = origin.produce("test", "0.123")
    print(msg)

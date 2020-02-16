import uuid
from confluent_kafka import Consumer
import json


if __name__ == "__main__":

    def acked(err, msg):
        """Delivery report callback called (from flush()) on successful or failed delivery of the message."""
        if err is not None:
            report = "failed to deliver message: {}".format(err.str())
        else:
            report = "produced to: {} [{}] @ {}".format(
                msg.topic(), msg.partition(), msg.offset())
        return report

    consumer = Consumer({
        'bootstrap.servers': '',
        'sasl.mechanisms': 'PLAIN',
        'security.protocol': 'SASL_SSL',
        'sasl.username': '',
        'sasl.password': '',
        'group.id': str(uuid.uuid1()),
        'auto.offset.reset': 'earliest'
    })

    consumer.subscribe(['test'])

    total_count = 0
    while True:
        msg = consumer.poll(timeout=1.0)
        if msg is None:
            continue
        if msg.error():
            print('error: {}'.format(msg.error()))
        else:
            record_key = msg.key()
            record_value = msg.value()
            data = json.loads(record_value)
            print("Consumed record with key {} and value {} and data {}".format(record_key, record_value, data))




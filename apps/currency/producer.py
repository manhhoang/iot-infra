from confluent_kafka import Producer


if __name__ == "__main__":

    def acked(err, msg):
        """Delivery report callback called (from flush()) on successful or failed delivery of the message."""
        if err is not None:
            report = "failed to deliver message: {}".format(err.str())
        else:
            report = "produced to: {} [{}] @ {}".format(
                msg.topic(), msg.partition(), msg.offset())
        return report

    producer = Producer({
        'bootstrap.servers': '',
        'sasl.mechanisms': 'PLAIN',
        'security.protocol': 'SASL_SSL',
        'sasl.username': '',
        'sasl.password': '',
    })
    msg = producer.produce('test', value='1', on_delivery=acked)
    producer.flush()
    print(msg)




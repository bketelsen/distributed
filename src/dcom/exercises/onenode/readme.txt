1. Start a nats server by running the "nats-docker.sh" script in this
directory, or by downloading NATS at http://nats.io/download/nats-io/gnatsd/
2. Build the publisher and subscriber applications on two different nodes
3. Run the subscriber first - ./subscriber -s nats://server.ip.address:4222 mytopic
4. Run the publisher - ./publisher -s nats://server.ip.address:4222 mytopic "Some Message"

Question:

What happens when you publish while there is no subscriber running?

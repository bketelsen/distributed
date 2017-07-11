1. Start a nats server by running the "nats-server1.sh" script in this directory
2. Build the publisher and subscriber applications on two different nodes
3. Run the subscriber first - ./subscriber -s nats://server.ip.address:4222 mytopic
4. Run the publisher - ./publisher -s nats://server.ip.address:4222 mytopic "Some Message"

Question:

What happens when you publish while there is no subscriber running?

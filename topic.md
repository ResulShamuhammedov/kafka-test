# run this command to create new topic:
```docker exec -it <kafka-container-id> kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test-topic```
## REST Microservices using Golang | OAuth API
OAuth API | DDD Approach | Apache Cassandra As Backend Database

## Author
Aditya Hajare ([Linkedin](https://in.linkedin.com/in/aditya-hajare)).

## Current Status
WIP (Work In Progress)!

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).

-------

### Install And Configure Cassandra:
- [How to setup Cassandra on MacOS](https://medium.com/@manishyadavv/how-to-install-cassandra-on-mac-os-d9338fcfcba4)
- Steps:
    ```sh
    # Install Python
    brew install python

    # Install Cassandra
    brew install cassandra

    # Install cqlsh
    pip install cql

    # Start Cassandra
    cassandra -f

    # Start cqlsh
    cqlsh
    ```
- In `cqlsh` shell, type following to create `keyspace`
    ```sh
    # List all keyspaces
    describe keyspaces;

    # Create new keyspace called "oauth" with single replica
    CREATE KEYSPACE oauth WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};
    ```
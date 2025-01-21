# golang-cassandra-example

- Running Cassandra
```bash
docker run --name cassandra-container -d -p 9042:9042 cassandra:latest
```

- Enter to the container
```sql
docker exec -it cassandra-container cqlsh
```

- Creating a keySapce
```sql
CREATE KEYSPACE userskeysapce 
WITH replication = { 
  'class': 'SimpleStrategy',
  'replication_factor': 1
};
```

- Creating a table
```sql
CREATE TABLE userskeysapce.users (
  id UUID PRIMARY KEY,
  name text,
  email text
);
```

- Cassandra commands
```sql
DESCRIBE KEYSPACES;               -- Get KEYSPACE list
DROP KEYSPACE userskeysapce;      -- Delete KEYSPACE
DESCRIBE KEYSPACE userskeysapce;  -- Getting information about a single KEYSPACE
```

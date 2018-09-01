# PersiaDB - Multi Model Database
If you check any databases you can feel that all of them just add a layer to exiting infrastructure and non of them can be good for all type of data model! So We decide to take logic part of database from Storage Engine and be in library to add to app logic layer instead an additional app.
In storage engines even newest one like [LevelDB](https://github.com/google/leveldb) and [RocksDB](https://github.com/facebook/rocksdb) have some fundemental problem e.g. they build in top of exiting FileSystems as defualts!!
Some NewSQLs like [Cockroach](https://github.com/cockroachdb/cockroach), [arangodb](https://github.com/arangodb/arangodb), [elassandra](https://github.com/strapdata/elassandra) have not same but god in some use cases.
It is just database library not standalone stateless or statefull application! this means it can scale with your app scalability architecture!

## Production Ready!?
This package is under development and not ready to use in real production. It can have breakable changes until version 1 release.
But we are glad to hear your experience or idea about this package.
We use MySQL as storage engine for start. when PersiaOS released we switch to it.

## Use Cases
- Multi Model Datastore
- Data Mining
- Machine Learning & AI (Artificial intelligence)

## Architecture
We need to store and [retrieve](https://en.wikipedia.org/wiki/Information_retrieval) data in multi model!
We have all futures like ACID, transaction in library layer. We don't support & suggest to use normalize data becuase of performance lack in data indexing.
We don't need row lock for transaction query becuase we don't use update data in SE layer.

### Rules
There is no update data! Just have write, get and delete in storage engine layer. UUID belong to the object can be used as version control of that record, So each object in each version have unique UUID. With this rule in every layer we can cache an object by its UUID and it is guaranty that always be same data on each UUID!

### Object data structure
- MetaData
- Data

### Replication Strategy
We choose replication strategy to indicate in storage engine layer. Developer can set needed replication storage and get related token that indicate where and how data must store!
Replication data can be any of this startegy to cover CAP theory:
- no replication
- replicate once on the same rack
- replicate once on a different rack, but same data center
- replicate once on a different data center
- replicate twice on two different data center
- replicate once on a different rack, and once on a different data center

### Logics
- Compress & uncompress data by MIME ???

### APIs
PersiaDB has very simple usage instructions! Just look at data-{}.go methods with their request and respond structures.

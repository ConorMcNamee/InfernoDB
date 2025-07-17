# How to build a database

The first step of building a database is going to be building a key value store system. We are going to want to store this KV store in memory. Than want to keep a append-only log of the system so that on crash we can rebuild the entire KV store. Provided we follow the ACID priniciples

 - A -> Atomicity
 - C -> Consistency
 - I -> Isolation
 - D -> Durability

We should be fine and we have a basic database (kinda)

Update: I now have a KV store which I can add data too, now I just need a log and then a recover system and I have a storage engine
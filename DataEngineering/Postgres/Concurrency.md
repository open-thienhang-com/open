# PostgreSQL Concurrency: Isolation and Locking

## ACID - 4 tính chất tạo nên thương hiệu của relational database

[![ACID][IMG-ACID]](https://www.geeksforgeeks.org/acid-properties-in-dbms/?ref=leftbar-rightbar)

---

1. Atomic:

2. Consistency:

3. Isolation:

4. Durability:


Dirty read

A transaction reads data written by a concurrent uncommitted transaction.

Nonrepeatable read

A transaction re-reads data it has previously read and finds that data has been modified by another transaction (that committed since the initial read).

Phantom read

A transaction re-executes a query returning a set of rows that satisfy a search condition and finds that the set of rows satisfying the condition has changed due to another recently committed transaction.

Serialization anomaly

The result of successfully committing a group of transactions is inconsistent with all possible orderings of running those transactions one at a time
---

### Những điểm cần lưu ý

| Thuộc tính  | Trách nhiệm của lớp xử lí |
| ------------- | ------------- |
| Atomicity  | Transaction Manager  |
| Consistency  | Application programmer  |
| Isolation  | Concurrency Control Manager  |
| Durability  | Recovery Manager  |

## Transaction isolation

Standard SQL Transaction isolation levels:

1. Read uncommitted.

2. Read committed.

3. Repeatable read.

4. Serializable.

> Nguồn tham khảo:

> <https://www.postgresql.org/docs/current/mvcc.html>

> <https://viblo.asia/p/014-postgresql-transaction-isolation-OeVKB67JKkW>

<!-- MARKDOWN LINKS & IMAGES -->
[IMG-ACID]: images/ACID-Properties.jpg

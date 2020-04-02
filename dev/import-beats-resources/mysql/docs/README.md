# MySQL Integration

This integration periodically fetches logs and metrics from [https://www.mysql.com/](MySQL) servers.

## Compatibility

The `error` and `slowlog` datasets were tested with logs from MySQL 5.5, 5.7 and 8.0, MariaDB 10.1, 10.2 and 10.3, and Percona 5.7 and 8.0.

The `galera_status` and `status` datasets were tested with MySQL and Percona 5.7 and 8.0 and are expected to work with all
versions >= 5.7.0. It is also tested with MariaDB 10.2, 10.3 and 10.4.

## Logs

### error

The `error` dataset collects the MySQL error logs.

{{fields "error"}}

### slowlog

The `slowlog` dataset collects the MySQL slow logs.

{{fields "slowlog"}}

## Metrics

### galera_status

The `galera_status` dataset periodically fetches metrics from [http://galeracluster.com/](Galera)-MySQL cluster servers.

An example event for `galera_status` looks as following:

```$json
TODO
```

The fields reported are:

{{fields "galera_status"}}

### status

The MySQL `status` dataset collects data from MySQL by running a `SHOW GLOBAL STATUS;` SQL query. This query returns a large number of metrics.

An example event for `status` looks as following:

```$json
{
   "@timestamp":"2020-04-02T09:49:06.809Z",
   "metricset":{
      "name":"status",
      "period":10000
   },
   "service":{
      "address":"127.0.0.1:3306",
      "type":"mysql"
   },
   "mysql":{
      "status":{
         "bytes":{
            "received":37122,
            "sent":1629640
         },
         "questions":464,
         "max_used_connections":3,
         "binlog":{
            "cache":{
               "disk_use":0,
               "use":0
            }
         },
         "open":{
            "tables":115,
            "files":14,
            "streams":0
         },
         "innodb":{
            "buffer_pool":{
               "bytes":{
                  "data":4898816,
                  "dirty":0
               },
               "pages":{
                  "misc":0,
                  "total":8191,
                  "data":299,
                  "dirty":0,
                  "flushed":36,
                  "free":7892
               },
               "read":{
                  "ahead":0,
                  "ahead_evicted":0,
                  "ahead_rnd":0,
                  "requests":1427
               },
               "pool":{
                  "reads":266,
                  "wait_free":0
               },
               "write_requests":325
            }
         },
         "aborted":{
            "connects":0,
            "clients":0
         },
         "queries":464,
         "opened_tables":122,
         "flush_commands":1,
         "threads":{
            "created":3,
            "connected":2,
            "running":2,
            "cached":1
         },
         "delayed":{
            "insert_threads":0,
            "writes":0,
            "errors":0
         },
         "connections":154,
         "handler":{
            "external_lock":542,
            "update":0,
            "mrr_init":0,
            "commit":5,
            "read":{
               "rnd_next":57819,
               "first":8,
               "key":6,
               "last":0,
               "next":1,
               "prev":0,
               "rnd":0
            },
            "savepoint":0,
            "write":0,
            "delete":0,
            "prepare":0,
            "rollback":0,
            "savepoint_rollback":0
         },
         "created":{
            "tmp":{
               "tables":0,
               "disk_tables":0,
               "files":6
            }
         },
         "command":{
            "update":0,
            "delete":0,
            "insert":0,
            "select":150
         }
      }
   },
   "ecs":{
      "version":"1.5.0"
   },
   "host":{
      "architecture":"x86_64",
      "os":{
         "build":"18G4032",
         "platform":"darwin",
         "version":"10.14.6",
         "family":"darwin",
         "name":"Mac OS X",
         "kernel":"18.7.0"
      },
      "id":"24F065F8-4274-521D-8DD5-5D27557E15B4",
      "ip":[
         "fe80::aede:48ff:fe00:1122",
         "fe80::1cf7:b917:ae2:f9ce",
         "192.168.0.13",
         "fe80::484c:baff:fe76:8f66",
         "fe80::b92d:5718:8a3f:297b",
         "fe80::8542:d769:a147:539f"
      ],
      "name":"elastic.local",
      "mac":[
         "ac:de:48:00:11:22",
         "a6:83:e7:ae:70:01",
         "a4:83:e7:ae:70:01",
         "06:83:e7:ae:70:01",
         "4a:4c:ba:76:8f:66",
         "a2:00:45:49:64:01",
         "a2:00:45:49:64:00",
         "a2:00:45:49:64:05",
         "a2:00:45:49:64:04",
         "a2:00:45:49:64:01"
      ],
      "hostname":"elastic.local"
   },
   "agent":{
      "id":"11839a5a-1fe0-498f-9994-ce3698038453",
      "version":"8.0.0",
      "type":"metricbeat",
      "ephemeral_id":"99850a99-7b9d-441f-88c3-7e0eacd83716",
      "hostname":"elastic.local"
   },
   "event":{
      "dataset":"mysql.status",
      "module":"mysql",
      "duration":2487295
   }
}
```

The fields reported are:

{{fields "status"}}

# etcdir
etcdir is Tool to store in etcd while maintaining directory structure.

## Example And Usage
For example local directory tree
```
etc
├── sample1
│   ├── file1
│   └── file2
└── sample2
    ├── file3
    └── file4
```

etcdir Command usage
```
./etcdir -d etc -n {node1, node2...}
```

```
/sample1/file1
/sample1/file2
/sample2/file3
/sample2/file4
```
Read file contents and store values ​​in etcd using directory structure as key


### Development
```
docker-compose up -d
```
localhost:2379

#### Test
TODO


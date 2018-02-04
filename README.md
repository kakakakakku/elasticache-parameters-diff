# elasticache-parameters-diff

`elasticache-parameters-diff` is CLI to get diff of Amazon ElastiCache Parameter Group.

## Getting started

```sh
$ go get -u github.com/kakakakakku/elasticache-parameters-diff
```

## Examples

Example1:

```sh
$ elasticache-parameters-diff --p1 default.redis3.2 --p2 default.redis3.2.cluster.on
===== Cache Parameter Group Name : default.redis3.2 =====

{
  AllowedValues: "yes,no",
  ChangeType: "requires-reboot",
  DataType: "string",
  Description: "Enable cluster mode",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "cluster-enabled",
  ParameterValue: "no",
  Source: "system"
}

===== Cache Parameter Group Name : default.redis3.2.cluster.on =====

{
  AllowedValues: "yes,no",
  ChangeType: "requires-reboot",
  DataType: "string",
  Description: "Enable cluster mode",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "cluster-enabled",
  ParameterValue: "yes",
  Source: "system"
}
```

Example2:

```sh
$ elasticache-parameters-diff --p1 default.redis3.2 --p2 sample
===== Cache Parameter Group Name : default.redis3.2 =====

{
  AllowedValues: "volatile-lru,allkeys-lru,volatile-random,allkeys-random,volatile-ttl,noeviction",
  ChangeType: "immediate",
  DataType: "string",
  Description: "Max memory policy.",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "maxmemory-policy",
  ParameterValue: "volatile-lru",
  Source: "system"
}
{
  AllowedValues: "0,20-",
  ChangeType: "immediate",
  DataType: "integer",
  Description: "Close connection if client is idle for a given number of seconds, or never if 0.",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "timeout",
  ParameterValue: "0",
  Source: "system"
}

===== Cache Parameter Group Name : sample =====

{
  AllowedValues: "volatile-lru,allkeys-lru,volatile-random,allkeys-random,volatile-ttl,noeviction",
  ChangeType: "immediate",
  DataType: "string",
  Description: "Max memory policy.",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "maxmemory-policy",
  ParameterValue: "allkeys-lru",
  Source: "user"
}
{
  AllowedValues: "0,20-",
  ChangeType: "immediate",
  DataType: "integer",
  Description: "Close connection if client is idle for a given number of seconds, or never if 0.",
  IsModifiable: true,
  MinimumEngineVersion: "3.2.4",
  ParameterName: "timeout",
  ParameterValue: "30",
  Source: "user"
}
```
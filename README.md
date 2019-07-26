# Multicast Hex Dump
Dump multicast data to stdout

Simple CLI tool to dump UDP Multicast traffic to console in hex format. This is useful for systems which do not have TCPDump and can deploy a single go binary executable.
```
go-multicast-dump -a 224.0.23.204:21000
```

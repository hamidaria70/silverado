# EAGLE

Eagle will be ran in each server that you can have ssh access and create a
result of system overview.

### To Build

To make a binary file of this project run the command below:

```bash
go build --ldflags '-linkmode external -extldflags "-static"' .
```

| Hostname | IP Address    | Up Time   | CPU Usage Percentage | Disk Usage Percentage | Memory Usage Percentage | Load Average 1 | Load Average 5 | Load average 15 |
| -------- | ------------- | --------- | -------------------- | --------------------- | ----------------------- | -------------- | -------------- | --------------- |
| office   | 192.168.1.116 | 119h5m43s | 5.04                 | 66.62                 | 15.5G                   | 0.85           | 0.69           | 0.81            |

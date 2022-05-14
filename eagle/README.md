# EAGLE


###Introduction

`Eagle` will be ran in each server that you can have ssh access and create a
result of system overview by `ansible`.

### To Build

To make a binary file of this project run the command below:

```bash
go build --ldflags '-linkmode external -extldflags "-static"' .
```

### To Use

1. Make sure that `ansible` is already installed on your host. If not , Run the
   commands below:

```bash
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

2. Make sure that you have `ssh access` to the servers **without password** 
   that you want to run `eagle`.

```bash
ssh-copy-id <USER>@<SERVER_IP>
```

3. You need to generate a host file for ansible to read hosts from it.So run:

```bash
chmod +x inventory/host-generator.sh
```

then

```bash
./inventory/host-generator.sh
```

follow the structure and generate the host file in inventory directory.

4. Last step to use eagle is to run:

```bash
ansible-playbook -i inventory/hosts playbooks/eagle.yml
```

5. And Finaly you can check the markdown view result in `result.md`

done.

### Example Result view

| Hostname | IP Address    | Up Time    | CPU Usage Percentage  | Disk Usage Percentage | Memory Usage Percentage | Load Average 1 | Load Average 5 | Load average 15 |
| -------- | ------------- | ---------- | --------------------- | --------------------- | ----------------------- | -------------- | -------------- | --------------- |
| office   | 192.168.1.116 | 119h45m38s | 5.34 % out of 8 cores | 66.62 % out of 109GGB | 30 % out of 15.5GGB     | 0.41           | 0.65           | 0.78            |


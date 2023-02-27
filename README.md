# DNSKiller
![alt-text](https://github.com/n0nexist/DNSKiller/blob/main/screenshot.png?raw=true)<br><br>
```DNSKiller``` is a <b>DNS</b> <ins>top-level</ins> domain and <ins>subdomain</ins> <b>bruteforcer</b> written in ```golang```.<br><br>

# Explaination
<i>A DNS top-level domain (TLD) and subdomain brute-force attack is a type of cyber attack that involves attempting to discover valid domain names by systematically generating and testing a large number of potential domain names.</i><br><br>

# How to download [build from source]
```
git clone https://github.com/n0nexist/DNSKiller
cd DNSKiller
go build DNSKiller.go
```
<br>

# How to download [linux pre-compiled binary]
<b>You can find the dowload link</b> <a href="https://github.com/n0nexist/DNSKiller/releases/latest">here</a>.
<br><br>

# Usage [examples]
<b>Basic usage</b>
```
DNSKiller.go google wordlists/sub.txt wordlists/top.txt 500
```
<b>Write to a log file</b>
```
DNSKiller.go google wordlists/sub.txt wordlists/top.txt 500 output.log
```

# grpc-vuln-service
## Service wrapper above nmap



## Installation
Install the dependencies and and start the server.
```sh
sudo apt-get update
sudo apt-get install nmap
git clone https://github.com/vulnersCom/nmap-vulners
cd nmap-vulners

mv vulners.nse /usr/share/namp/scripts/
mv http-vulners-regex.nse /usr/share/namp/scripts/
mv http-vulners-regex.json /usr/share/namp/nselib/data/
mv http-vulners-paths.txt  /usr/share/namp/nselib/data/

```
## Starting server
```sh
make start
```

## TODO
- Add cache for MRU caching most recently used targets
- Add timeout for scan request in .proto
- Add fix ci errors (static check) for deprecated function in nmap package 

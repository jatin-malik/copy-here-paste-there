# copy-here-paste-there
Stream your system clipboard over the network transparently in the background.

`copy-here-paste-there` ships as a single go binary. Build from source or just go install.

```
go install github.com/jatin-malik/copy-here-paste-there/cmd/copy-here-paste-there@latest
```

## Usage

Run `copy-here-paste-there --help` for the usage instructions. The CLI arguments are pretty self-explanatory.

The binary runs in one of the two modes, `server` and `client,` configurable via command-line argument `-mode`.


Run it as a server on node A
```
copy-here-paste-there -mode server -localport 8080 -log debug
```

Then, run it as a client on node B
```
copy-here-paste-there -mode client -host localhost -port 8080 -log debug
```
And that's it. Copy on node A and paste on node B. Or, copy on node B and paste on node A. It goes both ways.

### Over LAN

If you are communicating over the LAN, you will need your server's local IP. For MacOS, you can use `ifconfig` command.
Look for en0 or en1 ( commonly! ), then find the line that starts with inet. The IP address next to inet is your server's LAN IP.

```diff

en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
	options=6460<TSO4,TSO6,CHANNEL_IO,PARTIAL_CSUM,ZEROINVERT_CSUM>
	ether f8:4d:89:95:f0:8d
	inet6 fe80::1cac:baac:efb1:5c6%en0 prefixlen 64 secured scopeid 0x10
	inet6 2401:4900:8813:1bb5:1016:7162:bc82:13fb prefixlen 64 autoconf secured
	inet6 2401:4900:8813:1bb5:cc96:37ef:a420:7187 prefixlen 64 autoconf temporary
+	inet 192.168.1.7 netmask 0xffffff00 broadcast 192.168.1.255
	nd6 options=201<PERFORMNUD,DAD>
```  

### Over internet

If you are communicating over internet and do not have a public IP address, you can use something like [ngrok](https://download.ngrok.com/mac-os) to put your localhost online.

```
ngrok tcp 8080
```
This will give you a public endpoint that tunnels all the TCP traffic to your localhost on port 8080. Something like this
```
tcp://0.tcp.in.ngrok.io:10549 -> localhost:8080
```

You can then connect to this server from client as usual using

```
copy-here-paste-there -mode client -host "0.tcp.in.ngrok.io" -port 10549 -log debug
```

## Internals

Communication is done over TCP. Once the TCP connection is established between the client and server, it allows for bi-directional streaming of your system clipboard/pasteboard.
A watcher periodically ( every 2 seconds ) checks the system clipboard for any state change and if it observes a state update, it streams that over to the peer using a length-prefixed TCP message.







# Use RSocket to stream file without draining your infrastructure

## Where it started

File tranfer is the typical programming language exercise that almost every programmer has had. Yet, even today, it's not easy to share files, especially large files.

What's the limitation on email attachment? Maybe you can set up a FTP server. Or you can use dropbox. But those solutions require third party storage. 

When it comes to point to point file sharing, there are protocols like SCP and SFTP. However, they have certain limitations, mostly related to network. They are both over SSH, so when the network fluctuates, they might lose connection. SCP is way faster than SFTP. But that comes with cost. SCp will maximize the bandwith, and when there is a weak link in the network, the file transfer may be interrupted. 

We can keep arguing this for a long time. But what is a easy, reliable mechanism to transfer 100G file point to point?


## The demo

In this demo, we are gonna showcase a simple program that transmit file.
The code base is simple enough that doesn't require much explanation.
Under the `/server` folder is the server code. You may modify the port and then run
`go build`. It will give you a binary. Simply run `./server` and it's started!

Under the `/client` folder is the sender code. You again can modify the ip and port of the server before running `go build`.
And to run this client, use the command `./client f *file-name* `.

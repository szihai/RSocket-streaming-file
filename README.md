# Use RSocket to stream file without draining your infrastructure

## Where it started

File tranfer is the typical programming language exercise that almost every programmer has had. Yet, even today, it's not easy to share files, especially large files.

What's the limitation on email attachment? Maybe you can set up a FTP server. Or you can use dropbox. But those solutions require third party storage. 

When it comes to point to point file sharing, there are protocols like SCP and SFTP. However, they have certain limitations, mostly related to network. They are both over SSH, so when the network fluctuates, they might lose connection. SCP is way faster than SFTP. But that comes with cost. SCp will maximize the bandwith, and when there is a weak link in the network, the file transfer may be interrupted. 

We can keep arguing this for a long time. But what is a easy, reliable mechanism to transfer 100G file point to point?


## The demo

In this demo, we are gonna showcase a simple program that transmit file.
The code base is simple enough that doesn't require much explanation. 
After you clone the repo, please make sure you do `go get` of the missing packages.
Under the `/server` folder is the server code. You may modify the port and then run
`go build`. It will give you a binary. Simply run `./server` and it's started!

Under the `/client` folder is the sender code. You again can modify the ip and port of the server before running `go build`.
And to run this client, use the command `./client f *file-name* `.

## RSocket
[RSocket](http://rsocket.io) is an application protocol providing [Reactive Streams](https://www.reactivemanifesto.org/) semantics. As the reactive menifesto points out, resiliency is a key feature. When the network fluctuates, the protocol will try to stay put and keep the file transmission going. So such a simple program can actually withstand production use. 

## Fire and forget
RSocket defines four types of communication patterns. They are:
* Request response: which is the tranditional rpc call
* Request stream: like other streams, serial of calls
* Bi-directional steam: 2 way streams
* Fire and forget: It's rpc without response. This is what we use in this demo. For use cases like streaming file, sending logs or metrics, this is the perfect solution. Because users normally don't need or care about the ACK. We all know reactive streams is asynchronous, non-blocking. But the syntax can sometimes be tedious. Fire and forget provides the simple straightforward semantics to best catch the intention of such behavior.

## Improvement
As a user facing utility, this is a demo. It has many areas for improvement. Feel free to modify it as you wish. However, the RSocket protocol has made the mechanism steady and robust. 
* For large files   
The file size limit is caused by Golang's default `ReadFile` and `WriteFile`. If you need to process larger files, please find another utility to use. *But* that's not caused by RSocket.


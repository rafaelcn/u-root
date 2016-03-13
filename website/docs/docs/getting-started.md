#  Getting started with u-root:
---

1. Prerequisites
=============

To execute u-root and play with it, you need to have git, golang installed. 
On a Debian, Ubuntu or other .deb system, you should be able to get going with:

	sudo aptitude install git golang build-essential

U-root requires at least go1.5 and the go source tree needs to be present.


2. Github
======

We use Github! And we are located at:

	https://github.com/u-root/u-root

Clone the repository of u-root and fell free to help us, just be aware of what we are doing and what are our next goals.

To run u-root you will just need to either build* the `ramfs.go` or run it with this line below.
 
	go run scripts/ramfs.go -test



\* to build the `ramfs.go` use:

	go build scripts/ramfs.go




	

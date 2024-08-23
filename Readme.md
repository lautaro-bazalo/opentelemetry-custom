# Custom opentelemetry-collector POC

This is a point of code where I'm building a custom collector of opentelemetry with a custom receiver that simulate receive a accesslog and generate metrics. 


## aclreceiver module

First we had to initiliaze the module with a mock path to enable the import reference in the custom-collector.

In this moudle you will find three files
* `config.go`
* `factory.go`
* `aclreceiver.go`

Here the most important is file is the aclreceiver which implement the `start` and `shutdown` methods that handle the receiver.


## custom-collector 


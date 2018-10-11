Package to publish jobs to queues that are being used by any kind of worker. Available implementations at the moment:

- resque (Publish jobs to an existing resque system https://github.com/resque/resque) 
 


### **Installation**

``` got get github.com/ericbrisrubio/queuepublisher```

### **Example**

```
package main

import (
    publisher "github.com/ericbrisrubio/queuepublisher"
    _ "github.com/kavu/go-resque/go-redis" // Redis client from godis package
)


func main() {
    enqueuer := publisher.FactoryEnqueuer("resque", "<dbhost>", "<dbport>", "<dbname>", "<dbpass>", <use ssl connection>)
    enqueuer.PublishToQueue("<queueName>", "<className>", <data>)

}
```
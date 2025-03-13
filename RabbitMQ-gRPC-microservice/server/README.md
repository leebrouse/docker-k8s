# 对于rabbitMQ 和 gRPC 的简单微服务练习
## Basic funcation :
- 1. Pay the bill
- 2. Add to the Cart
- 3. Create the order 
- 4. Automatically start application using the makefile
## Each Dir funcation :
### Proto:
-  writing three .proto files(1. order.proto,2. payment.proto,3. shopping.proto) and use protoc commend to build .go file for using them in microservice
### Order :
- use Order which was created by the order.proto and implement the orderService interface,then connect with rabbitMQ 
### Payment :
- use Payment which was created by the payment.proto and implement the paymentService interface,then connect with rabbitMQ 
### Shopping :
- use Shopping which was created by the shopping.proto and implement the shoppingService interface,then connect with rabbitMQ 

    
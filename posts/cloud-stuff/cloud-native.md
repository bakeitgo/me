# Cloud native

* What is cloud native?

    * *Cloud-native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach. These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.*

* Pets vs Cattle - DevOps concept.

    * In traditional data centers, systems are treated as *Pets*. We scale it by adding more resources to the same machine (vertical scale) (CPU's RAM HardDisk etc.). If something is bad with service, we need to halt all services which belongs to it.

    * In *Cattles*, we scale up by adding more services (horizontal scale). Each service is a container or virtual machine. When something is unavailable, for e.g. container goes down, customer doesn't even notice it, because the new container goes up immediately (automated). In this method we don't care about repairing servers, if it fail or needs update, we delete them and create another (automated)

* How to construct cloud-based applications?

    * It's all described in [Twelve-Factor Application](https://12factor.net/)



* Microservices 

    * A small service (process) which makes one task in deterministic way. It's created with own dependencies, data storage, programming platform. 

    * It's communicating with other microservices via HTTP/HTTPS, gRPC, WebSockets or AMQP.

    * Each microservice deal with other microservices, all together interact to form an application.

    * Each is independent from each other, but they collaborate with selves.

* Why microservices?

    * They provide agility. We don't have to wait couple of months to deploy app, with microservices is less risky than monolithic application infrastructure. Why? Because we update small part of our application with microservices. They are independent. They scale well in this case. 

* Microservices challenges.

    * Communication - How front-end communicate with back-end? Will you allow direct communication or consider API gateway?

    * Resiliency - How to deal when Service A doesn't respond to Service B request? What if Service C is unavailable and other services calling it become blocked?

    * Distributed data - By design microservice encapsulate its own data and expose it via interface, how to query data or implement transaction across multiple services?

    * Secrets - How microservices will store and manage sensitive data which cant be exposed?


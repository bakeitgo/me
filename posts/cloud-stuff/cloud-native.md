# Cloud native

* What is cloud native?

    * *Cloud-native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach. These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.*

* Pets vs Cattle - DevOps concept.

    * In traditional data centers, systems are treated as *Pets*. We scale it by adding more resources to the same machine (vertical scale) (CPU's RAM HardDisk etc.). If something is bad with service, we need to halt all services which belongs to it.

    * In *Cattles*, we scale up by adding more services (horizontal scale). Each service is a container or virtual machine. When something is unavailable, for e.g. container goes down, customer doesn't even notice it, because the new container goes up immediately (automated). In this method we don't care about repairing servers, if it fail or needs update, we delete them and create another (automated)

* How to construct cloud-based applications?

    * It's all described in [Twelve-Factor Application](https://12factor.net/)


***
## Microservices 

* What is a *microservice*?

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
***
## Containers

* What is a container?

    * Code, Dependencies and Runtime is packaged into binary. It's a *container image*

    * Images are stored in container registry it acts like a repository / library for images.

    * Registry can be located in your local computer, data center or public cloud. *Docker* itself maintains a public registry via *Docker hub* 

    * When application starts, container is transformed from image to running container instance. Instance runs on any computer which have *container runtime engine* installed

    * Each container is independent and isolated from each other.

    * it embraces **Dependencies** principle from [Twelve-Factor Application](https://12factor.net/)
    ##
    *Factor #2 specifies that “Each microservice isolates and packages its own dependencies, embracing changes without impacting the entire system.”*


* Why containers?

    * They can be easily managed, they are isolated from infrastructure (they act as a single package), guarantee consistency across environments, we can deploy them in mostly in all environment that hosts *Docker runtime engine* (depends on CPU architecture), they are cheap because we don't have to pre-configure each environment with dependencies and runtime engines everytime. They have smaller size than VM's (it increases density or number of microservices which can run at one time)
***
## Differences between VMs and Containers

* VMs is an Hardware virtualisation, it have an hypervisor above itself (which can run multiple instances of hardware and provides ram rom cpu etc.)

    * We call it a full system virtualisation ( we work on hardware level )

    * Isolation - Imagine we have a server, and we want to split its resources, and we are creating different machines which splits it (our server looks like a bunch of servers).  They are relatively independent (communicating with e.g. vm 1 seems like communicating another server)

    * Interaction - we interact with hardwares which is managed by hypervisor

    * Flexibility - we have flexibility around managing how RAM, ROM, CPU threads our VM will use 

* Containers is an hardware virtualisation, on top have *kernel* (helps in communiucation between each other), on top we got *OS* (host OS), on top its every runned containers (it can be more than 1)

    * We call it an operating system level virtualisation (because all bases on OS)

    * Process isolation - our application see only things it needs (libraries, binaries etc.) (every container thinks have different host)

    * Interaction - two things here - namespaces (by it container thinks is operating on different OS), cgroups (they monitor and measure our resources to make sure we will never overload our OS)

    * Portability - Our container is defined in *single file* so its portable, we store it in repo and it is able to run it mostly everywhere (have to be compatible with cpu architecture ((ARM, x86))


### *Note: We can mix both approaches*

*** 
## What are containers and why do we care?

* First let explain concept from where it come from - *VMs*

    * Problem is we got individual machines / VMs and it is hard to maintain

    * VMs are ***EXTREMELY ERROR PRONE***

    * Servers as Pets
    ##
    1. Long lived
    1. 99.999% uptime
    1. If one went down, troubleshoot and restore

* Solution - *Containers*

    * What if we could bundle all our app code, support binaries, and configuration together and only do that once?

    * It's called an *image*, we put them on servers, if some breaks , we drop it down and setup new

    * Each one is copy of each other, so problem with hard managing is dropped

* 3 Phases of *Docker* 

    1. *Build Image* (package everything app needs to run)

    1. *Ship Image* (ship to runtimes in cloud or local machine)

    1. *Run image* (execute your application)

* What else it provides?

    * CI/CD (test and deploy consistently to different envs. (stage, uat, production))

    * Different versions of softwares

    * Roll forward (something is broke? -> ship new image)

* How to manage containers? 

    * To manage it we need *container orchestrator*

* What orchestration ensures?

    * Scheduling - provision container instances automatically

    * Affinity / anti-affinity - provision containers nearby or far apart from each other (helps availability and performance)

    * Health monitoring - detects failures

    * Failover - If container fails, setup new

    * Scaling - If new container is needed, provide it.

    * Networking - Containers need to communciate, right?

    * Service discovery - Containers need to know where others are located.

    * Rolling upgrades - Coordinate upgrades, with zero downtime. Rollback easily if needed.

* What points of [Twelve-Factor Application](https://12factor.net/) orchestration ensures?

    * Factor #9 *specifies that “Service instances should be disposable, favoring fast startups to increase scalability opportunities and graceful shutdowns to leave the system in a correct state.” Docker containers along with an orchestrator inherently satisfy this requirement."*

    * Factor #8 *specifies that “Services scale out across a large number of small identical processes (copies) as opposed to scaling-up a single large instance on the most powerful machine available.*

***
## Building / running docker image

* Commands

    * `docker build --tag hello-world .` build an image with tag name hello-world

    * `docker images` - show list of images

    * `docker ps` - check running containers list

    * `docker run <imgname>` - run image

    * `docker run -p 8080:80` - map 8080 port on host to 80 port on container

    * `docker run -d` - run container detached (headless)

    * `docker stop <name>` - stop container

    * `docker start <name>` - start container

    * `docker logs <name>` - show logs of container

    * `docker logs <name> -f` - show logs of container continously

***
## Repository for images

* *Docker Hub* is an images repo

    * We can do here auto builds (e.g. you push changes to github repo, image is build with new stuff)

* *Docker compose*

    * Run multi-container Docker apps. It uses YAML file to configure app services.

* Commands

    * `docker push <reponame>:<imgname>` - push img to repo

    * `docker pull <reponame>` - pull img from repo

    * `docker-compose up` - starts all containers

    * `docker-compose down` - shutdown containers
***

# Related to

* https://docs.docker.com/get-started/

* https://youtu.be/iqqDU2crIEQ?t=290

* https://docs.docker.com/compose/

  
#containers #docker #kubernetes #k8s #vm #virtualisation #cloudnative #cloud-native #pets #cattle #microservices #dockerhub #imagerepo #dockercompose #docker-compose

# Containers

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

    * Portability - Our container is defined in *single file* so its portable, we store it in repo and it is able to run it mostly everywhere (have to be compatible with cpu type ((ARM, x86))


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


# Related to

* https://docs.docker.com/get-started/
* https://youtu.be/iqqDU2crIEQ?t=290

#containers #docker #kubernetes #k8s #vm #virtualisation
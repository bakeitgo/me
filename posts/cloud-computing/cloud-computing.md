# Cloud computing / On-premise


* What is an ***on-premise*** environment?

	* On-premise is an local environment. It is an infrastructure where we need to manage, scale, and maintain by ourselves. In this type of environment there is huge problem with scalability, because for e.g. we cannot easily  add more servers when we need them, because it cost a lot of time. **Scalability** of *On-premise* env leads to huge cost, and time consuming practice. **Security** we need to maintain it by ourselves also. We need to obtain that our servers are secure in the IT sector, black hat attackers cannot easily access it, and secure physical access. **Backup** - we need to provide it by ourselves.

* What is *cloud computing* environment?

	* Cloud computing environment have huge scalability, we can easily add, delete servers when we need / don't need them, and we don't have to care about maintaining, manage and **STORE** our environment locally. **Security** and **backup** in this env. is handled by our provider.


* What about *cloud computing* environment types?

	* **Private cloud** - is like local environment (On-premise) but to name it like **Private cloud** it should be automated, have easy implementation of new instances, and provide easy maintaining of apps.

	* **Public cloud** - is an environment which is fully provided by cloud provider.

	* **Hybrid cloud** - Some shit we gotcha on a *on-premise* environment, like database center, other we got on *public cloud*. 

	* **Multi cloud** - It's like hybrid cloud but only using *public clouds*.


* What about infrastructure models?
	#
	* &#9748; ***WORD*** - **You maintain**
	* WORD - **Provider maintain**
	#
	| On-premise | IaaS | PaaS | SaaS |
	|------------|------|------|------|
	| &#9748; ***Application*** | &#9748; ***Application*** | &#9748; ***Application*** | Application |
	| &#9748; ***Data*** | &#9748; ***Data*** | &#9748; ***Data*** | Data |
	| &#9748; ***Databases*** | &#9748; ***Databases*** | Databases | Databases |
	| &#9748; ***OS*** | &#9748; ***OS*** | OS | OS |
	| &#9748; ***Virtualization*** | Virtualization | Virtualization | Virtualization |
	| &#9748; ***Servers*** | Servers | Servers | Servers |
	| &#9748; ***Storage*** | Storage | Storage | Storage |
	| &#9748; ***Networking*** | Networking | Networking | Networking |
	| &#9748; ***Data Center*** | Data Center | Data Center | Data Center |

	#
* Which to choose?	
	#
	* ***IaaS*** - If enterprise wants a virtualization, servers, storage managed by cloud, rest by SysOpses. 
	
	#### **Benefits** of using ***IaaS***:
	#
	##
	1. **Higher availability** - resistant due to power loss and other random circumstances

	1.  **Lower latency (Better performance)** - we can locate servers in multiple geographic points which are closer to user.

	1. **Comprehensive Security** - Via encryption and high level of security on-site, enterprises are more secure than on an On-premise infr. (*it depends, still enterprise can have better security advancements etc. but still they don't have to do it by **ourselve** and it's hard (probably even not possible) to do*)

	1. **Faster access to best-of-breed technology** - Providers compete with each other, so they serve latest technologies faster, why? Answer is simple, to gain more customers. 

	1. **Provides disaster recovery**. What? It prevents or minimize data loss due to unforeseen circumstances, it also includes *failover* and *fallback* procedures. **Failover** is a process to offload workloads to backup systems, so prod processes and end-user experience are disrupted as less as possible. **Fallback** is a data fallback to the original primary systems How? Via geographically-dispersed infrastructure.
	#
	* ***PaaS*** - Provides cloud platforms and runtime env to develop, run and manage applications without maintaining architecture, PaaS provides platform to do it. 
	#### **Benefits** of using Paas:
	#
	###

	1. **Faster time to market** - Ability to spin-up development, testing, prod env. in minutes, vs weeks / months.

	1. **Low-to no-risk testing and adoption of new technologies** - PaaS serve a wide range of resources, which means that we can test e.g. new OS, tools, languages, packages without investing in them.

	1. **Simplified collaboration** - Provides shared software development environment, giving development and operations teams access to all tools they need, from anywhere. Only internet connection is needed.

	1. **More scalable approach** - Ability to purchase additional capacity for app development.

	1. **Less to manage** - Infrastructure management, patches, updates and other maintain-need task are on the cloud provider shoulders.
	#
	* ***SaaS*** - If you want everything in the cloud, and don't own any IT equipment, this is a perfect fit. Everyone is using ***SaaS*** for e.g. *Slack*, *Dropbox* are *SaaS* apps. They manage everything, database, managingg user, server hardware, etc.
	#
	### Comparison of ***IaaS, PaaS, SaaS***:

	* Choose a ***SaaS*** CRM solution, offloading all day-to-day management to the third-party vendor, but also giving up all control over features and functionality, data storage, user access and security.

	* Choose a ***PaaS*** solution and build a custom CRM application. In this case, the company would offload management of infrastructure and application development resources to the cloud service provider. The customer would retain complete control over application features, but it would also assume responsibility for managing the application and associated data.

	* Build out backend IT infrastructure on the cloud using ***IaaS***, and use it to build its own development platform and application. The organization's IT team would have complete control over operating systems and server configurations, but also bear the burden of managing and maintaining them, along with the development platform and applications that run on them.

# Related to: 

* https://www.youtube.com/watch?v=M988_fsOSWo
* https://www.ibm.com/cloud/learn/iaas-paas-saas

#cloud #cloudcomputing #cloudnative #environment #servers #IaaS #PaaS #SaaS #onpremises
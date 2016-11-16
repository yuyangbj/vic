# Examples of Deploying a Virtual Container Host #

This topic provides examples of the options of the `vic-machine create` command to use when deploying virtual container hosts in different vSphere configurations.

- [Deploy to an ESXi Host with no Resource Pools and a Single Datastore](#esxi)
- [Deploy to a vCenter Server Cluster](#cluster)
- [Specify External, Management, Client, and Container Networks](#networks)
- [Configure a Non-DHCP Container Network](#ip-range)
- [Set a Static IP Address on the Different Networks](#static-ip)
- [Specify One or More Volume Stores](#volume-stores)
- [Deploy to a Standalone Host in vCenter Server](#standalone)
- [Deploy to a Resource Pool on an ESXi Host](#rp_host)
- [Deploy to a Resource Pool in a vCenter Server Cluster](#rp_cluster)
- [Use Auto-Generated Trusted CA Certificates](#auto_cert)
- [Use Custom Trusted CA Certificates](#custom_cert)
- [Limit Resource Use](#customized)
- [Authorize Access to an Insecure Private Registry Server](#registry)
- [Configure a Proxy Server](#proxy)

For simplicity, unless stated otherwise, these examples assume that the vSphere environment uses trusted certificates signed by a known Certificate Authority (CA), so the `--thumbprint` option is not specified. Similarly, all examples that do not relate explicitly to certificate use specify the `--tls-noverify` option.

For detailed descriptions of all of the `vic-machine create` options, see [Virtual Container Host Deployment Options](vch_installer_options.md).

<a name="esxi"></a>
## Deploy to an ESXi Host with no Resource Pools and a Single Datastore##

You can deploy a virtual container host directly on an ESXi host that is not managed by a vCenter Server instance. This example provides the minimum options required to deploy a virtual container host, namely the `--target` and `--user` options and an authentication option. The `vic-machine create` command prompts you for the password for the ESXi host and deploys a virtual container host with the default name `virtual-container-host`. If there is only one datastore on the host and there are no resource pools, you do not need to specify the `--image-store` or `--compute-resource` options. When deploying to an ESXi host, `vic-machine create` creates a standard virtual switch and a distributed port group, so you do not need to specify any network options if you do not have specific network requirements. The example uses the `--no-tlsverify` option to implement TLS authentication with self-signed untrusted certificates, with no client verification.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target <i>esxi_host_address</i>
--user root
--no-tlsverify
</pre>

<a name="cluster"></a>
## Deploy to a vCenter Server Cluster ##

If vCenter Server has more than one datacenter, you specify the datacenter in the `--target` option.

If vCenter Server manages more than one cluster, you use the `--compute-resource` option to specify the cluster on which to deploy the virtual container host.

When deploying a virtual container host to vCenter Server, you must use the `--bridge-network` option to specify an existing distributed port group for container VMs to use to communicate with each other. For information about how to create a distributed virtual switch and port group, see *Network Requirements* in [Environment Prerequisites for vSphere Integrated Containers Engine Installation](vic_installation_prereqs.md#networkreqs).

This example deploys a virtual container host with the following configuration:

- Provides the vCenter Single Sign-On user and password in the `--target` option. Note that the user name is wrapped in quotes, because it contains the `@` character. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system. 
- Deploys a virtual container host named `vch1` to the cluster `cluster1` in datacenter `dc1`. 
- Uses a distributed port group named `vic-bridge` for the bridge network. 
- Designates `datastore1` as the datastore in which to store container images, the files for the virtual container host appliance, and container VMs. 

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--name vch1
--no-tlsverify
</pre>

<a name="networks"></a>
## Specify External, Management, Client, and Container Networks ##

In addition to the mandatory bridge network, if your vCenter Server environment includes multiple networks, you can direct different types of traffic to different networks. 

- You can direct the traffic between the virtual container host, container VMs, and the internet to a specific network by specifying the `external-network` option. If you do not specify the `external-network` option, the virtual container host uses the default VM Network for external traffic.
- You can direct traffic between ESXi hosts, vCenter Server, and the virtual container host to a specific network by specifying the `--management-network` option. If you do not specify the `--management-network` option, the virtual container host uses the external network for management traffic.
- You can designate a specific network for use by the Docker API by specifying the `--client-network` option. If you do not specify the `--client-network` option, the Docker API uses the external network.
- You can designate a specific network for container VMs to use by specifying the `--container-network` option. Containers use this network if the container developer runs `docker run` or `docker create` with the `--net` option when they run or create a container. This option requires a distributed port group that must exist before you run `vic-machine create`. You cannot use the same distributed port group that you use for the bridge network. You can provide a descriptive name for the network, for use by Docker. If you do not specify a descriptive name, Docker uses the vSphere network name. For example, the descriptive name appears as an available network in the output of `docker info`. 

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, datacenter, cluster, image store, bridge network, and name for the virtual container host.
- Directs external, management, and Docker API traffic to network 1, network 2, and network 3 respectively. Note that the network names are wrapped in quotes, because they contain spaces. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system.
- Designates a distributed port group named `vic-containers` for use by container VMs that are run with the `--net` option.
- Gives the container network the name `vic-container-network`, for use by Docker.  

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--external-network 'network 1'
--management-network 'network 2'
--client-network 'network 3'
--container-network vic-containers:vic-container-network
--name vch1
--no-tlsverify
</pre>

For more information about the networking options, see the [Networking Options section](vch_installer_options.md#networking) in Virtual Container Host Deployment Options.

<a name="ip-range"></a>
## Configure a Non-DHCP Container Network ##

If the network that you designate as the container network in the `--container-network` option does not support DHCP, you can configure the gateway, DNS server, and a range of IP addresses for container VMs to use. 

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, datacenter, cluster, image store, bridge network, and name for the virtual container host.
- Uses the default VM Network for the external, management, and client networks.
- Designates a distributed port group named `vic-containers` for use by container VMs that are run with the `--net` option.
- Gives the container network the name `vic-container-network`, for use by Docker. 
- Specifies the gateway, two DNS servers, and a range of IP addresses on the container network for container VMs to use.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--container-network vic-containers:vic-container-network
--container-network-gateway vic-containers:<i>gateway_ip_address</i>/255.255.255.0
--container-network-dns vic-containers:<i>dns1_ip_address</i>
--container-network-dns vic-containers:<i>dns2_ip_address</i>
--container-network-ip-range vic-containers:192.168.100.0/24
--name vch1
--no-tlsverify
</pre>

For more information about the container network options, see the [container network section](vch_installer_options.md#container-network) in Virtual Container Host Deployment Options.

<a name="static-ip"></a>
## Set a Static IP Address on the Different Networks ##

If you specify networks for any or all of the external, management, and client networks, you can deploy the virtual container host so that the virtual container host endpoint VM has a static IP address on those networks. 

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, datacenter, cluster, image store, bridge network, and name for the virtual container host.
- Directs external, management, and Docker API traffic to network 1, network 2, and network 3 respectively. Note that the network names are wrapped in quotes, because they contain spaces. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system.
- Sets a DNS server for use by the virtual container host.
- Sets a static IP address for the virtual container host endpoint VM on each of the external, management, and client networks. 

**NOTE**: When you specify a static IP address for the virtual container host on the client network, `vic-machine create` uses this address as the Common Name with which to create auto-generated trusted certificates. In this case, full TLS authentication is implemented by default and you do not need to specify any authentication options. 

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--external-network 'network 1'
--external-network-gateway 192.168.1.1/24
--external-network-ip 192.168.1.10/24
--management-network 'network 2'
--management-network-gateway 192.168.2.1/24
--management-network-ip 192.168.2.10/24
--client-network 'network 3'
--client-network-gateway 192.168.3.1/24
--client-network-ip 192.168.3.10/24
--dns-server <i>dns_server_address</i>
--name vch1
</pre>

For more information about the networking options, see the [Options for Specifying a Static IP Address for the Virtual Container Host Endpoint VM](vch_installer_options.md#static-ip) in Virtual Container Host Deployment Options.

<a name="volume-stores"></a>
## Specify One or More Volume Stores ##

If container application developers will use the `docker volume create` command to create containers that use volumes, you must create volume stores when you deploy virtual container hosts. You specify volume stores in the `--volume-store` option. You can specify `--volume-store` multiple times to create multiple volume stores. 

When you create a volume store, you specify the name of the datastore to use and an optional path to a folder on that datastore. You also specify a descriptive name for that volume store for use by Docker.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, datacenter, cluster, bridge network, and name for the virtual container host.
- Specifies the `volumes` folder on `datastore 1` as the default volume store. Creating a volume store named `default` allows container application developers to create anonymous or named volumes by using `docker create -v`. 
- Specifies a second volume store named `volume_store_2` in the `volumes` folder on `datastore 2`. 
- Note that the datastore names are wrapped in quotes, because they contain spaces. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--bridge-network vic-bridge
--image-store 'datastore 1'
--volume-store 'datastore 1'/volumes:default</i>
--volume-store 'datastore 2'/volumes:volume_store_2</i>
--name vch1
--no-tlsverify
</pre>

For more information about volume stores, see the [volume-store section](vch_installer_options.md#volume-store) in Virtual Container Host Deployment Options. 

<a name="standalone"></a> 
## Deploy to a Standalone Host in vCenter Server ##

If vCenter Server manages multiple standalone ESXi hosts that are not part of a cluster, you use the `--compute-resource` option to specify the address of the ESXi host to which to deploy the virtual container host.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, bridge network, and name for the virtual container host.
- Deploys the virtual container host on the ESXi host with the FQDN `esxihost1.organization.company.com` in the datacenter `dc1`. You can also specify an IP address.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--image-store datastore1
--bridge-network vic-bridge
--compute-resource esxihost1.organization.company.com
--name vch1
--no-tlsverify
</pre>

<a name="rp_host"></a>
## Deploy to a Resource Pool on an ESXi Host ##

To deploy a virtual container host in a specific resource pool on an ESXi host that is not managed by vCenter Server, you specify the resource pool name in the `--compute-resource` option. 

This example deploys a virtual container host with the following configuration:

- Specifies the user name and password, and a name for the virtual container host.
- Designates `rp 1` as the resource pool in which to place the virtual container host. Note that the resource pool name is wrapped in quotes, because it contains a space. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target root:<i>password</i>@<i>esxi_host_address</i>
--compute-resource 'rp 1'
--name vch1
--no-tlsverify
</pre>

<a name="rp_cluster"></a>
## Deploy to a Resource Pool in a vCenter Server Cluster ##

To deploy a virtual container host in a resource pool in a vCenter Server cluster, you specify the names of the cluster and resource pool in the `compute-resource` option.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, datacenter, image store, bridge network, and name for the virtual container host.
- Designates `rp 1` in cluster `cluster 1` as the resource pool in which to place the virtual container host. Note that the resource pool and cluster names are wrapped in quotes, because they contain spaces. Use single quotes if you are using `vic-machine` on a Linux or Mac OS system and double quotes on a Windows system.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource 'cluster 1'/'rp 1'
--image-store datastore1
--bridge-network vic-bridge
--name vch1
--no-tlsverify
</pre>

<a name="auto_cert"></a>
##  Use Auto-Generated Trusted CA Certificates ##

You can deploy a virtual container host that implements two-way authentication with trusted auto-generated TLS certificates that are signed by a Certificate Authority (CA). To automatically generate a trusted CA certificate, you provide information that `vic-machine create` uses to populate the fields of a certificate request. At a minimum, you must specify the FQDN or the name of the domain in which the virtual container host will run in the `--tls-cname` option. `vic-machine create` uses the name as the Common Name in the certificate request. You can also optionally specify a CA file, an organization name, and a size for the certificate key. 

**NOTE**: Because the `--tls-cname` option requires an FQDN or domain name, you must have a DNS service running on the client network on which you deploy the virtual container host. However, if you specify a static IP address for the virtual container host endpoint VM on the client network, `vic-machine create` uses this address as the Common Name with which to create an auto-generated trusted certificate. In this case, full TLS authentication is implemented by default and you do not need to specify any authentication options, and DNS is not required on the client network.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, cluster, bridge network, and name for the virtual container host.
- Provides `vch1.example.org` as the FQDN for the virtual container host, for use as the Common Name in the certificate.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--tls-cname vch1.example.org
--name vch1
</pre>

For more information about using auto-generated CA certificates, see the [Security Options section](vch_installer_options.md#security) in Virtual Container Host Deployment Options.

<a name="custom_cert"></a>
## Use Custom Trusted CA Certificates ##

If your development environment uses custom CA certificates to authenticate connections between Docker clients and virtual container hosts, use the `--cert` and `--key` options to provide the paths to a custom X.509 certificate and its key when you deploy a virtual container host. The paths to the certificate and key files must be relative to the location from which you are running `vic-machine create`.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, cluster, bridge network, and name for the virtual container host.
- Provides the paths relative to the current location of the `*.pem` files for the custom CA certificate and key files.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--cert ../some/relative/path/<i>certificate_file</i>.pem
--key ../some/relative/path/<i>key_file</i>.pem
--name vch1
</pre>

For more information about using custom CA certificates, see the [Advanced Security Options section](vch_installer_options.md#adv-security) in Virtual Container Host Deployment Options.

<a name="customized"></a>
## Limit Resource Use ##

To limit the amount of system resources that the container VMs in a virtual container host can use, you can set resource limits on the virtual container host vApp. 

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, cluster, bridge network, and name for the virtual container host.
- Sets resource limits on the virtual container host by imposing memory and CPU reservations, limits, and shares.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--memory 1024
--memory-reservation 1024
--memory-shares low
--cpu 1024
--cpu-reservation 1024
--cpu-shares low
--name vch1
--no-tlsverify
</pre>

For more information about setting resource use limitations on virtual container hosts, see the [vApp Deployment Options](vch_installer_options.md#deployment) and [Advanced Resource Management Options](vch_installer_options.md#adv-mgmt) sections in Virtual Container Host Deployment Options.

<a name="registry"></a>
## Authorize Access to an Insecure Private Registry Server ##

An insecure private registry server is a private registry server for Docker images that is secured by self-signed certificates rather than by TLS. To authorize connections from a virtual container host to an insecure private registry server, set the `insecure-registry` option. You can specify `insecure-registry` multiple times to allow connections from the virtual container host to multiple insecure private registry servers.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, cluster, bridge network, and name for the virtual container host.
- Authorizes the virtual container host to pull Docker images from the insecure private registry servers located at the URLs <i>registry_URL_1</i> and <i>registry_URL_2</i>.
- The registry server at <i>registry_URL_2</i> listens for connections on port 5000. 

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--insecure-registry <i>registry_URL_1</i>
--insecure-registry <i>registry_URL_2:5000</i>
--name vch1
--no-tlsverify
</pre>

For more information about configuring virtual container hosts to connect to insecure private registry servers, see the section on the [`insecure-registry` option](vch_installer_options.md#registry) in Virtual Container Host Deployment Options.

**NOTE**: The current builds of vSphere Integrated Containers do not yet support private registry servers that you secure by using TLS certificates.

<a name="proxy"></a>
## Configure a Proxy Server ##

If your network access is controlled by a proxy server, you must   configure a virtual container host to connect to the proxy server when you deploy it, so that it can pull images from external sources.

This example deploys a virtual container host with the following configuration:

- Specifies the user name, password, image store, cluster, bridge network, and name for the virtual container host.
- Configures the virtual container host to access the network via an HTTPS proxy server.

<pre>vic-machine<i>-darwin</i><i>-linux</i><i>-windows</i> create
--target 'Administrator@vsphere.local':<i>password</i>@<i>vcenter_server_address</i>/dc1
--compute-resource cluster1
--image-store datastore1
--bridge-network vic-bridge
--https-proxy https://<i>proxy_server_address</i>:<i>port</i>
--name vch1
--no-tlsverify
</pre>
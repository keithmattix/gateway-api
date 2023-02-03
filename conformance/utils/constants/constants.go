package constants

const (
	// UnspecifiedIP constant for empty IP address
	UnspecifiedIP = "0.0.0.0"

	// DefaultClusterLocalDomain the default service domain suffix for Kubernetes, if not overridden in config.
	DefaultClusterLocalDomain = "cluster.local"

	// DefaultClusterSetLocalDomain is the default domain suffix for Kubernetes Multi-Cluster Services (MCS)
	// used for load balancing requests against endpoints across the ClusterSet (i.e. mesh).
	DefaultClusterSetLocalDomain = "clusterset.local"

	// KubeSystemNamespace is the system namespace where we place kubernetes system components.
	KubeSystemNamespace string = "kube-system"

	// KubePublicNamespace is the namespace where we place kubernetes public info (ConfigMaps).
	KubePublicNamespace string = "kube-public"

	// KubeNodeLeaseNamespace is the namespace for the lease objects associated with each kubernetes node.
	KubeNodeLeaseNamespace string = "kube-node-lease"

	// LocalPathStorageNamespace is the namespace for dynamically provisioning persistent local storage with
	// Kubernetes. Typically used with the Kind cluster: https://github.com/rancher/local-path-provisioner
	LocalPathStorageNamespace string = "local-path-storage"
)

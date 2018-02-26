package api

// Cluster represents high-level information about a LXD cluster.
type Cluster struct {
	ServerName string `json:"server_name" yaml:"server_name"`
}

// ClusterPut represents the fields required to bootstrap or join a LXD
// cluster.
//
// API extension: cluster
type ClusterPut struct {
	ServerName    string `json:"server_name" yaml:"server_name"`
	Address       string `json:"address" yaml:"address"`
	TargetAddress string `json:"target_address" yaml:"target_address"`
	TargetCert    string `json:"target_cert" yaml:"target_cert"`
	TargetCA      []byte `json:"target_ca" yaml:"target_ca"`
}

// ClusterMemberPost represents the fields required to rename a LXD node.
//
// API extension: clustering
type ClusterMemberPost struct {
	Name string `json:"name" yaml:"name"`
}

// ClusterMember represents the a LXD node in the cluster.
//
// API extension: clustering
type ClusterMember struct {
	Name     string `json:"name" yaml:"name"`
	URL      string `json:"url" yaml:"url"`
	Database bool   `json:"database" yaml:"database"`
	Status   string `json:"status" yaml:"status"`
	Message  string `json:"message" yaml:"message"`
}

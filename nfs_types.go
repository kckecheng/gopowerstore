/*
 *
 * Copyright 2020 Dell EMC Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopowerstore

type NFSExportDefaultAccessEnum string

const (
	No_Access      NFSExportDefaultAccessEnum = "No_Access"
	Read_Only      NFSExportDefaultAccessEnum = "Read_Only"
	Read_Write     NFSExportDefaultAccessEnum = "Read_Write"
	Root           NFSExportDefaultAccessEnum = "Root"
	Read_Only_Root NFSExportDefaultAccessEnum = "Read_Only_Root "
)

// NFSExportCreate details about creation of new NFS export
type NFSExportCreate struct {
	// NFS Export name
	Name string `json:"name"`
	// NFS Export description
	Description string `json:"description,omitempty"`
	// Unique identifier of the file system on which the NFS Export was created
	FileSystemID string `json:"file_system_id"`
	// Local path to a location within the file system.
	Path string `json:"path"`
}

// NFSExportModify details about modification of exiting NFS export
type NFSExportModify struct {
	// Hosts to add to the current read_write_root_hosts list. Hosts can be entered by Hostname, IP addresses
	AddHosts *[]string `json:"add_read_write_root_hosts,omitempty"`
	// An optional description for the host.
	// The description should not be more than 256 UTF-8 characters long and should not have any unprintable characters.
	Description *string `json:"description,omitempty"`
	// Hosts to remove from the current read_write_root_hosts list. Hosts can be entered by Hostname, IP addresses.
	RemoveHosts *[]string `json:"remove_read_write_root_hosts,omitempty"`
}

// NFSServerCreate details about creation of new NFS server
type NFSServerCreate struct {
	// Unique identifier of the NAS server
	NasServerID string `json:"nas_server_id"`
	// The name that will be used by NFS clients to connect to this NFS server
	// This name is required when using secure NFS
	HostName string `json:"host_name,omitempty"`
	// Indicates whether NFSv3 is enabled on the NAS server. When enabled, NFS shares can be accessed with NFSv3
	IsNFSv3Enabled bool `json:"is_nfsv3_enabled,omitempty"`
	// Indicates whether NFSv4 is enabled on the NAS server. When enabled, NFS shares can be accessed with NFSv4
	IsNFSv4Enabled bool `json:"is_nfsv4_enabled,omitempty"`
	// Indicates whether secure NFS is enabled on the NFS server
	IsSecureEnabled bool `json:"is_secure_enabled,omitempty"`
}

// Details about the NFSExport.
type NFSExport struct {
	// Unique id of the NFS Export
	ID string `json:"id,omitempty"`
	// Unique identifier of the file system on which the NFS Export was created
	FileSystemID string `json:"file_system_id,omitempty"`
	// NFS Export name
	Name string `json:"name,omitempty"`
	// NFS Export description
	Description string `json:"description,omitempty"`
	// Default access level for all hosts that can access the Export
	// [ No_Access, Read_Only, Read_Write, Root, Read_Only_Root ]
	DefaultAccess NFSExportDefaultAccessEnum `json:"default_access"`
	// Local path to a location within the file system.
	// With NFS, each export must have a unique local path.
	Path string `json:"path,omitempty"`
}

// Details about the file interface
type FileInterface struct {
	// Unique id of the file interface
	ID string `json:"id"`
	// Ip address of file interface
	IpAddress string `json:"ip_address"`
}

// Fields returns fields which must be requested to fill struct
func (n *NFSExport) Fields() []string {
	return []string{"description", "id", "name", "file_system_id", "default_access", "path"}
}

// Fields returns fields which must be requested to fill struct
func (n *FileInterface) Fields() []string {
	return []string{"id", "ip_address"}
}

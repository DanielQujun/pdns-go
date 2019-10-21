/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pdnsapi

// The RREntry object represents a single record.
type Record struct {

	// The content of this record
	Content string `json:"content"`

	// Whether or not this record is disabled. When unset, the record is not disabled
	Disabled bool `json:"disabled,omitempty"`

	// If set to true, the server will find the matching reverse zone and create a PTR there. Existing PTR records are replaced. If no matching reverse Zone, an error is thrown. Only valid in client bodies, only valid for A and AAAA types. Not returned by the server. This feature is deprecated and will be removed in 4.3.0.
	SetPtr bool `json:"set-ptr,omitempty"`
}

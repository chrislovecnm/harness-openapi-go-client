/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Template stable version update Response
type TemplateUpdateStableResponse struct {
	// Version which is set as stable for the given Template 
	StableVersion string `json:"stable_version,omitempty"`
}

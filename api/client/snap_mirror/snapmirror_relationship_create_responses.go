// Code generated by go-swagger; DO NOT EDIT.

package snap_mirror

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/ontap-go/api/models"
)

// SnapmirrorRelationshipCreateReader is a Reader for the SnapmirrorRelationshipCreate structure.
type SnapmirrorRelationshipCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnapmirrorRelationshipCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapmirrorRelationshipCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipCreateCreated creates a SnapmirrorRelationshipCreateCreated with default headers values
func NewSnapmirrorRelationshipCreateCreated() *SnapmirrorRelationshipCreateCreated {
	return &SnapmirrorRelationshipCreateCreated{}
}

/*
SnapmirrorRelationshipCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnapmirrorRelationshipCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship create created response has a 2xx status code
func (o *SnapmirrorRelationshipCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship create created response has a 3xx status code
func (o *SnapmirrorRelationshipCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship create created response has a 4xx status code
func (o *SnapmirrorRelationshipCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship create created response has a 5xx status code
func (o *SnapmirrorRelationshipCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship create created response a status code equal to that given
func (o *SnapmirrorRelationshipCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snapmirror relationship create created response
func (o *SnapmirrorRelationshipCreateCreated) Code() int {
	return 201
}

func (o *SnapmirrorRelationshipCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateCreated %s", 201, payload)
}

func (o *SnapmirrorRelationshipCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateCreated %s", 201, payload)
}

func (o *SnapmirrorRelationshipCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipCreateAccepted creates a SnapmirrorRelationshipCreateAccepted with default headers values
func NewSnapmirrorRelationshipCreateAccepted() *SnapmirrorRelationshipCreateAccepted {
	return &SnapmirrorRelationshipCreateAccepted{}
}

/*
SnapmirrorRelationshipCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship create accepted response has a 2xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship create accepted response has a 3xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship create accepted response has a 4xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship create accepted response has a 5xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship create accepted response a status code equal to that given
func (o *SnapmirrorRelationshipCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapmirror relationship create accepted response
func (o *SnapmirrorRelationshipCreateAccepted) Code() int {
	return 202
}

func (o *SnapmirrorRelationshipCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipCreateDefault creates a SnapmirrorRelationshipCreateDefault with default headers values
func NewSnapmirrorRelationshipCreateDefault(code int) *SnapmirrorRelationshipCreateDefault {
	return &SnapmirrorRelationshipCreateDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorRelationshipCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115542 | Invalid value given for the field. |
| 6619441 | The source volume cannot be the same as the destination volume. |
| 6619599 | Only the \"none\" tiering policy is supported when creating a destination volume with \"snapshot_locking_enabled\" set to true or if \"snapshot_locking_enabled\" is set to true on the source volume. |
| 6619637 | Relationship with specified destination volume already exists. |
| 6619699 | Schedule not found. |
| 6620374 | Internal error. Failed to get SVM information. |
| 6620478 | Internal error. Failed to check SnapMirror capability. |
| 6620654 | Invalid SnapMirror Consistency Group name. |
| 6621125 | The policy is not valid for relationships with FlexGroup volume endpoints. Only policies without Snapshot copy creation schedules are supported for these relationships. |
| 6621458 | The destination Consistency Group is the source of a SnapMirror Synchronous (SM-S) relationship. Sources of SM-S relationships cannot be the destination of any other SnapMirror relationship. |
| 6621782 | A property of the policy is not valid for relationships between these endpoints. |
| 6621834 | Object store configuration does not exist for specified SVM. |
| 6622088 | SnapMirror Asynchronous relationship is not supported on a Consistency Group volume that has Snapshot copy locking enabled. |
| 13303819 | Could not retrieve SnapMirror policy information. |
| 13303821 | Invalid SnapMirror policy UUID. |
| 13303841 | This operation is not supported for SnapMirror relationships between these endpoints. |
| 13303852 | destination.path provided does not contain \\\":\\\". |
| 13303853 | Restore relationships are not supported for SVM-DR endpoints. |
| 13303866 | Associating the specified SnapMirror policy with this SnapMirror relationship is not supported. |
| 13303868 | Create of destination endpoint and SnapMirror relationship failed. |
| 13303869 | Creating a destination endpoint is not supported for restore relationships. |
| 13303870 | A tiering policy cannot be specified if tiering is not being set to supported. |
| 13303871 | Storage service properties cannot be specified if the storage service is not being enabled. |
| 13303872 | Specified property requires a later effective cluster version. |
| 13303873 | Specifying a state when creating a relationship is only supported when creating a destination endpoint. |
| 13303874 | Specified state is not supported when creating this relationship. |
| 13303875 | Destination aggregates do not have sufficient space for hosting copies of source volumes. |
| 13303876 | Destination cluster does not have composite aggregates. |
| 13303877 | Source or destination cluster must be specified. |
| 13303878 | The specified fields do not match. |
| 13303879 | Source cluster name or UUID is needed to provision a destination SVM on the local cluster. |
| 13303880 | Source cluster must be remote for provisioning a destination SVM on the local cluster. |
| 13303881 | Network validation failed. |
| 13303882 | SVM validation failed. |
| 13303883 | Encryption is not enabled on the destination cluster. |
| 13303886 | SVM peer permission not found. |
| 13303887 | Synchronous SnapMirror relationships between FlexGroup volumes are not supported. |
| 13303888 | Synchronous SnapMirror relationships require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303889 | Asynchronous SnapMirror relationships between FlexGroup volumes require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303890 | Asynchronous SnapMirror relationships between FlexVol volumes require an effective cluster version of 9.3, 9.5, or later on both the source and destination clusters. |
| 13303891 | Creating a destination endpoint with storage service requires an effective cluster version of 9.7 or later. |
| 13303892 | Fetching remote information from the destination cluster failed. |
| 13303893 | Updating job description failed. |
| 13303894 | Destination volume name is invalid. It must contain the source volume name and have a suffix when creating a destination endpoint on a cluster with an effective cluster version of 9.6 or earlier. |
| 13303895 | Operation on the remote destination cluster is not supported. |
| 13303916 | FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303918 | No suitable destination aggregate type is available. |
| 13303919 | Only FabricPool enabled aggregates are available on the destination. |
| 13303920 | Only SnapLock aggregates are available on the destination. FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303921 | Unable to retrieve the SnapMirror capabilities of the destination cluster. |
| 13303922 | Specified source SVM is not a data SVM. |
| 13303923 | Specified destination SVM is not a data SVM. |
| 13303924 | Source SVM has an invalid Snapshot copy policy. |
| 13303925 | SnapMirror validation has failed. |
| 13303930 | The specified tiering policy is not supported for destination volumes of Synchronous relationships. |
| 13303938 | Fetching information from the local cluster failed. |
| 13303939 | Could not create an SVM peer relationship. |
| 13303944 | An SVM-DR relationship is not supported because the source SVM has CIFS configured and the associated SnapMirror policy has either the "identity_preservation" property not set or set to "exclude_network_and_protocol_config". |
| 13303949 | This SnapMirror policy is only supported for relationships with object store destination endpoints. |
| 13303966 | Consistency Group relationships require a policy of type \"sync\" with a sync_type of \"automated_failover\". |
| 13303967 | Consistency Group volume is not a FlexVol volume. |
| 13303968 | Unsupported volume type for the Consistency Group. |
| 13303969 | SnapMirror relationships between SVM endpoints and object store endpoints are not supported. |
| 13303970 | Unsupported policy type for the Consistency Group. |
| 13303971 | SnapMirror relationships between Consistency Group endpoints and object store endpoints are not supported. |
| 13303976 | Source or destination SVM is already part of an SVM-DR relation. |
| 13303977 | Destination Consistency Group volume UUIDs are not expected while provisioning the destination volumes. |
| 13303978 | Number of Consistency Group volume names and UUIDs does not match. |
| 13303979 | Number of Consistency Group volumes exceeds the allowed limit. |
| 13303980 | Number of source and destination Consistency Group volumes do not match. |
| 13303981 | ISCSI or FCP protocol is not configured. |
| 13303982 | SAN data interface is not configured on the SVM. |
| 13304021 | No suitable storage can be found meeting the specified requirements. No FabricPool enabled aggregates are available on the destination. |
| 13304022 | No suitable storage can be found meeting the specified requirements. No non-root, non-taken-over, non-SnapLock, non-composite aggregates are available on the destination. |
| 13304032 | In an "All SAN Array", an SVM-DR relationship is not supported when the associated SnapMirror policy does not have the "identity_preservation" property set to "exclude_network_and_protocol_config". |
| 13304080 | Specified UUID and name do not match. |
| 13304082 | Specified properties are mutually exclusive. |
| 13304083 | The specified property is not supported because all nodes in the cluster are not capable of supporting the property. |
| 13304093 | The property specified is not supported for the specified relationships. |
| 13304098 | This SnapMirror policy is not supported for SnapMirror relationhips with SnapLock volumes. |
| 13304099 | SnapLock Compliance Clock is not running on all nodes in the destination cluster. |
| 13304108 | Schedule not found in the Administrative SVM or the SVM for the relationship. |
| 13304112 | File restore from a Consistency Group asynchronous SnapMirror relationship endpoint is not supported. |
| 13304132 | Creating a destination endpoint is not supported with the \"backoff_level\" property. |
| 53411897 | The specified source volumes do not match the volumes contained in the source consistency group. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
##### Above error msgs have been added in sorted(ascending) order
##### of error codes, for example - 13303945 is error code for
##### SM_REST_SCHEDULE_CONFLICT_FOR_SM_OP and for 13303949 is err
##### code for SM_REST_POLICY_ONLY_SUPPORTED_FOR_OBJSTORE, so
##### SM_REST_SCHEDULE_CONFLICT_FOR_SM_OP has been placed before
##### SM_REST_POLICY_ONLY_SUPPORTED_FOR_OBJSTORE.
##### Also for private error msg, add private tags as  after adding the error msg at correct place(in sorted order).
##### Please make sure new error messages are being put at correct place
##### so that order is maintained.
*/
type SnapmirrorRelationshipCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror relationship create default response has a 2xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror relationship create default response has a 3xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror relationship create default response has a 4xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror relationship create default response has a 5xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror relationship create default response a status code equal to that given
func (o *SnapmirrorRelationshipCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror relationship create default response
func (o *SnapmirrorRelationshipCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirror_relationship_create default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirror_relationship_create default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
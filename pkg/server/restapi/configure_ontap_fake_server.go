// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/application"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/cloud"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/cluster"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/n_a_s"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/n_d_m_p"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/n_v_me"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/name_services"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/networking"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/object_store"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/s_a_n"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/s_vm"
	securityops "github.com/metal-stack/ontap-go/pkg/server/restapi/operations/security"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/snap_lock"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/snap_mirror"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/storage"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/support"
)

//go:generate swagger generate server --target ../../server --name OntapFakeServer --spec ../../../spec/ontap.yaml --principal interface{} --skip-models

func configureFlags(api *operations.OntapFakeServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OntapFakeServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()
	api.MultipartformProducer = runtime.DiscardProducer

	// Applies when the Authorization header is set with the Basic scheme
	if api.SimpleAuth == nil {
		api.SimpleAuth = func(user string, pass string) (interface{}, error) {
			return nil, errors.NotImplemented("basic auth  (simple) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// cluster.SoftwareUploadMaxParseMemory = 32 << 20

	if api.SecurityAccountCollectionGetHandler == nil {
		api.SecurityAccountCollectionGetHandler = securityops.AccountCollectionGetHandlerFunc(func(params securityops.AccountCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityAccountCreateHandler == nil {
		api.SecurityAccountCreateHandler = securityops.AccountCreateHandlerFunc(func(params securityops.AccountCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountCreate has not yet been implemented")
		})
	}
	if api.SecurityAccountDeleteHandler == nil {
		api.SecurityAccountDeleteHandler = securityops.AccountDeleteHandlerFunc(func(params securityops.AccountDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountDelete has not yet been implemented")
		})
	}
	if api.SecurityAccountDuoDeleteHandler == nil {
		api.SecurityAccountDuoDeleteHandler = securityops.AccountDuoDeleteHandlerFunc(func(params securityops.AccountDuoDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountDuoDelete has not yet been implemented")
		})
	}
	if api.SecurityAccountDuogroupDeleteHandler == nil {
		api.SecurityAccountDuogroupDeleteHandler = securityops.AccountDuogroupDeleteHandlerFunc(func(params securityops.AccountDuogroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountDuogroupDelete has not yet been implemented")
		})
	}
	if api.SecurityAccountGetHandler == nil {
		api.SecurityAccountGetHandler = securityops.AccountGetHandlerFunc(func(params securityops.AccountGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountGet has not yet been implemented")
		})
	}
	if api.SecurityAccountModifyHandler == nil {
		api.SecurityAccountModifyHandler = securityops.AccountModifyHandlerFunc(func(params securityops.AccountModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountModify has not yet been implemented")
		})
	}
	if api.SecurityAccountPasswordCreateHandler == nil {
		api.SecurityAccountPasswordCreateHandler = securityops.AccountPasswordCreateHandlerFunc(func(params securityops.AccountPasswordCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountPasswordCreate has not yet been implemented")
		})
	}
	if api.SecurityAccountPublickeyDeleteHandler == nil {
		api.SecurityAccountPublickeyDeleteHandler = securityops.AccountPublickeyDeleteHandlerFunc(func(params securityops.AccountPublickeyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountPublickeyDelete has not yet been implemented")
		})
	}
	if api.SecurityAccountTotpDeleteHandler == nil {
		api.SecurityAccountTotpDeleteHandler = securityops.AccountTotpDeleteHandlerFunc(func(params securityops.AccountTotpDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AccountTotpDelete has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryCollectionGetHandler == nil {
		api.NasActiveDirectoryCollectionGetHandler = n_a_s.ActiveDirectoryCollectionGetHandlerFunc(func(params n_a_s.ActiveDirectoryCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryCollectionGet has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryCreateHandler == nil {
		api.NasActiveDirectoryCreateHandler = n_a_s.ActiveDirectoryCreateHandlerFunc(func(params n_a_s.ActiveDirectoryCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryCreate has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryDeleteHandler == nil {
		api.NasActiveDirectoryDeleteHandler = n_a_s.ActiveDirectoryDeleteHandlerFunc(func(params n_a_s.ActiveDirectoryDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryDelete has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryGetHandler == nil {
		api.NasActiveDirectoryGetHandler = n_a_s.ActiveDirectoryGetHandlerFunc(func(params n_a_s.ActiveDirectoryGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryGet has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryModifyHandler == nil {
		api.NasActiveDirectoryModifyHandler = n_a_s.ActiveDirectoryModifyHandlerFunc(func(params n_a_s.ActiveDirectoryModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryModify has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryPreferredDcCollectionGetHandler == nil {
		api.NasActiveDirectoryPreferredDcCollectionGetHandler = n_a_s.ActiveDirectoryPreferredDcCollectionGetHandlerFunc(func(params n_a_s.ActiveDirectoryPreferredDcCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryPreferredDcCollectionGet has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryPreferredDcCreateHandler == nil {
		api.NasActiveDirectoryPreferredDcCreateHandler = n_a_s.ActiveDirectoryPreferredDcCreateHandlerFunc(func(params n_a_s.ActiveDirectoryPreferredDcCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryPreferredDcCreate has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryPreferredDcDeleteHandler == nil {
		api.NasActiveDirectoryPreferredDcDeleteHandler = n_a_s.ActiveDirectoryPreferredDcDeleteHandlerFunc(func(params n_a_s.ActiveDirectoryPreferredDcDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryPreferredDcDelete has not yet been implemented")
		})
	}
	if api.NasActiveDirectoryPreferredDcGetHandler == nil {
		api.NasActiveDirectoryPreferredDcGetHandler = n_a_s.ActiveDirectoryPreferredDcGetHandlerFunc(func(params n_a_s.ActiveDirectoryPreferredDcGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ActiveDirectoryPreferredDcGet has not yet been implemented")
		})
	}
	if api.StorageAggregateCollectionGetHandler == nil {
		api.StorageAggregateCollectionGetHandler = storage.AggregateCollectionGetHandlerFunc(func(params storage.AggregateCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregateCollectionGet has not yet been implemented")
		})
	}
	if api.StorageAggregateCreateHandler == nil {
		api.StorageAggregateCreateHandler = storage.AggregateCreateHandlerFunc(func(params storage.AggregateCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregateCreate has not yet been implemented")
		})
	}
	if api.StorageAggregateDeleteHandler == nil {
		api.StorageAggregateDeleteHandler = storage.AggregateDeleteHandlerFunc(func(params storage.AggregateDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregateDelete has not yet been implemented")
		})
	}
	if api.StorageAggregateGetHandler == nil {
		api.StorageAggregateGetHandler = storage.AggregateGetHandlerFunc(func(params storage.AggregateGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregateGet has not yet been implemented")
		})
	}
	if api.StorageAggregateModifyHandler == nil {
		api.StorageAggregateModifyHandler = storage.AggregateModifyHandlerFunc(func(params storage.AggregateModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregateModify has not yet been implemented")
		})
	}
	if api.StorageAggregatePerformanceMetricsCollectionGetHandler == nil {
		api.StorageAggregatePerformanceMetricsCollectionGetHandler = storage.AggregatePerformanceMetricsCollectionGetHandlerFunc(func(params storage.AggregatePerformanceMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.AggregatePerformanceMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreAllSvmBucketsCollectionGetHandler == nil {
		api.ObjectStoreAllSvmBucketsCollectionGetHandler = object_store.AllSvmBucketsCollectionGetHandlerFunc(func(params object_store.AllSvmBucketsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.AllSvmBucketsCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityAntiRansomwareSuspectCollectionGetHandler == nil {
		api.SecurityAntiRansomwareSuspectCollectionGetHandler = securityops.AntiRansomwareSuspectCollectionGetHandlerFunc(func(params securityops.AntiRansomwareSuspectCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AntiRansomwareSuspectCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityAntiRansomwareSuspectDeleteHandler == nil {
		api.SecurityAntiRansomwareSuspectDeleteHandler = securityops.AntiRansomwareSuspectDeleteHandlerFunc(func(params securityops.AntiRansomwareSuspectDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AntiRansomwareSuspectDelete has not yet been implemented")
		})
	}
	if api.ApplicationApplicationCollectionGetHandler == nil {
		api.ApplicationApplicationCollectionGetHandler = application.ApplicationCollectionGetHandlerFunc(func(params application.ApplicationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentCollectionGetHandler == nil {
		api.ApplicationApplicationComponentCollectionGetHandler = application.ApplicationComponentCollectionGetHandlerFunc(func(params application.ApplicationComponentCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentGetHandler == nil {
		api.ApplicationApplicationComponentGetHandler = application.ApplicationComponentGetHandlerFunc(func(params application.ApplicationComponentGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentSnapshotCollectionGetHandler == nil {
		api.ApplicationApplicationComponentSnapshotCollectionGetHandler = application.ApplicationComponentSnapshotCollectionGetHandlerFunc(func(params application.ApplicationComponentSnapshotCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentSnapshotCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentSnapshotCreateHandler == nil {
		api.ApplicationApplicationComponentSnapshotCreateHandler = application.ApplicationComponentSnapshotCreateHandlerFunc(func(params application.ApplicationComponentSnapshotCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentSnapshotCreate has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentSnapshotDeleteHandler == nil {
		api.ApplicationApplicationComponentSnapshotDeleteHandler = application.ApplicationComponentSnapshotDeleteHandlerFunc(func(params application.ApplicationComponentSnapshotDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentSnapshotDelete has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentSnapshotGetHandler == nil {
		api.ApplicationApplicationComponentSnapshotGetHandler = application.ApplicationComponentSnapshotGetHandlerFunc(func(params application.ApplicationComponentSnapshotGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentSnapshotGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationComponentSnapshotRestoreHandler == nil {
		api.ApplicationApplicationComponentSnapshotRestoreHandler = application.ApplicationComponentSnapshotRestoreHandlerFunc(func(params application.ApplicationComponentSnapshotRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationComponentSnapshotRestore has not yet been implemented")
		})
	}
	if api.ApplicationApplicationCreateHandler == nil {
		api.ApplicationApplicationCreateHandler = application.ApplicationCreateHandlerFunc(func(params application.ApplicationCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationCreate has not yet been implemented")
		})
	}
	if api.ApplicationApplicationDeleteHandler == nil {
		api.ApplicationApplicationDeleteHandler = application.ApplicationDeleteHandlerFunc(func(params application.ApplicationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationDelete has not yet been implemented")
		})
	}
	if api.ApplicationApplicationGetHandler == nil {
		api.ApplicationApplicationGetHandler = application.ApplicationGetHandlerFunc(func(params application.ApplicationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationModifyHandler == nil {
		api.ApplicationApplicationModifyHandler = application.ApplicationModifyHandlerFunc(func(params application.ApplicationModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationModify has not yet been implemented")
		})
	}
	if api.ApplicationApplicationSnapshotCollectionGetHandler == nil {
		api.ApplicationApplicationSnapshotCollectionGetHandler = application.ApplicationSnapshotCollectionGetHandlerFunc(func(params application.ApplicationSnapshotCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationSnapshotCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationSnapshotCreateHandler == nil {
		api.ApplicationApplicationSnapshotCreateHandler = application.ApplicationSnapshotCreateHandlerFunc(func(params application.ApplicationSnapshotCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationSnapshotCreate has not yet been implemented")
		})
	}
	if api.ApplicationApplicationSnapshotDeleteHandler == nil {
		api.ApplicationApplicationSnapshotDeleteHandler = application.ApplicationSnapshotDeleteHandlerFunc(func(params application.ApplicationSnapshotDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationSnapshotDelete has not yet been implemented")
		})
	}
	if api.ApplicationApplicationSnapshotGetHandler == nil {
		api.ApplicationApplicationSnapshotGetHandler = application.ApplicationSnapshotGetHandlerFunc(func(params application.ApplicationSnapshotGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationSnapshotGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationSnapshotRestoreHandler == nil {
		api.ApplicationApplicationSnapshotRestoreHandler = application.ApplicationSnapshotRestoreHandlerFunc(func(params application.ApplicationSnapshotRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationSnapshotRestore has not yet been implemented")
		})
	}
	if api.ApplicationApplicationTemplateCollectionGetHandler == nil {
		api.ApplicationApplicationTemplateCollectionGetHandler = application.ApplicationTemplateCollectionGetHandlerFunc(func(params application.ApplicationTemplateCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationTemplateCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationApplicationTemplateGetHandler == nil {
		api.ApplicationApplicationTemplateGetHandler = application.ApplicationTemplateGetHandlerFunc(func(params application.ApplicationTemplateGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ApplicationTemplateGet has not yet been implemented")
		})
	}
	if api.NasAuditCollectionGetHandler == nil {
		api.NasAuditCollectionGetHandler = n_a_s.AuditCollectionGetHandlerFunc(func(params n_a_s.AuditCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.AuditCollectionGet has not yet been implemented")
		})
	}
	if api.NasAuditCreateHandler == nil {
		api.NasAuditCreateHandler = n_a_s.AuditCreateHandlerFunc(func(params n_a_s.AuditCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.AuditCreate has not yet been implemented")
		})
	}
	if api.NasAuditDeleteHandler == nil {
		api.NasAuditDeleteHandler = n_a_s.AuditDeleteHandlerFunc(func(params n_a_s.AuditDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.AuditDelete has not yet been implemented")
		})
	}
	if api.NasAuditGetHandler == nil {
		api.NasAuditGetHandler = n_a_s.AuditGetHandlerFunc(func(params n_a_s.AuditGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.AuditGet has not yet been implemented")
		})
	}
	if api.SecurityAuditLogForwardingGetHandler == nil {
		api.SecurityAuditLogForwardingGetHandler = securityops.AuditLogForwardingGetHandlerFunc(func(params securityops.AuditLogForwardingGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AuditLogForwardingGet has not yet been implemented")
		})
	}
	if api.NasAuditModifyHandler == nil {
		api.NasAuditModifyHandler = n_a_s.AuditModifyHandlerFunc(func(params n_a_s.AuditModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.AuditModify has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateConfigurationCollectionGetHandler == nil {
		api.SupportAutoUpdateConfigurationCollectionGetHandler = support.AutoUpdateConfigurationCollectionGetHandlerFunc(func(params support.AutoUpdateConfigurationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateConfigurationCollectionGet has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateConfigurationGetHandler == nil {
		api.SupportAutoUpdateConfigurationGetHandler = support.AutoUpdateConfigurationGetHandlerFunc(func(params support.AutoUpdateConfigurationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateConfigurationGet has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateConfigurationModifyHandler == nil {
		api.SupportAutoUpdateConfigurationModifyHandler = support.AutoUpdateConfigurationModifyHandlerFunc(func(params support.AutoUpdateConfigurationModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateConfigurationModify has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateInfoGetHandler == nil {
		api.SupportAutoUpdateInfoGetHandler = support.AutoUpdateInfoGetHandlerFunc(func(params support.AutoUpdateInfoGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateInfoGet has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateInfoModifyHandler == nil {
		api.SupportAutoUpdateInfoModifyHandler = support.AutoUpdateInfoModifyHandlerFunc(func(params support.AutoUpdateInfoModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateInfoModify has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateStatusCollectionGetHandler == nil {
		api.SupportAutoUpdateStatusCollectionGetHandler = support.AutoUpdateStatusCollectionGetHandlerFunc(func(params support.AutoUpdateStatusCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateStatusCollectionGet has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateStatusGetHandler == nil {
		api.SupportAutoUpdateStatusGetHandler = support.AutoUpdateStatusGetHandlerFunc(func(params support.AutoUpdateStatusGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateStatusGet has not yet been implemented")
		})
	}
	if api.SupportAutoUpdateStatusModifyHandler == nil {
		api.SupportAutoUpdateStatusModifyHandler = support.AutoUpdateStatusModifyHandlerFunc(func(params support.AutoUpdateStatusModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutoUpdateStatusModify has not yet been implemented")
		})
	}
	if api.SupportAutosupportCreateHandler == nil {
		api.SupportAutosupportCreateHandler = support.AutosupportCreateHandlerFunc(func(params support.AutosupportCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutosupportCreate has not yet been implemented")
		})
	}
	if api.SupportAutosupportGetHandler == nil {
		api.SupportAutosupportGetHandler = support.AutosupportGetHandlerFunc(func(params support.AutosupportGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutosupportGet has not yet been implemented")
		})
	}
	if api.SupportAutosupportMessageCollectionGetHandler == nil {
		api.SupportAutosupportMessageCollectionGetHandler = support.AutosupportMessageCollectionGetHandlerFunc(func(params support.AutosupportMessageCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutosupportMessageCollectionGet has not yet been implemented")
		})
	}
	if api.SupportAutosupportMessageGetHandler == nil {
		api.SupportAutosupportMessageGetHandler = support.AutosupportMessageGetHandlerFunc(func(params support.AutosupportMessageGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutosupportMessageGet has not yet been implemented")
		})
	}
	if api.SupportAutosupportModifyHandler == nil {
		api.SupportAutosupportModifyHandler = support.AutosupportModifyHandlerFunc(func(params support.AutosupportModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.AutosupportModify has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsCollectionGetHandler == nil {
		api.SecurityAwsKmsCollectionGetHandler = securityops.AwsKmsCollectionGetHandlerFunc(func(params securityops.AwsKmsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsCreateHandler == nil {
		api.SecurityAwsKmsCreateHandler = securityops.AwsKmsCreateHandlerFunc(func(params securityops.AwsKmsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsCreate has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsDeleteHandler == nil {
		api.SecurityAwsKmsDeleteHandler = securityops.AwsKmsDeleteHandlerFunc(func(params securityops.AwsKmsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsDelete has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsGetHandler == nil {
		api.SecurityAwsKmsGetHandler = securityops.AwsKmsGetHandlerFunc(func(params securityops.AwsKmsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsGet has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsModifyHandler == nil {
		api.SecurityAwsKmsModifyHandler = securityops.AwsKmsModifyHandlerFunc(func(params securityops.AwsKmsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsModify has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsRekeyExternalHandler == nil {
		api.SecurityAwsKmsRekeyExternalHandler = securityops.AwsKmsRekeyExternalHandlerFunc(func(params securityops.AwsKmsRekeyExternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsRekeyExternal has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsRekeyInternalHandler == nil {
		api.SecurityAwsKmsRekeyInternalHandler = securityops.AwsKmsRekeyInternalHandlerFunc(func(params securityops.AwsKmsRekeyInternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsRekeyInternal has not yet been implemented")
		})
	}
	if api.SecurityAwsKmsRestoreHandler == nil {
		api.SecurityAwsKmsRestoreHandler = securityops.AwsKmsRestoreHandlerFunc(func(params securityops.AwsKmsRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AwsKmsRestore has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultCollectionGetHandler == nil {
		api.SecurityAzureKeyVaultCollectionGetHandler = securityops.AzureKeyVaultCollectionGetHandlerFunc(func(params securityops.AzureKeyVaultCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultCreateHandler == nil {
		api.SecurityAzureKeyVaultCreateHandler = securityops.AzureKeyVaultCreateHandlerFunc(func(params securityops.AzureKeyVaultCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultCreate has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultDeleteHandler == nil {
		api.SecurityAzureKeyVaultDeleteHandler = securityops.AzureKeyVaultDeleteHandlerFunc(func(params securityops.AzureKeyVaultDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultDelete has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultGetHandler == nil {
		api.SecurityAzureKeyVaultGetHandler = securityops.AzureKeyVaultGetHandlerFunc(func(params securityops.AzureKeyVaultGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultGet has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultModifyHandler == nil {
		api.SecurityAzureKeyVaultModifyHandler = securityops.AzureKeyVaultModifyHandlerFunc(func(params securityops.AzureKeyVaultModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultModify has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultRekeyExternalHandler == nil {
		api.SecurityAzureKeyVaultRekeyExternalHandler = securityops.AzureKeyVaultRekeyExternalHandlerFunc(func(params securityops.AzureKeyVaultRekeyExternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultRekeyExternal has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultRekeyInternalHandler == nil {
		api.SecurityAzureKeyVaultRekeyInternalHandler = securityops.AzureKeyVaultRekeyInternalHandlerFunc(func(params securityops.AzureKeyVaultRekeyInternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultRekeyInternal has not yet been implemented")
		})
	}
	if api.SecurityAzureKeyVaultRestoreHandler == nil {
		api.SecurityAzureKeyVaultRestoreHandler = securityops.AzureKeyVaultRestoreHandlerFunc(func(params securityops.AzureKeyVaultRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.AzureKeyVaultRestore has not yet been implemented")
		})
	}
	if api.ObjectStoreBucketsCollectionGetHandler == nil {
		api.ObjectStoreBucketsCollectionGetHandler = object_store.BucketsCollectionGetHandlerFunc(func(params object_store.BucketsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.BucketsCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterCapacityPoolCollectionGetHandler == nil {
		api.ClusterCapacityPoolCollectionGetHandler = cluster.CapacityPoolCollectionGetHandlerFunc(func(params cluster.CapacityPoolCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CapacityPoolCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterCapacityPoolGetHandler == nil {
		api.ClusterCapacityPoolGetHandler = cluster.CapacityPoolGetHandlerFunc(func(params cluster.CapacityPoolGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CapacityPoolGet has not yet been implemented")
		})
	}
	if api.ClusterChassisCollectionGetHandler == nil {
		api.ClusterChassisCollectionGetHandler = cluster.ChassisCollectionGetHandlerFunc(func(params cluster.ChassisCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ChassisCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterChassisGetHandler == nil {
		api.ClusterChassisGetHandler = cluster.ChassisGetHandlerFunc(func(params cluster.ChassisGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ChassisGet has not yet been implemented")
		})
	}
	if api.NasCifsCollectionPerformanceMetricsGetHandler == nil {
		api.NasCifsCollectionPerformanceMetricsGetHandler = n_a_s.CifsCollectionPerformanceMetricsGetHandlerFunc(func(params n_a_s.CifsCollectionPerformanceMetricsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsCollectionPerformanceMetricsGet has not yet been implemented")
		})
	}
	if api.NasCifsConnectionCollectionGetHandler == nil {
		api.NasCifsConnectionCollectionGetHandler = n_a_s.CifsConnectionCollectionGetHandlerFunc(func(params n_a_s.CifsConnectionCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsConnectionCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsDomainCollectionGetHandler == nil {
		api.NasCifsDomainCollectionGetHandler = n_a_s.CifsDomainCollectionGetHandlerFunc(func(params n_a_s.CifsDomainCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsDomainGetHandler == nil {
		api.NasCifsDomainGetHandler = n_a_s.CifsDomainGetHandlerFunc(func(params n_a_s.CifsDomainGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainGet has not yet been implemented")
		})
	}
	if api.NasCifsDomainModifyHandler == nil {
		api.NasCifsDomainModifyHandler = n_a_s.CifsDomainModifyHandlerFunc(func(params n_a_s.CifsDomainModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainModify has not yet been implemented")
		})
	}
	if api.NasCifsDomainPreferredDcCollectionGetHandler == nil {
		api.NasCifsDomainPreferredDcCollectionGetHandler = n_a_s.CifsDomainPreferredDcCollectionGetHandlerFunc(func(params n_a_s.CifsDomainPreferredDcCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainPreferredDcCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsDomainPreferredDcCreateHandler == nil {
		api.NasCifsDomainPreferredDcCreateHandler = n_a_s.CifsDomainPreferredDcCreateHandlerFunc(func(params n_a_s.CifsDomainPreferredDcCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainPreferredDcCreate has not yet been implemented")
		})
	}
	if api.NasCifsDomainPreferredDcDeleteHandler == nil {
		api.NasCifsDomainPreferredDcDeleteHandler = n_a_s.CifsDomainPreferredDcDeleteHandlerFunc(func(params n_a_s.CifsDomainPreferredDcDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainPreferredDcDelete has not yet been implemented")
		})
	}
	if api.NasCifsDomainPreferredDcGetHandler == nil {
		api.NasCifsDomainPreferredDcGetHandler = n_a_s.CifsDomainPreferredDcGetHandlerFunc(func(params n_a_s.CifsDomainPreferredDcGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsDomainPreferredDcGet has not yet been implemented")
		})
	}
	if api.NasCifsHomedirSearchPathGetHandler == nil {
		api.NasCifsHomedirSearchPathGetHandler = n_a_s.CifsHomedirSearchPathGetHandlerFunc(func(params n_a_s.CifsHomedirSearchPathGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsHomedirSearchPathGet has not yet been implemented")
		})
	}
	if api.NasCifsOpenFileCollectionGetHandler == nil {
		api.NasCifsOpenFileCollectionGetHandler = n_a_s.CifsOpenFileCollectionGetHandlerFunc(func(params n_a_s.CifsOpenFileCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsOpenFileCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsOpenFileDeleteHandler == nil {
		api.NasCifsOpenFileDeleteHandler = n_a_s.CifsOpenFileDeleteHandlerFunc(func(params n_a_s.CifsOpenFileDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsOpenFileDelete has not yet been implemented")
		})
	}
	if api.NasCifsOpenFileGetHandler == nil {
		api.NasCifsOpenFileGetHandler = n_a_s.CifsOpenFileGetHandlerFunc(func(params n_a_s.CifsOpenFileGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsOpenFileGet has not yet been implemented")
		})
	}
	if api.NasCifsSearchPathCollectionGetHandler == nil {
		api.NasCifsSearchPathCollectionGetHandler = n_a_s.CifsSearchPathCollectionGetHandlerFunc(func(params n_a_s.CifsSearchPathCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSearchPathCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsSearchPathCreateHandler == nil {
		api.NasCifsSearchPathCreateHandler = n_a_s.CifsSearchPathCreateHandlerFunc(func(params n_a_s.CifsSearchPathCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSearchPathCreate has not yet been implemented")
		})
	}
	if api.NasCifsSearchPathDeleteHandler == nil {
		api.NasCifsSearchPathDeleteHandler = n_a_s.CifsSearchPathDeleteHandlerFunc(func(params n_a_s.CifsSearchPathDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSearchPathDelete has not yet been implemented")
		})
	}
	if api.NasCifsSearchPathModifyHandler == nil {
		api.NasCifsSearchPathModifyHandler = n_a_s.CifsSearchPathModifyHandlerFunc(func(params n_a_s.CifsSearchPathModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSearchPathModify has not yet been implemented")
		})
	}
	if api.NasCifsServiceCollectionGetHandler == nil {
		api.NasCifsServiceCollectionGetHandler = n_a_s.CifsServiceCollectionGetHandlerFunc(func(params n_a_s.CifsServiceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsServiceCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsServiceCreateHandler == nil {
		api.NasCifsServiceCreateHandler = n_a_s.CifsServiceCreateHandlerFunc(func(params n_a_s.CifsServiceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsServiceCreate has not yet been implemented")
		})
	}
	if api.NasCifsServiceDeleteHandler == nil {
		api.NasCifsServiceDeleteHandler = n_a_s.CifsServiceDeleteHandlerFunc(func(params n_a_s.CifsServiceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsServiceDelete has not yet been implemented")
		})
	}
	if api.NasCifsServiceGetHandler == nil {
		api.NasCifsServiceGetHandler = n_a_s.CifsServiceGetHandlerFunc(func(params n_a_s.CifsServiceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsServiceGet has not yet been implemented")
		})
	}
	if api.NasCifsServiceModifyHandler == nil {
		api.NasCifsServiceModifyHandler = n_a_s.CifsServiceModifyHandlerFunc(func(params n_a_s.CifsServiceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsServiceModify has not yet been implemented")
		})
	}
	if api.NasCifsSessionCollectionGetHandler == nil {
		api.NasCifsSessionCollectionGetHandler = n_a_s.CifsSessionCollectionGetHandlerFunc(func(params n_a_s.CifsSessionCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSessionCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsSessionDeleteHandler == nil {
		api.NasCifsSessionDeleteHandler = n_a_s.CifsSessionDeleteHandlerFunc(func(params n_a_s.CifsSessionDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSessionDelete has not yet been implemented")
		})
	}
	if api.NasCifsSessionGetHandler == nil {
		api.NasCifsSessionGetHandler = n_a_s.CifsSessionGetHandlerFunc(func(params n_a_s.CifsSessionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSessionGet has not yet been implemented")
		})
	}
	if api.NasCifsShareACLCollectionGetHandler == nil {
		api.NasCifsShareACLCollectionGetHandler = n_a_s.CifsShareACLCollectionGetHandlerFunc(func(params n_a_s.CifsShareACLCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareACLCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsShareACLCreateHandler == nil {
		api.NasCifsShareACLCreateHandler = n_a_s.CifsShareACLCreateHandlerFunc(func(params n_a_s.CifsShareACLCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareACLCreate has not yet been implemented")
		})
	}
	if api.NasCifsShareACLDeleteHandler == nil {
		api.NasCifsShareACLDeleteHandler = n_a_s.CifsShareACLDeleteHandlerFunc(func(params n_a_s.CifsShareACLDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareACLDelete has not yet been implemented")
		})
	}
	if api.NasCifsShareACLGetHandler == nil {
		api.NasCifsShareACLGetHandler = n_a_s.CifsShareACLGetHandlerFunc(func(params n_a_s.CifsShareACLGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareACLGet has not yet been implemented")
		})
	}
	if api.NasCifsShareACLModifyHandler == nil {
		api.NasCifsShareACLModifyHandler = n_a_s.CifsShareACLModifyHandlerFunc(func(params n_a_s.CifsShareACLModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareACLModify has not yet been implemented")
		})
	}
	if api.NasCifsShareCollectionGetHandler == nil {
		api.NasCifsShareCollectionGetHandler = n_a_s.CifsShareCollectionGetHandlerFunc(func(params n_a_s.CifsShareCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsShareCreateHandler == nil {
		api.NasCifsShareCreateHandler = n_a_s.CifsShareCreateHandlerFunc(func(params n_a_s.CifsShareCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareCreate has not yet been implemented")
		})
	}
	if api.NasCifsShareDeleteHandler == nil {
		api.NasCifsShareDeleteHandler = n_a_s.CifsShareDeleteHandlerFunc(func(params n_a_s.CifsShareDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareDelete has not yet been implemented")
		})
	}
	if api.NasCifsShareGetHandler == nil {
		api.NasCifsShareGetHandler = n_a_s.CifsShareGetHandlerFunc(func(params n_a_s.CifsShareGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareGet has not yet been implemented")
		})
	}
	if api.NasCifsShareModifyHandler == nil {
		api.NasCifsShareModifyHandler = n_a_s.CifsShareModifyHandlerFunc(func(params n_a_s.CifsShareModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsShareModify has not yet been implemented")
		})
	}
	if api.NasCifsSymlinkMappingCollectionGetHandler == nil {
		api.NasCifsSymlinkMappingCollectionGetHandler = n_a_s.CifsSymlinkMappingCollectionGetHandlerFunc(func(params n_a_s.CifsSymlinkMappingCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSymlinkMappingCollectionGet has not yet been implemented")
		})
	}
	if api.NasCifsSymlinkMappingCreateHandler == nil {
		api.NasCifsSymlinkMappingCreateHandler = n_a_s.CifsSymlinkMappingCreateHandlerFunc(func(params n_a_s.CifsSymlinkMappingCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSymlinkMappingCreate has not yet been implemented")
		})
	}
	if api.NasCifsSymlinkMappingDeleteHandler == nil {
		api.NasCifsSymlinkMappingDeleteHandler = n_a_s.CifsSymlinkMappingDeleteHandlerFunc(func(params n_a_s.CifsSymlinkMappingDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSymlinkMappingDelete has not yet been implemented")
		})
	}
	if api.NasCifsSymlinkMappingGetHandler == nil {
		api.NasCifsSymlinkMappingGetHandler = n_a_s.CifsSymlinkMappingGetHandlerFunc(func(params n_a_s.CifsSymlinkMappingGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSymlinkMappingGet has not yet been implemented")
		})
	}
	if api.NasCifsSymlinkMappingModifyHandler == nil {
		api.NasCifsSymlinkMappingModifyHandler = n_a_s.CifsSymlinkMappingModifyHandlerFunc(func(params n_a_s.CifsSymlinkMappingModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.CifsSymlinkMappingModify has not yet been implemented")
		})
	}
	if api.NasClientLockCollectionGetHandler == nil {
		api.NasClientLockCollectionGetHandler = n_a_s.ClientLockCollectionGetHandlerFunc(func(params n_a_s.ClientLockCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ClientLockCollectionGet has not yet been implemented")
		})
	}
	if api.NasClientLockDeleteHandler == nil {
		api.NasClientLockDeleteHandler = n_a_s.ClientLockDeleteHandlerFunc(func(params n_a_s.ClientLockDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ClientLockDelete has not yet been implemented")
		})
	}
	if api.NasClientLockGetHandler == nil {
		api.NasClientLockGetHandler = n_a_s.ClientLockGetHandlerFunc(func(params n_a_s.ClientLockGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ClientLockGet has not yet been implemented")
		})
	}
	if api.StorageCloudStoreCollectionGetHandler == nil {
		api.StorageCloudStoreCollectionGetHandler = storage.CloudStoreCollectionGetHandlerFunc(func(params storage.CloudStoreCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.CloudStoreCollectionGet has not yet been implemented")
		})
	}
	if api.StorageCloudStoreCreateHandler == nil {
		api.StorageCloudStoreCreateHandler = storage.CloudStoreCreateHandlerFunc(func(params storage.CloudStoreCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.CloudStoreCreate has not yet been implemented")
		})
	}
	if api.StorageCloudStoreDeleteHandler == nil {
		api.StorageCloudStoreDeleteHandler = storage.CloudStoreDeleteHandlerFunc(func(params storage.CloudStoreDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.CloudStoreDelete has not yet been implemented")
		})
	}
	if api.StorageCloudStoreGetHandler == nil {
		api.StorageCloudStoreGetHandler = storage.CloudStoreGetHandlerFunc(func(params storage.CloudStoreGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.CloudStoreGet has not yet been implemented")
		})
	}
	if api.StorageCloudStoreModifyHandler == nil {
		api.StorageCloudStoreModifyHandler = storage.CloudStoreModifyHandlerFunc(func(params storage.CloudStoreModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.CloudStoreModify has not yet been implemented")
		})
	}
	if api.CloudCloudTargetCollectionGetHandler == nil {
		api.CloudCloudTargetCollectionGetHandler = cloud.CloudTargetCollectionGetHandlerFunc(func(params cloud.CloudTargetCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cloud.CloudTargetCollectionGet has not yet been implemented")
		})
	}
	if api.CloudCloudTargetCreateHandler == nil {
		api.CloudCloudTargetCreateHandler = cloud.CloudTargetCreateHandlerFunc(func(params cloud.CloudTargetCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cloud.CloudTargetCreate has not yet been implemented")
		})
	}
	if api.CloudCloudTargetDeleteHandler == nil {
		api.CloudCloudTargetDeleteHandler = cloud.CloudTargetDeleteHandlerFunc(func(params cloud.CloudTargetDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cloud.CloudTargetDelete has not yet been implemented")
		})
	}
	if api.CloudCloudTargetGetHandler == nil {
		api.CloudCloudTargetGetHandler = cloud.CloudTargetGetHandlerFunc(func(params cloud.CloudTargetGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cloud.CloudTargetGet has not yet been implemented")
		})
	}
	if api.CloudCloudTargetModifyHandler == nil {
		api.CloudCloudTargetModifyHandler = cloud.CloudTargetModifyHandlerFunc(func(params cloud.CloudTargetModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cloud.CloudTargetModify has not yet been implemented")
		})
	}
	if api.SecurityClusterAccountAdProxyCreateHandler == nil {
		api.SecurityClusterAccountAdProxyCreateHandler = securityops.ClusterAccountAdProxyCreateHandlerFunc(func(params securityops.ClusterAccountAdProxyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterAccountAdProxyCreate has not yet been implemented")
		})
	}
	if api.SecurityClusterAccountAdProxyDeleteHandler == nil {
		api.SecurityClusterAccountAdProxyDeleteHandler = securityops.ClusterAccountAdProxyDeleteHandlerFunc(func(params securityops.ClusterAccountAdProxyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterAccountAdProxyDelete has not yet been implemented")
		})
	}
	if api.SecurityClusterAccountAdProxyGetHandler == nil {
		api.SecurityClusterAccountAdProxyGetHandler = securityops.ClusterAccountAdProxyGetHandlerFunc(func(params securityops.ClusterAccountAdProxyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterAccountAdProxyGet has not yet been implemented")
		})
	}
	if api.SecurityClusterAccountAdProxyModifyHandler == nil {
		api.SecurityClusterAccountAdProxyModifyHandler = securityops.ClusterAccountAdProxyModifyHandlerFunc(func(params securityops.ClusterAccountAdProxyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterAccountAdProxyModify has not yet been implemented")
		})
	}
	if api.ClusterClusterCollectionPerformanceMetricsGetHandler == nil {
		api.ClusterClusterCollectionPerformanceMetricsGetHandler = cluster.ClusterCollectionPerformanceMetricsGetHandlerFunc(func(params cluster.ClusterCollectionPerformanceMetricsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterCollectionPerformanceMetricsGet has not yet been implemented")
		})
	}
	if api.ClusterClusterCreateHandler == nil {
		api.ClusterClusterCreateHandler = cluster.ClusterCreateHandlerFunc(func(params cluster.ClusterCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterCreate has not yet been implemented")
		})
	}
	if api.ClusterClusterGetHandler == nil {
		api.ClusterClusterGetHandler = cluster.ClusterGetHandlerFunc(func(params cluster.ClusterGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterGet has not yet been implemented")
		})
	}
	if api.SecurityClusterLdapCreateHandler == nil {
		api.SecurityClusterLdapCreateHandler = securityops.ClusterLdapCreateHandlerFunc(func(params securityops.ClusterLdapCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterLdapCreate has not yet been implemented")
		})
	}
	if api.SecurityClusterLdapDeleteHandler == nil {
		api.SecurityClusterLdapDeleteHandler = securityops.ClusterLdapDeleteHandlerFunc(func(params securityops.ClusterLdapDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterLdapDelete has not yet been implemented")
		})
	}
	if api.SecurityClusterLdapGetHandler == nil {
		api.SecurityClusterLdapGetHandler = securityops.ClusterLdapGetHandlerFunc(func(params securityops.ClusterLdapGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterLdapGet has not yet been implemented")
		})
	}
	if api.SecurityClusterLdapModifyHandler == nil {
		api.SecurityClusterLdapModifyHandler = securityops.ClusterLdapModifyHandlerFunc(func(params securityops.ClusterLdapModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterLdapModify has not yet been implemented")
		})
	}
	if api.ClusterClusterModifyHandler == nil {
		api.ClusterClusterModifyHandler = cluster.ClusterModifyHandlerFunc(func(params cluster.ClusterModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterModify has not yet been implemented")
		})
	}
	if api.NdmpClusterNdmpGetHandler == nil {
		api.NdmpClusterNdmpGetHandler = n_d_m_p.ClusterNdmpGetHandlerFunc(func(params n_d_m_p.ClusterNdmpGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.ClusterNdmpGet has not yet been implemented")
		})
	}
	if api.NdmpClusterNdmpModifyHandler == nil {
		api.NdmpClusterNdmpModifyHandler = n_d_m_p.ClusterNdmpModifyHandlerFunc(func(params n_d_m_p.ClusterNdmpModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.ClusterNdmpModify has not yet been implemented")
		})
	}
	if api.SecurityClusterNisCreateHandler == nil {
		api.SecurityClusterNisCreateHandler = securityops.ClusterNisCreateHandlerFunc(func(params securityops.ClusterNisCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterNisCreate has not yet been implemented")
		})
	}
	if api.SecurityClusterNisDeleteHandler == nil {
		api.SecurityClusterNisDeleteHandler = securityops.ClusterNisDeleteHandlerFunc(func(params securityops.ClusterNisDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterNisDelete has not yet been implemented")
		})
	}
	if api.SecurityClusterNisGetHandler == nil {
		api.SecurityClusterNisGetHandler = securityops.ClusterNisGetHandlerFunc(func(params securityops.ClusterNisGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterNisGet has not yet been implemented")
		})
	}
	if api.SecurityClusterNisModifyHandler == nil {
		api.SecurityClusterNisModifyHandler = securityops.ClusterNisModifyHandlerFunc(func(params securityops.ClusterNisModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.ClusterNisModify has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpKeysCollectionGetHandler == nil {
		api.ClusterClusterNtpKeysCollectionGetHandler = cluster.ClusterNtpKeysCollectionGetHandlerFunc(func(params cluster.ClusterNtpKeysCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpKeysCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpKeysCreateHandler == nil {
		api.ClusterClusterNtpKeysCreateHandler = cluster.ClusterNtpKeysCreateHandlerFunc(func(params cluster.ClusterNtpKeysCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpKeysCreate has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpKeysDeleteHandler == nil {
		api.ClusterClusterNtpKeysDeleteHandler = cluster.ClusterNtpKeysDeleteHandlerFunc(func(params cluster.ClusterNtpKeysDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpKeysDelete has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpKeysGetHandler == nil {
		api.ClusterClusterNtpKeysGetHandler = cluster.ClusterNtpKeysGetHandlerFunc(func(params cluster.ClusterNtpKeysGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpKeysGet has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpKeysModifyHandler == nil {
		api.ClusterClusterNtpKeysModifyHandler = cluster.ClusterNtpKeysModifyHandlerFunc(func(params cluster.ClusterNtpKeysModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpKeysModify has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpServersCollectionGetHandler == nil {
		api.ClusterClusterNtpServersCollectionGetHandler = cluster.ClusterNtpServersCollectionGetHandlerFunc(func(params cluster.ClusterNtpServersCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpServersCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpServersCreateHandler == nil {
		api.ClusterClusterNtpServersCreateHandler = cluster.ClusterNtpServersCreateHandlerFunc(func(params cluster.ClusterNtpServersCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpServersCreate has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpServersDeleteHandler == nil {
		api.ClusterClusterNtpServersDeleteHandler = cluster.ClusterNtpServersDeleteHandlerFunc(func(params cluster.ClusterNtpServersDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpServersDelete has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpServersGetHandler == nil {
		api.ClusterClusterNtpServersGetHandler = cluster.ClusterNtpServersGetHandlerFunc(func(params cluster.ClusterNtpServersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpServersGet has not yet been implemented")
		})
	}
	if api.ClusterClusterNtpServersModifyHandler == nil {
		api.ClusterClusterNtpServersModifyHandler = cluster.ClusterNtpServersModifyHandlerFunc(func(params cluster.ClusterNtpServersModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterNtpServersModify has not yet been implemented")
		})
	}
	if api.ClusterClusterPeerCollectionGetHandler == nil {
		api.ClusterClusterPeerCollectionGetHandler = cluster.ClusterPeerCollectionGetHandlerFunc(func(params cluster.ClusterPeerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterPeerCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterClusterPeerCreateHandler == nil {
		api.ClusterClusterPeerCreateHandler = cluster.ClusterPeerCreateHandlerFunc(func(params cluster.ClusterPeerCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterPeerCreate has not yet been implemented")
		})
	}
	if api.ClusterClusterPeerDeleteHandler == nil {
		api.ClusterClusterPeerDeleteHandler = cluster.ClusterPeerDeleteHandlerFunc(func(params cluster.ClusterPeerDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterPeerDelete has not yet been implemented")
		})
	}
	if api.ClusterClusterPeerGetHandler == nil {
		api.ClusterClusterPeerGetHandler = cluster.ClusterPeerGetHandlerFunc(func(params cluster.ClusterPeerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterPeerGet has not yet been implemented")
		})
	}
	if api.ClusterClusterPeerModifyHandler == nil {
		api.ClusterClusterPeerModifyHandler = cluster.ClusterPeerModifyHandlerFunc(func(params cluster.ClusterPeerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ClusterPeerModify has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupFileCollectionGetHandler == nil {
		api.SupportConfigurationBackupFileCollectionGetHandler = support.ConfigurationBackupFileCollectionGetHandlerFunc(func(params support.ConfigurationBackupFileCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupFileCollectionGet has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupFileCreateHandler == nil {
		api.SupportConfigurationBackupFileCreateHandler = support.ConfigurationBackupFileCreateHandlerFunc(func(params support.ConfigurationBackupFileCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupFileCreate has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupFileDeleteHandler == nil {
		api.SupportConfigurationBackupFileDeleteHandler = support.ConfigurationBackupFileDeleteHandlerFunc(func(params support.ConfigurationBackupFileDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupFileDelete has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupFileGetHandler == nil {
		api.SupportConfigurationBackupFileGetHandler = support.ConfigurationBackupFileGetHandlerFunc(func(params support.ConfigurationBackupFileGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupFileGet has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupGetHandler == nil {
		api.SupportConfigurationBackupGetHandler = support.ConfigurationBackupGetHandlerFunc(func(params support.ConfigurationBackupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupGet has not yet been implemented")
		})
	}
	if api.SupportConfigurationBackupModifyHandler == nil {
		api.SupportConfigurationBackupModifyHandler = support.ConfigurationBackupModifyHandlerFunc(func(params support.ConfigurationBackupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.ConfigurationBackupModify has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupCollectionGetHandler == nil {
		api.ApplicationConsistencyGroupCollectionGetHandler = application.ConsistencyGroupCollectionGetHandlerFunc(func(params application.ConsistencyGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupCreateHandler == nil {
		api.ApplicationConsistencyGroupCreateHandler = application.ConsistencyGroupCreateHandlerFunc(func(params application.ConsistencyGroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupCreate has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupDeleteHandler == nil {
		api.ApplicationConsistencyGroupDeleteHandler = application.ConsistencyGroupDeleteHandlerFunc(func(params application.ConsistencyGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupDelete has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupGetHandler == nil {
		api.ApplicationConsistencyGroupGetHandler = application.ConsistencyGroupGetHandlerFunc(func(params application.ConsistencyGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupGet has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupMetricsCollectionGetHandler == nil {
		api.ApplicationConsistencyGroupMetricsCollectionGetHandler = application.ConsistencyGroupMetricsCollectionGetHandlerFunc(func(params application.ConsistencyGroupMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupModifyHandler == nil {
		api.ApplicationConsistencyGroupModifyHandler = application.ConsistencyGroupModifyHandlerFunc(func(params application.ConsistencyGroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupModify has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupSnapshotCollectionGetHandler == nil {
		api.ApplicationConsistencyGroupSnapshotCollectionGetHandler = application.ConsistencyGroupSnapshotCollectionGetHandlerFunc(func(params application.ConsistencyGroupSnapshotCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupSnapshotCollectionGet has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupSnapshotCreateHandler == nil {
		api.ApplicationConsistencyGroupSnapshotCreateHandler = application.ConsistencyGroupSnapshotCreateHandlerFunc(func(params application.ConsistencyGroupSnapshotCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupSnapshotCreate has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupSnapshotDeleteHandler == nil {
		api.ApplicationConsistencyGroupSnapshotDeleteHandler = application.ConsistencyGroupSnapshotDeleteHandlerFunc(func(params application.ConsistencyGroupSnapshotDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupSnapshotDelete has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupSnapshotGetHandler == nil {
		api.ApplicationConsistencyGroupSnapshotGetHandler = application.ConsistencyGroupSnapshotGetHandlerFunc(func(params application.ConsistencyGroupSnapshotGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupSnapshotGet has not yet been implemented")
		})
	}
	if api.ApplicationConsistencyGroupSnapshotModifyHandler == nil {
		api.ApplicationConsistencyGroupSnapshotModifyHandler = application.ConsistencyGroupSnapshotModifyHandlerFunc(func(params application.ConsistencyGroupSnapshotModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation application.ConsistencyGroupSnapshotModify has not yet been implemented")
		})
	}
	if api.SupportCoredumpCollectionGetHandler == nil {
		api.SupportCoredumpCollectionGetHandler = support.CoredumpCollectionGetHandlerFunc(func(params support.CoredumpCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.CoredumpCollectionGet has not yet been implemented")
		})
	}
	if api.SupportCoredumpDeleteHandler == nil {
		api.SupportCoredumpDeleteHandler = support.CoredumpDeleteHandlerFunc(func(params support.CoredumpDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.CoredumpDelete has not yet been implemented")
		})
	}
	if api.SupportCoredumpGetHandler == nil {
		api.SupportCoredumpGetHandler = support.CoredumpGetHandlerFunc(func(params support.CoredumpGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.CoredumpGet has not yet been implemented")
		})
	}
	if api.ClusterCounterRowCollectionGetHandler == nil {
		api.ClusterCounterRowCollectionGetHandler = cluster.CounterRowCollectionGetHandlerFunc(func(params cluster.CounterRowCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CounterRowCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterCounterRowGetHandler == nil {
		api.ClusterCounterRowGetHandler = cluster.CounterRowGetHandlerFunc(func(params cluster.CounterRowGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CounterRowGet has not yet been implemented")
		})
	}
	if api.ClusterCounterTableCollectionGetHandler == nil {
		api.ClusterCounterTableCollectionGetHandler = cluster.CounterTableCollectionGetHandlerFunc(func(params cluster.CounterTableCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CounterTableCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterCounterTableGetHandler == nil {
		api.ClusterCounterTableGetHandler = cluster.CounterTableGetHandlerFunc(func(params cluster.CounterTableGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.CounterTableGet has not yet been implemented")
		})
	}
	if api.SecurityCreateCertificateSigningRequestHandler == nil {
		api.SecurityCreateCertificateSigningRequestHandler = securityops.CreateCertificateSigningRequestHandlerFunc(func(params securityops.CreateCertificateSigningRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.CreateCertificateSigningRequest has not yet been implemented")
		})
	}
	if api.StorageDiskCollectionGetHandler == nil {
		api.StorageDiskCollectionGetHandler = storage.DiskCollectionGetHandlerFunc(func(params storage.DiskCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.DiskCollectionGet has not yet been implemented")
		})
	}
	if api.StorageDiskGetHandler == nil {
		api.StorageDiskGetHandler = storage.DiskGetHandlerFunc(func(params storage.DiskGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.DiskGet has not yet been implemented")
		})
	}
	if api.StorageDiskModifyHandler == nil {
		api.StorageDiskModifyHandler = storage.DiskModifyHandlerFunc(func(params storage.DiskModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.DiskModify has not yet been implemented")
		})
	}
	if api.NameServicesDNSCollectionGetHandler == nil {
		api.NameServicesDNSCollectionGetHandler = name_services.DNSCollectionGetHandlerFunc(func(params name_services.DNSCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.DNSCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesDNSCreateHandler == nil {
		api.NameServicesDNSCreateHandler = name_services.DNSCreateHandlerFunc(func(params name_services.DNSCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.DNSCreate has not yet been implemented")
		})
	}
	if api.NameServicesDNSDeleteHandler == nil {
		api.NameServicesDNSDeleteHandler = name_services.DNSDeleteHandlerFunc(func(params name_services.DNSDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.DNSDelete has not yet been implemented")
		})
	}
	if api.NameServicesDNSGetHandler == nil {
		api.NameServicesDNSGetHandler = name_services.DNSGetHandlerFunc(func(params name_services.DNSGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.DNSGet has not yet been implemented")
		})
	}
	if api.NameServicesDNSModifyHandler == nil {
		api.NameServicesDNSModifyHandler = name_services.DNSModifyHandlerFunc(func(params name_services.DNSModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.DNSModify has not yet been implemented")
		})
	}
	if api.SecurityDuoCollectionGetHandler == nil {
		api.SecurityDuoCollectionGetHandler = securityops.DuoCollectionGetHandlerFunc(func(params securityops.DuoCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuoCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityDuoCreateHandler == nil {
		api.SecurityDuoCreateHandler = securityops.DuoCreateHandlerFunc(func(params securityops.DuoCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuoCreate has not yet been implemented")
		})
	}
	if api.SecurityDuoGetHandler == nil {
		api.SecurityDuoGetHandler = securityops.DuoGetHandlerFunc(func(params securityops.DuoGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuoGet has not yet been implemented")
		})
	}
	if api.SecurityDuoModifyHandler == nil {
		api.SecurityDuoModifyHandler = securityops.DuoModifyHandlerFunc(func(params securityops.DuoModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuoModify has not yet been implemented")
		})
	}
	if api.SecurityDuogroupCollectionGetHandler == nil {
		api.SecurityDuogroupCollectionGetHandler = securityops.DuogroupCollectionGetHandlerFunc(func(params securityops.DuogroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuogroupCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityDuogroupCreateHandler == nil {
		api.SecurityDuogroupCreateHandler = securityops.DuogroupCreateHandlerFunc(func(params securityops.DuogroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuogroupCreate has not yet been implemented")
		})
	}
	if api.SecurityDuogroupGetHandler == nil {
		api.SecurityDuogroupGetHandler = securityops.DuogroupGetHandlerFunc(func(params securityops.DuogroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuogroupGet has not yet been implemented")
		})
	}
	if api.SecurityDuogroupModifyHandler == nil {
		api.SecurityDuogroupModifyHandler = securityops.DuogroupModifyHandlerFunc(func(params securityops.DuogroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.DuogroupModify has not yet been implemented")
		})
	}
	if api.NasEffectivePermissionGetHandler == nil {
		api.NasEffectivePermissionGetHandler = n_a_s.EffectivePermissionGetHandlerFunc(func(params n_a_s.EffectivePermissionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.EffectivePermissionGet has not yet been implemented")
		})
	}
	if api.SupportEmsApplicationLogsCreateHandler == nil {
		api.SupportEmsApplicationLogsCreateHandler = support.EmsApplicationLogsCreateHandlerFunc(func(params support.EmsApplicationLogsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsApplicationLogsCreate has not yet been implemented")
		})
	}
	if api.SupportEmsConfigGetHandler == nil {
		api.SupportEmsConfigGetHandler = support.EmsConfigGetHandlerFunc(func(params support.EmsConfigGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsConfigGet has not yet been implemented")
		})
	}
	if api.SupportEmsConfigModifyHandler == nil {
		api.SupportEmsConfigModifyHandler = support.EmsConfigModifyHandlerFunc(func(params support.EmsConfigModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsConfigModify has not yet been implemented")
		})
	}
	if api.SupportEmsDestinationCollectionGetHandler == nil {
		api.SupportEmsDestinationCollectionGetHandler = support.EmsDestinationCollectionGetHandlerFunc(func(params support.EmsDestinationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsDestinationCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsDestinationCreateHandler == nil {
		api.SupportEmsDestinationCreateHandler = support.EmsDestinationCreateHandlerFunc(func(params support.EmsDestinationCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsDestinationCreate has not yet been implemented")
		})
	}
	if api.SupportEmsDestinationDeleteHandler == nil {
		api.SupportEmsDestinationDeleteHandler = support.EmsDestinationDeleteHandlerFunc(func(params support.EmsDestinationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsDestinationDelete has not yet been implemented")
		})
	}
	if api.SupportEmsDestinationGetHandler == nil {
		api.SupportEmsDestinationGetHandler = support.EmsDestinationGetHandlerFunc(func(params support.EmsDestinationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsDestinationGet has not yet been implemented")
		})
	}
	if api.SupportEmsDestinationModifyHandler == nil {
		api.SupportEmsDestinationModifyHandler = support.EmsDestinationModifyHandlerFunc(func(params support.EmsDestinationModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsDestinationModify has not yet been implemented")
		})
	}
	if api.SupportEmsEventCollectionGetHandler == nil {
		api.SupportEmsEventCollectionGetHandler = support.EmsEventCollectionGetHandlerFunc(func(params support.EmsEventCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsEventCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsFilterCollectionGetHandler == nil {
		api.SupportEmsFilterCollectionGetHandler = support.EmsFilterCollectionGetHandlerFunc(func(params support.EmsFilterCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsFilterCreateHandler == nil {
		api.SupportEmsFilterCreateHandler = support.EmsFilterCreateHandlerFunc(func(params support.EmsFilterCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterCreate has not yet been implemented")
		})
	}
	if api.SupportEmsFilterDeleteHandler == nil {
		api.SupportEmsFilterDeleteHandler = support.EmsFilterDeleteHandlerFunc(func(params support.EmsFilterDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterDelete has not yet been implemented")
		})
	}
	if api.SupportEmsFilterGetHandler == nil {
		api.SupportEmsFilterGetHandler = support.EmsFilterGetHandlerFunc(func(params support.EmsFilterGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterGet has not yet been implemented")
		})
	}
	if api.SupportEmsFilterModifyHandler == nil {
		api.SupportEmsFilterModifyHandler = support.EmsFilterModifyHandlerFunc(func(params support.EmsFilterModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterModify has not yet been implemented")
		})
	}
	if api.SupportEmsFilterRuleCollectionGetHandler == nil {
		api.SupportEmsFilterRuleCollectionGetHandler = support.EmsFilterRuleCollectionGetHandlerFunc(func(params support.EmsFilterRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterRuleCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsFilterRuleDeleteHandler == nil {
		api.SupportEmsFilterRuleDeleteHandler = support.EmsFilterRuleDeleteHandlerFunc(func(params support.EmsFilterRuleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterRuleDelete has not yet been implemented")
		})
	}
	if api.SupportEmsFilterRuleGetHandler == nil {
		api.SupportEmsFilterRuleGetHandler = support.EmsFilterRuleGetHandlerFunc(func(params support.EmsFilterRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterRuleGet has not yet been implemented")
		})
	}
	if api.SupportEmsFilterRuleModifyHandler == nil {
		api.SupportEmsFilterRuleModifyHandler = support.EmsFilterRuleModifyHandlerFunc(func(params support.EmsFilterRuleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFilterRuleModify has not yet been implemented")
		})
	}
	if api.SupportEmsFiltersRulesCreateHandler == nil {
		api.SupportEmsFiltersRulesCreateHandler = support.EmsFiltersRulesCreateHandlerFunc(func(params support.EmsFiltersRulesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsFiltersRulesCreate has not yet been implemented")
		})
	}
	if api.SupportEmsMessageCollectionGetHandler == nil {
		api.SupportEmsMessageCollectionGetHandler = support.EmsMessageCollectionGetHandlerFunc(func(params support.EmsMessageCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsMessageCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsRoleConfigCollectionGetHandler == nil {
		api.SupportEmsRoleConfigCollectionGetHandler = support.EmsRoleConfigCollectionGetHandlerFunc(func(params support.EmsRoleConfigCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsRoleConfigCollectionGet has not yet been implemented")
		})
	}
	if api.SupportEmsRoleConfigCreateHandler == nil {
		api.SupportEmsRoleConfigCreateHandler = support.EmsRoleConfigCreateHandlerFunc(func(params support.EmsRoleConfigCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsRoleConfigCreate has not yet been implemented")
		})
	}
	if api.SupportEmsRoleConfigDeleteHandler == nil {
		api.SupportEmsRoleConfigDeleteHandler = support.EmsRoleConfigDeleteHandlerFunc(func(params support.EmsRoleConfigDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsRoleConfigDelete has not yet been implemented")
		})
	}
	if api.SupportEmsRoleConfigGetHandler == nil {
		api.SupportEmsRoleConfigGetHandler = support.EmsRoleConfigGetHandlerFunc(func(params support.EmsRoleConfigGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsRoleConfigGet has not yet been implemented")
		})
	}
	if api.SupportEmsRoleConfigModifyHandler == nil {
		api.SupportEmsRoleConfigModifyHandler = support.EmsRoleConfigModifyHandlerFunc(func(params support.EmsRoleConfigModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.EmsRoleConfigModify has not yet been implemented")
		})
	}
	if api.NasExportPolicyCollectionGetHandler == nil {
		api.NasExportPolicyCollectionGetHandler = n_a_s.ExportPolicyCollectionGetHandlerFunc(func(params n_a_s.ExportPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasExportPolicyCreateHandler == nil {
		api.NasExportPolicyCreateHandler = n_a_s.ExportPolicyCreateHandlerFunc(func(params n_a_s.ExportPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportPolicyCreate has not yet been implemented")
		})
	}
	if api.NasExportPolicyDeleteHandler == nil {
		api.NasExportPolicyDeleteHandler = n_a_s.ExportPolicyDeleteHandlerFunc(func(params n_a_s.ExportPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportPolicyDelete has not yet been implemented")
		})
	}
	if api.NasExportPolicyGetHandler == nil {
		api.NasExportPolicyGetHandler = n_a_s.ExportPolicyGetHandlerFunc(func(params n_a_s.ExportPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportPolicyGet has not yet been implemented")
		})
	}
	if api.NasExportPolicyModifyHandler == nil {
		api.NasExportPolicyModifyHandler = n_a_s.ExportPolicyModifyHandlerFunc(func(params n_a_s.ExportPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportPolicyModify has not yet been implemented")
		})
	}
	if api.NasExportRuleClientsCreateHandler == nil {
		api.NasExportRuleClientsCreateHandler = n_a_s.ExportRuleClientsCreateHandlerFunc(func(params n_a_s.ExportRuleClientsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleClientsCreate has not yet been implemented")
		})
	}
	if api.NasExportRuleClientsDeleteHandler == nil {
		api.NasExportRuleClientsDeleteHandler = n_a_s.ExportRuleClientsDeleteHandlerFunc(func(params n_a_s.ExportRuleClientsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleClientsDelete has not yet been implemented")
		})
	}
	if api.NasExportRuleClientsGetHandler == nil {
		api.NasExportRuleClientsGetHandler = n_a_s.ExportRuleClientsGetHandlerFunc(func(params n_a_s.ExportRuleClientsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleClientsGet has not yet been implemented")
		})
	}
	if api.NasExportRuleCollectionGetHandler == nil {
		api.NasExportRuleCollectionGetHandler = n_a_s.ExportRuleCollectionGetHandlerFunc(func(params n_a_s.ExportRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleCollectionGet has not yet been implemented")
		})
	}
	if api.NasExportRuleCreateHandler == nil {
		api.NasExportRuleCreateHandler = n_a_s.ExportRuleCreateHandlerFunc(func(params n_a_s.ExportRuleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleCreate has not yet been implemented")
		})
	}
	if api.NasExportRuleDeleteHandler == nil {
		api.NasExportRuleDeleteHandler = n_a_s.ExportRuleDeleteHandlerFunc(func(params n_a_s.ExportRuleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleDelete has not yet been implemented")
		})
	}
	if api.NasExportRuleGetHandler == nil {
		api.NasExportRuleGetHandler = n_a_s.ExportRuleGetHandlerFunc(func(params n_a_s.ExportRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleGet has not yet been implemented")
		})
	}
	if api.NasExportRuleModifyHandler == nil {
		api.NasExportRuleModifyHandler = n_a_s.ExportRuleModifyHandlerFunc(func(params n_a_s.ExportRuleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ExportRuleModify has not yet been implemented")
		})
	}
	if api.NetworkingFabricCollectionGetHandler == nil {
		api.NetworkingFabricCollectionGetHandler = networking.FabricCollectionGetHandlerFunc(func(params networking.FabricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FabricCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingFabricGetHandler == nil {
		api.NetworkingFabricGetHandler = networking.FabricGetHandlerFunc(func(params networking.FabricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FabricGet has not yet been implemented")
		})
	}
	if api.NetworkingFcInterfaceCollectionGetHandler == nil {
		api.NetworkingFcInterfaceCollectionGetHandler = networking.FcInterfaceCollectionGetHandlerFunc(func(params networking.FcInterfaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcInterfaceCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingFcInterfaceCreateHandler == nil {
		api.NetworkingFcInterfaceCreateHandler = networking.FcInterfaceCreateHandlerFunc(func(params networking.FcInterfaceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcInterfaceCreate has not yet been implemented")
		})
	}
	if api.NetworkingFcInterfaceDeleteHandler == nil {
		api.NetworkingFcInterfaceDeleteHandler = networking.FcInterfaceDeleteHandlerFunc(func(params networking.FcInterfaceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcInterfaceDelete has not yet been implemented")
		})
	}
	if api.NetworkingFcInterfaceGetHandler == nil {
		api.NetworkingFcInterfaceGetHandler = networking.FcInterfaceGetHandlerFunc(func(params networking.FcInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcInterfaceGet has not yet been implemented")
		})
	}
	if api.NetworkingFcInterfaceModifyHandler == nil {
		api.NetworkingFcInterfaceModifyHandler = networking.FcInterfaceModifyHandlerFunc(func(params networking.FcInterfaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcInterfaceModify has not yet been implemented")
		})
	}
	if api.SanFcLoginCollectionGetHandler == nil {
		api.SanFcLoginCollectionGetHandler = s_a_n.FcLoginCollectionGetHandlerFunc(func(params s_a_n.FcLoginCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcLoginCollectionGet has not yet been implemented")
		})
	}
	if api.SanFcLoginGetHandler == nil {
		api.SanFcLoginGetHandler = s_a_n.FcLoginGetHandlerFunc(func(params s_a_n.FcLoginGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcLoginGet has not yet been implemented")
		})
	}
	if api.NetworkingFcPortCollectionGetHandler == nil {
		api.NetworkingFcPortCollectionGetHandler = networking.FcPortCollectionGetHandlerFunc(func(params networking.FcPortCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcPortCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingFcPortGetHandler == nil {
		api.NetworkingFcPortGetHandler = networking.FcPortGetHandlerFunc(func(params networking.FcPortGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcPortGet has not yet been implemented")
		})
	}
	if api.NetworkingFcPortModifyHandler == nil {
		api.NetworkingFcPortModifyHandler = networking.FcPortModifyHandlerFunc(func(params networking.FcPortModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcPortModify has not yet been implemented")
		})
	}
	if api.NetworkingFcSwitchCollectionGetHandler == nil {
		api.NetworkingFcSwitchCollectionGetHandler = networking.FcSwitchCollectionGetHandlerFunc(func(params networking.FcSwitchCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcSwitchCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingFcSwitchGetHandler == nil {
		api.NetworkingFcSwitchGetHandler = networking.FcSwitchGetHandlerFunc(func(params networking.FcSwitchGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcSwitchGet has not yet been implemented")
		})
	}
	if api.NetworkingFcZoneCollectionGetHandler == nil {
		api.NetworkingFcZoneCollectionGetHandler = networking.FcZoneCollectionGetHandlerFunc(func(params networking.FcZoneCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcZoneCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingFcZoneGetHandler == nil {
		api.NetworkingFcZoneGetHandler = networking.FcZoneGetHandlerFunc(func(params networking.FcZoneGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.FcZoneGet has not yet been implemented")
		})
	}
	if api.SanFcpServiceCollectionGetHandler == nil {
		api.SanFcpServiceCollectionGetHandler = s_a_n.FcpServiceCollectionGetHandlerFunc(func(params s_a_n.FcpServiceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcpServiceCollectionGet has not yet been implemented")
		})
	}
	if api.SanFcpServiceCreateHandler == nil {
		api.SanFcpServiceCreateHandler = s_a_n.FcpServiceCreateHandlerFunc(func(params s_a_n.FcpServiceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcpServiceCreate has not yet been implemented")
		})
	}
	if api.SanFcpServiceDeleteHandler == nil {
		api.SanFcpServiceDeleteHandler = s_a_n.FcpServiceDeleteHandlerFunc(func(params s_a_n.FcpServiceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcpServiceDelete has not yet been implemented")
		})
	}
	if api.SanFcpServiceGetHandler == nil {
		api.SanFcpServiceGetHandler = s_a_n.FcpServiceGetHandlerFunc(func(params s_a_n.FcpServiceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcpServiceGet has not yet been implemented")
		})
	}
	if api.SanFcpServiceModifyHandler == nil {
		api.SanFcpServiceModifyHandler = s_a_n.FcpServiceModifyHandlerFunc(func(params s_a_n.FcpServiceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.FcpServiceModify has not yet been implemented")
		})
	}
	if api.StorageFileCloneCreateHandler == nil {
		api.StorageFileCloneCreateHandler = storage.FileCloneCreateHandlerFunc(func(params storage.FileCloneCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileCloneCreate has not yet been implemented")
		})
	}
	if api.StorageFileCopyCreateHandler == nil {
		api.StorageFileCopyCreateHandler = storage.FileCopyCreateHandlerFunc(func(params storage.FileCopyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileCopyCreate has not yet been implemented")
		})
	}
	if api.StorageFileDeleteHandler == nil {
		api.StorageFileDeleteHandler = storage.FileDeleteHandlerFunc(func(params storage.FileDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileDelete has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityACLCreateHandler == nil {
		api.NasFileDirectorySecurityACLCreateHandler = n_a_s.FileDirectorySecurityACLCreateHandlerFunc(func(params n_a_s.FileDirectorySecurityACLCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityACLCreate has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityACLDeleteHandler == nil {
		api.NasFileDirectorySecurityACLDeleteHandler = n_a_s.FileDirectorySecurityACLDeleteHandlerFunc(func(params n_a_s.FileDirectorySecurityACLDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityACLDelete has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityACLModifyHandler == nil {
		api.NasFileDirectorySecurityACLModifyHandler = n_a_s.FileDirectorySecurityACLModifyHandlerFunc(func(params n_a_s.FileDirectorySecurityACLModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityACLModify has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityCreateHandler == nil {
		api.NasFileDirectorySecurityCreateHandler = n_a_s.FileDirectorySecurityCreateHandlerFunc(func(params n_a_s.FileDirectorySecurityCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityCreate has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityDeleteHandler == nil {
		api.NasFileDirectorySecurityDeleteHandler = n_a_s.FileDirectorySecurityDeleteHandlerFunc(func(params n_a_s.FileDirectorySecurityDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityDelete has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityGetHandler == nil {
		api.NasFileDirectorySecurityGetHandler = n_a_s.FileDirectorySecurityGetHandlerFunc(func(params n_a_s.FileDirectorySecurityGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityGet has not yet been implemented")
		})
	}
	if api.NasFileDirectorySecurityModifyHandler == nil {
		api.NasFileDirectorySecurityModifyHandler = n_a_s.FileDirectorySecurityModifyHandlerFunc(func(params n_a_s.FileDirectorySecurityModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FileDirectorySecurityModify has not yet been implemented")
		})
	}
	if api.StorageFileInfoCollectionGetHandler == nil {
		api.StorageFileInfoCollectionGetHandler = storage.FileInfoCollectionGetHandlerFunc(func(params storage.FileInfoCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileInfoCollectionGet has not yet been implemented")
		})
	}
	if api.StorageFileInfoCreateHandler == nil {
		api.StorageFileInfoCreateHandler = storage.FileInfoCreateHandlerFunc(func(params storage.FileInfoCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileInfoCreate has not yet been implemented")
		})
	}
	if api.StorageFileInfoModifyHandler == nil {
		api.StorageFileInfoModifyHandler = storage.FileInfoModifyHandlerFunc(func(params storage.FileInfoModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileInfoModify has not yet been implemented")
		})
	}
	if api.StorageFileMoveCollectionGetHandler == nil {
		api.StorageFileMoveCollectionGetHandler = storage.FileMoveCollectionGetHandlerFunc(func(params storage.FileMoveCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileMoveCollectionGet has not yet been implemented")
		})
	}
	if api.StorageFileMoveCreateHandler == nil {
		api.StorageFileMoveCreateHandler = storage.FileMoveCreateHandlerFunc(func(params storage.FileMoveCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileMoveCreate has not yet been implemented")
		})
	}
	if api.StorageFileMoveGetHandler == nil {
		api.StorageFileMoveGetHandler = storage.FileMoveGetHandlerFunc(func(params storage.FileMoveGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FileMoveGet has not yet been implemented")
		})
	}
	if api.ClusterFirmwareHistoryCollectionGetHandler == nil {
		api.ClusterFirmwareHistoryCollectionGetHandler = cluster.FirmwareHistoryCollectionGetHandlerFunc(func(params cluster.FirmwareHistoryCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.FirmwareHistoryCollectionGet has not yet been implemented")
		})
	}
	if api.StorageFlexcacheCollectionGetHandler == nil {
		api.StorageFlexcacheCollectionGetHandler = storage.FlexcacheCollectionGetHandlerFunc(func(params storage.FlexcacheCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheCollectionGet has not yet been implemented")
		})
	}
	if api.StorageFlexcacheCreateHandler == nil {
		api.StorageFlexcacheCreateHandler = storage.FlexcacheCreateHandlerFunc(func(params storage.FlexcacheCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheCreate has not yet been implemented")
		})
	}
	if api.StorageFlexcacheDeleteHandler == nil {
		api.StorageFlexcacheDeleteHandler = storage.FlexcacheDeleteHandlerFunc(func(params storage.FlexcacheDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheDelete has not yet been implemented")
		})
	}
	if api.StorageFlexcacheGetHandler == nil {
		api.StorageFlexcacheGetHandler = storage.FlexcacheGetHandlerFunc(func(params storage.FlexcacheGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheGet has not yet been implemented")
		})
	}
	if api.StorageFlexcacheModifyHandler == nil {
		api.StorageFlexcacheModifyHandler = storage.FlexcacheModifyHandlerFunc(func(params storage.FlexcacheModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheModify has not yet been implemented")
		})
	}
	if api.StorageFlexcacheOriginCollectionGetHandler == nil {
		api.StorageFlexcacheOriginCollectionGetHandler = storage.FlexcacheOriginCollectionGetHandlerFunc(func(params storage.FlexcacheOriginCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheOriginCollectionGet has not yet been implemented")
		})
	}
	if api.StorageFlexcacheOriginGetHandler == nil {
		api.StorageFlexcacheOriginGetHandler = storage.FlexcacheOriginGetHandlerFunc(func(params storage.FlexcacheOriginGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheOriginGet has not yet been implemented")
		})
	}
	if api.StorageFlexcacheOriginModifyHandler == nil {
		api.StorageFlexcacheOriginModifyHandler = storage.FlexcacheOriginModifyHandlerFunc(func(params storage.FlexcacheOriginModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.FlexcacheOriginModify has not yet been implemented")
		})
	}
	if api.NasFpolicyCollectionGetHandler == nil {
		api.NasFpolicyCollectionGetHandler = n_a_s.FpolicyCollectionGetHandlerFunc(func(params n_a_s.FpolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyConnectionCollectionGetHandler == nil {
		api.NasFpolicyConnectionCollectionGetHandler = n_a_s.FpolicyConnectionCollectionGetHandlerFunc(func(params n_a_s.FpolicyConnectionCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyConnectionCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyConnectionGetHandler == nil {
		api.NasFpolicyConnectionGetHandler = n_a_s.FpolicyConnectionGetHandlerFunc(func(params n_a_s.FpolicyConnectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyConnectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyConnectionModifyHandler == nil {
		api.NasFpolicyConnectionModifyHandler = n_a_s.FpolicyConnectionModifyHandlerFunc(func(params n_a_s.FpolicyConnectionModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyConnectionModify has not yet been implemented")
		})
	}
	if api.NasFpolicyCreateHandler == nil {
		api.NasFpolicyCreateHandler = n_a_s.FpolicyCreateHandlerFunc(func(params n_a_s.FpolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyCreate has not yet been implemented")
		})
	}
	if api.NasFpolicyDeleteHandler == nil {
		api.NasFpolicyDeleteHandler = n_a_s.FpolicyDeleteHandlerFunc(func(params n_a_s.FpolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyDelete has not yet been implemented")
		})
	}
	if api.NasFpolicyEngineCollectionGetHandler == nil {
		api.NasFpolicyEngineCollectionGetHandler = n_a_s.FpolicyEngineCollectionGetHandlerFunc(func(params n_a_s.FpolicyEngineCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEngineCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyEngineCreateHandler == nil {
		api.NasFpolicyEngineCreateHandler = n_a_s.FpolicyEngineCreateHandlerFunc(func(params n_a_s.FpolicyEngineCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEngineCreate has not yet been implemented")
		})
	}
	if api.NasFpolicyEngineDeleteHandler == nil {
		api.NasFpolicyEngineDeleteHandler = n_a_s.FpolicyEngineDeleteHandlerFunc(func(params n_a_s.FpolicyEngineDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEngineDelete has not yet been implemented")
		})
	}
	if api.NasFpolicyEngineGetHandler == nil {
		api.NasFpolicyEngineGetHandler = n_a_s.FpolicyEngineGetHandlerFunc(func(params n_a_s.FpolicyEngineGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEngineGet has not yet been implemented")
		})
	}
	if api.NasFpolicyEngineModifyHandler == nil {
		api.NasFpolicyEngineModifyHandler = n_a_s.FpolicyEngineModifyHandlerFunc(func(params n_a_s.FpolicyEngineModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEngineModify has not yet been implemented")
		})
	}
	if api.NasFpolicyEventCollectionGetHandler == nil {
		api.NasFpolicyEventCollectionGetHandler = n_a_s.FpolicyEventCollectionGetHandlerFunc(func(params n_a_s.FpolicyEventCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEventCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyEventCreateHandler == nil {
		api.NasFpolicyEventCreateHandler = n_a_s.FpolicyEventCreateHandlerFunc(func(params n_a_s.FpolicyEventCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEventCreate has not yet been implemented")
		})
	}
	if api.NasFpolicyEventDeleteHandler == nil {
		api.NasFpolicyEventDeleteHandler = n_a_s.FpolicyEventDeleteHandlerFunc(func(params n_a_s.FpolicyEventDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEventDelete has not yet been implemented")
		})
	}
	if api.NasFpolicyEventModifyHandler == nil {
		api.NasFpolicyEventModifyHandler = n_a_s.FpolicyEventModifyHandlerFunc(func(params n_a_s.FpolicyEventModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEventModify has not yet been implemented")
		})
	}
	if api.NasFpolicyEventsGetHandler == nil {
		api.NasFpolicyEventsGetHandler = n_a_s.FpolicyEventsGetHandlerFunc(func(params n_a_s.FpolicyEventsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyEventsGet has not yet been implemented")
		})
	}
	if api.NasFpolicyGetHandler == nil {
		api.NasFpolicyGetHandler = n_a_s.FpolicyGetHandlerFunc(func(params n_a_s.FpolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyGet has not yet been implemented")
		})
	}
	if api.NasFpolicyPersistentStoreCollectionGetHandler == nil {
		api.NasFpolicyPersistentStoreCollectionGetHandler = n_a_s.FpolicyPersistentStoreCollectionGetHandlerFunc(func(params n_a_s.FpolicyPersistentStoreCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPersistentStoreCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyPersistentStoreCreateHandler == nil {
		api.NasFpolicyPersistentStoreCreateHandler = n_a_s.FpolicyPersistentStoreCreateHandlerFunc(func(params n_a_s.FpolicyPersistentStoreCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPersistentStoreCreate has not yet been implemented")
		})
	}
	if api.NasFpolicyPersistentStoreDeleteHandler == nil {
		api.NasFpolicyPersistentStoreDeleteHandler = n_a_s.FpolicyPersistentStoreDeleteHandlerFunc(func(params n_a_s.FpolicyPersistentStoreDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPersistentStoreDelete has not yet been implemented")
		})
	}
	if api.NasFpolicyPersistentStoreGetHandler == nil {
		api.NasFpolicyPersistentStoreGetHandler = n_a_s.FpolicyPersistentStoreGetHandlerFunc(func(params n_a_s.FpolicyPersistentStoreGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPersistentStoreGet has not yet been implemented")
		})
	}
	if api.NasFpolicyPersistentStoreModifyHandler == nil {
		api.NasFpolicyPersistentStoreModifyHandler = n_a_s.FpolicyPersistentStoreModifyHandlerFunc(func(params n_a_s.FpolicyPersistentStoreModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPersistentStoreModify has not yet been implemented")
		})
	}
	if api.NasFpolicyPolicyCollectionGetHandler == nil {
		api.NasFpolicyPolicyCollectionGetHandler = n_a_s.FpolicyPolicyCollectionGetHandlerFunc(func(params n_a_s.FpolicyPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasFpolicyPolicyCreateHandler == nil {
		api.NasFpolicyPolicyCreateHandler = n_a_s.FpolicyPolicyCreateHandlerFunc(func(params n_a_s.FpolicyPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPolicyCreate has not yet been implemented")
		})
	}
	if api.NasFpolicyPolicyDeleteHandler == nil {
		api.NasFpolicyPolicyDeleteHandler = n_a_s.FpolicyPolicyDeleteHandlerFunc(func(params n_a_s.FpolicyPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPolicyDelete has not yet been implemented")
		})
	}
	if api.NasFpolicyPolicyGetHandler == nil {
		api.NasFpolicyPolicyGetHandler = n_a_s.FpolicyPolicyGetHandlerFunc(func(params n_a_s.FpolicyPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPolicyGet has not yet been implemented")
		})
	}
	if api.NasFpolicyPolicyModifyHandler == nil {
		api.NasFpolicyPolicyModifyHandler = n_a_s.FpolicyPolicyModifyHandlerFunc(func(params n_a_s.FpolicyPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.FpolicyPolicyModify has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsCollectionGetHandler == nil {
		api.SecurityGcpKmsCollectionGetHandler = securityops.GcpKmsCollectionGetHandlerFunc(func(params securityops.GcpKmsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsCreateHandler == nil {
		api.SecurityGcpKmsCreateHandler = securityops.GcpKmsCreateHandlerFunc(func(params securityops.GcpKmsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsCreate has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsDeleteHandler == nil {
		api.SecurityGcpKmsDeleteHandler = securityops.GcpKmsDeleteHandlerFunc(func(params securityops.GcpKmsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsDelete has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsGetHandler == nil {
		api.SecurityGcpKmsGetHandler = securityops.GcpKmsGetHandlerFunc(func(params securityops.GcpKmsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsGet has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsModifyHandler == nil {
		api.SecurityGcpKmsModifyHandler = securityops.GcpKmsModifyHandlerFunc(func(params securityops.GcpKmsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsModify has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsRekeyExternalHandler == nil {
		api.SecurityGcpKmsRekeyExternalHandler = securityops.GcpKmsRekeyExternalHandlerFunc(func(params securityops.GcpKmsRekeyExternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsRekeyExternal has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsRekeyInternalHandler == nil {
		api.SecurityGcpKmsRekeyInternalHandler = securityops.GcpKmsRekeyInternalHandlerFunc(func(params securityops.GcpKmsRekeyInternalParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsRekeyInternal has not yet been implemented")
		})
	}
	if api.SecurityGcpKmsRestoreHandler == nil {
		api.SecurityGcpKmsRestoreHandler = securityops.GcpKmsRestoreHandlerFunc(func(params securityops.GcpKmsRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.GcpKmsRestore has not yet been implemented")
		})
	}
	if api.NameServicesGlobalCacheSettingGetHandler == nil {
		api.NameServicesGlobalCacheSettingGetHandler = name_services.GlobalCacheSettingGetHandlerFunc(func(params name_services.GlobalCacheSettingGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.GlobalCacheSettingGet has not yet been implemented")
		})
	}
	if api.NameServicesGlobalCacheSettingModifyHandler == nil {
		api.NameServicesGlobalCacheSettingModifyHandler = name_services.GlobalCacheSettingModifyHandlerFunc(func(params name_services.GlobalCacheSettingModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.GlobalCacheSettingModify has not yet been implemented")
		})
	}
	if api.NameServicesGroupMembershipSettingsCollectionGetHandler == nil {
		api.NameServicesGroupMembershipSettingsCollectionGetHandler = name_services.GroupMembershipSettingsCollectionGetHandlerFunc(func(params name_services.GroupMembershipSettingsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.GroupMembershipSettingsCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesGroupMembershipSettingsGetHandler == nil {
		api.NameServicesGroupMembershipSettingsGetHandler = name_services.GroupMembershipSettingsGetHandlerFunc(func(params name_services.GroupMembershipSettingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.GroupMembershipSettingsGet has not yet been implemented")
		})
	}
	if api.NameServicesGroupMembershipSettingsModifyHandler == nil {
		api.NameServicesGroupMembershipSettingsModifyHandler = name_services.GroupMembershipSettingsModifyHandlerFunc(func(params name_services.GroupMembershipSettingsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.GroupMembershipSettingsModify has not yet been implemented")
		})
	}
	if api.NasGroupPoliciesToBeAppliedModifyHandler == nil {
		api.NasGroupPoliciesToBeAppliedModifyHandler = n_a_s.GroupPoliciesToBeAppliedModifyHandlerFunc(func(params n_a_s.GroupPoliciesToBeAppliedModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPoliciesToBeAppliedModify has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectCentralAccessPolicyCollectionGetHandler == nil {
		api.NasGroupPolicyObjectCentralAccessPolicyCollectionGetHandler = n_a_s.GroupPolicyObjectCentralAccessPolicyCollectionGetHandlerFunc(func(params n_a_s.GroupPolicyObjectCentralAccessPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectCentralAccessPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectCentralAccessPolicyGetHandler == nil {
		api.NasGroupPolicyObjectCentralAccessPolicyGetHandler = n_a_s.GroupPolicyObjectCentralAccessPolicyGetHandlerFunc(func(params n_a_s.GroupPolicyObjectCentralAccessPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectCentralAccessPolicyGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectCentralAccessRuleCollectionGetHandler == nil {
		api.NasGroupPolicyObjectCentralAccessRuleCollectionGetHandler = n_a_s.GroupPolicyObjectCentralAccessRuleCollectionGetHandlerFunc(func(params n_a_s.GroupPolicyObjectCentralAccessRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectCentralAccessRuleCollectionGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectCentralAccessRuleGetHandler == nil {
		api.NasGroupPolicyObjectCentralAccessRuleGetHandler = n_a_s.GroupPolicyObjectCentralAccessRuleGetHandlerFunc(func(params n_a_s.GroupPolicyObjectCentralAccessRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectCentralAccessRuleGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectCollectionGetHandler == nil {
		api.NasGroupPolicyObjectCollectionGetHandler = n_a_s.GroupPolicyObjectCollectionGetHandlerFunc(func(params n_a_s.GroupPolicyObjectCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectCollectionGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectGetHandler == nil {
		api.NasGroupPolicyObjectGetHandler = n_a_s.GroupPolicyObjectGetHandlerFunc(func(params n_a_s.GroupPolicyObjectGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectRestrictedGroupCollectionGetHandler == nil {
		api.NasGroupPolicyObjectRestrictedGroupCollectionGetHandler = n_a_s.GroupPolicyObjectRestrictedGroupCollectionGetHandlerFunc(func(params n_a_s.GroupPolicyObjectRestrictedGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectRestrictedGroupCollectionGet has not yet been implemented")
		})
	}
	if api.NasGroupPolicyObjectRestrictedGroupGetHandler == nil {
		api.NasGroupPolicyObjectRestrictedGroupGetHandler = n_a_s.GroupPolicyObjectRestrictedGroupGetHandlerFunc(func(params n_a_s.GroupPolicyObjectRestrictedGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.GroupPolicyObjectRestrictedGroupGet has not yet been implemented")
		})
	}
	if api.NameServicesHostRecordGetHandler == nil {
		api.NameServicesHostRecordGetHandler = name_services.HostRecordGetHandlerFunc(func(params name_services.HostRecordGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.HostRecordGet has not yet been implemented")
		})
	}
	if api.NameServicesHostsSettingsCollectionGetHandler == nil {
		api.NameServicesHostsSettingsCollectionGetHandler = name_services.HostsSettingsCollectionGetHandlerFunc(func(params name_services.HostsSettingsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.HostsSettingsCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesHostsSettingsGetHandler == nil {
		api.NameServicesHostsSettingsGetHandler = name_services.HostsSettingsGetHandlerFunc(func(params name_services.HostsSettingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.HostsSettingsGet has not yet been implemented")
		})
	}
	if api.NameServicesHostsSettingsModifyHandler == nil {
		api.NameServicesHostsSettingsModifyHandler = name_services.HostsSettingsModifyHandlerFunc(func(params name_services.HostsSettingsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.HostsSettingsModify has not yet been implemented")
		})
	}
	if api.NetworkingHTTPProxyCollectionGetHandler == nil {
		api.NetworkingHTTPProxyCollectionGetHandler = networking.HTTPProxyCollectionGetHandlerFunc(func(params networking.HTTPProxyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.HTTPProxyCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingHTTPProxyCreateHandler == nil {
		api.NetworkingHTTPProxyCreateHandler = networking.HTTPProxyCreateHandlerFunc(func(params networking.HTTPProxyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.HTTPProxyCreate has not yet been implemented")
		})
	}
	if api.NetworkingHTTPProxyDeleteHandler == nil {
		api.NetworkingHTTPProxyDeleteHandler = networking.HTTPProxyDeleteHandlerFunc(func(params networking.HTTPProxyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.HTTPProxyDelete has not yet been implemented")
		})
	}
	if api.NetworkingHTTPProxyGetHandler == nil {
		api.NetworkingHTTPProxyGetHandler = networking.HTTPProxyGetHandlerFunc(func(params networking.HTTPProxyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.HTTPProxyGet has not yet been implemented")
		})
	}
	if api.NetworkingHTTPProxyModifyHandler == nil {
		api.NetworkingHTTPProxyModifyHandler = networking.HTTPProxyModifyHandlerFunc(func(params networking.HTTPProxyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.HTTPProxyModify has not yet been implemented")
		})
	}
	if api.SanIgroupCollectionGetHandler == nil {
		api.SanIgroupCollectionGetHandler = s_a_n.IgroupCollectionGetHandlerFunc(func(params s_a_n.IgroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupCollectionGet has not yet been implemented")
		})
	}
	if api.SanIgroupCreateHandler == nil {
		api.SanIgroupCreateHandler = s_a_n.IgroupCreateHandlerFunc(func(params s_a_n.IgroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupCreate has not yet been implemented")
		})
	}
	if api.SanIgroupDeleteHandler == nil {
		api.SanIgroupDeleteHandler = s_a_n.IgroupDeleteHandlerFunc(func(params s_a_n.IgroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupDelete has not yet been implemented")
		})
	}
	if api.SanIgroupGetHandler == nil {
		api.SanIgroupGetHandler = s_a_n.IgroupGetHandlerFunc(func(params s_a_n.IgroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupGet has not yet been implemented")
		})
	}
	if api.SanIgroupInitiatorCollectionGetHandler == nil {
		api.SanIgroupInitiatorCollectionGetHandler = s_a_n.IgroupInitiatorCollectionGetHandlerFunc(func(params s_a_n.IgroupInitiatorCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupInitiatorCollectionGet has not yet been implemented")
		})
	}
	if api.SanIgroupInitiatorCreateHandler == nil {
		api.SanIgroupInitiatorCreateHandler = s_a_n.IgroupInitiatorCreateHandlerFunc(func(params s_a_n.IgroupInitiatorCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupInitiatorCreate has not yet been implemented")
		})
	}
	if api.SanIgroupInitiatorDeleteHandler == nil {
		api.SanIgroupInitiatorDeleteHandler = s_a_n.IgroupInitiatorDeleteHandlerFunc(func(params s_a_n.IgroupInitiatorDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupInitiatorDelete has not yet been implemented")
		})
	}
	if api.SanIgroupInitiatorGetHandler == nil {
		api.SanIgroupInitiatorGetHandler = s_a_n.IgroupInitiatorGetHandlerFunc(func(params s_a_n.IgroupInitiatorGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupInitiatorGet has not yet been implemented")
		})
	}
	if api.SanIgroupInitiatorModifyHandler == nil {
		api.SanIgroupInitiatorModifyHandler = s_a_n.IgroupInitiatorModifyHandlerFunc(func(params s_a_n.IgroupInitiatorModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupInitiatorModify has not yet been implemented")
		})
	}
	if api.SanIgroupModifyHandler == nil {
		api.SanIgroupModifyHandler = s_a_n.IgroupModifyHandlerFunc(func(params s_a_n.IgroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupModify has not yet been implemented")
		})
	}
	if api.SanIgroupNestedCollectionGetHandler == nil {
		api.SanIgroupNestedCollectionGetHandler = s_a_n.IgroupNestedCollectionGetHandlerFunc(func(params s_a_n.IgroupNestedCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupNestedCollectionGet has not yet been implemented")
		})
	}
	if api.SanIgroupNestedCreateHandler == nil {
		api.SanIgroupNestedCreateHandler = s_a_n.IgroupNestedCreateHandlerFunc(func(params s_a_n.IgroupNestedCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupNestedCreate has not yet been implemented")
		})
	}
	if api.SanIgroupNestedDeleteHandler == nil {
		api.SanIgroupNestedDeleteHandler = s_a_n.IgroupNestedDeleteHandlerFunc(func(params s_a_n.IgroupNestedDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupNestedDelete has not yet been implemented")
		})
	}
	if api.SanIgroupNestedGetHandler == nil {
		api.SanIgroupNestedGetHandler = s_a_n.IgroupNestedGetHandlerFunc(func(params s_a_n.IgroupNestedGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IgroupNestedGet has not yet been implemented")
		})
	}
	if api.SanInitiatorCollectionGetHandler == nil {
		api.SanInitiatorCollectionGetHandler = s_a_n.InitiatorCollectionGetHandlerFunc(func(params s_a_n.InitiatorCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.InitiatorCollectionGet has not yet been implemented")
		})
	}
	if api.SanInitiatorGetHandler == nil {
		api.SanInitiatorGetHandler = s_a_n.InitiatorGetHandlerFunc(func(params s_a_n.InitiatorGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.InitiatorGet has not yet been implemented")
		})
	}
	if api.NetworkingInterfacesMetricsCollectionGetHandler == nil {
		api.NetworkingInterfacesMetricsCollectionGetHandler = networking.InterfacesMetricsCollectionGetHandlerFunc(func(params networking.InterfacesMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.InterfacesMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingIPServicePolicyCreateHandler == nil {
		api.NetworkingIPServicePolicyCreateHandler = networking.IPServicePolicyCreateHandlerFunc(func(params networking.IPServicePolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPServicePolicyCreate has not yet been implemented")
		})
	}
	if api.NetworkingIPServicePolicyDeleteHandler == nil {
		api.NetworkingIPServicePolicyDeleteHandler = networking.IPServicePolicyDeleteHandlerFunc(func(params networking.IPServicePolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPServicePolicyDelete has not yet been implemented")
		})
	}
	if api.NetworkingIPServicePolicyModifyHandler == nil {
		api.NetworkingIPServicePolicyModifyHandler = networking.IPServicePolicyModifyHandlerFunc(func(params networking.IPServicePolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPServicePolicyModify has not yet been implemented")
		})
	}
	if api.NetworkingIPSubnetCollectionGetHandler == nil {
		api.NetworkingIPSubnetCollectionGetHandler = networking.IPSubnetCollectionGetHandlerFunc(func(params networking.IPSubnetCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPSubnetCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingIPSubnetCreateHandler == nil {
		api.NetworkingIPSubnetCreateHandler = networking.IPSubnetCreateHandlerFunc(func(params networking.IPSubnetCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPSubnetCreate has not yet been implemented")
		})
	}
	if api.NetworkingIPSubnetDeleteHandler == nil {
		api.NetworkingIPSubnetDeleteHandler = networking.IPSubnetDeleteHandlerFunc(func(params networking.IPSubnetDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPSubnetDelete has not yet been implemented")
		})
	}
	if api.NetworkingIPSubnetGetHandler == nil {
		api.NetworkingIPSubnetGetHandler = networking.IPSubnetGetHandlerFunc(func(params networking.IPSubnetGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPSubnetGet has not yet been implemented")
		})
	}
	if api.NetworkingIPSubnetModifyHandler == nil {
		api.NetworkingIPSubnetModifyHandler = networking.IPSubnetModifyHandlerFunc(func(params networking.IPSubnetModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IPSubnetModify has not yet been implemented")
		})
	}
	if api.SecurityIpsecCaCertificateCollectionGetHandler == nil {
		api.SecurityIpsecCaCertificateCollectionGetHandler = securityops.IpsecCaCertificateCollectionGetHandlerFunc(func(params securityops.IpsecCaCertificateCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecCaCertificateCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityIpsecCaCertificateCreateHandler == nil {
		api.SecurityIpsecCaCertificateCreateHandler = securityops.IpsecCaCertificateCreateHandlerFunc(func(params securityops.IpsecCaCertificateCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecCaCertificateCreate has not yet been implemented")
		})
	}
	if api.SecurityIpsecCaCertificateDeleteHandler == nil {
		api.SecurityIpsecCaCertificateDeleteHandler = securityops.IpsecCaCertificateDeleteHandlerFunc(func(params securityops.IpsecCaCertificateDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecCaCertificateDelete has not yet been implemented")
		})
	}
	if api.SecurityIpsecCaCertificateGetHandler == nil {
		api.SecurityIpsecCaCertificateGetHandler = securityops.IpsecCaCertificateGetHandlerFunc(func(params securityops.IpsecCaCertificateGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecCaCertificateGet has not yet been implemented")
		})
	}
	if api.SecurityIpsecGetHandler == nil {
		api.SecurityIpsecGetHandler = securityops.IpsecGetHandlerFunc(func(params securityops.IpsecGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecGet has not yet been implemented")
		})
	}
	if api.SecurityIpsecModifyHandler == nil {
		api.SecurityIpsecModifyHandler = securityops.IpsecModifyHandlerFunc(func(params securityops.IpsecModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecModify has not yet been implemented")
		})
	}
	if api.SecurityIpsecPolicyCollectionGetHandler == nil {
		api.SecurityIpsecPolicyCollectionGetHandler = securityops.IpsecPolicyCollectionGetHandlerFunc(func(params securityops.IpsecPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityIpsecPolicyCreateHandler == nil {
		api.SecurityIpsecPolicyCreateHandler = securityops.IpsecPolicyCreateHandlerFunc(func(params securityops.IpsecPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecPolicyCreate has not yet been implemented")
		})
	}
	if api.SecurityIpsecPolicyDeleteHandler == nil {
		api.SecurityIpsecPolicyDeleteHandler = securityops.IpsecPolicyDeleteHandlerFunc(func(params securityops.IpsecPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecPolicyDelete has not yet been implemented")
		})
	}
	if api.SecurityIpsecPolicyGetHandler == nil {
		api.SecurityIpsecPolicyGetHandler = securityops.IpsecPolicyGetHandlerFunc(func(params securityops.IpsecPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecPolicyGet has not yet been implemented")
		})
	}
	if api.SecurityIpsecPolicyModifyHandler == nil {
		api.SecurityIpsecPolicyModifyHandler = securityops.IpsecPolicyModifyHandlerFunc(func(params securityops.IpsecPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.IpsecPolicyModify has not yet been implemented")
		})
	}
	if api.NetworkingIpspaceDeleteHandler == nil {
		api.NetworkingIpspaceDeleteHandler = networking.IpspaceDeleteHandlerFunc(func(params networking.IpspaceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IpspaceDelete has not yet been implemented")
		})
	}
	if api.NetworkingIpspaceGetHandler == nil {
		api.NetworkingIpspaceGetHandler = networking.IpspaceGetHandlerFunc(func(params networking.IpspaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IpspaceGet has not yet been implemented")
		})
	}
	if api.NetworkingIpspaceModifyHandler == nil {
		api.NetworkingIpspaceModifyHandler = networking.IpspaceModifyHandlerFunc(func(params networking.IpspaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IpspaceModify has not yet been implemented")
		})
	}
	if api.NetworkingIpspacesCreateHandler == nil {
		api.NetworkingIpspacesCreateHandler = networking.IpspacesCreateHandlerFunc(func(params networking.IpspacesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IpspacesCreate has not yet been implemented")
		})
	}
	if api.NetworkingIpspacesGetHandler == nil {
		api.NetworkingIpspacesGetHandler = networking.IpspacesGetHandlerFunc(func(params networking.IpspacesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.IpspacesGet has not yet been implemented")
		})
	}
	if api.SanIscsiCredentialsCollectionGetHandler == nil {
		api.SanIscsiCredentialsCollectionGetHandler = s_a_n.IscsiCredentialsCollectionGetHandlerFunc(func(params s_a_n.IscsiCredentialsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiCredentialsCollectionGet has not yet been implemented")
		})
	}
	if api.SanIscsiCredentialsCreateHandler == nil {
		api.SanIscsiCredentialsCreateHandler = s_a_n.IscsiCredentialsCreateHandlerFunc(func(params s_a_n.IscsiCredentialsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiCredentialsCreate has not yet been implemented")
		})
	}
	if api.SanIscsiCredentialsDeleteHandler == nil {
		api.SanIscsiCredentialsDeleteHandler = s_a_n.IscsiCredentialsDeleteHandlerFunc(func(params s_a_n.IscsiCredentialsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiCredentialsDelete has not yet been implemented")
		})
	}
	if api.SanIscsiCredentialsGetHandler == nil {
		api.SanIscsiCredentialsGetHandler = s_a_n.IscsiCredentialsGetHandlerFunc(func(params s_a_n.IscsiCredentialsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiCredentialsGet has not yet been implemented")
		})
	}
	if api.SanIscsiCredentialsModifyHandler == nil {
		api.SanIscsiCredentialsModifyHandler = s_a_n.IscsiCredentialsModifyHandlerFunc(func(params s_a_n.IscsiCredentialsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiCredentialsModify has not yet been implemented")
		})
	}
	if api.SanIscsiServiceCollectionGetHandler == nil {
		api.SanIscsiServiceCollectionGetHandler = s_a_n.IscsiServiceCollectionGetHandlerFunc(func(params s_a_n.IscsiServiceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiServiceCollectionGet has not yet been implemented")
		})
	}
	if api.SanIscsiServiceCreateHandler == nil {
		api.SanIscsiServiceCreateHandler = s_a_n.IscsiServiceCreateHandlerFunc(func(params s_a_n.IscsiServiceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiServiceCreate has not yet been implemented")
		})
	}
	if api.SanIscsiServiceDeleteHandler == nil {
		api.SanIscsiServiceDeleteHandler = s_a_n.IscsiServiceDeleteHandlerFunc(func(params s_a_n.IscsiServiceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiServiceDelete has not yet been implemented")
		})
	}
	if api.SanIscsiServiceGetHandler == nil {
		api.SanIscsiServiceGetHandler = s_a_n.IscsiServiceGetHandlerFunc(func(params s_a_n.IscsiServiceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiServiceGet has not yet been implemented")
		})
	}
	if api.SanIscsiServiceModifyHandler == nil {
		api.SanIscsiServiceModifyHandler = s_a_n.IscsiServiceModifyHandlerFunc(func(params s_a_n.IscsiServiceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiServiceModify has not yet been implemented")
		})
	}
	if api.SanIscsiSessionCollectionGetHandler == nil {
		api.SanIscsiSessionCollectionGetHandler = s_a_n.IscsiSessionCollectionGetHandlerFunc(func(params s_a_n.IscsiSessionCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiSessionCollectionGet has not yet been implemented")
		})
	}
	if api.SanIscsiSessionGetHandler == nil {
		api.SanIscsiSessionGetHandler = s_a_n.IscsiSessionGetHandlerFunc(func(params s_a_n.IscsiSessionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.IscsiSessionGet has not yet been implemented")
		})
	}
	if api.ClusterJobCollectionGetHandler == nil {
		api.ClusterJobCollectionGetHandler = cluster.JobCollectionGetHandlerFunc(func(params cluster.JobCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.JobCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterJobGetHandler == nil {
		api.ClusterJobGetHandler = cluster.JobGetHandlerFunc(func(params cluster.JobGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.JobGet has not yet been implemented")
		})
	}
	if api.ClusterJobModifyHandler == nil {
		api.ClusterJobModifyHandler = cluster.JobModifyHandlerFunc(func(params cluster.JobModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.JobModify has not yet been implemented")
		})
	}
	if api.NasKerberosInterfaceCollectionGetHandler == nil {
		api.NasKerberosInterfaceCollectionGetHandler = n_a_s.KerberosInterfaceCollectionGetHandlerFunc(func(params n_a_s.KerberosInterfaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosInterfaceCollectionGet has not yet been implemented")
		})
	}
	if api.NasKerberosInterfaceGetHandler == nil {
		api.NasKerberosInterfaceGetHandler = n_a_s.KerberosInterfaceGetHandlerFunc(func(params n_a_s.KerberosInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosInterfaceGet has not yet been implemented")
		})
	}
	if api.NasKerberosInterfaceModifyHandler == nil {
		api.NasKerberosInterfaceModifyHandler = n_a_s.KerberosInterfaceModifyHandlerFunc(func(params n_a_s.KerberosInterfaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosInterfaceModify has not yet been implemented")
		})
	}
	if api.NasKerberosRealmCollectionGetHandler == nil {
		api.NasKerberosRealmCollectionGetHandler = n_a_s.KerberosRealmCollectionGetHandlerFunc(func(params n_a_s.KerberosRealmCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosRealmCollectionGet has not yet been implemented")
		})
	}
	if api.NasKerberosRealmCreateHandler == nil {
		api.NasKerberosRealmCreateHandler = n_a_s.KerberosRealmCreateHandlerFunc(func(params n_a_s.KerberosRealmCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosRealmCreate has not yet been implemented")
		})
	}
	if api.NasKerberosRealmDeleteHandler == nil {
		api.NasKerberosRealmDeleteHandler = n_a_s.KerberosRealmDeleteHandlerFunc(func(params n_a_s.KerberosRealmDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosRealmDelete has not yet been implemented")
		})
	}
	if api.NasKerberosRealmGetHandler == nil {
		api.NasKerberosRealmGetHandler = n_a_s.KerberosRealmGetHandlerFunc(func(params n_a_s.KerberosRealmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosRealmGet has not yet been implemented")
		})
	}
	if api.NasKerberosRealmModifyHandler == nil {
		api.NasKerberosRealmModifyHandler = n_a_s.KerberosRealmModifyHandlerFunc(func(params n_a_s.KerberosRealmModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.KerberosRealmModify has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerAuthKeyCollectionGetHandler == nil {
		api.SecurityKeyManagerAuthKeyCollectionGetHandler = securityops.KeyManagerAuthKeyCollectionGetHandlerFunc(func(params securityops.KeyManagerAuthKeyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerAuthKeyCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerAuthKeyCreateHandler == nil {
		api.SecurityKeyManagerAuthKeyCreateHandler = securityops.KeyManagerAuthKeyCreateHandlerFunc(func(params securityops.KeyManagerAuthKeyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerAuthKeyCreate has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerAuthKeyDeleteHandler == nil {
		api.SecurityKeyManagerAuthKeyDeleteHandler = securityops.KeyManagerAuthKeyDeleteHandlerFunc(func(params securityops.KeyManagerAuthKeyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerAuthKeyDelete has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerAuthKeyGetHandler == nil {
		api.SecurityKeyManagerAuthKeyGetHandler = securityops.KeyManagerAuthKeyGetHandlerFunc(func(params securityops.KeyManagerAuthKeyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerAuthKeyGet has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerConfigGetHandler == nil {
		api.SecurityKeyManagerConfigGetHandler = securityops.KeyManagerConfigGetHandlerFunc(func(params securityops.KeyManagerConfigGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerConfigGet has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerConfigModifyHandler == nil {
		api.SecurityKeyManagerConfigModifyHandler = securityops.KeyManagerConfigModifyHandlerFunc(func(params securityops.KeyManagerConfigModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerConfigModify has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerKeysCollectionGetHandler == nil {
		api.SecurityKeyManagerKeysCollectionGetHandler = securityops.KeyManagerKeysCollectionGetHandlerFunc(func(params securityops.KeyManagerKeysCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerKeysCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityKeyManagerKeysGetHandler == nil {
		api.SecurityKeyManagerKeysGetHandler = securityops.KeyManagerKeysGetHandlerFunc(func(params securityops.KeyManagerKeysGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.KeyManagerKeysGet has not yet been implemented")
		})
	}
	if api.NameServicesLdapCollectionGetHandler == nil {
		api.NameServicesLdapCollectionGetHandler = name_services.LdapCollectionGetHandlerFunc(func(params name_services.LdapCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesLdapCreateHandler == nil {
		api.NameServicesLdapCreateHandler = name_services.LdapCreateHandlerFunc(func(params name_services.LdapCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapCreate has not yet been implemented")
		})
	}
	if api.NameServicesLdapDeleteHandler == nil {
		api.NameServicesLdapDeleteHandler = name_services.LdapDeleteHandlerFunc(func(params name_services.LdapDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapDelete has not yet been implemented")
		})
	}
	if api.NameServicesLdapGetHandler == nil {
		api.NameServicesLdapGetHandler = name_services.LdapGetHandlerFunc(func(params name_services.LdapGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapGet has not yet been implemented")
		})
	}
	if api.NameServicesLdapModifyHandler == nil {
		api.NameServicesLdapModifyHandler = name_services.LdapModifyHandlerFunc(func(params name_services.LdapModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapModify has not yet been implemented")
		})
	}
	if api.NameServicesLdapSchemaCollectionGetHandler == nil {
		api.NameServicesLdapSchemaCollectionGetHandler = name_services.LdapSchemaCollectionGetHandlerFunc(func(params name_services.LdapSchemaCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapSchemaCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesLdapSchemaCreateHandler == nil {
		api.NameServicesLdapSchemaCreateHandler = name_services.LdapSchemaCreateHandlerFunc(func(params name_services.LdapSchemaCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapSchemaCreate has not yet been implemented")
		})
	}
	if api.NameServicesLdapSchemaDeleteHandler == nil {
		api.NameServicesLdapSchemaDeleteHandler = name_services.LdapSchemaDeleteHandlerFunc(func(params name_services.LdapSchemaDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapSchemaDelete has not yet been implemented")
		})
	}
	if api.NameServicesLdapSchemaGetHandler == nil {
		api.NameServicesLdapSchemaGetHandler = name_services.LdapSchemaGetHandlerFunc(func(params name_services.LdapSchemaGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapSchemaGet has not yet been implemented")
		})
	}
	if api.NameServicesLdapSchemaModifyHandler == nil {
		api.NameServicesLdapSchemaModifyHandler = name_services.LdapSchemaModifyHandlerFunc(func(params name_services.LdapSchemaModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LdapSchemaModify has not yet been implemented")
		})
	}
	if api.ClusterLicenseCreateHandler == nil {
		api.ClusterLicenseCreateHandler = cluster.LicenseCreateHandlerFunc(func(params cluster.LicenseCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseCreate has not yet been implemented")
		})
	}
	if api.ClusterLicenseDeleteHandler == nil {
		api.ClusterLicenseDeleteHandler = cluster.LicenseDeleteHandlerFunc(func(params cluster.LicenseDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseDelete has not yet been implemented")
		})
	}
	if api.ClusterLicenseGetHandler == nil {
		api.ClusterLicenseGetHandler = cluster.LicenseGetHandlerFunc(func(params cluster.LicenseGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseGet has not yet been implemented")
		})
	}
	if api.ClusterLicenseManagerCollectionGetHandler == nil {
		api.ClusterLicenseManagerCollectionGetHandler = cluster.LicenseManagerCollectionGetHandlerFunc(func(params cluster.LicenseManagerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseManagerCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterLicenseManagerGetHandler == nil {
		api.ClusterLicenseManagerGetHandler = cluster.LicenseManagerGetHandlerFunc(func(params cluster.LicenseManagerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseManagerGet has not yet been implemented")
		})
	}
	if api.ClusterLicenseManagerModifyHandler == nil {
		api.ClusterLicenseManagerModifyHandler = cluster.LicenseManagerModifyHandlerFunc(func(params cluster.LicenseManagerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicenseManagerModify has not yet been implemented")
		})
	}
	if api.ClusterLicensesGetHandler == nil {
		api.ClusterLicensesGetHandler = cluster.LicensesGetHandlerFunc(func(params cluster.LicensesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.LicensesGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupCollectionGetHandler == nil {
		api.NasLocalCifsGroupCollectionGetHandler = n_a_s.LocalCifsGroupCollectionGetHandlerFunc(func(params n_a_s.LocalCifsGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupCollectionGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupCreateHandler == nil {
		api.NasLocalCifsGroupCreateHandler = n_a_s.LocalCifsGroupCreateHandlerFunc(func(params n_a_s.LocalCifsGroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupCreate has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupDeleteHandler == nil {
		api.NasLocalCifsGroupDeleteHandler = n_a_s.LocalCifsGroupDeleteHandlerFunc(func(params n_a_s.LocalCifsGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupDelete has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupGetHandler == nil {
		api.NasLocalCifsGroupGetHandler = n_a_s.LocalCifsGroupGetHandlerFunc(func(params n_a_s.LocalCifsGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupMembersBulkDeleteHandler == nil {
		api.NasLocalCifsGroupMembersBulkDeleteHandler = n_a_s.LocalCifsGroupMembersBulkDeleteHandlerFunc(func(params n_a_s.LocalCifsGroupMembersBulkDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupMembersBulkDelete has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupMembersCollectionGetHandler == nil {
		api.NasLocalCifsGroupMembersCollectionGetHandler = n_a_s.LocalCifsGroupMembersCollectionGetHandlerFunc(func(params n_a_s.LocalCifsGroupMembersCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupMembersCollectionGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupMembersCreateHandler == nil {
		api.NasLocalCifsGroupMembersCreateHandler = n_a_s.LocalCifsGroupMembersCreateHandlerFunc(func(params n_a_s.LocalCifsGroupMembersCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupMembersCreate has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupMembersDeleteHandler == nil {
		api.NasLocalCifsGroupMembersDeleteHandler = n_a_s.LocalCifsGroupMembersDeleteHandlerFunc(func(params n_a_s.LocalCifsGroupMembersDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupMembersDelete has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupMembersGetHandler == nil {
		api.NasLocalCifsGroupMembersGetHandler = n_a_s.LocalCifsGroupMembersGetHandlerFunc(func(params n_a_s.LocalCifsGroupMembersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupMembersGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsGroupModifyHandler == nil {
		api.NasLocalCifsGroupModifyHandler = n_a_s.LocalCifsGroupModifyHandlerFunc(func(params n_a_s.LocalCifsGroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsGroupModify has not yet been implemented")
		})
	}
	if api.NasLocalCifsUserCollectionGetHandler == nil {
		api.NasLocalCifsUserCollectionGetHandler = n_a_s.LocalCifsUserCollectionGetHandlerFunc(func(params n_a_s.LocalCifsUserCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUserCollectionGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsUserCreateHandler == nil {
		api.NasLocalCifsUserCreateHandler = n_a_s.LocalCifsUserCreateHandlerFunc(func(params n_a_s.LocalCifsUserCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUserCreate has not yet been implemented")
		})
	}
	if api.NasLocalCifsUserDeleteHandler == nil {
		api.NasLocalCifsUserDeleteHandler = n_a_s.LocalCifsUserDeleteHandlerFunc(func(params n_a_s.LocalCifsUserDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUserDelete has not yet been implemented")
		})
	}
	if api.NasLocalCifsUserGetHandler == nil {
		api.NasLocalCifsUserGetHandler = n_a_s.LocalCifsUserGetHandlerFunc(func(params n_a_s.LocalCifsUserGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUserGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsUserModifyHandler == nil {
		api.NasLocalCifsUserModifyHandler = n_a_s.LocalCifsUserModifyHandlerFunc(func(params n_a_s.LocalCifsUserModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUserModify has not yet been implemented")
		})
	}
	if api.NasLocalCifsUsersAndGroupsImportCreateHandler == nil {
		api.NasLocalCifsUsersAndGroupsImportCreateHandler = n_a_s.LocalCifsUsersAndGroupsImportCreateHandlerFunc(func(params n_a_s.LocalCifsUsersAndGroupsImportCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUsersAndGroupsImportCreate has not yet been implemented")
		})
	}
	if api.NasLocalCifsUsersAndGroupsImportGetHandler == nil {
		api.NasLocalCifsUsersAndGroupsImportGetHandler = n_a_s.LocalCifsUsersAndGroupsImportGetHandlerFunc(func(params n_a_s.LocalCifsUsersAndGroupsImportGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUsersAndGroupsImportGet has not yet been implemented")
		})
	}
	if api.NasLocalCifsUsersAndGroupsImportModifyHandler == nil {
		api.NasLocalCifsUsersAndGroupsImportModifyHandler = n_a_s.LocalCifsUsersAndGroupsImportModifyHandlerFunc(func(params n_a_s.LocalCifsUsersAndGroupsImportModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.LocalCifsUsersAndGroupsImportModify has not yet been implemented")
		})
	}
	if api.NameServicesLocalHostCollectionGetHandler == nil {
		api.NameServicesLocalHostCollectionGetHandler = name_services.LocalHostCollectionGetHandlerFunc(func(params name_services.LocalHostCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LocalHostCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesLocalHostCreateHandler == nil {
		api.NameServicesLocalHostCreateHandler = name_services.LocalHostCreateHandlerFunc(func(params name_services.LocalHostCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LocalHostCreate has not yet been implemented")
		})
	}
	if api.NameServicesLocalHostDeleteHandler == nil {
		api.NameServicesLocalHostDeleteHandler = name_services.LocalHostDeleteHandlerFunc(func(params name_services.LocalHostDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LocalHostDelete has not yet been implemented")
		})
	}
	if api.NameServicesLocalHostGetHandler == nil {
		api.NameServicesLocalHostGetHandler = name_services.LocalHostGetHandlerFunc(func(params name_services.LocalHostGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LocalHostGet has not yet been implemented")
		})
	}
	if api.NameServicesLocalHostModifyHandler == nil {
		api.NameServicesLocalHostModifyHandler = name_services.LocalHostModifyHandlerFunc(func(params name_services.LocalHostModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.LocalHostModify has not yet been implemented")
		})
	}
	if api.SecurityLoginMessagesCollectionGetHandler == nil {
		api.SecurityLoginMessagesCollectionGetHandler = securityops.LoginMessagesCollectionGetHandlerFunc(func(params securityops.LoginMessagesCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.LoginMessagesCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityLoginMessagesGetHandler == nil {
		api.SecurityLoginMessagesGetHandler = securityops.LoginMessagesGetHandlerFunc(func(params securityops.LoginMessagesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.LoginMessagesGet has not yet been implemented")
		})
	}
	if api.SecurityLoginMessagesModifyHandler == nil {
		api.SecurityLoginMessagesModifyHandler = securityops.LoginMessagesModifyHandlerFunc(func(params securityops.LoginMessagesModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.LoginMessagesModify has not yet been implemented")
		})
	}
	if api.SanLunAttributeCollectionGetHandler == nil {
		api.SanLunAttributeCollectionGetHandler = s_a_n.LunAttributeCollectionGetHandlerFunc(func(params s_a_n.LunAttributeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunAttributeCollectionGet has not yet been implemented")
		})
	}
	if api.SanLunAttributeCreateHandler == nil {
		api.SanLunAttributeCreateHandler = s_a_n.LunAttributeCreateHandlerFunc(func(params s_a_n.LunAttributeCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunAttributeCreate has not yet been implemented")
		})
	}
	if api.SanLunAttributeDeleteHandler == nil {
		api.SanLunAttributeDeleteHandler = s_a_n.LunAttributeDeleteHandlerFunc(func(params s_a_n.LunAttributeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunAttributeDelete has not yet been implemented")
		})
	}
	if api.SanLunAttributeGetHandler == nil {
		api.SanLunAttributeGetHandler = s_a_n.LunAttributeGetHandlerFunc(func(params s_a_n.LunAttributeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunAttributeGet has not yet been implemented")
		})
	}
	if api.SanLunAttributeModifyHandler == nil {
		api.SanLunAttributeModifyHandler = s_a_n.LunAttributeModifyHandlerFunc(func(params s_a_n.LunAttributeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunAttributeModify has not yet been implemented")
		})
	}
	if api.SanLunCollectionGetHandler == nil {
		api.SanLunCollectionGetHandler = s_a_n.LunCollectionGetHandlerFunc(func(params s_a_n.LunCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunCollectionGet has not yet been implemented")
		})
	}
	if api.SanLunCreateHandler == nil {
		api.SanLunCreateHandler = s_a_n.LunCreateHandlerFunc(func(params s_a_n.LunCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunCreate has not yet been implemented")
		})
	}
	if api.SanLunDeleteHandler == nil {
		api.SanLunDeleteHandler = s_a_n.LunDeleteHandlerFunc(func(params s_a_n.LunDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunDelete has not yet been implemented")
		})
	}
	if api.SanLunGetHandler == nil {
		api.SanLunGetHandler = s_a_n.LunGetHandlerFunc(func(params s_a_n.LunGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunGet has not yet been implemented")
		})
	}
	if api.SanLunMapCollectionGetHandler == nil {
		api.SanLunMapCollectionGetHandler = s_a_n.LunMapCollectionGetHandlerFunc(func(params s_a_n.LunMapCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapCollectionGet has not yet been implemented")
		})
	}
	if api.SanLunMapCreateHandler == nil {
		api.SanLunMapCreateHandler = s_a_n.LunMapCreateHandlerFunc(func(params s_a_n.LunMapCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapCreate has not yet been implemented")
		})
	}
	if api.SanLunMapDeleteHandler == nil {
		api.SanLunMapDeleteHandler = s_a_n.LunMapDeleteHandlerFunc(func(params s_a_n.LunMapDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapDelete has not yet been implemented")
		})
	}
	if api.SanLunMapGetHandler == nil {
		api.SanLunMapGetHandler = s_a_n.LunMapGetHandlerFunc(func(params s_a_n.LunMapGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapGet has not yet been implemented")
		})
	}
	if api.SanLunMapReportingNodeCollectionGetHandler == nil {
		api.SanLunMapReportingNodeCollectionGetHandler = s_a_n.LunMapReportingNodeCollectionGetHandlerFunc(func(params s_a_n.LunMapReportingNodeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapReportingNodeCollectionGet has not yet been implemented")
		})
	}
	if api.SanLunMapReportingNodeCreateHandler == nil {
		api.SanLunMapReportingNodeCreateHandler = s_a_n.LunMapReportingNodeCreateHandlerFunc(func(params s_a_n.LunMapReportingNodeCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapReportingNodeCreate has not yet been implemented")
		})
	}
	if api.SanLunMapReportingNodeDeleteHandler == nil {
		api.SanLunMapReportingNodeDeleteHandler = s_a_n.LunMapReportingNodeDeleteHandlerFunc(func(params s_a_n.LunMapReportingNodeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapReportingNodeDelete has not yet been implemented")
		})
	}
	if api.SanLunMapReportingNodeGetHandler == nil {
		api.SanLunMapReportingNodeGetHandler = s_a_n.LunMapReportingNodeGetHandlerFunc(func(params s_a_n.LunMapReportingNodeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunMapReportingNodeGet has not yet been implemented")
		})
	}
	if api.SanLunModifyHandler == nil {
		api.SanLunModifyHandler = s_a_n.LunModifyHandlerFunc(func(params s_a_n.LunModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.LunModify has not yet been implemented")
		})
	}
	if api.ClusterMediatorCollectionGetHandler == nil {
		api.ClusterMediatorCollectionGetHandler = cluster.MediatorCollectionGetHandlerFunc(func(params cluster.MediatorCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MediatorCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMediatorCreateHandler == nil {
		api.ClusterMediatorCreateHandler = cluster.MediatorCreateHandlerFunc(func(params cluster.MediatorCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MediatorCreate has not yet been implemented")
		})
	}
	if api.ClusterMediatorDeleteHandler == nil {
		api.ClusterMediatorDeleteHandler = cluster.MediatorDeleteHandlerFunc(func(params cluster.MediatorDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MediatorDelete has not yet been implemented")
		})
	}
	if api.ClusterMediatorGetHandler == nil {
		api.ClusterMediatorGetHandler = cluster.MediatorGetHandlerFunc(func(params cluster.MediatorGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MediatorGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterCreateHandler == nil {
		api.ClusterMetroclusterCreateHandler = cluster.MetroclusterCreateHandlerFunc(func(params cluster.MetroclusterCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterCreate has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDiagnosticsCreateHandler == nil {
		api.ClusterMetroclusterDiagnosticsCreateHandler = cluster.MetroclusterDiagnosticsCreateHandlerFunc(func(params cluster.MetroclusterDiagnosticsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDiagnosticsCreate has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDiagnosticsGetHandler == nil {
		api.ClusterMetroclusterDiagnosticsGetHandler = cluster.MetroclusterDiagnosticsGetHandlerFunc(func(params cluster.MetroclusterDiagnosticsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDiagnosticsGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDrGroupCollectionGetHandler == nil {
		api.ClusterMetroclusterDrGroupCollectionGetHandler = cluster.MetroclusterDrGroupCollectionGetHandlerFunc(func(params cluster.MetroclusterDrGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDrGroupCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDrGroupCreateHandler == nil {
		api.ClusterMetroclusterDrGroupCreateHandler = cluster.MetroclusterDrGroupCreateHandlerFunc(func(params cluster.MetroclusterDrGroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDrGroupCreate has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDrGroupDeleteHandler == nil {
		api.ClusterMetroclusterDrGroupDeleteHandler = cluster.MetroclusterDrGroupDeleteHandlerFunc(func(params cluster.MetroclusterDrGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDrGroupDelete has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterDrGroupGetHandler == nil {
		api.ClusterMetroclusterDrGroupGetHandler = cluster.MetroclusterDrGroupGetHandlerFunc(func(params cluster.MetroclusterDrGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterDrGroupGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterGetHandler == nil {
		api.ClusterMetroclusterGetHandler = cluster.MetroclusterGetHandlerFunc(func(params cluster.MetroclusterGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterInterconnectCollectionGetHandler == nil {
		api.ClusterMetroclusterInterconnectCollectionGetHandler = cluster.MetroclusterInterconnectCollectionGetHandlerFunc(func(params cluster.MetroclusterInterconnectCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterInterconnectCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterInterconnectGetHandler == nil {
		api.ClusterMetroclusterInterconnectGetHandler = cluster.MetroclusterInterconnectGetHandlerFunc(func(params cluster.MetroclusterInterconnectGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterInterconnectGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterInterconnectModifyHandler == nil {
		api.ClusterMetroclusterInterconnectModifyHandler = cluster.MetroclusterInterconnectModifyHandlerFunc(func(params cluster.MetroclusterInterconnectModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterInterconnectModify has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterModifyHandler == nil {
		api.ClusterMetroclusterModifyHandler = cluster.MetroclusterModifyHandlerFunc(func(params cluster.MetroclusterModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterModify has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterNodeCollectionGetHandler == nil {
		api.ClusterMetroclusterNodeCollectionGetHandler = cluster.MetroclusterNodeCollectionGetHandlerFunc(func(params cluster.MetroclusterNodeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterNodeCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterNodeGetHandler == nil {
		api.ClusterMetroclusterNodeGetHandler = cluster.MetroclusterNodeGetHandlerFunc(func(params cluster.MetroclusterNodeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterNodeGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterOperationCollectionGetHandler == nil {
		api.ClusterMetroclusterOperationCollectionGetHandler = cluster.MetroclusterOperationCollectionGetHandlerFunc(func(params cluster.MetroclusterOperationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterOperationCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterOperationGetHandler == nil {
		api.ClusterMetroclusterOperationGetHandler = cluster.MetroclusterOperationGetHandlerFunc(func(params cluster.MetroclusterOperationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterOperationGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterSvmCollectionGetHandler == nil {
		api.ClusterMetroclusterSvmCollectionGetHandler = cluster.MetroclusterSvmCollectionGetHandlerFunc(func(params cluster.MetroclusterSvmCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterSvmCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterMetroclusterSvmGetHandler == nil {
		api.ClusterMetroclusterSvmGetHandler = cluster.MetroclusterSvmGetHandlerFunc(func(params cluster.MetroclusterSvmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.MetroclusterSvmGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyApprovalGroupCollectionGetHandler == nil {
		api.SecurityMultiAdminVerifyApprovalGroupCollectionGetHandler = securityops.MultiAdminVerifyApprovalGroupCollectionGetHandlerFunc(func(params securityops.MultiAdminVerifyApprovalGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyApprovalGroupCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyApprovalGroupCreateHandler == nil {
		api.SecurityMultiAdminVerifyApprovalGroupCreateHandler = securityops.MultiAdminVerifyApprovalGroupCreateHandlerFunc(func(params securityops.MultiAdminVerifyApprovalGroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyApprovalGroupCreate has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyApprovalGroupDeleteHandler == nil {
		api.SecurityMultiAdminVerifyApprovalGroupDeleteHandler = securityops.MultiAdminVerifyApprovalGroupDeleteHandlerFunc(func(params securityops.MultiAdminVerifyApprovalGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyApprovalGroupDelete has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyApprovalGroupGetHandler == nil {
		api.SecurityMultiAdminVerifyApprovalGroupGetHandler = securityops.MultiAdminVerifyApprovalGroupGetHandlerFunc(func(params securityops.MultiAdminVerifyApprovalGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyApprovalGroupGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyApprovalGroupModifyHandler == nil {
		api.SecurityMultiAdminVerifyApprovalGroupModifyHandler = securityops.MultiAdminVerifyApprovalGroupModifyHandlerFunc(func(params securityops.MultiAdminVerifyApprovalGroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyApprovalGroupModify has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyConfigGetHandler == nil {
		api.SecurityMultiAdminVerifyConfigGetHandler = securityops.MultiAdminVerifyConfigGetHandlerFunc(func(params securityops.MultiAdminVerifyConfigGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyConfigGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyConfigModifyHandler == nil {
		api.SecurityMultiAdminVerifyConfigModifyHandler = securityops.MultiAdminVerifyConfigModifyHandlerFunc(func(params securityops.MultiAdminVerifyConfigModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyConfigModify has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRequestCollectionGetHandler == nil {
		api.SecurityMultiAdminVerifyRequestCollectionGetHandler = securityops.MultiAdminVerifyRequestCollectionGetHandlerFunc(func(params securityops.MultiAdminVerifyRequestCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRequestCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRequestCreateHandler == nil {
		api.SecurityMultiAdminVerifyRequestCreateHandler = securityops.MultiAdminVerifyRequestCreateHandlerFunc(func(params securityops.MultiAdminVerifyRequestCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRequestCreate has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRequestDeleteHandler == nil {
		api.SecurityMultiAdminVerifyRequestDeleteHandler = securityops.MultiAdminVerifyRequestDeleteHandlerFunc(func(params securityops.MultiAdminVerifyRequestDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRequestDelete has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRequestGetHandler == nil {
		api.SecurityMultiAdminVerifyRequestGetHandler = securityops.MultiAdminVerifyRequestGetHandlerFunc(func(params securityops.MultiAdminVerifyRequestGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRequestGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRequestModifyHandler == nil {
		api.SecurityMultiAdminVerifyRequestModifyHandler = securityops.MultiAdminVerifyRequestModifyHandlerFunc(func(params securityops.MultiAdminVerifyRequestModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRequestModify has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRuleCollectionGetHandler == nil {
		api.SecurityMultiAdminVerifyRuleCollectionGetHandler = securityops.MultiAdminVerifyRuleCollectionGetHandlerFunc(func(params securityops.MultiAdminVerifyRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRuleCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRuleCreateHandler == nil {
		api.SecurityMultiAdminVerifyRuleCreateHandler = securityops.MultiAdminVerifyRuleCreateHandlerFunc(func(params securityops.MultiAdminVerifyRuleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRuleCreate has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRuleDeleteHandler == nil {
		api.SecurityMultiAdminVerifyRuleDeleteHandler = securityops.MultiAdminVerifyRuleDeleteHandlerFunc(func(params securityops.MultiAdminVerifyRuleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRuleDelete has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRuleGetHandler == nil {
		api.SecurityMultiAdminVerifyRuleGetHandler = securityops.MultiAdminVerifyRuleGetHandlerFunc(func(params securityops.MultiAdminVerifyRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRuleGet has not yet been implemented")
		})
	}
	if api.SecurityMultiAdminVerifyRuleModifyHandler == nil {
		api.SecurityMultiAdminVerifyRuleModifyHandler = securityops.MultiAdminVerifyRuleModifyHandlerFunc(func(params securityops.MultiAdminVerifyRuleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.MultiAdminVerifyRuleModify has not yet been implemented")
		})
	}
	if api.NameServicesNameMappingCollectionGetHandler == nil {
		api.NameServicesNameMappingCollectionGetHandler = name_services.NameMappingCollectionGetHandlerFunc(func(params name_services.NameMappingCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NameMappingCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesNameMappingCreateHandler == nil {
		api.NameServicesNameMappingCreateHandler = name_services.NameMappingCreateHandlerFunc(func(params name_services.NameMappingCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NameMappingCreate has not yet been implemented")
		})
	}
	if api.NameServicesNameMappingDeleteHandler == nil {
		api.NameServicesNameMappingDeleteHandler = name_services.NameMappingDeleteHandlerFunc(func(params name_services.NameMappingDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NameMappingDelete has not yet been implemented")
		})
	}
	if api.NameServicesNameMappingModifyHandler == nil {
		api.NameServicesNameMappingModifyHandler = name_services.NameMappingModifyHandlerFunc(func(params name_services.NameMappingModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NameMappingModify has not yet been implemented")
		})
	}
	if api.NameServicesNameMappingPositionGetHandler == nil {
		api.NameServicesNameMappingPositionGetHandler = name_services.NameMappingPositionGetHandlerFunc(func(params name_services.NameMappingPositionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NameMappingPositionGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeCollectionGetHandler == nil {
		api.NdmpNdmpNodeCollectionGetHandler = n_d_m_p.NdmpNodeCollectionGetHandlerFunc(func(params n_d_m_p.NdmpNodeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeCollectionGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeGetHandler == nil {
		api.NdmpNdmpNodeGetHandler = n_d_m_p.NdmpNodeGetHandlerFunc(func(params n_d_m_p.NdmpNodeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeModifyHandler == nil {
		api.NdmpNdmpNodeModifyHandler = n_d_m_p.NdmpNodeModifyHandlerFunc(func(params n_d_m_p.NdmpNodeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeModify has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeSessionDeleteHandler == nil {
		api.NdmpNdmpNodeSessionDeleteHandler = n_d_m_p.NdmpNodeSessionDeleteHandlerFunc(func(params n_d_m_p.NdmpNodeSessionDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeSessionDelete has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeSessionGetHandler == nil {
		api.NdmpNdmpNodeSessionGetHandler = n_d_m_p.NdmpNodeSessionGetHandlerFunc(func(params n_d_m_p.NdmpNodeSessionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeSessionGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpNodeSessionsCollectionGetHandler == nil {
		api.NdmpNdmpNodeSessionsCollectionGetHandler = n_d_m_p.NdmpNodeSessionsCollectionGetHandlerFunc(func(params n_d_m_p.NdmpNodeSessionsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpNodeSessionsCollectionGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpPasswordGetHandler == nil {
		api.NdmpNdmpPasswordGetHandler = n_d_m_p.NdmpPasswordGetHandlerFunc(func(params n_d_m_p.NdmpPasswordGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpPasswordGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpSvmCollectionGetHandler == nil {
		api.NdmpNdmpSvmCollectionGetHandler = n_d_m_p.NdmpSvmCollectionGetHandlerFunc(func(params n_d_m_p.NdmpSvmCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpSvmCollectionGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpSvmGetHandler == nil {
		api.NdmpNdmpSvmGetHandler = n_d_m_p.NdmpSvmGetHandlerFunc(func(params n_d_m_p.NdmpSvmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpSvmGet has not yet been implemented")
		})
	}
	if api.NdmpNdmpSvmModifyHandler == nil {
		api.NdmpNdmpSvmModifyHandler = n_d_m_p.NdmpSvmModifyHandlerFunc(func(params n_d_m_p.NdmpSvmModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_d_m_p.NdmpSvmModify has not yet been implemented")
		})
	}
	if api.NasNetbiosCollectionGetHandler == nil {
		api.NasNetbiosCollectionGetHandler = n_a_s.NetbiosCollectionGetHandlerFunc(func(params n_a_s.NetbiosCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NetbiosCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesNetgroupFileDeleteHandler == nil {
		api.NameServicesNetgroupFileDeleteHandler = name_services.NetgroupFileDeleteHandlerFunc(func(params name_services.NetgroupFileDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NetgroupFileDelete has not yet been implemented")
		})
	}
	if api.NameServicesNetgroupFileGetHandler == nil {
		api.NameServicesNetgroupFileGetHandler = name_services.NetgroupFileGetHandlerFunc(func(params name_services.NetgroupFileGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NetgroupFileGet has not yet been implemented")
		})
	}
	if api.NameServicesNetgroupsSettingsCollectionGetHandler == nil {
		api.NameServicesNetgroupsSettingsCollectionGetHandler = name_services.NetgroupsSettingsCollectionGetHandlerFunc(func(params name_services.NetgroupsSettingsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NetgroupsSettingsCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesNetgroupsSettingsGetHandler == nil {
		api.NameServicesNetgroupsSettingsGetHandler = name_services.NetgroupsSettingsGetHandlerFunc(func(params name_services.NetgroupsSettingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NetgroupsSettingsGet has not yet been implemented")
		})
	}
	if api.NameServicesNetgroupsSettingsModifyHandler == nil {
		api.NameServicesNetgroupsSettingsModifyHandler = name_services.NetgroupsSettingsModifyHandlerFunc(func(params name_services.NetgroupsSettingsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NetgroupsSettingsModify has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetBroadcastDomainDeleteHandler == nil {
		api.NetworkingNetworkEthernetBroadcastDomainDeleteHandler = networking.NetworkEthernetBroadcastDomainDeleteHandlerFunc(func(params networking.NetworkEthernetBroadcastDomainDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetBroadcastDomainDelete has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetBroadcastDomainGetHandler == nil {
		api.NetworkingNetworkEthernetBroadcastDomainGetHandler = networking.NetworkEthernetBroadcastDomainGetHandlerFunc(func(params networking.NetworkEthernetBroadcastDomainGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetBroadcastDomainGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetBroadcastDomainModifyHandler == nil {
		api.NetworkingNetworkEthernetBroadcastDomainModifyHandler = networking.NetworkEthernetBroadcastDomainModifyHandlerFunc(func(params networking.NetworkEthernetBroadcastDomainModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetBroadcastDomainModify has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetBroadcastDomainsCreateHandler == nil {
		api.NetworkingNetworkEthernetBroadcastDomainsCreateHandler = networking.NetworkEthernetBroadcastDomainsCreateHandlerFunc(func(params networking.NetworkEthernetBroadcastDomainsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetBroadcastDomainsCreate has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetBroadcastDomainsGetHandler == nil {
		api.NetworkingNetworkEthernetBroadcastDomainsGetHandler = networking.NetworkEthernetBroadcastDomainsGetHandlerFunc(func(params networking.NetworkEthernetBroadcastDomainsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetBroadcastDomainsGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetPortDeleteHandler == nil {
		api.NetworkingNetworkEthernetPortDeleteHandler = networking.NetworkEthernetPortDeleteHandlerFunc(func(params networking.NetworkEthernetPortDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetPortDelete has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetPortGetHandler == nil {
		api.NetworkingNetworkEthernetPortGetHandler = networking.NetworkEthernetPortGetHandlerFunc(func(params networking.NetworkEthernetPortGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetPortGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetPortModifyHandler == nil {
		api.NetworkingNetworkEthernetPortModifyHandler = networking.NetworkEthernetPortModifyHandlerFunc(func(params networking.NetworkEthernetPortModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetPortModify has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetPortsCreateHandler == nil {
		api.NetworkingNetworkEthernetPortsCreateHandler = networking.NetworkEthernetPortsCreateHandlerFunc(func(params networking.NetworkEthernetPortsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetPortsCreate has not yet been implemented")
		})
	}
	if api.NetworkingNetworkEthernetPortsGetHandler == nil {
		api.NetworkingNetworkEthernetPortsGetHandler = networking.NetworkEthernetPortsGetHandlerFunc(func(params networking.NetworkEthernetPortsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkEthernetPortsGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPBgpPeerGroupDeleteHandler == nil {
		api.NetworkingNetworkIPBgpPeerGroupDeleteHandler = networking.NetworkIPBgpPeerGroupDeleteHandlerFunc(func(params networking.NetworkIPBgpPeerGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPBgpPeerGroupDelete has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPBgpPeerGroupGetHandler == nil {
		api.NetworkingNetworkIPBgpPeerGroupGetHandler = networking.NetworkIPBgpPeerGroupGetHandlerFunc(func(params networking.NetworkIPBgpPeerGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPBgpPeerGroupGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPBgpPeerGroupModifyHandler == nil {
		api.NetworkingNetworkIPBgpPeerGroupModifyHandler = networking.NetworkIPBgpPeerGroupModifyHandlerFunc(func(params networking.NetworkIPBgpPeerGroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPBgpPeerGroupModify has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPBgpPeerGroupsCreateHandler == nil {
		api.NetworkingNetworkIPBgpPeerGroupsCreateHandler = networking.NetworkIPBgpPeerGroupsCreateHandlerFunc(func(params networking.NetworkIPBgpPeerGroupsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPBgpPeerGroupsCreate has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPBgpPeerGroupsGetHandler == nil {
		api.NetworkingNetworkIPBgpPeerGroupsGetHandler = networking.NetworkIPBgpPeerGroupsGetHandlerFunc(func(params networking.NetworkIPBgpPeerGroupsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPBgpPeerGroupsGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPInterfaceDeleteHandler == nil {
		api.NetworkingNetworkIPInterfaceDeleteHandler = networking.NetworkIPInterfaceDeleteHandlerFunc(func(params networking.NetworkIPInterfaceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPInterfaceDelete has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPInterfaceGetHandler == nil {
		api.NetworkingNetworkIPInterfaceGetHandler = networking.NetworkIPInterfaceGetHandlerFunc(func(params networking.NetworkIPInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPInterfaceGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPInterfaceModifyHandler == nil {
		api.NetworkingNetworkIPInterfaceModifyHandler = networking.NetworkIPInterfaceModifyHandlerFunc(func(params networking.NetworkIPInterfaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPInterfaceModify has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPInterfacesCreateHandler == nil {
		api.NetworkingNetworkIPInterfacesCreateHandler = networking.NetworkIPInterfacesCreateHandlerFunc(func(params networking.NetworkIPInterfacesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPInterfacesCreate has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPInterfacesGetHandler == nil {
		api.NetworkingNetworkIPInterfacesGetHandler = networking.NetworkIPInterfacesGetHandlerFunc(func(params networking.NetworkIPInterfacesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPInterfacesGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPRouteDeleteHandler == nil {
		api.NetworkingNetworkIPRouteDeleteHandler = networking.NetworkIPRouteDeleteHandlerFunc(func(params networking.NetworkIPRouteDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPRouteDelete has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPRouteGetHandler == nil {
		api.NetworkingNetworkIPRouteGetHandler = networking.NetworkIPRouteGetHandlerFunc(func(params networking.NetworkIPRouteGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPRouteGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPRoutesCreateHandler == nil {
		api.NetworkingNetworkIPRoutesCreateHandler = networking.NetworkIPRoutesCreateHandlerFunc(func(params networking.NetworkIPRoutesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPRoutesCreate has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPRoutesGetHandler == nil {
		api.NetworkingNetworkIPRoutesGetHandler = networking.NetworkIPRoutesGetHandlerFunc(func(params networking.NetworkIPRoutesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPRoutesGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPServicePoliciesGetHandler == nil {
		api.NetworkingNetworkIPServicePoliciesGetHandler = networking.NetworkIPServicePoliciesGetHandlerFunc(func(params networking.NetworkIPServicePoliciesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPServicePoliciesGet has not yet been implemented")
		})
	}
	if api.NetworkingNetworkIPServicePolicyGetHandler == nil {
		api.NetworkingNetworkIPServicePolicyGetHandler = networking.NetworkIPServicePolicyGetHandlerFunc(func(params networking.NetworkIPServicePolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.NetworkIPServicePolicyGet has not yet been implemented")
		})
	}
	if api.NasNfsClientsCacheGetHandler == nil {
		api.NasNfsClientsCacheGetHandler = n_a_s.NfsClientsCacheGetHandlerFunc(func(params n_a_s.NfsClientsCacheGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsClientsCacheGet has not yet been implemented")
		})
	}
	if api.NasNfsClientsCacheModifyHandler == nil {
		api.NasNfsClientsCacheModifyHandler = n_a_s.NfsClientsCacheModifyHandlerFunc(func(params n_a_s.NfsClientsCacheModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsClientsCacheModify has not yet been implemented")
		})
	}
	if api.NasNfsClientsGetHandler == nil {
		api.NasNfsClientsGetHandler = n_a_s.NfsClientsGetHandlerFunc(func(params n_a_s.NfsClientsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsClientsGet has not yet been implemented")
		})
	}
	if api.NasNfsClientsMapGetHandler == nil {
		api.NasNfsClientsMapGetHandler = n_a_s.NfsClientsMapGetHandlerFunc(func(params n_a_s.NfsClientsMapGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsClientsMapGet has not yet been implemented")
		})
	}
	if api.NasNfsCollectionGetHandler == nil {
		api.NasNfsCollectionGetHandler = n_a_s.NfsCollectionGetHandlerFunc(func(params n_a_s.NfsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsCollectionGet has not yet been implemented")
		})
	}
	if api.NasNfsCollectionPerformanceMetricsGetHandler == nil {
		api.NasNfsCollectionPerformanceMetricsGetHandler = n_a_s.NfsCollectionPerformanceMetricsGetHandlerFunc(func(params n_a_s.NfsCollectionPerformanceMetricsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsCollectionPerformanceMetricsGet has not yet been implemented")
		})
	}
	if api.NasNfsCreateHandler == nil {
		api.NasNfsCreateHandler = n_a_s.NfsCreateHandlerFunc(func(params n_a_s.NfsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsCreate has not yet been implemented")
		})
	}
	if api.NasNfsDeleteHandler == nil {
		api.NasNfsDeleteHandler = n_a_s.NfsDeleteHandlerFunc(func(params n_a_s.NfsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsDelete has not yet been implemented")
		})
	}
	if api.NasNfsGetHandler == nil {
		api.NasNfsGetHandler = n_a_s.NfsGetHandlerFunc(func(params n_a_s.NfsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsGet has not yet been implemented")
		})
	}
	if api.NasNfsModifyHandler == nil {
		api.NasNfsModifyHandler = n_a_s.NfsModifyHandlerFunc(func(params n_a_s.NfsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsModify has not yet been implemented")
		})
	}
	if api.NasNfsTLSInterfaceCollectionGetHandler == nil {
		api.NasNfsTLSInterfaceCollectionGetHandler = n_a_s.NfsTLSInterfaceCollectionGetHandlerFunc(func(params n_a_s.NfsTLSInterfaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsTLSInterfaceCollectionGet has not yet been implemented")
		})
	}
	if api.NasNfsTLSInterfaceGetHandler == nil {
		api.NasNfsTLSInterfaceGetHandler = n_a_s.NfsTLSInterfaceGetHandlerFunc(func(params n_a_s.NfsTLSInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsTLSInterfaceGet has not yet been implemented")
		})
	}
	if api.NasNfsTLSInterfaceModifyHandler == nil {
		api.NasNfsTLSInterfaceModifyHandler = n_a_s.NfsTLSInterfaceModifyHandlerFunc(func(params n_a_s.NfsTLSInterfaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.NfsTLSInterfaceModify has not yet been implemented")
		})
	}
	if api.NameServicesNisCollectionGetHandler == nil {
		api.NameServicesNisCollectionGetHandler = name_services.NisCollectionGetHandlerFunc(func(params name_services.NisCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NisCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesNisCreateHandler == nil {
		api.NameServicesNisCreateHandler = name_services.NisCreateHandlerFunc(func(params name_services.NisCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NisCreate has not yet been implemented")
		})
	}
	if api.NameServicesNisDeleteHandler == nil {
		api.NameServicesNisDeleteHandler = name_services.NisDeleteHandlerFunc(func(params name_services.NisDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NisDelete has not yet been implemented")
		})
	}
	if api.NameServicesNisGetHandler == nil {
		api.NameServicesNisGetHandler = name_services.NisGetHandlerFunc(func(params name_services.NisGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NisGet has not yet been implemented")
		})
	}
	if api.NameServicesNisModifyHandler == nil {
		api.NameServicesNisModifyHandler = name_services.NisModifyHandlerFunc(func(params name_services.NisModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.NisModify has not yet been implemented")
		})
	}
	if api.ClusterNodeDeleteHandler == nil {
		api.ClusterNodeDeleteHandler = cluster.NodeDeleteHandlerFunc(func(params cluster.NodeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodeDelete has not yet been implemented")
		})
	}
	if api.ClusterNodeGetHandler == nil {
		api.ClusterNodeGetHandler = cluster.NodeGetHandlerFunc(func(params cluster.NodeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodeGet has not yet been implemented")
		})
	}
	if api.ClusterNodeMetricsCollectionGetHandler == nil {
		api.ClusterNodeMetricsCollectionGetHandler = cluster.NodeMetricsCollectionGetHandlerFunc(func(params cluster.NodeMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodeMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterNodeModifyHandler == nil {
		api.ClusterNodeModifyHandler = cluster.NodeModifyHandlerFunc(func(params cluster.NodeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodeModify has not yet been implemented")
		})
	}
	if api.ClusterNodesCreateHandler == nil {
		api.ClusterNodesCreateHandler = cluster.NodesCreateHandlerFunc(func(params cluster.NodesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodesCreate has not yet been implemented")
		})
	}
	if api.ClusterNodesGetHandler == nil {
		api.ClusterNodesGetHandler = cluster.NodesGetHandlerFunc(func(params cluster.NodesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.NodesGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeInterfaceCollectionGetHandler == nil {
		api.NvMeNvmeInterfaceCollectionGetHandler = n_v_me.NvmeInterfaceCollectionGetHandlerFunc(func(params n_v_me.NvmeInterfaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeInterfaceCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeInterfaceGetHandler == nil {
		api.NvMeNvmeInterfaceGetHandler = n_v_me.NvmeInterfaceGetHandlerFunc(func(params n_v_me.NvmeInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeInterfaceGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeNamespaceCollectionGetHandler == nil {
		api.NvMeNvmeNamespaceCollectionGetHandler = n_v_me.NvmeNamespaceCollectionGetHandlerFunc(func(params n_v_me.NvmeNamespaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeNamespaceCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeNamespaceCreateHandler == nil {
		api.NvMeNvmeNamespaceCreateHandler = n_v_me.NvmeNamespaceCreateHandlerFunc(func(params n_v_me.NvmeNamespaceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeNamespaceCreate has not yet been implemented")
		})
	}
	if api.NvMeNvmeNamespaceDeleteHandler == nil {
		api.NvMeNvmeNamespaceDeleteHandler = n_v_me.NvmeNamespaceDeleteHandlerFunc(func(params n_v_me.NvmeNamespaceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeNamespaceDelete has not yet been implemented")
		})
	}
	if api.NvMeNvmeNamespaceGetHandler == nil {
		api.NvMeNvmeNamespaceGetHandler = n_v_me.NvmeNamespaceGetHandlerFunc(func(params n_v_me.NvmeNamespaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeNamespaceGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeNamespaceModifyHandler == nil {
		api.NvMeNvmeNamespaceModifyHandler = n_v_me.NvmeNamespaceModifyHandlerFunc(func(params n_v_me.NvmeNamespaceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeNamespaceModify has not yet been implemented")
		})
	}
	if api.NvMeNvmeServiceCollectionGetHandler == nil {
		api.NvMeNvmeServiceCollectionGetHandler = n_v_me.NvmeServiceCollectionGetHandlerFunc(func(params n_v_me.NvmeServiceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeServiceCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeServiceCreateHandler == nil {
		api.NvMeNvmeServiceCreateHandler = n_v_me.NvmeServiceCreateHandlerFunc(func(params n_v_me.NvmeServiceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeServiceCreate has not yet been implemented")
		})
	}
	if api.NvMeNvmeServiceDeleteHandler == nil {
		api.NvMeNvmeServiceDeleteHandler = n_v_me.NvmeServiceDeleteHandlerFunc(func(params n_v_me.NvmeServiceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeServiceDelete has not yet been implemented")
		})
	}
	if api.NvMeNvmeServiceGetHandler == nil {
		api.NvMeNvmeServiceGetHandler = n_v_me.NvmeServiceGetHandlerFunc(func(params n_v_me.NvmeServiceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeServiceGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeServiceModifyHandler == nil {
		api.NvMeNvmeServiceModifyHandler = n_v_me.NvmeServiceModifyHandlerFunc(func(params n_v_me.NvmeServiceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeServiceModify has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemCollectionGetHandler == nil {
		api.NvMeNvmeSubsystemCollectionGetHandler = n_v_me.NvmeSubsystemCollectionGetHandlerFunc(func(params n_v_me.NvmeSubsystemCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemControllerCollectionGetHandler == nil {
		api.NvMeNvmeSubsystemControllerCollectionGetHandler = n_v_me.NvmeSubsystemControllerCollectionGetHandlerFunc(func(params n_v_me.NvmeSubsystemControllerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemControllerCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemControllerGetHandler == nil {
		api.NvMeNvmeSubsystemControllerGetHandler = n_v_me.NvmeSubsystemControllerGetHandlerFunc(func(params n_v_me.NvmeSubsystemControllerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemControllerGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemCreateHandler == nil {
		api.NvMeNvmeSubsystemCreateHandler = n_v_me.NvmeSubsystemCreateHandlerFunc(func(params n_v_me.NvmeSubsystemCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemCreate has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemDeleteHandler == nil {
		api.NvMeNvmeSubsystemDeleteHandler = n_v_me.NvmeSubsystemDeleteHandlerFunc(func(params n_v_me.NvmeSubsystemDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemDelete has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemGetHandler == nil {
		api.NvMeNvmeSubsystemGetHandler = n_v_me.NvmeSubsystemGetHandlerFunc(func(params n_v_me.NvmeSubsystemGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemHostCollectionGetHandler == nil {
		api.NvMeNvmeSubsystemHostCollectionGetHandler = n_v_me.NvmeSubsystemHostCollectionGetHandlerFunc(func(params n_v_me.NvmeSubsystemHostCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemHostCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemHostCreateHandler == nil {
		api.NvMeNvmeSubsystemHostCreateHandler = n_v_me.NvmeSubsystemHostCreateHandlerFunc(func(params n_v_me.NvmeSubsystemHostCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemHostCreate has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemHostDeleteHandler == nil {
		api.NvMeNvmeSubsystemHostDeleteHandler = n_v_me.NvmeSubsystemHostDeleteHandlerFunc(func(params n_v_me.NvmeSubsystemHostDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemHostDelete has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemHostGetHandler == nil {
		api.NvMeNvmeSubsystemHostGetHandler = n_v_me.NvmeSubsystemHostGetHandlerFunc(func(params n_v_me.NvmeSubsystemHostGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemHostGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemMapCollectionGetHandler == nil {
		api.NvMeNvmeSubsystemMapCollectionGetHandler = n_v_me.NvmeSubsystemMapCollectionGetHandlerFunc(func(params n_v_me.NvmeSubsystemMapCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemMapCollectionGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemMapCreateHandler == nil {
		api.NvMeNvmeSubsystemMapCreateHandler = n_v_me.NvmeSubsystemMapCreateHandlerFunc(func(params n_v_me.NvmeSubsystemMapCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemMapCreate has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemMapDeleteHandler == nil {
		api.NvMeNvmeSubsystemMapDeleteHandler = n_v_me.NvmeSubsystemMapDeleteHandlerFunc(func(params n_v_me.NvmeSubsystemMapDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemMapDelete has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemMapGetHandler == nil {
		api.NvMeNvmeSubsystemMapGetHandler = n_v_me.NvmeSubsystemMapGetHandlerFunc(func(params n_v_me.NvmeSubsystemMapGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemMapGet has not yet been implemented")
		})
	}
	if api.NvMeNvmeSubsystemModifyHandler == nil {
		api.NvMeNvmeSubsystemModifyHandler = n_v_me.NvmeSubsystemModifyHandlerFunc(func(params n_v_me.NvmeSubsystemModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.NvmeSubsystemModify has not yet been implemented")
		})
	}
	if api.NetworkingPerformanceFcInterfaceMetricCollectionGetHandler == nil {
		api.NetworkingPerformanceFcInterfaceMetricCollectionGetHandler = networking.PerformanceFcInterfaceMetricCollectionGetHandlerFunc(func(params networking.PerformanceFcInterfaceMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.PerformanceFcInterfaceMetricCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingPerformanceFcInterfaceMetricGetHandler == nil {
		api.NetworkingPerformanceFcInterfaceMetricGetHandler = networking.PerformanceFcInterfaceMetricGetHandlerFunc(func(params networking.PerformanceFcInterfaceMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.PerformanceFcInterfaceMetricGet has not yet been implemented")
		})
	}
	if api.NetworkingPerformanceFcPortMetricCollectionGetHandler == nil {
		api.NetworkingPerformanceFcPortMetricCollectionGetHandler = networking.PerformanceFcPortMetricCollectionGetHandlerFunc(func(params networking.PerformanceFcPortMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.PerformanceFcPortMetricCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingPerformanceFcPortMetricGetHandler == nil {
		api.NetworkingPerformanceFcPortMetricGetHandler = networking.PerformanceFcPortMetricGetHandlerFunc(func(params networking.PerformanceFcPortMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.PerformanceFcPortMetricGet has not yet been implemented")
		})
	}
	if api.SanPerformanceFcpMetricCollectionGetHandler == nil {
		api.SanPerformanceFcpMetricCollectionGetHandler = s_a_n.PerformanceFcpMetricCollectionGetHandlerFunc(func(params s_a_n.PerformanceFcpMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceFcpMetricCollectionGet has not yet been implemented")
		})
	}
	if api.SanPerformanceFcpMetricGetHandler == nil {
		api.SanPerformanceFcpMetricGetHandler = s_a_n.PerformanceFcpMetricGetHandlerFunc(func(params s_a_n.PerformanceFcpMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceFcpMetricGet has not yet been implemented")
		})
	}
	if api.SanPerformanceIscsiMetricCollectionGetHandler == nil {
		api.SanPerformanceIscsiMetricCollectionGetHandler = s_a_n.PerformanceIscsiMetricCollectionGetHandlerFunc(func(params s_a_n.PerformanceIscsiMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceIscsiMetricCollectionGet has not yet been implemented")
		})
	}
	if api.SanPerformanceIscsiMetricGetHandler == nil {
		api.SanPerformanceIscsiMetricGetHandler = s_a_n.PerformanceIscsiMetricGetHandlerFunc(func(params s_a_n.PerformanceIscsiMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceIscsiMetricGet has not yet been implemented")
		})
	}
	if api.SanPerformanceLunMetricCollectionGetHandler == nil {
		api.SanPerformanceLunMetricCollectionGetHandler = s_a_n.PerformanceLunMetricCollectionGetHandlerFunc(func(params s_a_n.PerformanceLunMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceLunMetricCollectionGet has not yet been implemented")
		})
	}
	if api.SanPerformanceLunMetricGetHandler == nil {
		api.SanPerformanceLunMetricGetHandler = s_a_n.PerformanceLunMetricGetHandlerFunc(func(params s_a_n.PerformanceLunMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PerformanceLunMetricGet has not yet been implemented")
		})
	}
	if api.NvMePerformanceNamespaceMetricCollectionGetHandler == nil {
		api.NvMePerformanceNamespaceMetricCollectionGetHandler = n_v_me.PerformanceNamespaceMetricCollectionGetHandlerFunc(func(params n_v_me.PerformanceNamespaceMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.PerformanceNamespaceMetricCollectionGet has not yet been implemented")
		})
	}
	if api.NvMePerformanceNamespaceMetricGetHandler == nil {
		api.NvMePerformanceNamespaceMetricGetHandler = n_v_me.PerformanceNamespaceMetricGetHandlerFunc(func(params n_v_me.PerformanceNamespaceMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.PerformanceNamespaceMetricGet has not yet been implemented")
		})
	}
	if api.NvMePerformanceNvmeMetricCollectionGetHandler == nil {
		api.NvMePerformanceNvmeMetricCollectionGetHandler = n_v_me.PerformanceNvmeMetricCollectionGetHandlerFunc(func(params n_v_me.PerformanceNvmeMetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.PerformanceNvmeMetricCollectionGet has not yet been implemented")
		})
	}
	if api.NvMePerformanceNvmeMetricGetHandler == nil {
		api.NvMePerformanceNvmeMetricGetHandler = n_v_me.PerformanceNvmeMetricGetHandlerFunc(func(params n_v_me.PerformanceNvmeMetricGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_v_me.PerformanceNvmeMetricGet has not yet been implemented")
		})
	}
	if api.ObjectStorePerformanceS3MetricCollectionGetHandler == nil {
		api.ObjectStorePerformanceS3MetricCollectionGetHandler = object_store.PerformanceS3MetricCollectionGetHandlerFunc(func(params object_store.PerformanceS3MetricCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.PerformanceS3MetricCollectionGet has not yet been implemented")
		})
	}
	if api.StoragePlexCollectionGetHandler == nil {
		api.StoragePlexCollectionGetHandler = storage.PlexCollectionGetHandlerFunc(func(params storage.PlexCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.PlexCollectionGet has not yet been implemented")
		})
	}
	if api.StoragePlexGetHandler == nil {
		api.StoragePlexGetHandler = storage.PlexGetHandlerFunc(func(params storage.PlexGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.PlexGet has not yet been implemented")
		})
	}
	if api.NasPoliciesAndRulesToBeAppliedCollectionGetHandler == nil {
		api.NasPoliciesAndRulesToBeAppliedCollectionGetHandler = n_a_s.PoliciesAndRulesToBeAppliedCollectionGetHandlerFunc(func(params n_a_s.PoliciesAndRulesToBeAppliedCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.PoliciesAndRulesToBeAppliedCollectionGet has not yet been implemented")
		})
	}
	if api.NasPoliciesAndRulesToBeAppliedGetHandler == nil {
		api.NasPoliciesAndRulesToBeAppliedGetHandler = n_a_s.PoliciesAndRulesToBeAppliedGetHandlerFunc(func(params n_a_s.PoliciesAndRulesToBeAppliedGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.PoliciesAndRulesToBeAppliedGet has not yet been implemented")
		})
	}
	if api.StoragePortCollectionGetHandler == nil {
		api.StoragePortCollectionGetHandler = storage.PortCollectionGetHandlerFunc(func(params storage.PortCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.PortCollectionGet has not yet been implemented")
		})
	}
	if api.StoragePortGetHandler == nil {
		api.StoragePortGetHandler = storage.PortGetHandlerFunc(func(params storage.PortGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.PortGet has not yet been implemented")
		})
	}
	if api.NetworkingPortMetricsCollectionGetHandler == nil {
		api.NetworkingPortMetricsCollectionGetHandler = networking.PortMetricsCollectionGetHandlerFunc(func(params networking.PortMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.PortMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.SanPortsetCollectionGetHandler == nil {
		api.SanPortsetCollectionGetHandler = s_a_n.PortsetCollectionGetHandlerFunc(func(params s_a_n.PortsetCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetCollectionGet has not yet been implemented")
		})
	}
	if api.SanPortsetCreateHandler == nil {
		api.SanPortsetCreateHandler = s_a_n.PortsetCreateHandlerFunc(func(params s_a_n.PortsetCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetCreate has not yet been implemented")
		})
	}
	if api.SanPortsetDeleteHandler == nil {
		api.SanPortsetDeleteHandler = s_a_n.PortsetDeleteHandlerFunc(func(params s_a_n.PortsetDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetDelete has not yet been implemented")
		})
	}
	if api.SanPortsetGetHandler == nil {
		api.SanPortsetGetHandler = s_a_n.PortsetGetHandlerFunc(func(params s_a_n.PortsetGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetGet has not yet been implemented")
		})
	}
	if api.SanPortsetInterfaceCollectionGetHandler == nil {
		api.SanPortsetInterfaceCollectionGetHandler = s_a_n.PortsetInterfaceCollectionGetHandlerFunc(func(params s_a_n.PortsetInterfaceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetInterfaceCollectionGet has not yet been implemented")
		})
	}
	if api.SanPortsetInterfaceCreateHandler == nil {
		api.SanPortsetInterfaceCreateHandler = s_a_n.PortsetInterfaceCreateHandlerFunc(func(params s_a_n.PortsetInterfaceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetInterfaceCreate has not yet been implemented")
		})
	}
	if api.SanPortsetInterfaceDeleteHandler == nil {
		api.SanPortsetInterfaceDeleteHandler = s_a_n.PortsetInterfaceDeleteHandlerFunc(func(params s_a_n.PortsetInterfaceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetInterfaceDelete has not yet been implemented")
		})
	}
	if api.SanPortsetInterfaceGetHandler == nil {
		api.SanPortsetInterfaceGetHandler = s_a_n.PortsetInterfaceGetHandlerFunc(func(params s_a_n.PortsetInterfaceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.PortsetInterfaceGet has not yet been implemented")
		})
	}
	if api.SecurityPublickeyCollectionGetHandler == nil {
		api.SecurityPublickeyCollectionGetHandler = securityops.PublickeyCollectionGetHandlerFunc(func(params securityops.PublickeyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.PublickeyCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityPublickeyCreateHandler == nil {
		api.SecurityPublickeyCreateHandler = securityops.PublickeyCreateHandlerFunc(func(params securityops.PublickeyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.PublickeyCreate has not yet been implemented")
		})
	}
	if api.SecurityPublickeyGetHandler == nil {
		api.SecurityPublickeyGetHandler = securityops.PublickeyGetHandlerFunc(func(params securityops.PublickeyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.PublickeyGet has not yet been implemented")
		})
	}
	if api.SecurityPublickeyModifyHandler == nil {
		api.SecurityPublickeyModifyHandler = securityops.PublickeyModifyHandlerFunc(func(params securityops.PublickeyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.PublickeyModify has not yet been implemented")
		})
	}
	if api.StorageQosOptionGetHandler == nil {
		api.StorageQosOptionGetHandler = storage.QosOptionGetHandlerFunc(func(params storage.QosOptionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosOptionGet has not yet been implemented")
		})
	}
	if api.StorageQosOptionModifyHandler == nil {
		api.StorageQosOptionModifyHandler = storage.QosOptionModifyHandlerFunc(func(params storage.QosOptionModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosOptionModify has not yet been implemented")
		})
	}
	if api.StorageQosPolicyCollectionGetHandler == nil {
		api.StorageQosPolicyCollectionGetHandler = storage.QosPolicyCollectionGetHandlerFunc(func(params storage.QosPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.StorageQosPolicyCreateHandler == nil {
		api.StorageQosPolicyCreateHandler = storage.QosPolicyCreateHandlerFunc(func(params storage.QosPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosPolicyCreate has not yet been implemented")
		})
	}
	if api.StorageQosPolicyDeleteHandler == nil {
		api.StorageQosPolicyDeleteHandler = storage.QosPolicyDeleteHandlerFunc(func(params storage.QosPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosPolicyDelete has not yet been implemented")
		})
	}
	if api.StorageQosPolicyGetHandler == nil {
		api.StorageQosPolicyGetHandler = storage.QosPolicyGetHandlerFunc(func(params storage.QosPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosPolicyGet has not yet been implemented")
		})
	}
	if api.StorageQosPolicyModifyHandler == nil {
		api.StorageQosPolicyModifyHandler = storage.QosPolicyModifyHandlerFunc(func(params storage.QosPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosPolicyModify has not yet been implemented")
		})
	}
	if api.StorageQosWorkloadCollectionGetHandler == nil {
		api.StorageQosWorkloadCollectionGetHandler = storage.QosWorkloadCollectionGetHandlerFunc(func(params storage.QosWorkloadCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosWorkloadCollectionGet has not yet been implemented")
		})
	}
	if api.StorageQosWorkloadGetHandler == nil {
		api.StorageQosWorkloadGetHandler = storage.QosWorkloadGetHandlerFunc(func(params storage.QosWorkloadGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QosWorkloadGet has not yet been implemented")
		})
	}
	if api.StorageQtreeCollectionGetHandler == nil {
		api.StorageQtreeCollectionGetHandler = storage.QtreeCollectionGetHandlerFunc(func(params storage.QtreeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QtreeCollectionGet has not yet been implemented")
		})
	}
	if api.StorageQtreeCreateHandler == nil {
		api.StorageQtreeCreateHandler = storage.QtreeCreateHandlerFunc(func(params storage.QtreeCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QtreeCreate has not yet been implemented")
		})
	}
	if api.StorageQtreeDeleteHandler == nil {
		api.StorageQtreeDeleteHandler = storage.QtreeDeleteHandlerFunc(func(params storage.QtreeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QtreeDelete has not yet been implemented")
		})
	}
	if api.StorageQtreeGetHandler == nil {
		api.StorageQtreeGetHandler = storage.QtreeGetHandlerFunc(func(params storage.QtreeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QtreeGet has not yet been implemented")
		})
	}
	if api.StorageQtreeModifyHandler == nil {
		api.StorageQtreeModifyHandler = storage.QtreeModifyHandlerFunc(func(params storage.QtreeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QtreeModify has not yet been implemented")
		})
	}
	if api.StorageQuotaReportCollectionGetHandler == nil {
		api.StorageQuotaReportCollectionGetHandler = storage.QuotaReportCollectionGetHandlerFunc(func(params storage.QuotaReportCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaReportCollectionGet has not yet been implemented")
		})
	}
	if api.StorageQuotaReportGetHandler == nil {
		api.StorageQuotaReportGetHandler = storage.QuotaReportGetHandlerFunc(func(params storage.QuotaReportGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaReportGet has not yet been implemented")
		})
	}
	if api.StorageQuotaRuleCollectionGetHandler == nil {
		api.StorageQuotaRuleCollectionGetHandler = storage.QuotaRuleCollectionGetHandlerFunc(func(params storage.QuotaRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaRuleCollectionGet has not yet been implemented")
		})
	}
	if api.StorageQuotaRuleCreateHandler == nil {
		api.StorageQuotaRuleCreateHandler = storage.QuotaRuleCreateHandlerFunc(func(params storage.QuotaRuleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaRuleCreate has not yet been implemented")
		})
	}
	if api.StorageQuotaRuleDeleteHandler == nil {
		api.StorageQuotaRuleDeleteHandler = storage.QuotaRuleDeleteHandlerFunc(func(params storage.QuotaRuleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaRuleDelete has not yet been implemented")
		})
	}
	if api.StorageQuotaRuleGetHandler == nil {
		api.StorageQuotaRuleGetHandler = storage.QuotaRuleGetHandlerFunc(func(params storage.QuotaRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaRuleGet has not yet been implemented")
		})
	}
	if api.StorageQuotaRuleModifyHandler == nil {
		api.StorageQuotaRuleModifyHandler = storage.QuotaRuleModifyHandlerFunc(func(params storage.QuotaRuleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.QuotaRuleModify has not yet been implemented")
		})
	}
	if api.ClusterResourceTagCollectionGetHandler == nil {
		api.ClusterResourceTagCollectionGetHandler = cluster.ResourceTagCollectionGetHandlerFunc(func(params cluster.ResourceTagCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterResourceTagGetHandler == nil {
		api.ClusterResourceTagGetHandler = cluster.ResourceTagGetHandlerFunc(func(params cluster.ResourceTagGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagGet has not yet been implemented")
		})
	}
	if api.ClusterResourceTagResourceCollectionGetHandler == nil {
		api.ClusterResourceTagResourceCollectionGetHandler = cluster.ResourceTagResourceCollectionGetHandlerFunc(func(params cluster.ResourceTagResourceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagResourceCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterResourceTagResourceCreateHandler == nil {
		api.ClusterResourceTagResourceCreateHandler = cluster.ResourceTagResourceCreateHandlerFunc(func(params cluster.ResourceTagResourceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagResourceCreate has not yet been implemented")
		})
	}
	if api.ClusterResourceTagResourceDeleteHandler == nil {
		api.ClusterResourceTagResourceDeleteHandler = cluster.ResourceTagResourceDeleteHandlerFunc(func(params cluster.ResourceTagResourceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagResourceDelete has not yet been implemented")
		})
	}
	if api.ClusterResourceTagResourceGetHandler == nil {
		api.ClusterResourceTagResourceGetHandler = cluster.ResourceTagResourceGetHandlerFunc(func(params cluster.ResourceTagResourceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ResourceTagResourceGet has not yet been implemented")
		})
	}
	if api.SecurityRoleCollectionGetHandler == nil {
		api.SecurityRoleCollectionGetHandler = securityops.RoleCollectionGetHandlerFunc(func(params securityops.RoleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RoleCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityRoleCreateHandler == nil {
		api.SecurityRoleCreateHandler = securityops.RoleCreateHandlerFunc(func(params securityops.RoleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RoleCreate has not yet been implemented")
		})
	}
	if api.SecurityRoleDeleteHandler == nil {
		api.SecurityRoleDeleteHandler = securityops.RoleDeleteHandlerFunc(func(params securityops.RoleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RoleDelete has not yet been implemented")
		})
	}
	if api.SecurityRoleGetHandler == nil {
		api.SecurityRoleGetHandler = securityops.RoleGetHandlerFunc(func(params securityops.RoleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RoleGet has not yet been implemented")
		})
	}
	if api.SecurityRolePrivilegeCollectionGetHandler == nil {
		api.SecurityRolePrivilegeCollectionGetHandler = securityops.RolePrivilegeCollectionGetHandlerFunc(func(params securityops.RolePrivilegeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RolePrivilegeCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityRolePrivilegeCreateHandler == nil {
		api.SecurityRolePrivilegeCreateHandler = securityops.RolePrivilegeCreateHandlerFunc(func(params securityops.RolePrivilegeCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RolePrivilegeCreate has not yet been implemented")
		})
	}
	if api.SecurityRolePrivilegeDeleteHandler == nil {
		api.SecurityRolePrivilegeDeleteHandler = securityops.RolePrivilegeDeleteHandlerFunc(func(params securityops.RolePrivilegeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RolePrivilegeDelete has not yet been implemented")
		})
	}
	if api.SecurityRolePrivilegeGetHandler == nil {
		api.SecurityRolePrivilegeGetHandler = securityops.RolePrivilegeGetHandlerFunc(func(params securityops.RolePrivilegeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RolePrivilegeGet has not yet been implemented")
		})
	}
	if api.SecurityRolePrivilegeModifyHandler == nil {
		api.SecurityRolePrivilegeModifyHandler = securityops.RolePrivilegeModifyHandlerFunc(func(params securityops.RolePrivilegeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.RolePrivilegeModify has not yet been implemented")
		})
	}
	if api.NasS3AuditCreateHandler == nil {
		api.NasS3AuditCreateHandler = n_a_s.S3AuditCreateHandlerFunc(func(params n_a_s.S3AuditCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.S3AuditCreate has not yet been implemented")
		})
	}
	if api.NasS3AuditDeleteHandler == nil {
		api.NasS3AuditDeleteHandler = n_a_s.S3AuditDeleteHandlerFunc(func(params n_a_s.S3AuditDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.S3AuditDelete has not yet been implemented")
		})
	}
	if api.NasS3AuditGetHandler == nil {
		api.NasS3AuditGetHandler = n_a_s.S3AuditGetHandlerFunc(func(params n_a_s.S3AuditGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.S3AuditGet has not yet been implemented")
		})
	}
	if api.NasS3AuditModifyHandler == nil {
		api.NasS3AuditModifyHandler = n_a_s.S3AuditModifyHandlerFunc(func(params n_a_s.S3AuditModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.S3AuditModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketCreateHandler == nil {
		api.ObjectStoreS3BucketCreateHandler = object_store.S3BucketCreateHandlerFunc(func(params object_store.S3BucketCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketDeleteHandler == nil {
		api.ObjectStoreS3BucketDeleteHandler = object_store.S3BucketDeleteHandlerFunc(func(params object_store.S3BucketDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketGetHandler == nil {
		api.ObjectStoreS3BucketGetHandler = object_store.S3BucketGetHandlerFunc(func(params object_store.S3BucketGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketLifecycleRuleCollectionGetHandler == nil {
		api.ObjectStoreS3BucketLifecycleRuleCollectionGetHandler = object_store.S3BucketLifecycleRuleCollectionGetHandlerFunc(func(params object_store.S3BucketLifecycleRuleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketLifecycleRuleCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketLifecycleRuleCreateHandler == nil {
		api.ObjectStoreS3BucketLifecycleRuleCreateHandler = object_store.S3BucketLifecycleRuleCreateHandlerFunc(func(params object_store.S3BucketLifecycleRuleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketLifecycleRuleCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketLifecycleRuleDeleteHandler == nil {
		api.ObjectStoreS3BucketLifecycleRuleDeleteHandler = object_store.S3BucketLifecycleRuleDeleteHandlerFunc(func(params object_store.S3BucketLifecycleRuleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketLifecycleRuleDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketLifecycleRuleGetHandler == nil {
		api.ObjectStoreS3BucketLifecycleRuleGetHandler = object_store.S3BucketLifecycleRuleGetHandlerFunc(func(params object_store.S3BucketLifecycleRuleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketLifecycleRuleGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketLifecycleRuleModifyHandler == nil {
		api.ObjectStoreS3BucketLifecycleRuleModifyHandler = object_store.S3BucketLifecycleRuleModifyHandlerFunc(func(params object_store.S3BucketLifecycleRuleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketLifecycleRuleModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketModifyHandler == nil {
		api.ObjectStoreS3BucketModifyHandler = object_store.S3BucketModifyHandlerFunc(func(params object_store.S3BucketModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketSvmCreateHandler == nil {
		api.ObjectStoreS3BucketSvmCreateHandler = object_store.S3BucketSvmCreateHandlerFunc(func(params object_store.S3BucketSvmCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketSvmCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketSvmDeleteHandler == nil {
		api.ObjectStoreS3BucketSvmDeleteHandler = object_store.S3BucketSvmDeleteHandlerFunc(func(params object_store.S3BucketSvmDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketSvmDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketSvmGetHandler == nil {
		api.ObjectStoreS3BucketSvmGetHandler = object_store.S3BucketSvmGetHandlerFunc(func(params object_store.S3BucketSvmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketSvmGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3BucketSvmModifyHandler == nil {
		api.ObjectStoreS3BucketSvmModifyHandler = object_store.S3BucketSvmModifyHandlerFunc(func(params object_store.S3BucketSvmModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3BucketSvmModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3GroupCollectionGetHandler == nil {
		api.ObjectStoreS3GroupCollectionGetHandler = object_store.S3GroupCollectionGetHandlerFunc(func(params object_store.S3GroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3GroupCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3GroupCreateHandler == nil {
		api.ObjectStoreS3GroupCreateHandler = object_store.S3GroupCreateHandlerFunc(func(params object_store.S3GroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3GroupCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3GroupDeleteHandler == nil {
		api.ObjectStoreS3GroupDeleteHandler = object_store.S3GroupDeleteHandlerFunc(func(params object_store.S3GroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3GroupDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3GroupGetHandler == nil {
		api.ObjectStoreS3GroupGetHandler = object_store.S3GroupGetHandlerFunc(func(params object_store.S3GroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3GroupGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3GroupModifyHandler == nil {
		api.ObjectStoreS3GroupModifyHandler = object_store.S3GroupModifyHandlerFunc(func(params object_store.S3GroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3GroupModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3PolicyCollectionGetHandler == nil {
		api.ObjectStoreS3PolicyCollectionGetHandler = object_store.S3PolicyCollectionGetHandlerFunc(func(params object_store.S3PolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3PolicyCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3PolicyCreateHandler == nil {
		api.ObjectStoreS3PolicyCreateHandler = object_store.S3PolicyCreateHandlerFunc(func(params object_store.S3PolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3PolicyCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3PolicyDeleteHandler == nil {
		api.ObjectStoreS3PolicyDeleteHandler = object_store.S3PolicyDeleteHandlerFunc(func(params object_store.S3PolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3PolicyDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3PolicyGetHandler == nil {
		api.ObjectStoreS3PolicyGetHandler = object_store.S3PolicyGetHandlerFunc(func(params object_store.S3PolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3PolicyGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3PolicyModifyHandler == nil {
		api.ObjectStoreS3PolicyModifyHandler = object_store.S3PolicyModifyHandlerFunc(func(params object_store.S3PolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3PolicyModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3ServiceCollectionGetHandler == nil {
		api.ObjectStoreS3ServiceCollectionGetHandler = object_store.S3ServiceCollectionGetHandlerFunc(func(params object_store.S3ServiceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3ServiceCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3ServiceCreateHandler == nil {
		api.ObjectStoreS3ServiceCreateHandler = object_store.S3ServiceCreateHandlerFunc(func(params object_store.S3ServiceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3ServiceCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3ServiceDeleteHandler == nil {
		api.ObjectStoreS3ServiceDeleteHandler = object_store.S3ServiceDeleteHandlerFunc(func(params object_store.S3ServiceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3ServiceDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3ServiceGetHandler == nil {
		api.ObjectStoreS3ServiceGetHandler = object_store.S3ServiceGetHandlerFunc(func(params object_store.S3ServiceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3ServiceGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3ServiceModifyHandler == nil {
		api.ObjectStoreS3ServiceModifyHandler = object_store.S3ServiceModifyHandlerFunc(func(params object_store.S3ServiceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3ServiceModify has not yet been implemented")
		})
	}
	if api.ObjectStoreS3UserCollectionGetHandler == nil {
		api.ObjectStoreS3UserCollectionGetHandler = object_store.S3UserCollectionGetHandlerFunc(func(params object_store.S3UserCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3UserCollectionGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3UserCreateHandler == nil {
		api.ObjectStoreS3UserCreateHandler = object_store.S3UserCreateHandlerFunc(func(params object_store.S3UserCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3UserCreate has not yet been implemented")
		})
	}
	if api.ObjectStoreS3UserDeleteHandler == nil {
		api.ObjectStoreS3UserDeleteHandler = object_store.S3UserDeleteHandlerFunc(func(params object_store.S3UserDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3UserDelete has not yet been implemented")
		})
	}
	if api.ObjectStoreS3UserGetHandler == nil {
		api.ObjectStoreS3UserGetHandler = object_store.S3UserGetHandlerFunc(func(params object_store.S3UserGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3UserGet has not yet been implemented")
		})
	}
	if api.ObjectStoreS3UserModifyHandler == nil {
		api.ObjectStoreS3UserModifyHandler = object_store.S3UserModifyHandlerFunc(func(params object_store.S3UserModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation object_store.S3UserModify has not yet been implemented")
		})
	}
	if api.ClusterScheduleCollectionGetHandler == nil {
		api.ClusterScheduleCollectionGetHandler = cluster.ScheduleCollectionGetHandlerFunc(func(params cluster.ScheduleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ScheduleCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterScheduleCreateHandler == nil {
		api.ClusterScheduleCreateHandler = cluster.ScheduleCreateHandlerFunc(func(params cluster.ScheduleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ScheduleCreate has not yet been implemented")
		})
	}
	if api.ClusterScheduleDeleteHandler == nil {
		api.ClusterScheduleDeleteHandler = cluster.ScheduleDeleteHandlerFunc(func(params cluster.ScheduleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ScheduleDelete has not yet been implemented")
		})
	}
	if api.ClusterScheduleGetHandler == nil {
		api.ClusterScheduleGetHandler = cluster.ScheduleGetHandlerFunc(func(params cluster.ScheduleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ScheduleGet has not yet been implemented")
		})
	}
	if api.ClusterScheduleModifyHandler == nil {
		api.ClusterScheduleModifyHandler = cluster.ScheduleModifyHandlerFunc(func(params cluster.ScheduleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.ScheduleModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityAssociationCollectionGetHandler == nil {
		api.SecuritySecurityAssociationCollectionGetHandler = securityops.SecurityAssociationCollectionGetHandlerFunc(func(params securityops.SecurityAssociationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityAssociationCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityAssociationGetHandler == nil {
		api.SecuritySecurityAssociationGetHandler = securityops.SecurityAssociationGetHandlerFunc(func(params securityops.SecurityAssociationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityAssociationGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityAuditGetHandler == nil {
		api.SecuritySecurityAuditGetHandler = securityops.SecurityAuditGetHandlerFunc(func(params securityops.SecurityAuditGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityAuditGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityAuditLogCollectionGetHandler == nil {
		api.SecuritySecurityAuditLogCollectionGetHandler = securityops.SecurityAuditLogCollectionGetHandlerFunc(func(params securityops.SecurityAuditLogCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityAuditLogCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityAuditModifyHandler == nil {
		api.SecuritySecurityAuditModifyHandler = securityops.SecurityAuditModifyHandlerFunc(func(params securityops.SecurityAuditModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityAuditModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityCertificateCollectionGetHandler == nil {
		api.SecuritySecurityCertificateCollectionGetHandler = securityops.SecurityCertificateCollectionGetHandlerFunc(func(params securityops.SecurityCertificateCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityCertificateCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityCertificateCreateHandler == nil {
		api.SecuritySecurityCertificateCreateHandler = securityops.SecurityCertificateCreateHandlerFunc(func(params securityops.SecurityCertificateCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityCertificateCreate has not yet been implemented")
		})
	}
	if api.SecuritySecurityCertificateDeleteHandler == nil {
		api.SecuritySecurityCertificateDeleteHandler = securityops.SecurityCertificateDeleteHandlerFunc(func(params securityops.SecurityCertificateDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityCertificateDelete has not yet been implemented")
		})
	}
	if api.SecuritySecurityCertificateGetHandler == nil {
		api.SecuritySecurityCertificateGetHandler = securityops.SecurityCertificateGetHandlerFunc(func(params securityops.SecurityCertificateGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityCertificateGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityCertificateSignHandler == nil {
		api.SecuritySecurityCertificateSignHandler = securityops.SecurityCertificateSignHandlerFunc(func(params securityops.SecurityCertificateSignParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityCertificateSign has not yet been implemented")
		})
	}
	if api.SecuritySecurityConfigGetHandler == nil {
		api.SecuritySecurityConfigGetHandler = securityops.SecurityConfigGetHandlerFunc(func(params securityops.SecurityConfigGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityConfigGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityConfigModifyHandler == nil {
		api.SecuritySecurityConfigModifyHandler = securityops.SecurityConfigModifyHandlerFunc(func(params securityops.SecurityConfigModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityConfigModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerCollectionGetHandler == nil {
		api.SecuritySecurityKeyManagerCollectionGetHandler = securityops.SecurityKeyManagerCollectionGetHandlerFunc(func(params securityops.SecurityKeyManagerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerCreateHandler == nil {
		api.SecuritySecurityKeyManagerCreateHandler = securityops.SecurityKeyManagerCreateHandlerFunc(func(params securityops.SecurityKeyManagerCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerCreate has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerDeleteHandler == nil {
		api.SecuritySecurityKeyManagerDeleteHandler = securityops.SecurityKeyManagerDeleteHandlerFunc(func(params securityops.SecurityKeyManagerDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerDelete has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerGetHandler == nil {
		api.SecuritySecurityKeyManagerGetHandler = securityops.SecurityKeyManagerGetHandlerFunc(func(params securityops.SecurityKeyManagerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerKeyServersCollectionGetHandler == nil {
		api.SecuritySecurityKeyManagerKeyServersCollectionGetHandler = securityops.SecurityKeyManagerKeyServersCollectionGetHandlerFunc(func(params securityops.SecurityKeyManagerKeyServersCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerKeyServersCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerKeyServersCreateHandler == nil {
		api.SecuritySecurityKeyManagerKeyServersCreateHandler = securityops.SecurityKeyManagerKeyServersCreateHandlerFunc(func(params securityops.SecurityKeyManagerKeyServersCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerKeyServersCreate has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerKeyServersDeleteHandler == nil {
		api.SecuritySecurityKeyManagerKeyServersDeleteHandler = securityops.SecurityKeyManagerKeyServersDeleteHandlerFunc(func(params securityops.SecurityKeyManagerKeyServersDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerKeyServersDelete has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerKeyServersGetHandler == nil {
		api.SecuritySecurityKeyManagerKeyServersGetHandler = securityops.SecurityKeyManagerKeyServersGetHandlerFunc(func(params securityops.SecurityKeyManagerKeyServersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerKeyServersGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerKeyServersModifyHandler == nil {
		api.SecuritySecurityKeyManagerKeyServersModifyHandler = securityops.SecurityKeyManagerKeyServersModifyHandlerFunc(func(params securityops.SecurityKeyManagerKeyServersModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerKeyServersModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerMigrateHandler == nil {
		api.SecuritySecurityKeyManagerMigrateHandler = securityops.SecurityKeyManagerMigrateHandlerFunc(func(params securityops.SecurityKeyManagerMigrateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerMigrate has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerModifyHandler == nil {
		api.SecuritySecurityKeyManagerModifyHandler = securityops.SecurityKeyManagerModifyHandlerFunc(func(params securityops.SecurityKeyManagerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeyManagerRestoreHandler == nil {
		api.SecuritySecurityKeyManagerRestoreHandler = securityops.SecurityKeyManagerRestoreHandlerFunc(func(params securityops.SecurityKeyManagerRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeyManagerRestore has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeystoreCollectionGetHandler == nil {
		api.SecuritySecurityKeystoreCollectionGetHandler = securityops.SecurityKeystoreCollectionGetHandlerFunc(func(params securityops.SecurityKeystoreCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeystoreCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeystoreDeleteHandler == nil {
		api.SecuritySecurityKeystoreDeleteHandler = securityops.SecurityKeystoreDeleteHandlerFunc(func(params securityops.SecurityKeystoreDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeystoreDelete has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeystoreGetHandler == nil {
		api.SecuritySecurityKeystoreGetHandler = securityops.SecurityKeystoreGetHandlerFunc(func(params securityops.SecurityKeystoreGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeystoreGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityKeystoreModifyHandler == nil {
		api.SecuritySecurityKeystoreModifyHandler = securityops.SecurityKeystoreModifyHandlerFunc(func(params securityops.SecurityKeystoreModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityKeystoreModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityLogForwardingCreateHandler == nil {
		api.SecuritySecurityLogForwardingCreateHandler = securityops.SecurityLogForwardingCreateHandlerFunc(func(params securityops.SecurityLogForwardingCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityLogForwardingCreate has not yet been implemented")
		})
	}
	if api.SecuritySecurityLogForwardingDeleteHandler == nil {
		api.SecuritySecurityLogForwardingDeleteHandler = securityops.SecurityLogForwardingDeleteHandlerFunc(func(params securityops.SecurityLogForwardingDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityLogForwardingDelete has not yet been implemented")
		})
	}
	if api.SecuritySecurityLogForwardingGetHandler == nil {
		api.SecuritySecurityLogForwardingGetHandler = securityops.SecurityLogForwardingGetHandlerFunc(func(params securityops.SecurityLogForwardingGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityLogForwardingGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityLogForwardingModifyHandler == nil {
		api.SecuritySecurityLogForwardingModifyHandler = securityops.SecurityLogForwardingModifyHandlerFunc(func(params securityops.SecurityLogForwardingModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityLogForwardingModify has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2CollectionGetHandler == nil {
		api.SecuritySecurityOauth2CollectionGetHandler = securityops.SecurityOauth2CollectionGetHandlerFunc(func(params securityops.SecurityOauth2CollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2CollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2CreateHandler == nil {
		api.SecuritySecurityOauth2CreateHandler = securityops.SecurityOauth2CreateHandlerFunc(func(params securityops.SecurityOauth2CreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2Create has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2DeleteHandler == nil {
		api.SecuritySecurityOauth2DeleteHandler = securityops.SecurityOauth2DeleteHandlerFunc(func(params securityops.SecurityOauth2DeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2Delete has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2GetHandler == nil {
		api.SecuritySecurityOauth2GetHandler = securityops.SecurityOauth2GetHandlerFunc(func(params securityops.SecurityOauth2GetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2Get has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2GlobalGetHandler == nil {
		api.SecuritySecurityOauth2GlobalGetHandler = securityops.SecurityOauth2GlobalGetHandlerFunc(func(params securityops.SecurityOauth2GlobalGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2GlobalGet has not yet been implemented")
		})
	}
	if api.SecuritySecurityOauth2GlobalModifyHandler == nil {
		api.SecuritySecurityOauth2GlobalModifyHandler = securityops.SecurityOauth2GlobalModifyHandlerFunc(func(params securityops.SecurityOauth2GlobalModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecurityOauth2GlobalModify has not yet been implemented")
		})
	}
	if api.SecuritySecuritySamlSpCreateHandler == nil {
		api.SecuritySecuritySamlSpCreateHandler = securityops.SecuritySamlSpCreateHandlerFunc(func(params securityops.SecuritySamlSpCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecuritySamlSpCreate has not yet been implemented")
		})
	}
	if api.SecuritySecuritySamlSpDeleteHandler == nil {
		api.SecuritySecuritySamlSpDeleteHandler = securityops.SecuritySamlSpDeleteHandlerFunc(func(params securityops.SecuritySamlSpDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecuritySamlSpDelete has not yet been implemented")
		})
	}
	if api.SecuritySecuritySamlSpGetHandler == nil {
		api.SecuritySecuritySamlSpGetHandler = securityops.SecuritySamlSpGetHandlerFunc(func(params securityops.SecuritySamlSpGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecuritySamlSpGet has not yet been implemented")
		})
	}
	if api.SecuritySecuritySamlSpModifyHandler == nil {
		api.SecuritySecuritySamlSpModifyHandler = securityops.SecuritySamlSpModifyHandlerFunc(func(params securityops.SecuritySamlSpModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SecuritySamlSpModify has not yet been implemented")
		})
	}
	if api.ClusterSensorsCollectionGetHandler == nil {
		api.ClusterSensorsCollectionGetHandler = cluster.SensorsCollectionGetHandlerFunc(func(params cluster.SensorsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SensorsCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterSensorsGetHandler == nil {
		api.ClusterSensorsGetHandler = cluster.SensorsGetHandlerFunc(func(params cluster.SensorsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SensorsGet has not yet been implemented")
		})
	}
	if api.NasShadowcopyCollectionGetHandler == nil {
		api.NasShadowcopyCollectionGetHandler = n_a_s.ShadowcopyCollectionGetHandlerFunc(func(params n_a_s.ShadowcopyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopyCollectionGet has not yet been implemented")
		})
	}
	if api.NasShadowcopyGetHandler == nil {
		api.NasShadowcopyGetHandler = n_a_s.ShadowcopyGetHandlerFunc(func(params n_a_s.ShadowcopyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopyGet has not yet been implemented")
		})
	}
	if api.NasShadowcopyModifyHandler == nil {
		api.NasShadowcopyModifyHandler = n_a_s.ShadowcopyModifyHandlerFunc(func(params n_a_s.ShadowcopyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopyModify has not yet been implemented")
		})
	}
	if api.NasShadowcopySetCollectionGetHandler == nil {
		api.NasShadowcopySetCollectionGetHandler = n_a_s.ShadowcopySetCollectionGetHandlerFunc(func(params n_a_s.ShadowcopySetCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopySetCollectionGet has not yet been implemented")
		})
	}
	if api.NasShadowcopySetGetHandler == nil {
		api.NasShadowcopySetGetHandler = n_a_s.ShadowcopySetGetHandlerFunc(func(params n_a_s.ShadowcopySetGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopySetGet has not yet been implemented")
		})
	}
	if api.NasShadowcopySetModifyHandler == nil {
		api.NasShadowcopySetModifyHandler = n_a_s.ShadowcopySetModifyHandlerFunc(func(params n_a_s.ShadowcopySetModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.ShadowcopySetModify has not yet been implemented")
		})
	}
	if api.StorageShelfCollectionGetHandler == nil {
		api.StorageShelfCollectionGetHandler = storage.ShelfCollectionGetHandlerFunc(func(params storage.ShelfCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.ShelfCollectionGet has not yet been implemented")
		})
	}
	if api.StorageShelfGetHandler == nil {
		api.StorageShelfGetHandler = storage.ShelfGetHandlerFunc(func(params storage.ShelfGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.ShelfGet has not yet been implemented")
		})
	}
	if api.StorageShelfModifyHandler == nil {
		api.StorageShelfModifyHandler = storage.ShelfModifyHandlerFunc(func(params storage.ShelfModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.ShelfModify has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockComplianceClockCollectionGetHandler == nil {
		api.SnapLockSnaplockComplianceClockCollectionGetHandler = snap_lock.SnaplockComplianceClockCollectionGetHandlerFunc(func(params snap_lock.SnaplockComplianceClockCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockComplianceClockCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockComplianceClockCreateHandler == nil {
		api.SnapLockSnaplockComplianceClockCreateHandler = snap_lock.SnaplockComplianceClockCreateHandlerFunc(func(params snap_lock.SnaplockComplianceClockCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockComplianceClockCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockComplianceClockGetHandler == nil {
		api.SnapLockSnaplockComplianceClockGetHandler = snap_lock.SnaplockComplianceClockGetHandlerFunc(func(params snap_lock.SnaplockComplianceClockGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockComplianceClockGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFilePrivilegedDeleteHandler == nil {
		api.SnapLockSnaplockFilePrivilegedDeleteHandler = snap_lock.SnaplockFilePrivilegedDeleteHandlerFunc(func(params snap_lock.SnaplockFilePrivilegedDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFilePrivilegedDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFileRetentionGetHandler == nil {
		api.SnapLockSnaplockFileRetentionGetHandler = snap_lock.SnaplockFileRetentionGetHandlerFunc(func(params snap_lock.SnaplockFileRetentionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFileRetentionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFileRetentionTimeModifyHandler == nil {
		api.SnapLockSnaplockFileRetentionTimeModifyHandler = snap_lock.SnaplockFileRetentionTimeModifyHandlerFunc(func(params snap_lock.SnaplockFileRetentionTimeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFileRetentionTimeModify has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFingerprintOperationCollectionGetHandler == nil {
		api.SnapLockSnaplockFingerprintOperationCollectionGetHandler = snap_lock.SnaplockFingerprintOperationCollectionGetHandlerFunc(func(params snap_lock.SnaplockFingerprintOperationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFingerprintOperationCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFingerprintOperationCreateHandler == nil {
		api.SnapLockSnaplockFingerprintOperationCreateHandler = snap_lock.SnaplockFingerprintOperationCreateHandlerFunc(func(params snap_lock.SnaplockFingerprintOperationCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFingerprintOperationCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFingerprintOperationDeleteHandler == nil {
		api.SnapLockSnaplockFingerprintOperationDeleteHandler = snap_lock.SnaplockFingerprintOperationDeleteHandlerFunc(func(params snap_lock.SnaplockFingerprintOperationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFingerprintOperationDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockFingerprintOperationGetHandler == nil {
		api.SnapLockSnaplockFingerprintOperationGetHandler = snap_lock.SnaplockFingerprintOperationGetHandlerFunc(func(params snap_lock.SnaplockFingerprintOperationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockFingerprintOperationGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldBeginHandler == nil {
		api.SnapLockSnaplockLegalHoldBeginHandler = snap_lock.SnaplockLegalHoldBeginHandlerFunc(func(params snap_lock.SnaplockLegalHoldBeginParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldBegin has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldCollectionGetHandler == nil {
		api.SnapLockSnaplockLegalHoldCollectionGetHandler = snap_lock.SnaplockLegalHoldCollectionGetHandlerFunc(func(params snap_lock.SnaplockLegalHoldCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldCreateHandler == nil {
		api.SnapLockSnaplockLegalHoldCreateHandler = snap_lock.SnaplockLegalHoldCreateHandlerFunc(func(params snap_lock.SnaplockLegalHoldCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldDeleteHandler == nil {
		api.SnapLockSnaplockLegalHoldDeleteHandler = snap_lock.SnaplockLegalHoldDeleteHandlerFunc(func(params snap_lock.SnaplockLegalHoldDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldFilesGetHandler == nil {
		api.SnapLockSnaplockLegalHoldFilesGetHandler = snap_lock.SnaplockLegalHoldFilesGetHandlerFunc(func(params snap_lock.SnaplockLegalHoldFilesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldFilesGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldGetHandler == nil {
		api.SnapLockSnaplockLegalHoldGetHandler = snap_lock.SnaplockLegalHoldGetHandlerFunc(func(params snap_lock.SnaplockLegalHoldGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldInstanceGetHandler == nil {
		api.SnapLockSnaplockLegalHoldInstanceGetHandler = snap_lock.SnaplockLegalHoldInstanceGetHandlerFunc(func(params snap_lock.SnaplockLegalHoldInstanceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldInstanceGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLegalHoldOperationDeleteHandler == nil {
		api.SnapLockSnaplockLegalHoldOperationDeleteHandler = snap_lock.SnaplockLegalHoldOperationDeleteHandlerFunc(func(params snap_lock.SnaplockLegalHoldOperationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLegalHoldOperationDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLogCollectionGetHandler == nil {
		api.SnapLockSnaplockLogCollectionGetHandler = snap_lock.SnaplockLogCollectionGetHandlerFunc(func(params snap_lock.SnaplockLogCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLogCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLogCreateHandler == nil {
		api.SnapLockSnaplockLogCreateHandler = snap_lock.SnaplockLogCreateHandlerFunc(func(params snap_lock.SnaplockLogCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLogCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLogDeleteHandler == nil {
		api.SnapLockSnaplockLogDeleteHandler = snap_lock.SnaplockLogDeleteHandlerFunc(func(params snap_lock.SnaplockLogDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLogDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLogGetHandler == nil {
		api.SnapLockSnaplockLogGetHandler = snap_lock.SnaplockLogGetHandlerFunc(func(params snap_lock.SnaplockLogGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLogGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockLogModifyHandler == nil {
		api.SnapLockSnaplockLogModifyHandler = snap_lock.SnaplockLogModifyHandlerFunc(func(params snap_lock.SnaplockLogModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockLogModify has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionOperationCollectionGetHandler == nil {
		api.SnapLockSnaplockRetentionOperationCollectionGetHandler = snap_lock.SnaplockRetentionOperationCollectionGetHandlerFunc(func(params snap_lock.SnaplockRetentionOperationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionOperationCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionOperationCreateHandler == nil {
		api.SnapLockSnaplockRetentionOperationCreateHandler = snap_lock.SnaplockRetentionOperationCreateHandlerFunc(func(params snap_lock.SnaplockRetentionOperationCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionOperationCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionOperationDeleteHandler == nil {
		api.SnapLockSnaplockRetentionOperationDeleteHandler = snap_lock.SnaplockRetentionOperationDeleteHandlerFunc(func(params snap_lock.SnaplockRetentionOperationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionOperationDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionOperationGetHandler == nil {
		api.SnapLockSnaplockRetentionOperationGetHandler = snap_lock.SnaplockRetentionOperationGetHandlerFunc(func(params snap_lock.SnaplockRetentionOperationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionOperationGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionPolicyCollectionGetHandler == nil {
		api.SnapLockSnaplockRetentionPolicyCollectionGetHandler = snap_lock.SnaplockRetentionPolicyCollectionGetHandlerFunc(func(params snap_lock.SnaplockRetentionPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionPolicyCreateHandler == nil {
		api.SnapLockSnaplockRetentionPolicyCreateHandler = snap_lock.SnaplockRetentionPolicyCreateHandlerFunc(func(params snap_lock.SnaplockRetentionPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionPolicyCreate has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionPolicyDeleteHandler == nil {
		api.SnapLockSnaplockRetentionPolicyDeleteHandler = snap_lock.SnaplockRetentionPolicyDeleteHandlerFunc(func(params snap_lock.SnaplockRetentionPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionPolicyDelete has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionPolicyGetHandler == nil {
		api.SnapLockSnaplockRetentionPolicyGetHandler = snap_lock.SnaplockRetentionPolicyGetHandlerFunc(func(params snap_lock.SnaplockRetentionPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionPolicyGet has not yet been implemented")
		})
	}
	if api.SnapLockSnaplockRetentionPolicyModifyHandler == nil {
		api.SnapLockSnaplockRetentionPolicyModifyHandler = snap_lock.SnaplockRetentionPolicyModifyHandlerFunc(func(params snap_lock.SnaplockRetentionPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_lock.SnaplockRetentionPolicyModify has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorPoliciesGetHandler == nil {
		api.SnapMirrorSnapmirrorPoliciesGetHandler = snap_mirror.SnapmirrorPoliciesGetHandlerFunc(func(params snap_mirror.SnapmirrorPoliciesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorPoliciesGet has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorPolicyCreateHandler == nil {
		api.SnapMirrorSnapmirrorPolicyCreateHandler = snap_mirror.SnapmirrorPolicyCreateHandlerFunc(func(params snap_mirror.SnapmirrorPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorPolicyCreate has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorPolicyDeleteHandler == nil {
		api.SnapMirrorSnapmirrorPolicyDeleteHandler = snap_mirror.SnapmirrorPolicyDeleteHandlerFunc(func(params snap_mirror.SnapmirrorPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorPolicyDelete has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorPolicyGetHandler == nil {
		api.SnapMirrorSnapmirrorPolicyGetHandler = snap_mirror.SnapmirrorPolicyGetHandlerFunc(func(params snap_mirror.SnapmirrorPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorPolicyGet has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorPolicyModifyHandler == nil {
		api.SnapMirrorSnapmirrorPolicyModifyHandler = snap_mirror.SnapmirrorPolicyModifyHandlerFunc(func(params snap_mirror.SnapmirrorPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorPolicyModify has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipCreateHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipCreateHandler = snap_mirror.SnapmirrorRelationshipCreateHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipCreate has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipDeleteHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipDeleteHandler = snap_mirror.SnapmirrorRelationshipDeleteHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipDelete has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipGetHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipGetHandler = snap_mirror.SnapmirrorRelationshipGetHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipGet has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipModifyHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipModifyHandler = snap_mirror.SnapmirrorRelationshipModifyHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipModify has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipTransferCreateHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipTransferCreateHandler = snap_mirror.SnapmirrorRelationshipTransferCreateHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipTransferCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipTransferCreate has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipTransferGetHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipTransferGetHandler = snap_mirror.SnapmirrorRelationshipTransferGetHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipTransferGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipTransferGet has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipTransferModifyHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipTransferModifyHandler = snap_mirror.SnapmirrorRelationshipTransferModifyHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipTransferModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipTransferModify has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipTransfersGetHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipTransfersGetHandler = snap_mirror.SnapmirrorRelationshipTransfersGetHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipTransfersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipTransfersGet has not yet been implemented")
		})
	}
	if api.SnapMirrorSnapmirrorRelationshipsGetHandler == nil {
		api.SnapMirrorSnapmirrorRelationshipsGetHandler = snap_mirror.SnapmirrorRelationshipsGetHandlerFunc(func(params snap_mirror.SnapmirrorRelationshipsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation snap_mirror.SnapmirrorRelationshipsGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotCollectionGetHandler == nil {
		api.StorageSnapshotCollectionGetHandler = storage.SnapshotCollectionGetHandlerFunc(func(params storage.SnapshotCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotCollectionGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotCreateHandler == nil {
		api.StorageSnapshotCreateHandler = storage.SnapshotCreateHandlerFunc(func(params storage.SnapshotCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotCreate has not yet been implemented")
		})
	}
	if api.StorageSnapshotDeleteHandler == nil {
		api.StorageSnapshotDeleteHandler = storage.SnapshotDeleteHandlerFunc(func(params storage.SnapshotDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotDelete has not yet been implemented")
		})
	}
	if api.StorageSnapshotGetHandler == nil {
		api.StorageSnapshotGetHandler = storage.SnapshotGetHandlerFunc(func(params storage.SnapshotGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotModifyHandler == nil {
		api.StorageSnapshotModifyHandler = storage.SnapshotModifyHandlerFunc(func(params storage.SnapshotModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotModify has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyCollectionGetHandler == nil {
		api.StorageSnapshotPolicyCollectionGetHandler = storage.SnapshotPolicyCollectionGetHandlerFunc(func(params storage.SnapshotPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyCreateHandler == nil {
		api.StorageSnapshotPolicyCreateHandler = storage.SnapshotPolicyCreateHandlerFunc(func(params storage.SnapshotPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyCreate has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyDeleteHandler == nil {
		api.StorageSnapshotPolicyDeleteHandler = storage.SnapshotPolicyDeleteHandlerFunc(func(params storage.SnapshotPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyDelete has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyGetHandler == nil {
		api.StorageSnapshotPolicyGetHandler = storage.SnapshotPolicyGetHandlerFunc(func(params storage.SnapshotPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyModifyHandler == nil {
		api.StorageSnapshotPolicyModifyHandler = storage.SnapshotPolicyModifyHandlerFunc(func(params storage.SnapshotPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyModify has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyScheduleCollectionGetHandler == nil {
		api.StorageSnapshotPolicyScheduleCollectionGetHandler = storage.SnapshotPolicyScheduleCollectionGetHandlerFunc(func(params storage.SnapshotPolicyScheduleCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyScheduleCollectionGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyScheduleCreateHandler == nil {
		api.StorageSnapshotPolicyScheduleCreateHandler = storage.SnapshotPolicyScheduleCreateHandlerFunc(func(params storage.SnapshotPolicyScheduleCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyScheduleCreate has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyScheduleDeleteHandler == nil {
		api.StorageSnapshotPolicyScheduleDeleteHandler = storage.SnapshotPolicyScheduleDeleteHandlerFunc(func(params storage.SnapshotPolicyScheduleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyScheduleDelete has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyScheduleGetHandler == nil {
		api.StorageSnapshotPolicyScheduleGetHandler = storage.SnapshotPolicyScheduleGetHandlerFunc(func(params storage.SnapshotPolicyScheduleGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyScheduleGet has not yet been implemented")
		})
	}
	if api.StorageSnapshotPolicyScheduleModifyHandler == nil {
		api.StorageSnapshotPolicyScheduleModifyHandler = storage.SnapshotPolicyScheduleModifyHandlerFunc(func(params storage.SnapshotPolicyScheduleModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SnapshotPolicyScheduleModify has not yet been implemented")
		})
	}
	if api.SupportSnmpGetHandler == nil {
		api.SupportSnmpGetHandler = support.SnmpGetHandlerFunc(func(params support.SnmpGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpGet has not yet been implemented")
		})
	}
	if api.SupportSnmpModifyHandler == nil {
		api.SupportSnmpModifyHandler = support.SnmpModifyHandlerFunc(func(params support.SnmpModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpModify has not yet been implemented")
		})
	}
	if api.SupportSnmpTraphostsCollectionGetHandler == nil {
		api.SupportSnmpTraphostsCollectionGetHandler = support.SnmpTraphostsCollectionGetHandlerFunc(func(params support.SnmpTraphostsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpTraphostsCollectionGet has not yet been implemented")
		})
	}
	if api.SupportSnmpTraphostsCreateHandler == nil {
		api.SupportSnmpTraphostsCreateHandler = support.SnmpTraphostsCreateHandlerFunc(func(params support.SnmpTraphostsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpTraphostsCreate has not yet been implemented")
		})
	}
	if api.SupportSnmpTraphostsDeleteHandler == nil {
		api.SupportSnmpTraphostsDeleteHandler = support.SnmpTraphostsDeleteHandlerFunc(func(params support.SnmpTraphostsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpTraphostsDelete has not yet been implemented")
		})
	}
	if api.SupportSnmpTraphostsGetHandler == nil {
		api.SupportSnmpTraphostsGetHandler = support.SnmpTraphostsGetHandlerFunc(func(params support.SnmpTraphostsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpTraphostsGet has not yet been implemented")
		})
	}
	if api.SupportSnmpUsersCollectionGetHandler == nil {
		api.SupportSnmpUsersCollectionGetHandler = support.SnmpUsersCollectionGetHandlerFunc(func(params support.SnmpUsersCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpUsersCollectionGet has not yet been implemented")
		})
	}
	if api.SupportSnmpUsersCreateHandler == nil {
		api.SupportSnmpUsersCreateHandler = support.SnmpUsersCreateHandlerFunc(func(params support.SnmpUsersCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpUsersCreate has not yet been implemented")
		})
	}
	if api.SupportSnmpUsersDeleteHandler == nil {
		api.SupportSnmpUsersDeleteHandler = support.SnmpUsersDeleteHandlerFunc(func(params support.SnmpUsersDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpUsersDelete has not yet been implemented")
		})
	}
	if api.SupportSnmpUsersGetHandler == nil {
		api.SupportSnmpUsersGetHandler = support.SnmpUsersGetHandlerFunc(func(params support.SnmpUsersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpUsersGet has not yet been implemented")
		})
	}
	if api.SupportSnmpUsersModifyHandler == nil {
		api.SupportSnmpUsersModifyHandler = support.SnmpUsersModifyHandlerFunc(func(params support.SnmpUsersModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation support.SnmpUsersModify has not yet been implemented")
		})
	}
	if api.ClusterSoftwareDownloadGetHandler == nil {
		api.ClusterSoftwareDownloadGetHandler = cluster.SoftwareDownloadGetHandlerFunc(func(params cluster.SoftwareDownloadGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwareDownloadGet has not yet been implemented")
		})
	}
	if api.ClusterSoftwareGetHandler == nil {
		api.ClusterSoftwareGetHandler = cluster.SoftwareGetHandlerFunc(func(params cluster.SoftwareGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwareGet has not yet been implemented")
		})
	}
	if api.ClusterSoftwareHistoryCollectionGetHandler == nil {
		api.ClusterSoftwareHistoryCollectionGetHandler = cluster.SoftwareHistoryCollectionGetHandlerFunc(func(params cluster.SoftwareHistoryCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwareHistoryCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterSoftwareModifyHandler == nil {
		api.ClusterSoftwareModifyHandler = cluster.SoftwareModifyHandlerFunc(func(params cluster.SoftwareModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwareModify has not yet been implemented")
		})
	}
	if api.ClusterSoftwarePackageCreateHandler == nil {
		api.ClusterSoftwarePackageCreateHandler = cluster.SoftwarePackageCreateHandlerFunc(func(params cluster.SoftwarePackageCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwarePackageCreate has not yet been implemented")
		})
	}
	if api.ClusterSoftwarePackageDeleteHandler == nil {
		api.ClusterSoftwarePackageDeleteHandler = cluster.SoftwarePackageDeleteHandlerFunc(func(params cluster.SoftwarePackageDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwarePackageDelete has not yet been implemented")
		})
	}
	if api.ClusterSoftwarePackageGetHandler == nil {
		api.ClusterSoftwarePackageGetHandler = cluster.SoftwarePackageGetHandlerFunc(func(params cluster.SoftwarePackageGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwarePackageGet has not yet been implemented")
		})
	}
	if api.ClusterSoftwarePackagesCollectionGetHandler == nil {
		api.ClusterSoftwarePackagesCollectionGetHandler = cluster.SoftwarePackagesCollectionGetHandlerFunc(func(params cluster.SoftwarePackagesCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwarePackagesCollectionGet has not yet been implemented")
		})
	}
	if api.ClusterSoftwareUploadHandler == nil {
		api.ClusterSoftwareUploadHandler = cluster.SoftwareUploadHandlerFunc(func(params cluster.SoftwareUploadParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.SoftwareUpload has not yet been implemented")
		})
	}
	if api.StorageSplitLoadCollectionGetHandler == nil {
		api.StorageSplitLoadCollectionGetHandler = storage.SplitLoadCollectionGetHandlerFunc(func(params storage.SplitLoadCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SplitLoadCollectionGet has not yet been implemented")
		})
	}
	if api.StorageSplitLoadGetHandler == nil {
		api.StorageSplitLoadGetHandler = storage.SplitLoadGetHandlerFunc(func(params storage.SplitLoadGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SplitLoadGet has not yet been implemented")
		})
	}
	if api.StorageSplitLoadModifyHandler == nil {
		api.StorageSplitLoadModifyHandler = storage.SplitLoadModifyHandlerFunc(func(params storage.SplitLoadModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SplitLoadModify has not yet been implemented")
		})
	}
	if api.StorageSplitStatusCollectionGetHandler == nil {
		api.StorageSplitStatusCollectionGetHandler = storage.SplitStatusCollectionGetHandlerFunc(func(params storage.SplitStatusCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SplitStatusCollectionGet has not yet been implemented")
		})
	}
	if api.StorageSplitStatusGetHandler == nil {
		api.StorageSplitStatusGetHandler = storage.SplitStatusGetHandlerFunc(func(params storage.SplitStatusGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.SplitStatusGet has not yet been implemented")
		})
	}
	if api.SecuritySSHGetHandler == nil {
		api.SecuritySSHGetHandler = securityops.SSHGetHandlerFunc(func(params securityops.SSHGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SSHGet has not yet been implemented")
		})
	}
	if api.SecuritySSHModifyHandler == nil {
		api.SecuritySSHModifyHandler = securityops.SSHModifyHandlerFunc(func(params securityops.SSHModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SSHModify has not yet been implemented")
		})
	}
	if api.StorageStartDirectoryRestoreHandler == nil {
		api.StorageStartDirectoryRestoreHandler = storage.StartDirectoryRestoreHandlerFunc(func(params storage.StartDirectoryRestoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StartDirectoryRestore has not yet been implemented")
		})
	}
	if api.StorageStorageBridgeCollectionGetHandler == nil {
		api.StorageStorageBridgeCollectionGetHandler = storage.StorageBridgeCollectionGetHandlerFunc(func(params storage.StorageBridgeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageBridgeCollectionGet has not yet been implemented")
		})
	}
	if api.StorageStorageBridgeGetHandler == nil {
		api.StorageStorageBridgeGetHandler = storage.StorageBridgeGetHandlerFunc(func(params storage.StorageBridgeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageBridgeGet has not yet been implemented")
		})
	}
	if api.StorageStorageClusterGetHandler == nil {
		api.StorageStorageClusterGetHandler = storage.StorageClusterGetHandlerFunc(func(params storage.StorageClusterGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageClusterGet has not yet been implemented")
		})
	}
	if api.StorageStoragePoolCollectionGetHandler == nil {
		api.StorageStoragePoolCollectionGetHandler = storage.StoragePoolCollectionGetHandlerFunc(func(params storage.StoragePoolCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePoolCollectionGet has not yet been implemented")
		})
	}
	if api.StorageStoragePoolCreateHandler == nil {
		api.StorageStoragePoolCreateHandler = storage.StoragePoolCreateHandlerFunc(func(params storage.StoragePoolCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePoolCreate has not yet been implemented")
		})
	}
	if api.StorageStoragePoolDeleteHandler == nil {
		api.StorageStoragePoolDeleteHandler = storage.StoragePoolDeleteHandlerFunc(func(params storage.StoragePoolDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePoolDelete has not yet been implemented")
		})
	}
	if api.StorageStoragePoolGetHandler == nil {
		api.StorageStoragePoolGetHandler = storage.StoragePoolGetHandlerFunc(func(params storage.StoragePoolGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePoolGet has not yet been implemented")
		})
	}
	if api.StorageStoragePoolModifyHandler == nil {
		api.StorageStoragePoolModifyHandler = storage.StoragePoolModifyHandlerFunc(func(params storage.StoragePoolModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePoolModify has not yet been implemented")
		})
	}
	if api.StorageStoragePortModifyHandler == nil {
		api.StorageStoragePortModifyHandler = storage.StoragePortModifyHandlerFunc(func(params storage.StoragePortModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StoragePortModify has not yet been implemented")
		})
	}
	if api.StorageStorageSwitchCollectionGetHandler == nil {
		api.StorageStorageSwitchCollectionGetHandler = storage.StorageSwitchCollectionGetHandlerFunc(func(params storage.StorageSwitchCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageSwitchCollectionGet has not yet been implemented")
		})
	}
	if api.StorageStorageSwitchGetHandler == nil {
		api.StorageStorageSwitchGetHandler = storage.StorageSwitchGetHandlerFunc(func(params storage.StorageSwitchGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.StorageSwitchGet has not yet been implemented")
		})
	}
	if api.SVMSvmCollectionGetHandler == nil {
		api.SVMSvmCollectionGetHandler = s_vm.SvmCollectionGetHandlerFunc(func(params s_vm.SvmCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmCollectionGet has not yet been implemented")
		})
	}
	if api.SVMSvmCreateHandler == nil {
		api.SVMSvmCreateHandler = s_vm.SvmCreateHandlerFunc(func(params s_vm.SvmCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmCreate has not yet been implemented")
		})
	}
	if api.SVMSvmDeleteHandler == nil {
		api.SVMSvmDeleteHandler = s_vm.SvmDeleteHandlerFunc(func(params s_vm.SvmDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmDelete has not yet been implemented")
		})
	}
	if api.SVMSvmGetHandler == nil {
		api.SVMSvmGetHandler = s_vm.SvmGetHandlerFunc(func(params s_vm.SvmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmGet has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationCollectionGetHandler == nil {
		api.SVMSvmMigrationCollectionGetHandler = s_vm.SvmMigrationCollectionGetHandlerFunc(func(params s_vm.SvmMigrationCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationCollectionGet has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationCreateHandler == nil {
		api.SVMSvmMigrationCreateHandler = s_vm.SvmMigrationCreateHandlerFunc(func(params s_vm.SvmMigrationCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationCreate has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationDeleteHandler == nil {
		api.SVMSvmMigrationDeleteHandler = s_vm.SvmMigrationDeleteHandlerFunc(func(params s_vm.SvmMigrationDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationDelete has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationGetHandler == nil {
		api.SVMSvmMigrationGetHandler = s_vm.SvmMigrationGetHandlerFunc(func(params s_vm.SvmMigrationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationGet has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationModifyHandler == nil {
		api.SVMSvmMigrationModifyHandler = s_vm.SvmMigrationModifyHandlerFunc(func(params s_vm.SvmMigrationModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationModify has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationVolumeCollectionGetHandler == nil {
		api.SVMSvmMigrationVolumeCollectionGetHandler = s_vm.SvmMigrationVolumeCollectionGetHandlerFunc(func(params s_vm.SvmMigrationVolumeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationVolumeCollectionGet has not yet been implemented")
		})
	}
	if api.SVMSvmMigrationVolumeGetHandler == nil {
		api.SVMSvmMigrationVolumeGetHandler = s_vm.SvmMigrationVolumeGetHandlerFunc(func(params s_vm.SvmMigrationVolumeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmMigrationVolumeGet has not yet been implemented")
		})
	}
	if api.SVMSvmModifyHandler == nil {
		api.SVMSvmModifyHandler = s_vm.SvmModifyHandlerFunc(func(params s_vm.SvmModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmModify has not yet been implemented")
		})
	}
	if api.SVMSvmPeerCollectionGetHandler == nil {
		api.SVMSvmPeerCollectionGetHandler = s_vm.SvmPeerCollectionGetHandlerFunc(func(params s_vm.SvmPeerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerCollectionGet has not yet been implemented")
		})
	}
	if api.SVMSvmPeerCreateHandler == nil {
		api.SVMSvmPeerCreateHandler = s_vm.SvmPeerCreateHandlerFunc(func(params s_vm.SvmPeerCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerCreate has not yet been implemented")
		})
	}
	if api.SVMSvmPeerDeleteHandler == nil {
		api.SVMSvmPeerDeleteHandler = s_vm.SvmPeerDeleteHandlerFunc(func(params s_vm.SvmPeerDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerDelete has not yet been implemented")
		})
	}
	if api.SVMSvmPeerInstanceGetHandler == nil {
		api.SVMSvmPeerInstanceGetHandler = s_vm.SvmPeerInstanceGetHandlerFunc(func(params s_vm.SvmPeerInstanceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerInstanceGet has not yet been implemented")
		})
	}
	if api.SVMSvmPeerModifyHandler == nil {
		api.SVMSvmPeerModifyHandler = s_vm.SvmPeerModifyHandlerFunc(func(params s_vm.SvmPeerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerModify has not yet been implemented")
		})
	}
	if api.SVMSvmPeerPermissionCollectionGetHandler == nil {
		api.SVMSvmPeerPermissionCollectionGetHandler = s_vm.SvmPeerPermissionCollectionGetHandlerFunc(func(params s_vm.SvmPeerPermissionCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerPermissionCollectionGet has not yet been implemented")
		})
	}
	if api.SVMSvmPeerPermissionCreateHandler == nil {
		api.SVMSvmPeerPermissionCreateHandler = s_vm.SvmPeerPermissionCreateHandlerFunc(func(params s_vm.SvmPeerPermissionCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerPermissionCreate has not yet been implemented")
		})
	}
	if api.SVMSvmPeerPermissionDeleteHandler == nil {
		api.SVMSvmPeerPermissionDeleteHandler = s_vm.SvmPeerPermissionDeleteHandlerFunc(func(params s_vm.SvmPeerPermissionDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerPermissionDelete has not yet been implemented")
		})
	}
	if api.SVMSvmPeerPermissionInstanceGetHandler == nil {
		api.SVMSvmPeerPermissionInstanceGetHandler = s_vm.SvmPeerPermissionInstanceGetHandlerFunc(func(params s_vm.SvmPeerPermissionInstanceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerPermissionInstanceGet has not yet been implemented")
		})
	}
	if api.SVMSvmPeerPermissionModifyHandler == nil {
		api.SVMSvmPeerPermissionModifyHandler = s_vm.SvmPeerPermissionModifyHandlerFunc(func(params s_vm.SvmPeerPermissionModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.SvmPeerPermissionModify has not yet been implemented")
		})
	}
	if api.SecuritySvmSSHServerCollectionGetHandler == nil {
		api.SecuritySvmSSHServerCollectionGetHandler = securityops.SvmSSHServerCollectionGetHandlerFunc(func(params securityops.SvmSSHServerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SvmSSHServerCollectionGet has not yet been implemented")
		})
	}
	if api.SecuritySvmSSHServerGetHandler == nil {
		api.SecuritySvmSSHServerGetHandler = securityops.SvmSSHServerGetHandlerFunc(func(params securityops.SvmSSHServerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SvmSSHServerGet has not yet been implemented")
		})
	}
	if api.SecuritySvmSSHServerModifyHandler == nil {
		api.SecuritySvmSSHServerModifyHandler = securityops.SvmSSHServerModifyHandlerFunc(func(params securityops.SvmSSHServerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.SvmSSHServerModify has not yet been implemented")
		})
	}
	if api.NetworkingSwitchCollectionGetHandler == nil {
		api.NetworkingSwitchCollectionGetHandler = networking.SwitchCollectionGetHandlerFunc(func(params networking.SwitchCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingSwitchCreateHandler == nil {
		api.NetworkingSwitchCreateHandler = networking.SwitchCreateHandlerFunc(func(params networking.SwitchCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchCreate has not yet been implemented")
		})
	}
	if api.NetworkingSwitchDeleteHandler == nil {
		api.NetworkingSwitchDeleteHandler = networking.SwitchDeleteHandlerFunc(func(params networking.SwitchDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchDelete has not yet been implemented")
		})
	}
	if api.NetworkingSwitchGetHandler == nil {
		api.NetworkingSwitchGetHandler = networking.SwitchGetHandlerFunc(func(params networking.SwitchGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchGet has not yet been implemented")
		})
	}
	if api.NetworkingSwitchModifyHandler == nil {
		api.NetworkingSwitchModifyHandler = networking.SwitchModifyHandlerFunc(func(params networking.SwitchModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchModify has not yet been implemented")
		})
	}
	if api.NetworkingSwitchPortCollectionGetHandler == nil {
		api.NetworkingSwitchPortCollectionGetHandler = networking.SwitchPortCollectionGetHandlerFunc(func(params networking.SwitchPortCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchPortCollectionGet has not yet been implemented")
		})
	}
	if api.NetworkingSwitchPortGetHandler == nil {
		api.NetworkingSwitchPortGetHandler = networking.SwitchPortGetHandlerFunc(func(params networking.SwitchPortGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation networking.SwitchPortGet has not yet been implemented")
		})
	}
	if api.StorageTapeDeviceCollectionGetHandler == nil {
		api.StorageTapeDeviceCollectionGetHandler = storage.TapeDeviceCollectionGetHandlerFunc(func(params storage.TapeDeviceCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TapeDeviceCollectionGet has not yet been implemented")
		})
	}
	if api.StorageTapeDeviceGetHandler == nil {
		api.StorageTapeDeviceGetHandler = storage.TapeDeviceGetHandlerFunc(func(params storage.TapeDeviceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TapeDeviceGet has not yet been implemented")
		})
	}
	if api.StorageTapeDeviceModifyHandler == nil {
		api.StorageTapeDeviceModifyHandler = storage.TapeDeviceModifyHandlerFunc(func(params storage.TapeDeviceModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TapeDeviceModify has not yet been implemented")
		})
	}
	if api.StorageTokenCollectionGetHandler == nil {
		api.StorageTokenCollectionGetHandler = storage.TokenCollectionGetHandlerFunc(func(params storage.TokenCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TokenCollectionGet has not yet been implemented")
		})
	}
	if api.StorageTokenCreateHandler == nil {
		api.StorageTokenCreateHandler = storage.TokenCreateHandlerFunc(func(params storage.TokenCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TokenCreate has not yet been implemented")
		})
	}
	if api.StorageTokenDeleteHandler == nil {
		api.StorageTokenDeleteHandler = storage.TokenDeleteHandlerFunc(func(params storage.TokenDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TokenDelete has not yet been implemented")
		})
	}
	if api.StorageTokenGetHandler == nil {
		api.StorageTokenGetHandler = storage.TokenGetHandlerFunc(func(params storage.TokenGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TokenGet has not yet been implemented")
		})
	}
	if api.StorageTokenModifyHandler == nil {
		api.StorageTokenModifyHandler = storage.TokenModifyHandlerFunc(func(params storage.TokenModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TokenModify has not yet been implemented")
		})
	}
	if api.StorageTopMetricsClientCollectionGetHandler == nil {
		api.StorageTopMetricsClientCollectionGetHandler = storage.TopMetricsClientCollectionGetHandlerFunc(func(params storage.TopMetricsClientCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TopMetricsClientCollectionGet has not yet been implemented")
		})
	}
	if api.StorageTopMetricsDirectoryCollectionGetHandler == nil {
		api.StorageTopMetricsDirectoryCollectionGetHandler = storage.TopMetricsDirectoryCollectionGetHandlerFunc(func(params storage.TopMetricsDirectoryCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TopMetricsDirectoryCollectionGet has not yet been implemented")
		})
	}
	if api.StorageTopMetricsFileCollectionGetHandler == nil {
		api.StorageTopMetricsFileCollectionGetHandler = storage.TopMetricsFileCollectionGetHandlerFunc(func(params storage.TopMetricsFileCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TopMetricsFileCollectionGet has not yet been implemented")
		})
	}
	if api.SVMTopMetricsSvmClientCollectionGetHandler == nil {
		api.SVMTopMetricsSvmClientCollectionGetHandler = s_vm.TopMetricsSvmClientCollectionGetHandlerFunc(func(params s_vm.TopMetricsSvmClientCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.TopMetricsSvmClientCollectionGet has not yet been implemented")
		})
	}
	if api.SVMTopMetricsSvmDirectoryCollectionGetHandler == nil {
		api.SVMTopMetricsSvmDirectoryCollectionGetHandler = s_vm.TopMetricsSvmDirectoryCollectionGetHandlerFunc(func(params s_vm.TopMetricsSvmDirectoryCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.TopMetricsSvmDirectoryCollectionGet has not yet been implemented")
		})
	}
	if api.SVMTopMetricsSvmFileCollectionGetHandler == nil {
		api.SVMTopMetricsSvmFileCollectionGetHandler = s_vm.TopMetricsSvmFileCollectionGetHandlerFunc(func(params s_vm.TopMetricsSvmFileCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.TopMetricsSvmFileCollectionGet has not yet been implemented")
		})
	}
	if api.SVMTopMetricsSvmUserCollectionGetHandler == nil {
		api.SVMTopMetricsSvmUserCollectionGetHandler = s_vm.TopMetricsSvmUserCollectionGetHandlerFunc(func(params s_vm.TopMetricsSvmUserCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.TopMetricsSvmUserCollectionGet has not yet been implemented")
		})
	}
	if api.StorageTopMetricsUserCollectionGetHandler == nil {
		api.StorageTopMetricsUserCollectionGetHandler = storage.TopMetricsUserCollectionGetHandlerFunc(func(params storage.TopMetricsUserCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.TopMetricsUserCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityTotpCollectionGetHandler == nil {
		api.SecurityTotpCollectionGetHandler = securityops.TotpCollectionGetHandlerFunc(func(params securityops.TotpCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.TotpCollectionGet has not yet been implemented")
		})
	}
	if api.SecurityTotpCreateHandler == nil {
		api.SecurityTotpCreateHandler = securityops.TotpCreateHandlerFunc(func(params securityops.TotpCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.TotpCreate has not yet been implemented")
		})
	}
	if api.SecurityTotpGetHandler == nil {
		api.SecurityTotpGetHandler = securityops.TotpGetHandlerFunc(func(params securityops.TotpGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.TotpGet has not yet been implemented")
		})
	}
	if api.SecurityTotpModifyHandler == nil {
		api.SecurityTotpModifyHandler = securityops.TotpModifyHandlerFunc(func(params securityops.TotpModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation security.TotpModify has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupCollectionGetHandler == nil {
		api.NameServicesUnixGroupCollectionGetHandler = name_services.UnixGroupCollectionGetHandlerFunc(func(params name_services.UnixGroupCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupCreateHandler == nil {
		api.NameServicesUnixGroupCreateHandler = name_services.UnixGroupCreateHandlerFunc(func(params name_services.UnixGroupCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupCreate has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupDeleteHandler == nil {
		api.NameServicesUnixGroupDeleteHandler = name_services.UnixGroupDeleteHandlerFunc(func(params name_services.UnixGroupDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupDelete has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupGetHandler == nil {
		api.NameServicesUnixGroupGetHandler = name_services.UnixGroupGetHandlerFunc(func(params name_services.UnixGroupGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupModifyHandler == nil {
		api.NameServicesUnixGroupModifyHandler = name_services.UnixGroupModifyHandlerFunc(func(params name_services.UnixGroupModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupModify has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupSettingsCollectionGetHandler == nil {
		api.NameServicesUnixGroupSettingsCollectionGetHandler = name_services.UnixGroupSettingsCollectionGetHandlerFunc(func(params name_services.UnixGroupSettingsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupSettingsCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupSettingsGetHandler == nil {
		api.NameServicesUnixGroupSettingsGetHandler = name_services.UnixGroupSettingsGetHandlerFunc(func(params name_services.UnixGroupSettingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupSettingsGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupSettingsModifyHandler == nil {
		api.NameServicesUnixGroupSettingsModifyHandler = name_services.UnixGroupSettingsModifyHandlerFunc(func(params name_services.UnixGroupSettingsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupSettingsModify has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupUserDeleteHandler == nil {
		api.NameServicesUnixGroupUserDeleteHandler = name_services.UnixGroupUserDeleteHandlerFunc(func(params name_services.UnixGroupUserDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupUserDelete has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupUsersCollectionGetHandler == nil {
		api.NameServicesUnixGroupUsersCollectionGetHandler = name_services.UnixGroupUsersCollectionGetHandlerFunc(func(params name_services.UnixGroupUsersCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupUsersCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupUsersCreateHandler == nil {
		api.NameServicesUnixGroupUsersCreateHandler = name_services.UnixGroupUsersCreateHandlerFunc(func(params name_services.UnixGroupUsersCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupUsersCreate has not yet been implemented")
		})
	}
	if api.NameServicesUnixGroupUsersGetHandler == nil {
		api.NameServicesUnixGroupUsersGetHandler = name_services.UnixGroupUsersGetHandlerFunc(func(params name_services.UnixGroupUsersGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixGroupUsersGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserCollectionGetHandler == nil {
		api.NameServicesUnixUserCollectionGetHandler = name_services.UnixUserCollectionGetHandlerFunc(func(params name_services.UnixUserCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserCreateHandler == nil {
		api.NameServicesUnixUserCreateHandler = name_services.UnixUserCreateHandlerFunc(func(params name_services.UnixUserCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserCreate has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserDeleteHandler == nil {
		api.NameServicesUnixUserDeleteHandler = name_services.UnixUserDeleteHandlerFunc(func(params name_services.UnixUserDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserDelete has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserGetHandler == nil {
		api.NameServicesUnixUserGetHandler = name_services.UnixUserGetHandlerFunc(func(params name_services.UnixUserGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserModifyHandler == nil {
		api.NameServicesUnixUserModifyHandler = name_services.UnixUserModifyHandlerFunc(func(params name_services.UnixUserModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserModify has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserSettingsCollectionGetHandler == nil {
		api.NameServicesUnixUserSettingsCollectionGetHandler = name_services.UnixUserSettingsCollectionGetHandlerFunc(func(params name_services.UnixUserSettingsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserSettingsCollectionGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserSettingsGetHandler == nil {
		api.NameServicesUnixUserSettingsGetHandler = name_services.UnixUserSettingsGetHandlerFunc(func(params name_services.UnixUserSettingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserSettingsGet has not yet been implemented")
		})
	}
	if api.NameServicesUnixUserSettingsModifyHandler == nil {
		api.NameServicesUnixUserSettingsModifyHandler = name_services.UnixUserSettingsModifyHandlerFunc(func(params name_services.UnixUserSettingsModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation name_services.UnixUserSettingsModify has not yet been implemented")
		})
	}
	if api.NasUserGroupPrivilegesCollectionGetHandler == nil {
		api.NasUserGroupPrivilegesCollectionGetHandler = n_a_s.UserGroupPrivilegesCollectionGetHandlerFunc(func(params n_a_s.UserGroupPrivilegesCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.UserGroupPrivilegesCollectionGet has not yet been implemented")
		})
	}
	if api.NasUserGroupPrivilegesCreateHandler == nil {
		api.NasUserGroupPrivilegesCreateHandler = n_a_s.UserGroupPrivilegesCreateHandlerFunc(func(params n_a_s.UserGroupPrivilegesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.UserGroupPrivilegesCreate has not yet been implemented")
		})
	}
	if api.NasUserGroupPrivilegesGetHandler == nil {
		api.NasUserGroupPrivilegesGetHandler = n_a_s.UserGroupPrivilegesGetHandlerFunc(func(params n_a_s.UserGroupPrivilegesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.UserGroupPrivilegesGet has not yet been implemented")
		})
	}
	if api.NasUserGroupPrivilegesModifyHandler == nil {
		api.NasUserGroupPrivilegesModifyHandler = n_a_s.UserGroupPrivilegesModifyHandlerFunc(func(params n_a_s.UserGroupPrivilegesModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.UserGroupPrivilegesModify has not yet been implemented")
		})
	}
	if api.StorageVolumeCollectionGetHandler == nil {
		api.StorageVolumeCollectionGetHandler = storage.VolumeCollectionGetHandlerFunc(func(params storage.VolumeCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeCollectionGet has not yet been implemented")
		})
	}
	if api.StorageVolumeCreateHandler == nil {
		api.StorageVolumeCreateHandler = storage.VolumeCreateHandlerFunc(func(params storage.VolumeCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeCreate has not yet been implemented")
		})
	}
	if api.StorageVolumeDeleteHandler == nil {
		api.StorageVolumeDeleteHandler = storage.VolumeDeleteHandlerFunc(func(params storage.VolumeDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeDelete has not yet been implemented")
		})
	}
	if api.StorageVolumeEfficiencyPolicyCollectionGetHandler == nil {
		api.StorageVolumeEfficiencyPolicyCollectionGetHandler = storage.VolumeEfficiencyPolicyCollectionGetHandlerFunc(func(params storage.VolumeEfficiencyPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeEfficiencyPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.StorageVolumeEfficiencyPolicyCreateHandler == nil {
		api.StorageVolumeEfficiencyPolicyCreateHandler = storage.VolumeEfficiencyPolicyCreateHandlerFunc(func(params storage.VolumeEfficiencyPolicyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeEfficiencyPolicyCreate has not yet been implemented")
		})
	}
	if api.StorageVolumeEfficiencyPolicyDeleteHandler == nil {
		api.StorageVolumeEfficiencyPolicyDeleteHandler = storage.VolumeEfficiencyPolicyDeleteHandlerFunc(func(params storage.VolumeEfficiencyPolicyDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeEfficiencyPolicyDelete has not yet been implemented")
		})
	}
	if api.StorageVolumeEfficiencyPolicyGetHandler == nil {
		api.StorageVolumeEfficiencyPolicyGetHandler = storage.VolumeEfficiencyPolicyGetHandlerFunc(func(params storage.VolumeEfficiencyPolicyGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeEfficiencyPolicyGet has not yet been implemented")
		})
	}
	if api.StorageVolumeEfficiencyPolicyModifyHandler == nil {
		api.StorageVolumeEfficiencyPolicyModifyHandler = storage.VolumeEfficiencyPolicyModifyHandlerFunc(func(params storage.VolumeEfficiencyPolicyModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeEfficiencyPolicyModify has not yet been implemented")
		})
	}
	if api.StorageVolumeGetHandler == nil {
		api.StorageVolumeGetHandler = storage.VolumeGetHandlerFunc(func(params storage.VolumeGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeGet has not yet been implemented")
		})
	}
	if api.StorageVolumeMetricsCollectionGetHandler == nil {
		api.StorageVolumeMetricsCollectionGetHandler = storage.VolumeMetricsCollectionGetHandlerFunc(func(params storage.VolumeMetricsCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeMetricsCollectionGet has not yet been implemented")
		})
	}
	if api.StorageVolumeModifyHandler == nil {
		api.StorageVolumeModifyHandler = storage.VolumeModifyHandlerFunc(func(params storage.VolumeModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation storage.VolumeModify has not yet been implemented")
		})
	}
	if api.NasVscanCollectionGetHandler == nil {
		api.NasVscanCollectionGetHandler = n_a_s.VscanCollectionGetHandlerFunc(func(params n_a_s.VscanCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanCollectionGet has not yet been implemented")
		})
	}
	if api.NasVscanConfigDeleteHandler == nil {
		api.NasVscanConfigDeleteHandler = n_a_s.VscanConfigDeleteHandlerFunc(func(params n_a_s.VscanConfigDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanConfigDelete has not yet been implemented")
		})
	}
	if api.NasVscanCreateHandler == nil {
		api.NasVscanCreateHandler = n_a_s.VscanCreateHandlerFunc(func(params n_a_s.VscanCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanCreate has not yet been implemented")
		})
	}
	if api.NasVscanEventCollectionGetHandler == nil {
		api.NasVscanEventCollectionGetHandler = n_a_s.VscanEventCollectionGetHandlerFunc(func(params n_a_s.VscanEventCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanEventCollectionGet has not yet been implemented")
		})
	}
	if api.NasVscanGetHandler == nil {
		api.NasVscanGetHandler = n_a_s.VscanGetHandlerFunc(func(params n_a_s.VscanGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanGet has not yet been implemented")
		})
	}
	if api.NasVscanModifyHandler == nil {
		api.NasVscanModifyHandler = n_a_s.VscanModifyHandlerFunc(func(params n_a_s.VscanModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanModify has not yet been implemented")
		})
	}
	if api.NasVscanOnAccessCreateHandler == nil {
		api.NasVscanOnAccessCreateHandler = n_a_s.VscanOnAccessCreateHandlerFunc(func(params n_a_s.VscanOnAccessCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnAccessCreate has not yet been implemented")
		})
	}
	if api.NasVscanOnAccessDeleteHandler == nil {
		api.NasVscanOnAccessDeleteHandler = n_a_s.VscanOnAccessDeleteHandlerFunc(func(params n_a_s.VscanOnAccessDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnAccessDelete has not yet been implemented")
		})
	}
	if api.NasVscanOnAccessGetHandler == nil {
		api.NasVscanOnAccessGetHandler = n_a_s.VscanOnAccessGetHandlerFunc(func(params n_a_s.VscanOnAccessGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnAccessGet has not yet been implemented")
		})
	}
	if api.NasVscanOnAccessModifyHandler == nil {
		api.NasVscanOnAccessModifyHandler = n_a_s.VscanOnAccessModifyHandlerFunc(func(params n_a_s.VscanOnAccessModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnAccessModify has not yet been implemented")
		})
	}
	if api.NasVscanOnAccessPolicyCollectionGetHandler == nil {
		api.NasVscanOnAccessPolicyCollectionGetHandler = n_a_s.VscanOnAccessPolicyCollectionGetHandlerFunc(func(params n_a_s.VscanOnAccessPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnAccessPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasVscanOnDemandCreateHandler == nil {
		api.NasVscanOnDemandCreateHandler = n_a_s.VscanOnDemandCreateHandlerFunc(func(params n_a_s.VscanOnDemandCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnDemandCreate has not yet been implemented")
		})
	}
	if api.NasVscanOnDemandDeleteHandler == nil {
		api.NasVscanOnDemandDeleteHandler = n_a_s.VscanOnDemandDeleteHandlerFunc(func(params n_a_s.VscanOnDemandDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnDemandDelete has not yet been implemented")
		})
	}
	if api.NasVscanOnDemandGetHandler == nil {
		api.NasVscanOnDemandGetHandler = n_a_s.VscanOnDemandGetHandlerFunc(func(params n_a_s.VscanOnDemandGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnDemandGet has not yet been implemented")
		})
	}
	if api.NasVscanOnDemandModifyHandler == nil {
		api.NasVscanOnDemandModifyHandler = n_a_s.VscanOnDemandModifyHandlerFunc(func(params n_a_s.VscanOnDemandModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnDemandModify has not yet been implemented")
		})
	}
	if api.NasVscanOnDemandPolicyCollectionGetHandler == nil {
		api.NasVscanOnDemandPolicyCollectionGetHandler = n_a_s.VscanOnDemandPolicyCollectionGetHandlerFunc(func(params n_a_s.VscanOnDemandPolicyCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanOnDemandPolicyCollectionGet has not yet been implemented")
		})
	}
	if api.NasVscanScannerCollectionGetHandler == nil {
		api.NasVscanScannerCollectionGetHandler = n_a_s.VscanScannerCollectionGetHandlerFunc(func(params n_a_s.VscanScannerCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanScannerCollectionGet has not yet been implemented")
		})
	}
	if api.NasVscanScannerCreateHandler == nil {
		api.NasVscanScannerCreateHandler = n_a_s.VscanScannerCreateHandlerFunc(func(params n_a_s.VscanScannerCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanScannerCreate has not yet been implemented")
		})
	}
	if api.NasVscanScannerDeleteHandler == nil {
		api.NasVscanScannerDeleteHandler = n_a_s.VscanScannerDeleteHandlerFunc(func(params n_a_s.VscanScannerDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanScannerDelete has not yet been implemented")
		})
	}
	if api.NasVscanScannerModifyHandler == nil {
		api.NasVscanScannerModifyHandler = n_a_s.VscanScannerModifyHandlerFunc(func(params n_a_s.VscanScannerModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanScannerModify has not yet been implemented")
		})
	}
	if api.NasVscanScannerPoolGetHandler == nil {
		api.NasVscanScannerPoolGetHandler = n_a_s.VscanScannerPoolGetHandlerFunc(func(params n_a_s.VscanScannerPoolGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanScannerPoolGet has not yet been implemented")
		})
	}
	if api.NasVscanServerStatusGetHandler == nil {
		api.NasVscanServerStatusGetHandler = n_a_s.VscanServerStatusGetHandlerFunc(func(params n_a_s.VscanServerStatusGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation n_a_s.VscanServerStatusGet has not yet been implemented")
		})
	}
	if api.SanVvolBindingCollectionGetHandler == nil {
		api.SanVvolBindingCollectionGetHandler = s_a_n.VvolBindingCollectionGetHandlerFunc(func(params s_a_n.VvolBindingCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.VvolBindingCollectionGet has not yet been implemented")
		})
	}
	if api.SanVvolBindingCreateHandler == nil {
		api.SanVvolBindingCreateHandler = s_a_n.VvolBindingCreateHandlerFunc(func(params s_a_n.VvolBindingCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.VvolBindingCreate has not yet been implemented")
		})
	}
	if api.SanVvolBindingDeleteHandler == nil {
		api.SanVvolBindingDeleteHandler = s_a_n.VvolBindingDeleteHandlerFunc(func(params s_a_n.VvolBindingDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.VvolBindingDelete has not yet been implemented")
		})
	}
	if api.SanVvolBindingGetHandler == nil {
		api.SanVvolBindingGetHandler = s_a_n.VvolBindingGetHandlerFunc(func(params s_a_n.VvolBindingGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.VvolBindingGet has not yet been implemented")
		})
	}
	if api.ClusterWebGetHandler == nil {
		api.ClusterWebGetHandler = cluster.WebGetHandlerFunc(func(params cluster.WebGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.WebGet has not yet been implemented")
		})
	}
	if api.ClusterWebModifyHandler == nil {
		api.ClusterWebModifyHandler = cluster.WebModifyHandlerFunc(func(params cluster.WebModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation cluster.WebModify has not yet been implemented")
		})
	}
	if api.SVMWebSvmGetHandler == nil {
		api.SVMWebSvmGetHandler = s_vm.WebSvmGetHandlerFunc(func(params s_vm.WebSvmGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.WebSvmGet has not yet been implemented")
		})
	}
	if api.SVMWebSvmModifyHandler == nil {
		api.SVMWebSvmModifyHandler = s_vm.WebSvmModifyHandlerFunc(func(params s_vm.WebSvmModifyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_vm.WebSvmModify has not yet been implemented")
		})
	}
	if api.SanWwpnAliasCollectionGetHandler == nil {
		api.SanWwpnAliasCollectionGetHandler = s_a_n.WwpnAliasCollectionGetHandlerFunc(func(params s_a_n.WwpnAliasCollectionGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.WwpnAliasCollectionGet has not yet been implemented")
		})
	}
	if api.SanWwpnAliasCreateHandler == nil {
		api.SanWwpnAliasCreateHandler = s_a_n.WwpnAliasCreateHandlerFunc(func(params s_a_n.WwpnAliasCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.WwpnAliasCreate has not yet been implemented")
		})
	}
	if api.SanWwpnAliasDeleteHandler == nil {
		api.SanWwpnAliasDeleteHandler = s_a_n.WwpnAliasDeleteHandlerFunc(func(params s_a_n.WwpnAliasDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.WwpnAliasDelete has not yet been implemented")
		})
	}
	if api.SanWwpnAliasGetHandler == nil {
		api.SanWwpnAliasGetHandler = s_a_n.WwpnAliasGetHandlerFunc(func(params s_a_n.WwpnAliasGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation s_a_n.WwpnAliasGet has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

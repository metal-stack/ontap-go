// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/pkg/pointer"
	"github.com/metal-stack/ontap-go/pkg/server/models"

	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/cluster"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/networking"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/s_vm"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/security"
	"github.com/metal-stack/ontap-go/pkg/server/restapi/operations/storage"
)

//go:generate swagger generate server --target ../../server --name OntapFakeServer --spec ../../../spec/ontap.yaml --principal interface{} --skip-models

func configureFlags(api *operations.OntapFakeServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

type Store struct {
	clusters map[string]*models.Cluster
	svms     map[string]*models.Svm
	users    map[string]*models.Account
	// Ontap Api is weird
	nodes     map[string]*models.NodeResponseInlineRecordsInlineArrayItem
	aggregate map[string]*models.Aggregate
}

var GlobalStore = &Store{
	clusters:  make(map[string]*models.Cluster),
	svms:      make(map[string]*models.Svm),
	users:     make(map[string]*models.Account),
	nodes:     make(map[string]*models.NodeResponseInlineRecordsInlineArrayItem),
	aggregate: make(map[string]*models.Aggregate),
}

func InitStore() *Store {

	return &Store{
		clusters:  make(map[string]*models.Cluster),
		svms:      make(map[string]*models.Svm),
		users:     make(map[string]*models.Account),
		nodes:     make(map[string]*models.NodeResponseInlineRecordsInlineArrayItem),
		aggregate: make(map[string]*models.Aggregate),
	}
}

func (s *Store) GetCluster(uuid string) *models.Cluster {
	return s.clusters[uuid]
}

func (s *Store) SetCluster(cluster *models.Cluster) {
	if cluster != nil {
		s.clusters[*cluster.Name] = cluster
	}
}

func (s *Store) GetAccount(uuid string) *models.Account {
	return s.users[uuid]
}

func (s *Store) SetAccount(account *models.Account) {
	if account != nil {
		// can be buggy we create svmAdmin names for every svm, same name everywhere
		s.users[*account.Name] = account
	}
}

func (s *Store) GetSvm(uuid string) *models.Svm {
	return s.svms[uuid]
}

func (s *Store) SetSvm(svm *models.Svm) {
	if svm != nil && svm.UUID != nil {
		s.svms[*svm.UUID] = svm
	}
}

func (s *Store) GetNode(uuid string) *models.NodeResponseInlineRecordsInlineArrayItem {
	return s.nodes[uuid]
}

func (s *Store) SetNode(node *models.NodeResponseInlineRecordsInlineArrayItem) {
	if node != nil && node.UUID != nil {
		s.nodes[node.UUID.String()] = node
	}
}

func (s *Store) GetAggregate(uuid string) *models.Aggregate {
	return s.aggregate[uuid]
}

func (s *Store) SetAggregate(node *models.Aggregate) {
	if node != nil && node.UUID != nil {
		s.aggregate[*node.UUID] = node
	}
}

func uuidPtr(s string) *strfmt.UUID {
	uuid := strfmt.UUID(s)
	return &uuid
}

func configureAPI(api *operations.OntapFakeServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	//not used yet
	api.Logger = log.Printf

	api.UseSwaggerUI()
	//initalize memory store
	memoryStore := InitStore()

	memoryStore.SetCluster(&models.Cluster{
		Statistics: &models.ClusterInlineStatistics{
			Status: pointer.Pointer("ok"),
		},
		Name: pointer.Pointer("testCluster"),
	})

	memoryStore.SetSvm(&models.Svm{
		Name: pointer.Pointer("clusterSVM"),
		UUID: pointer.Pointer("12345678-1234-1234-1234-123456789abc"),
	})

	memoryStore.SetNode(&models.NodeResponseInlineRecordsInlineArrayItem{
		Name: pointer.Pointer("NodeA"),
		UUID: uuidPtr("12345678-1234-1234-1234-123456789abd"),
	})

	memoryStore.SetNode(&models.NodeResponseInlineRecordsInlineArrayItem{
		Name: pointer.Pointer("NodeB"),
		UUID: uuidPtr("12345678-1234-1234-1234-123456789abe"),
	})

	memoryStore.SetAggregate(&models.Aggregate{
		UUID: pointer.Pointer("aggrUUID"),
	})

	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()
	api.MultipartformProducer = runtime.DiscardProducer

	// Applies when the Authorization header is set with the Basic scheme
	// Override the generated SimpleAuth function
	api.SimpleAuth = func(user string, pass string) (interface{}, error) {
		// Accept any credentials for testing
		return user, nil
	}

	api.ClusterClusterGetHandler = cluster.ClusterGetHandlerFunc(func(params cluster.ClusterGetParams, principal interface{}) middleware.Responder {
		cgok := cluster.NewClusterGetOK()

		cluster := &models.Cluster{
			Name: pointer.Pointer("test"),
			Statistics: &models.ClusterInlineStatistics{
				Status: pointer.Pointer("ok"),
			},
		}

		cgok.SetPayload(cluster)
		return cgok
	})

	api.SVMSvmGetHandler = s_vm.SvmGetHandlerFunc(func(params s_vm.SvmGetParams, principal interface{}) middleware.Responder {
		cgokpayload := memoryStore.GetSvm("svmUUID")

		response := s_vm.NewSvmGetOK()
		response.SetPayload(cgokpayload)
		return response
	})

	api.SVMSvmCollectionGetHandler = s_vm.SvmCollectionGetHandlerFunc(func(params s_vm.SvmCollectionGetParams, principal interface{}) middleware.Responder {
		var allSvms []*models.Svm
		for _, svm := range memoryStore.svms {
			allSvms = append(allSvms, svm)
		}

		cgokpayload := &models.SvmResponse{
			SvmResponseInlineRecords: allSvms,
		}

		response := s_vm.NewSvmCollectionGetOK()
		response.SetPayload(cgokpayload)
		return response
	})

	api.SVMSvmCreateHandler = s_vm.SvmCreateHandlerFunc(func(params s_vm.SvmCreateParams, principal interface{}) middleware.Responder {

		aggregate := memoryStore.GetAggregate("aggrUUID")
		svmAggr := []*models.SvmInlineAggregatesInlineArrayItem{
			&models.SvmInlineAggregatesInlineArrayItem{
				UUID: aggregate.UUID,
			},
		}

		memoryStore.SetSvm(&models.Svm{
			Name:                params.Info.Name,
			SvmInlineAggregates: svmAggr,
			Nvme: &models.SvmInlineNvme{
				Allowed: pointer.Pointer(true),
				Enabled: pointer.Pointer(true),
			},
			State: pointer.Pointer("running"),
			UUID:  pointer.Pointer("svmUUID"),
		})

		return nil
	})

	api.ClusterNodesGetHandler = cluster.NodesGetHandlerFunc(func(params cluster.NodesGetParams, principal interface{}) middleware.Responder {
		var allNodes []*models.NodeResponseInlineRecordsInlineArrayItem
		for _, node := range memoryStore.nodes {
			allNodes = append(allNodes, node)
		}

		cgokpayload := &models.NodeResponse{
			NodeResponseInlineRecords: allNodes,
		}

		response := cluster.NewNodesGetOK()
		response.SetPayload(cgokpayload)
		return response
	})

	api.StorageAggregateCollectionGetHandler = storage.AggregateCollectionGetHandlerFunc(func(params storage.AggregateCollectionGetParams, principal interface{}) middleware.Responder {
		var allAggregates []*models.Aggregate
		for _, aggregate := range memoryStore.aggregate {
			allAggregates = append(allAggregates, aggregate)
		}

		cgokpayload := &models.AggregateResponse{
			AggregateResponseInlineRecords: allAggregates,
		}

		response := storage.NewAggregateCollectionGetOK()
		response.SetPayload(cgokpayload)
		return response
	})

	api.NetworkingNetworkIPInterfacesCreateHandler = networking.NetworkIPInterfacesCreateHandlerFunc(func(params networking.NetworkIPInterfacesCreateParams, principal interface{}) middleware.Responder {

		svm := memoryStore.GetSvm("svmUUID")

		memoryStore.SetSvm(svm)

		response := networking.NewNetworkIPInterfacesCreateCreated()
		return response
	})

	api.NetworkingNetworkIPBgpPeerGroupsGetHandler = networking.NetworkIPBgpPeerGroupsGetHandlerFunc(func(params networking.NetworkIPBgpPeerGroupsGetParams, principal interface{}) middleware.Responder {

		var neighbors int64 = 4

		response := networking.NewNetworkIPBgpPeerGroupsGetOK()
		response.SetPayload(&models.BgpPeerGroupResponse{
			NumRecords: &neighbors,
		})
		return response
	})

	api.SecurityAccountCreateHandler = security.AccountCreateHandlerFunc(func(params security.AccountCreateParams, principal interface{}) middleware.Responder {

		account := &models.Account{
			Name:  params.Info.Name,
			Owner: params.Info.Owner,
		}
		memoryStore.SetAccount(account)

		return nil
	})

	api.SecurityAccountPasswordCreateHandler = security.AccountPasswordCreateHandlerFunc(func(params security.AccountPasswordCreateParams, principal interface{}) middleware.Responder {

		return nil
	})

	api.NetworkingNetworkIPInterfacesGetHandler = networking.NetworkIPInterfacesGetHandlerFunc(func(params networking.NetworkIPInterfacesGetParams, principal interface{}) middleware.Responder {

		interfaces := make(map[string]string)
		// We aren't fixing ip mismatch issue so this is fine
		interfaces["datalif+0"] = "0.0.0.0"
		interfaces["datalif+1"] = "0.0.0.0"
		interfaces["managementlif"] = "0.0.0.0"

		return nil
	})

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

package handler

import (
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	prometheusapi "github.com/prometheus/client_golang/api"

	"github.com/kubermatic/kubermatic/api/pkg/handler/auth"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
)

// Routing represents an object which binds endpoints to http handlers.
type Routing struct {
	datacenters           map[string]provider.DatacenterMeta
	cloudProviders        provider.CloudRegistry
	sshKeyProvider        provider.SSHKeyProvider
	userProvider          provider.UserProvider
	projectProvider       provider.ProjectProvider
	logger                log.Logger
	oidcAuthenticator     auth.OIDCAuthenticator
	oidcIssuer            auth.OIDCIssuerVerifier
	clusterProviders      map[string]provider.ClusterProvider
	updateManager         common.UpdateManager
	prometheusClient      prometheusapi.Client
	projectMemberProvider provider.ProjectMemberProvider
	userProjectMapper     provider.ProjectMemberMapper
}

// NewRouting creates a new Routing.
func NewRouting(
	datacenters map[string]provider.DatacenterMeta,
	newClusterProviders map[string]provider.ClusterProvider,
	cloudProviders map[string]provider.CloudProvider,
	newSSHKeyProvider provider.SSHKeyProvider,
	userProvider provider.UserProvider,
	projectProvider provider.ProjectProvider,
	oidcAuthenticator auth.OIDCAuthenticator,
	oidcIssuerVerifier auth.OIDCIssuerVerifier,
	updateManager common.UpdateManager,
	prometheusClient prometheusapi.Client,
	projectMemberProvider provider.ProjectMemberProvider,
	userProjectMapper provider.ProjectMemberMapper,
) Routing {
	return Routing{
		datacenters:           datacenters,
		clusterProviders:      newClusterProviders,
		sshKeyProvider:        newSSHKeyProvider,
		userProvider:          userProvider,
		projectProvider:       projectProvider,
		cloudProviders:        cloudProviders,
		logger:                log.NewLogfmtLogger(os.Stderr),
		oidcAuthenticator:     oidcAuthenticator,
		oidcIssuer:            oidcIssuerVerifier,
		updateManager:         updateManager,
		prometheusClient:      prometheusClient,
		projectMemberProvider: projectMemberProvider,
		userProjectMapper:     userProjectMapper,
	}
}

func (r Routing) defaultServerOptions() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(r.logger),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerBefore(r.oidcAuthenticator.Extractor()),
	}
}

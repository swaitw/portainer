package system

import (
	"net/http"

	portainer "github.com/portainer/portainer/api"
	statusutil "github.com/portainer/portainer/api/internal/nodes"
	"github.com/portainer/portainer/api/internal/snapshot"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"

	"github.com/rs/zerolog/log"
)

type nodesCountResponse struct {
	Nodes int `json:"nodes"`
}

// @id systemNodesCount
// @summary Retrieve the count of nodes
// @description **Access policy**: authenticated
// @security ApiKeyAuth
// @security jwt
// @tags system
// @produce json
// @success 200 {object} nodesCountResponse "Success"
// @failure 500 "Server error"
// @router /system/nodes [get]
func (handler *Handler) systemNodesCount(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	endpoints, err := handler.dataStore.Endpoint().Endpoints()
	if err != nil {
		return httperror.InternalServerError("Failed to get environment list", err)
	}

	var nodes int

	for _, endpoint := range endpoints {
		if err := snapshot.FillSnapshotData(handler.dataStore, &endpoint); err != nil {
			return httperror.InternalServerError("Unable to add snapshot data", err)
		}

		nodes += statusutil.NodesCount([]portainer.Endpoint{endpoint})
	}

	return response.JSON(w, &nodesCountResponse{Nodes: nodes})
}

// @id statusNodesCount
// @summary Retrieve the count of nodes
// @deprecated
// @description Deprecated: use the `/system/nodes` endpoint instead.
// @description **Access policy**: authenticated
// @security ApiKeyAuth
// @security jwt
// @tags status
// @produce json
// @success 200 {object} nodesCountResponse "Success"
// @failure 500 "Server error"
// @router /status/nodes [get]
func (handler *Handler) statusNodesCountDeprecated(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	log.Warn().Msg("The /status/nodes endpoint is deprecated, please use the /system/nodes endpoint instead")

	return handler.systemNodesCount(w, r)
}

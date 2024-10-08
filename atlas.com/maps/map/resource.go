package _map

import (
	"atlas-maps/map/character"
	"atlas-maps/rest"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/Chronicle20/atlas-rest/server"
	"github.com/gorilla/mux"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	getCharactersInMap = "get_characters_in_map"
)

func InitResource(si jsonapi.ServerInformation) server.RouteInitializer {
	return func(router *mux.Router, l logrus.FieldLogger) {
		r := router.PathPrefix("/worlds").Subrouter()
		r.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/characters", rest.RegisterHandler(l)(si)(getCharactersInMap, handleGetCharactersInMap)).Methods(http.MethodGet)
	}
}

func handleGetCharactersInMap(d *rest.HandlerDependency, c *rest.HandlerContext) http.HandlerFunc {
	return rest.ParseWorldId(d.Logger(), func(worldId byte) http.HandlerFunc {
		return rest.ParseChannelId(d.Logger(), func(channelId byte) http.HandlerFunc {
			return rest.ParseMapId(d.Logger(), func(mapId uint32) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					ids, err := character.GetCharactersInMap(d.Context())(worldId, channelId, mapId)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					res, err := model.SliceMap(Transform)(model.FixedProvider(ids))(model.ParallelMap())()
					if err != nil {
						d.Logger().WithError(err).Errorf("Creating REST model.")
						w.WriteHeader(http.StatusInternalServerError)
						return
					}

					server.Marshal[[]RestModel](d.Logger())(w)(c.ServerInformation())(res)
				}
			})
		})
	})
}

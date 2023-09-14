package domain

import (
	"fmt"
	"time"

	"github.com/gobeam/stringy"
	"github.com/stackvista/toposender/internal/config"
	"github.com/stackvista/toposender/internal/stackstate"
)

func ToTopologyMessage(cfg *config.Config, t *Topology) *stackstate.TopologyMessage {
	nameToId := make(map[string]string)
	for _, c := range t.Components {
		nameToId[c.Name] = generateComponentIdentifier(t.Name, c)
	}

	tm := stackstate.NewTopologyMessage()
	tm.ApiKey = cfg.ApiKey
	tm.InternalHostname = cfg.HostName
	tm.CollectionTimestamp = time.Now().UnixMilli()
	tm.Topologies = []stackstate.Topology{
		{
			StartSnapshot: true,
			StopSnapshot:  true,
			Instance: stackstate.Instance{
				Type: stringy.New(t.Name).SnakeCase().ToLower(),
				Url:  fmt.Sprintf("https://%s", cfg.HostName),
			},
			Components: ToComponentMessages(nameToId, t.Components),
			Relations:  ToRelationMessages(nameToId, t.Relations),
		},
	}
	return tm
}

func generateComponentIdentifier(name string, c Component) string {
	snakeName := stringy.New(name).SnakeCase().ToLower()
	snakeComponent := stringy.New(c.Name).SnakeCase().ToLower()
	return fmt.Sprintf("%s:%s", snakeName, snakeComponent)
}

func ToComponentMessages(nameToId map[string]string, components []Component) []stackstate.Component {
	componentMessages := make([]stackstate.Component, len(components))
	for i, c := range components {
		data := map[string]interface{}{}
		labels := []string{}
		for k, v := range c.Labels {
			labels = append(labels, fmt.Sprintf("%s:%s", k, v))
		}
		data["labels"] = labels
		data["layer"] = c.Layer
		data["domain"] = c.Domain
		data["environment"] = c.Environment
		data["name"] = c.Name

		componentMessages[i] = stackstate.Component{
			ExternalId: nameToId[c.Name],
			Type:       stackstate.ComponentType{Name: c.Type},
			Data:       data,
		}
	}
	return componentMessages
}

func ToRelationMessages(nameToId map[string]string, relations []Relation) []stackstate.Relation {
	relationMessages := make([]stackstate.Relation, len(relations))
	for i, r := range relations {
		sourceId := nameToId[r.Source]
		targetId := nameToId[r.Target]
		relationMessages[i] = stackstate.Relation{
			ExternalId: fmt.Sprintf("%s->%s", sourceId, targetId),
			SourceId:   sourceId,
			TargetId:   targetId,
			Type:       stackstate.RelationType{Name: r.Type},
			Data:       map[string]interface{}{},
		}
	}
	return relationMessages
}

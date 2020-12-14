package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsCdContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cd = &appsCdContainer{path: "assets/apps/cd"}

func (c *appsCdContainer) Spinnaker(opts ...attr.Attribute) *node.Node {
	return node.New("spinnaker", attr.AssetImage("assets/apps/cd/spinnaker.png"), attr.Shape(attr.None))
}

func (c *appsCdContainer) TektonCli(opts ...attr.Attribute) *node.Node {
	return node.New("tekton-cli", attr.AssetImage("assets/apps/cd/tekton-cli.png"), attr.Shape(attr.None))
}

func (c *appsCdContainer) Tekton(opts ...attr.Attribute) *node.Node {
	return node.New("tekton", attr.AssetImage("assets/apps/cd/tekton.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.cd"
	Register(namespace, "Spinnaker", Cd.Spinnaker)
	Register(namespace, "TektonCli", Cd.TektonCli)
	Register(namespace, "Tekton", Cd.Tekton)
}
package clientgen

import (
	"io"
	"strings"
	"text/template"

	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/pkg/parser"
	"github.com/kcp-dev/code-generator/pkg/util"
)

type FakeGroup struct {
	// Group is the group in this client.
	Group types.GroupVersionInfo

	// Kinds are the kinds in the group.
	Kinds []parser.Kind

	// PackagePath is the package under which this client-set will be exposed.
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	PackagePath string

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string
}

func (g *FakeGroup) WriteContent(w io.Writer) error {
	templ, err := template.New("fakeGroup").Funcs(template.FuncMap{
		"upperFirst": util.UpperFirst,
		"lowerFirst": util.LowerFirst,
		"toLower":    strings.ToLower,
	}).Parse(fakeGroup)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"group":                          g.Group,
		"kinds":                          g.Kinds,
		"packagePath":                    g.PackagePath,
		"singleClusterClientPackagePath": g.SingleClusterClientPackagePath,
	}
	return templ.Execute(w, m)
}

var fakeGroup = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package {{.group.Version.PackageName}}

import (
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/rest"
	kcp{{.group.PackageAlias}} "{{.packagePath}}/typed/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
	{{.group.PackageAlias}} "{{.singleClusterClientPackagePath}}/typed/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
)

var _ kcp{{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}ClusterInterface = (*{{.group.GroupGoName}}{{.group.Version}}ClusterClient)(nil)

type {{.group.GroupGoName}}{{.group.Version}}ClusterClient struct {
	*kcptesting.Fake 
}

func (c *{{.group.GroupGoName}}{{.group.Version}}ClusterClient) Cluster(cluster logicalcluster.Name) {{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &{{.group.GroupGoName}}{{.group.Version}}Client{Fake: c.Fake, Cluster: cluster}
}

{{ range .kinds}}
func (c *{{$.group.GroupGoName}}{{$.group.Version}}ClusterClient) {{.Plural}}() kcp{{$.group.PackageAlias}}.{{.String}}ClusterInterface {
	return &{{.Plural | lowerFirst}}ClusterClient{Fake: c.Fake}
}
{{end -}}

var _ {{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Interface = (*{{.group.GroupGoName}}{{.group.Version}}Client)(nil)

type {{.group.GroupGoName}}{{.group.Version}}Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *{{.group.GroupGoName}}{{.group.Version}}Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

{{ range .kinds}}
func (c *{{$.group.GroupGoName}}{{$.group.Version}}Client) {{.Plural}}({{if .IsNamespaced}}namespace string{{end}}) {{$.group.PackageAlias}}.{{.String}}Interface {
	return &{{.Plural | lowerFirst}}Client{Fake: c.Fake, Cluster: c.Cluster{{if .IsNamespaced}}, Namespace: namespace{{end}}}
}
{{end -}}
`

package listergen

import (
	"io"
	"text/template"

	clientgentypes "k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/pkg/parser"
)

type Lister struct {
	// Group is:
	// - the name of the API group, e.g. "authorization",
	// - the version and package path of the API, e.g. "v1" and "k8s.io/api/rbac/v1"
	Group clientgentypes.GroupVersionInfo
	// Kind is the kind for which we are generating listers, e.g. "ClusterRole"
	Kind parser.Kind

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string

	// SingleClusterListerPackagePath is the fully qualified Go package name under which the (pre-existing)
	// listers for single-cluster contexts are defined. Option. e.g. "k8s.io/client-go/listers"
	SingleClusterListerPackagePath string
}

func (l *Lister) WriteContent(w io.Writer) error {
	templ, err := template.New("lister").Funcs(templateFuncs).Parse(lister)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"group":                          l.Group,
		"kind":                           &l.Kind,
		"apiPackagePath":                 l.APIPackagePath,
		"singleClusterListerPackagePath": l.SingleClusterListerPackagePath,
		"useUpstreamInterfaces":          l.SingleClusterListerPackagePath != "",
	}
	return templ.Execute(w, m)
}

var lister = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package {{.group.Version.PackageName}}

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"	
	"github.com/kcp-dev/logicalcluster/v2"
	
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"

	{{.group.PackageAlias}} "{{.apiPackagePath}}/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
	{{if .useUpstreamInterfaces -}}
	{{.group.PackageAlias}}listers "{{.singleClusterListerPackagePath}}/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
	{{end -}}
)

// {{.kind.String}}ClusterLister can list {{.kind.Plural}} across all workspaces, or scope down to a {{.kind.String}}Lister for one workspace.
type {{.kind.String}}ClusterLister interface {
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
{{if not .useUpstreamInterfaces -}}
	Cluster(cluster logicalcluster.Name) {{.kind.String}}Lister
{{else -}}
	Cluster(cluster logicalcluster.Name){{.group.PackageAlias}}listers.{{.kind.String}}Lister
{{end -}}	
}

type {{.kind.String | lowerFirst }}ClusterLister struct {
	indexer cache.Indexer
}

// New{{.kind.String}}ClusterLister returns a new {{.kind.String}}ClusterLister.
func New{{.kind.String}}ClusterLister(indexer cache.Indexer) *{{.kind.String | lowerFirst}}ClusterLister {
	return &{{.kind.String | lowerFirst}}ClusterLister{indexer: indexer}
}

// List lists all {{.kind.Plural}} in the indexer across all workspaces.
func (s *{{.kind.String | lowerFirst}}ClusterLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get {{.kind.Plural}}.
{{if not .useUpstreamInterfaces -}}
func (s *{{.kind.String | lowerFirst}}ClusterLister) Cluster(cluster logicalcluster.Name) {{.kind.String}}Lister {
{{else -}}
func (s *{{.kind.String | lowerFirst}}ClusterLister) Cluster(cluster logicalcluster.Name){{.group.PackageAlias}}listers.{{.kind.String}}Lister {
{{end -}}
	return &{{.kind.String | lowerFirst}}Lister{indexer: s.indexer, cluster: cluster}
}

{{if not .useUpstreamInterfaces -}}
type {{.kind.String}}Lister interface {
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
{{ if  not .kind.IsNamespaced -}}
	Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error)
{{else -}}
	{{.kind.Plural}}(namespace string) {{.kind.String}}NamespaceLister
{{end -}}
}
{{end -}}

{{if not .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}Lister can list all {{.kind.Plural}} inside a workspace{{ if .kind.IsNamespaced }} or scope down to a {{.kind.String}}Lister for one namespace{{end}}.
{{else -}}
// {{.kind.String | lowerFirst}}Lister implements the {{.group.PackageAlias}}listers.{{.kind.String}}Lister interface.
{{end -}}
type {{.kind.String | lowerFirst}}Lister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all {{.kind.Plural}} in the indexer for a workspace.
func (s *{{.kind.String | lowerFirst}}Lister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*{{.group.PackageAlias}}.{{.kind.String}})
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

{{ if  not .kind.IsNamespaced -}}
// Get retrieves the {{.kind.String}} from the indexer for a given workspace and name.
func (s *{{.kind.String | lowerFirst}}Lister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.String}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ else -}}
// {{.kind.Plural}} returns an object that can list and get {{.kind.Plural}} in one namespace.
{{if not .useUpstreamInterfaces -}}
func (s *{{.kind.String | lowerFirst}}Lister) {{.kind.Plural}}(namespace string) {{.kind.String}}NamespaceLister {
{{else -}}
func (s *{{.kind.String | lowerFirst}}Lister) {{.kind.Plural}}(namespace string) {{.group.PackageAlias}}listers.{{.kind.String}}NamespaceLister {
{{end -}}
	return &{{.kind.String | lowerFirst}}NamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

{{if not .useUpstreamInterfaces -}}
type {{.kind.String}}NamespaceLister interface {
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
	Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error)
	}
{{end -}}

{{ if not .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}NamespaceLister helps list and get {{.kind.Plural}}.
// All objects returned here must be treated as read-only.
{{ end -}}
{{ if .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}NamespaceLister implements the {{.group.PackageAlias}}listers.{{.kind.String}}NamespaceLister interface.
{{ end -}}
type {{.kind.String | lowerFirst}}NamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all {{.kind.Plural}} in the indexer for a given workspace and namespace.
func (s *{{.kind.String | lowerFirst}}NamespaceLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*{{.group.PackageAlias}}.{{.kind.String}})
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the {{.kind.String}} from the indexer for a given workspace, namespace and name.
func (s *{{.kind.String | lowerFirst}}NamespaceLister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.String}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ end -}}
`
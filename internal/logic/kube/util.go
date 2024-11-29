package kube

import (
	"github.com/gogf/gf/v2/errors/gerror"
	kruiseapps "github.com/openkruise/kruise-api/apps/v1alpha1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

func YamlStrToApiObject(strYaml string) (interface{}, error) {
	jsonData, err := yaml.ToJSON([]byte(strYaml))
	if err != nil {
		return nil, gerror.Wrap(err, "str yaml convert to json failed")
	}
	//g.Dump(jsonData)
	// CRD scheme added here
	kruiseapps.AddToScheme(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	apiObject, _, err := decode([]byte(jsonData), nil, nil)
	if err != nil {
		return nil, gerror.Wrap(err, "scheme.codecs.decode failed")
	}
	return apiObject, err
}

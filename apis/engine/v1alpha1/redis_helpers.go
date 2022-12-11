package v1alpha1

import (
	"fmt"

	"kubevault.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
	"kmodules.xyz/client-go/tools/clusterid"
)

func (_ RedisRole) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceRedisRoles))
}

const DefaultRedisDatabasePlugin = "redis-database-plugin"

func (r RedisRole) RoleName() string {
	cluster := "-"
	if clusterid.ClusterName() != "" {
		cluster = clusterid.ClusterName()
	}
	return fmt.Sprintf("k8s.%s.%s.%s", cluster, r.Namespace, r.Name)
}

func (r RedisRole) IsValid() error {
	return nil
}

func (r *RedisConfiguration) SetDefaults() {
	if r == nil {
		return
	}

	// If user doesn't specify the list of AllowedRoles
	// It is set to "*" (allow all)
	if r.AllowedRoles == nil || len(r.AllowedRoles) == 0 {
		r.AllowedRoles = []string{"*"}
	}

	if r.PluginName == "" {
		r.PluginName = DefaultRedisDatabasePlugin
	}
}

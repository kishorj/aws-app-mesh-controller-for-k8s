package virtualservice

import (
	appmesh "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	"github.com/aws/aws-app-mesh-controller-for-k8s/pkg/references"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ReferenceKindVirtualNode   = "VirtualNode"
	ReferenceKindVirtualRouter = "VirtualRouter"
)

func ExtractVirtualNodeReferences(vs *appmesh.VirtualService) []appmesh.VirtualNodeReference {
	if vs.Spec.Provider == nil || vs.Spec.Provider.VirtualNode == nil {
		return nil
	}
	vnRef := vs.Spec.Provider.VirtualNode.VirtualNodeRef
	return []appmesh.VirtualNodeReference{vnRef}
}

func ExtractVirtualRouterReferences(vs *appmesh.VirtualService) []appmesh.VirtualRouterReference {
	if vs.Spec.Provider == nil || vs.Spec.Provider.VirtualRouter == nil {
		return nil
	}
	vrRef := vs.Spec.Provider.VirtualRouter.VirtualRouterRef
	return []appmesh.VirtualRouterReference{vrRef}
}

func VirtualNodeReferenceIndexFunc(obj runtime.Object) []types.NamespacedName {
	vs := obj.(*appmesh.VirtualService)
	vnRefs := ExtractVirtualNodeReferences(vs)

	var vnKeys []types.NamespacedName
	for _, vnRef := range vnRefs {
		vnKey := references.ObjectKeyForVirtualNodeReference(vs, vnRef)
		vnKeys = append(vnKeys, vnKey)
	}
	return vnKeys
}

func VirtualRouterReferenceIndexFunc(obj runtime.Object) []types.NamespacedName {
	vs := obj.(*appmesh.VirtualService)
	vrRefs := ExtractVirtualRouterReferences(vs)

	var vrKeys []types.NamespacedName
	for _, vrRef := range vrRefs {
		vrKey := references.ObjectKeyForVirtualRouterReference(vs, vrRef)
		vrKeys = append(vrKeys, vrKey)
	}
	return vrKeys
}

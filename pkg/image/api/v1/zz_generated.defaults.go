// +build !ignore_autogenerated_openshift

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/kubernetes/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ImageStream{}, func(obj interface{}) { SetObjectDefaults_ImageStream(obj.(*ImageStream)) })
	scheme.AddTypeDefaultingFunc(&ImageStreamImport{}, func(obj interface{}) { SetObjectDefaults_ImageStreamImport(obj.(*ImageStreamImport)) })
	scheme.AddTypeDefaultingFunc(&ImageStreamList{}, func(obj interface{}) { SetObjectDefaults_ImageStreamList(obj.(*ImageStreamList)) })
	scheme.AddTypeDefaultingFunc(&ImageStreamTag{}, func(obj interface{}) { SetObjectDefaults_ImageStreamTag(obj.(*ImageStreamTag)) })
	scheme.AddTypeDefaultingFunc(&ImageStreamTagList{}, func(obj interface{}) { SetObjectDefaults_ImageStreamTagList(obj.(*ImageStreamTagList)) })
	return nil
}

func SetObjectDefaults_ImageStream(in *ImageStream) {
	for i := range in.Spec.Tags {
		a := &in.Spec.Tags[i]
		SetDefaults_TagReferencePolicy(&a.ReferencePolicy)
	}
}

func SetObjectDefaults_ImageStreamImport(in *ImageStreamImport) {
	if in.Spec.Repository != nil {
		SetDefaults_TagReferencePolicy(&in.Spec.Repository.ReferencePolicy)
	}
	for i := range in.Spec.Images {
		a := &in.Spec.Images[i]
		SetDefaults_ImageImportSpec(a)
		SetDefaults_TagReferencePolicy(&a.ReferencePolicy)
	}
	if in.Status.Import != nil {
		SetObjectDefaults_ImageStream(in.Status.Import)
	}
}

func SetObjectDefaults_ImageStreamList(in *ImageStreamList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ImageStream(a)
	}
}

func SetObjectDefaults_ImageStreamTag(in *ImageStreamTag) {
	if in.Tag != nil {
		SetDefaults_TagReferencePolicy(&in.Tag.ReferencePolicy)
	}
}

func SetObjectDefaults_ImageStreamTagList(in *ImageStreamTagList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ImageStreamTag(a)
	}
}

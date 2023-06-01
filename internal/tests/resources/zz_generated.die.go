//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 VMware Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by diegen. DO NOT EDIT.

package resources

import (
	"dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	meta "github.com/fluxcd/pkg/apis/meta"
	"github.com/garethjevans/monorepository-controller/api/v1alpha1"
	apis "github.com/vmware-labs/reconciler-runtime/apis"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	jsonpath "k8s.io/client-go/util/jsonpath"
	osx "os"
	reflectx "reflect"
	yaml "sigs.k8s.io/yaml"
)

var MonoRepositoryBlank = (&MonoRepositoryDie{}).DieFeed(v1alpha1.MonoRepository{})

type MonoRepositoryDie struct {
	v1.FrozenObjectMeta
	mutable bool
	r       v1alpha1.MonoRepository
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *MonoRepositoryDie) DieImmutable(immutable bool) *MonoRepositoryDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *MonoRepositoryDie) DieFeed(r v1alpha1.MonoRepository) *MonoRepositoryDie {
	if d.mutable {
		d.FrozenObjectMeta = v1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &MonoRepositoryDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *MonoRepositoryDie) DieFeedPtr(r *v1alpha1.MonoRepository) *MonoRepositoryDie {
	if r == nil {
		r = &v1alpha1.MonoRepository{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *MonoRepositoryDie) DieFeedJSON(j []byte) *MonoRepositoryDie {
	r := v1alpha1.MonoRepository{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *MonoRepositoryDie) DieFeedYAML(y []byte) *MonoRepositoryDie {
	r := v1alpha1.MonoRepository{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *MonoRepositoryDie) DieFeedYAMLFile(name string) *MonoRepositoryDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositoryDie) DieFeedRawExtension(raw runtime.RawExtension) *MonoRepositoryDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *MonoRepositoryDie) DieRelease() v1alpha1.MonoRepository {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *MonoRepositoryDie) DieReleasePtr() *v1alpha1.MonoRepository {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object. Panics on error.
func (d *MonoRepositoryDie) DieReleaseUnstructured() *unstructured.Unstructured {
	r := d.DieReleasePtr()
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *MonoRepositoryDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *MonoRepositoryDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositoryDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *MonoRepositoryDie) DieStamp(fn func(r *v1alpha1.MonoRepository)) *MonoRepositoryDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *MonoRepositoryDie) DieStampAt(jp string, fn interface{}) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *MonoRepositoryDie) DeepCopy() *MonoRepositoryDie {
	r := *d.r.DeepCopy()
	return &MonoRepositoryDie{
		FrozenObjectMeta: v1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*MonoRepositoryDie)(nil)

func (d *MonoRepositoryDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *MonoRepositoryDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *MonoRepositoryDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *MonoRepositoryDie) UnmarshalJSON(b []byte) error {
	if d == MonoRepositoryBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &v1alpha1.MonoRepository{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *MonoRepositoryDie) APIVersion(v string) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *MonoRepositoryDie) Kind(v string) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *MonoRepositoryDie) MetadataDie(fn func(d *v1.ObjectMetaDie)) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		d := v1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *MonoRepositoryDie) SpecDie(fn func(d *MonoRepositorySpecDie)) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		d := MonoRepositorySpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

// StatusDie stamps the resource's status field with a mutable die.
func (d *MonoRepositoryDie) StatusDie(fn func(d *MonoRepositoryStatusDie)) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		d := MonoRepositoryStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

func (d *MonoRepositoryDie) Spec(v v1alpha1.MonoRepositorySpec) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		r.Spec = v
	})
}

func (d *MonoRepositoryDie) Status(v v1alpha1.MonoRepositoryStatus) *MonoRepositoryDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepository) {
		r.Status = v
	})
}

var MonoRepositorySpecBlank = (&MonoRepositorySpecDie{}).DieFeed(v1alpha1.MonoRepositorySpec{})

type MonoRepositorySpecDie struct {
	mutable bool
	r       v1alpha1.MonoRepositorySpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *MonoRepositorySpecDie) DieImmutable(immutable bool) *MonoRepositorySpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *MonoRepositorySpecDie) DieFeed(r v1alpha1.MonoRepositorySpec) *MonoRepositorySpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &MonoRepositorySpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *MonoRepositorySpecDie) DieFeedPtr(r *v1alpha1.MonoRepositorySpec) *MonoRepositorySpecDie {
	if r == nil {
		r = &v1alpha1.MonoRepositorySpec{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *MonoRepositorySpecDie) DieFeedJSON(j []byte) *MonoRepositorySpecDie {
	r := v1alpha1.MonoRepositorySpec{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *MonoRepositorySpecDie) DieFeedYAML(y []byte) *MonoRepositorySpecDie {
	r := v1alpha1.MonoRepositorySpec{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *MonoRepositorySpecDie) DieFeedYAMLFile(name string) *MonoRepositorySpecDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositorySpecDie) DieFeedRawExtension(raw runtime.RawExtension) *MonoRepositorySpecDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *MonoRepositorySpecDie) DieRelease() v1alpha1.MonoRepositorySpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *MonoRepositorySpecDie) DieReleasePtr() *v1alpha1.MonoRepositorySpec {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *MonoRepositorySpecDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *MonoRepositorySpecDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositorySpecDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *MonoRepositorySpecDie) DieStamp(fn func(r *v1alpha1.MonoRepositorySpec)) *MonoRepositorySpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *MonoRepositorySpecDie) DieStampAt(jp string, fn interface{}) *MonoRepositorySpecDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositorySpec) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *MonoRepositorySpecDie) DeepCopy() *MonoRepositorySpecDie {
	r := *d.r.DeepCopy()
	return &MonoRepositorySpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *MonoRepositorySpecDie) SourceRef(v v1alpha1.SourceRef) *MonoRepositorySpecDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositorySpec) {
		r.SourceRef = v
	})
}

func (d *MonoRepositorySpecDie) Include(v string) *MonoRepositorySpecDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositorySpec) {
		r.Include = v
	})
}

var MonoRepositoryStatusBlank = (&MonoRepositoryStatusDie{}).DieFeed(v1alpha1.MonoRepositoryStatus{})

type MonoRepositoryStatusDie struct {
	mutable bool
	r       v1alpha1.MonoRepositoryStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *MonoRepositoryStatusDie) DieImmutable(immutable bool) *MonoRepositoryStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *MonoRepositoryStatusDie) DieFeed(r v1alpha1.MonoRepositoryStatus) *MonoRepositoryStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &MonoRepositoryStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *MonoRepositoryStatusDie) DieFeedPtr(r *v1alpha1.MonoRepositoryStatus) *MonoRepositoryStatusDie {
	if r == nil {
		r = &v1alpha1.MonoRepositoryStatus{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *MonoRepositoryStatusDie) DieFeedJSON(j []byte) *MonoRepositoryStatusDie {
	r := v1alpha1.MonoRepositoryStatus{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *MonoRepositoryStatusDie) DieFeedYAML(y []byte) *MonoRepositoryStatusDie {
	r := v1alpha1.MonoRepositoryStatus{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *MonoRepositoryStatusDie) DieFeedYAMLFile(name string) *MonoRepositoryStatusDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositoryStatusDie) DieFeedRawExtension(raw runtime.RawExtension) *MonoRepositoryStatusDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *MonoRepositoryStatusDie) DieRelease() v1alpha1.MonoRepositoryStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *MonoRepositoryStatusDie) DieReleasePtr() *v1alpha1.MonoRepositoryStatus {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *MonoRepositoryStatusDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *MonoRepositoryStatusDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *MonoRepositoryStatusDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *MonoRepositoryStatusDie) DieStamp(fn func(r *v1alpha1.MonoRepositoryStatus)) *MonoRepositoryStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *MonoRepositoryStatusDie) DieStampAt(jp string, fn interface{}) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			args := []reflectx.Value{cv}
			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *MonoRepositoryStatusDie) DeepCopy() *MonoRepositoryStatusDie {
	r := *d.r.DeepCopy()
	return &MonoRepositoryStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *MonoRepositoryStatusDie) Status(v apis.Status) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.Status = v
	})
}

// URL is the dynamic fetch link for the latest Artifact. It is provided on a "best effort" basis, and using the precise GitRepositoryStatus.Artifact data is recommended.
func (d *MonoRepositoryStatusDie) URL(v string) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.URL = v
	})
}

// Artifact represents the last successful GitRepository reconciliation.
func (d *MonoRepositoryStatusDie) Artifact(v *v1alpha1.Artifact) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.Artifact = v
	})
}

// ObservedInclude is the observed list of GitRepository resources used to calculate the checksum for this artifact
func (d *MonoRepositoryStatusDie) ObservedInclude(v string) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.ObservedInclude = v
	})
}

// ObservedFileList is the file list used to calculate the checksum for this artifact
func (d *MonoRepositoryStatusDie) ObservedFileList(v string) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.ObservedFileList = v
	})
}

func (d *MonoRepositoryStatusDie) ReconcileRequestStatus(v meta.ReconcileRequestStatus) *MonoRepositoryStatusDie {
	return d.DieStamp(func(r *v1alpha1.MonoRepositoryStatus) {
		r.ReconcileRequestStatus = v
	})
}

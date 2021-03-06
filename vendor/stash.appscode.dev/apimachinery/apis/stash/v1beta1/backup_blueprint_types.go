/*
Copyright The Stash Authors.

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

package v1beta1

import (
	"stash.appscode.dev/apimachinery/apis/stash/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

const (
	ResourceKindBackupBlueprint     = "BackupBlueprint"
	ResourcePluralBackupBlueprint   = "backupblueprints"
	ResourceSingularBackupBlueprint = "backupblueprint"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=backupblueprints,singular=backupblueprint,scope=Cluster,shortName=bb,categories={stash,appscode}
// +kubebuilder:printcolumn:name="Task",type="string",JSONPath=".spec.task.name"
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type BackupBlueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              BackupBlueprintSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type BackupBlueprintSpec struct {
	// RepositorySpec is used to create Repository crd for respective workload
	v1alpha1.RepositorySpec `json:",inline" protobuf:"bytes,1,opt,name=repositorySpec"`
	Schedule                string `json:"schedule,omitempty" protobuf:"bytes,2,opt,name=schedule"`
	// Task specify the Task crd that specifies steps for backup process
	// +optional
	Task TaskRef `json:"task,omitempty" protobuf:"bytes,3,opt,name=task"`
	// RetentionPolicy indicates the policy to follow to clean old backup snapshots
	RetentionPolicy v1alpha1.RetentionPolicy `json:"retentionPolicy" protobuf:"bytes,4,opt,name=retentionPolicy"`
	// RuntimeSettings allow to specify Resources, NodeSelector, Affinity, Toleration, ReadinessProbe etc.
	// +optional
	RuntimeSettings ofst.RuntimeSettings `json:"runtimeSettings,omitempty" protobuf:"bytes,5,opt,name=runtimeSettings"`
	// Temp directory configuration for functions/sidecar
	// An `EmptyDir` will always be mounted at /tmp with this settings
	// +optional
	TempDir EmptyDirSettings `json:"tempDir,omitempty" protobuf:"bytes,6,opt,name=tempDir"`
	// InterimVolumeTemplate specifies a template for a volume to hold targeted data temporarily
	// before uploading to backend or inserting into target. It is only usable for job model.
	// Don't specify it in sidecar model.
	// +optional
	InterimVolumeTemplate *ofst.PersistentVolumeClaim `json:"interimVolumeTemplate,omitempty" protobuf:"bytes,7,opt,name=interimVolumeTemplate"`
	// BackupHistoryLimit specifies the number of BackupSession and it's associate resources to keep.
	// This is helpful for debugging purpose.
	// Default: 1
	// +optional
	BackupHistoryLimit *int32 `json:"backupHistoryLimit,omitempty" protobuf:"varint,8,opt,name=backupHistoryLimit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackupBlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BackupBlueprint `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

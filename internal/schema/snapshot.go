/*
copyright 2020 the Escli authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package schema

import "time"

type SnapshotSnapshotsMetadata struct {
	SnapshotSnapshotMetadata []SnapshotSnapshotMetadata `json:"snapshots"`
}

type SnapshotSnapshotMetadata struct {
	Name               string        `json:"name,omitempty"`
	Snapshot           string        `json:"snapshot,omitempty"`
	UUID               string        `json:"uuid,omitempty"`
	VersionID          int           `json:"version_id,omitempty"`
	Version            string        `json:"version,omitempty"`
	Indices            []string      `json:"indices,omitempty"`
	IncludeGlobalState bool          `json:"include_global_state,omitempty"`
	State              string        `json:"state,omitempty"`
	StartTime          time.Time     `json:"start_time,omitempty"`
	StartTimeInMillis  int64         `json:"start_time_in_millis,omitempty"`
	EndTime            time.Time     `json:"end_time,omitempty"`
	EndTimeInMillis    int64         `json:"end_time_in_millis,omitempty"`
	DurationInMillis   int           `json:"duration_in_millis,omitempty"`
	Failures           []interface{} `json:"failures,omitempty"`
	Shards             struct {
		Total      int `json:"total"`
		Failed     int `json:"failed"`
		Successful int `json:"successful"`
	} `json:"shards,omitempty"`
}

type SnapshotRequestBody struct {
	Indices string `json:"indices"`
}

type SnapshotRepositoryMetadata struct {
	Type string `json:"type"`
	Settings SnapshotRepositorySettingsMetadata `json:"settings"`
}

type SnapshotRepositorySettingsMetadata struct {
	Bucket string `json:"bucket"`
	BasePath string `json:"base_path"`
	StorageClass string `json:"storage_class"`
	Region string `json:"region"`
}

type SnapshotSnapshotsIndicesMetadata struct {
	Snapshots []SnapshotSnapshotS3Metadata `json:"snapshots"`
	Indices map[string] SnapshotIndexMetadata `json:"indices"`
}

type SnapshotIndexMetadata struct {
	ID string `json:"id"`
	Snapshots []string `json:"snapshots"`
}

type SnapshotSnapshotS3Metadata struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	State int `json:"state"`
}

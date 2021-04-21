// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START pubsub_generated_pubsublite_AdminClient_UpdateTopic]

package main

import (
	"context"
	"time"

	"cloud.google.com/go/pubsublite"
)

func main() {
	ctx := context.Background()
	// NOTE: region must correspond to the zone of the topic.
	admin, err := pubsublite.NewAdminClient(ctx, "region")
	if err != nil {
		// TODO: Handle error.
	}

	updateConfig := pubsublite.TopicConfigToUpdate{
		Name:                       "projects/my-project/locations/zone/topics/my-topic",
		PartitionCount:             3, // Only increases currently supported.
		PublishCapacityMiBPerSec:   8,
		SubscribeCapacityMiBPerSec: 16,
		RetentionDuration:          24 * time.Hour, // Garbage collect messages older than 24 hours.
	}
	_, err = admin.UpdateTopic(ctx, updateConfig)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END pubsub_generated_pubsublite_AdminClient_UpdateTopic]
package main

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/google/logger"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
)

// listAllInstances prints all instances present in a project, grouped by their zone.
func ListAllInstances(logger *logger.Logger) error {
	// projectID := "your_project_id"
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	// Use the `MaxResults` parameter to limit the number of results that the API returns per response page.
	req := &computepb.AggregatedListInstancesRequest{
		Project:    "devcastops",
		MaxResults: proto.Uint32(3),
	}

	it := instancesClient.AggregatedList(ctx, req)
	logger.Infof("Instances found:\n")
	// Despite using the `MaxResults` parameter, you don't need to handle the pagination
	// yourself. The returned iterator object handles pagination
	// automatically, returning separated pages as you iterate over the results.
	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		instances := pair.Value.Instances
		if len(instances) > 0 {
			logger.Infof("%s\n", pair.Key)
			for _, instance := range instances {
				logger.Infof("- %s %s\n", instance.GetName(), instance.GetMachineType())
			}
		}
	}
	return nil
}

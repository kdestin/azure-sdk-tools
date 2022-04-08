//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package test_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/CreateOrUpdateASimpleGalleryImageVersionWithVMAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := test.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		test.GalleryImageVersion{
			Location: to.Ptr("<location>"),
			Properties: &test.GalleryImageVersionProperties{
				PublishingProfile: &test.GalleryImageVersionPublishingProfile{
					TargetRegions: []*test.TargetRegion{
						{
							Name: to.Ptr("<name>"),
							Encryption: &test.EncryptionImages{
								DataDiskImages: []*test.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &test.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name: to.Ptr("<name>"),
							Encryption: &test.EncryptionImages{
								DataDiskImages: []*test.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &test.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(test.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &test.GalleryImageVersionStorageProfile{
					Source: &test.GalleryArtifactVersionSource{
						ID: to.Ptr("<id>"),
					},
				},
			},
		},
		&test.GalleryImageVersionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/UpdateASimpleGalleryImageVersion.json
func ExampleGalleryImageVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := test.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		test.GalleryImageVersionUpdate{
			Properties: &test.GalleryImageVersionProperties{
				PublishingProfile: &test.GalleryImageVersionPublishingProfile{
					TargetRegions: []*test.TargetRegion{
						{
							Name:                 to.Ptr("<name>"),
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name:                 to.Ptr("<name>"),
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(test.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &test.GalleryImageVersionStorageProfile{
					Source: &test.GalleryArtifactVersionSource{
						ID: to.Ptr("<id>"),
					},
				},
			},
		},
		&test.GalleryImageVersionsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/GetAGalleryImageVersionWithReplicationStatus.json
func ExampleGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := test.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&test.GalleryImageVersionsClientGetOptions{Expand: to.Ptr(test.ReplicationStatusTypesReplicationStatus)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/DeleteAGalleryImageVersion.json
func ExampleGalleryImageVersionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := test.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&test.GalleryImageVersionsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/ListGalleryImageVersionsInAGalleryImage.json
func ExampleGalleryImageVersionsClient_ListByGalleryImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := test.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByGalleryImage("<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

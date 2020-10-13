// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package vpc_link

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.VPCLink{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetVpcLinkWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.Name != nil {
		ko.Spec.Name = resp.Name
	}
	if resp.SecurityGroupIds != nil {
		f2 := []*string{}
		for _, f2iter := range resp.SecurityGroupIds {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		ko.Spec.SecurityGroupIDs = f2
	}
	if resp.SubnetIds != nil {
		f3 := []*string{}
		for _, f3iter := range resp.SubnetIds {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		ko.Spec.SubnetIDs = f3
	}
	if resp.Tags != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.Tags {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		ko.Spec.Tags = f4
	}
	if resp.VpcLinkId != nil {
		ko.Status.VPCLinkID = resp.VpcLinkId
	}
	if resp.VpcLinkStatus != nil {
		ko.Status.VPCLinkStatus = resp.VpcLinkStatus
	}
	if resp.VpcLinkStatusMessage != nil {
		ko.Status.VPCLinkStatusMessage = resp.VpcLinkStatusMessage
	}
	if resp.VpcLinkVersion != nil {
		ko.Status.VPCLinkVersion = resp.VpcLinkVersion
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.VPCLinkID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetVpcLinkInput, error) {
	res := &svcsdk.GetVpcLinkInput{}

	if r.ko.Status.VPCLinkID != nil {
		res.SetVpcLinkId(*r.ko.Status.VPCLinkID)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetVpcLinksInput, error) {
	res := &svcsdk.GetVpcLinksInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateVpcLinkWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.VpcLinkId != nil {
		ko.Status.VPCLinkID = resp.VpcLinkId
	}
	if resp.VpcLinkStatus != nil {
		ko.Status.VPCLinkStatus = resp.VpcLinkStatus
	}
	if resp.VpcLinkStatusMessage != nil {
		ko.Status.VPCLinkStatusMessage = resp.VpcLinkStatusMessage
	}
	if resp.VpcLinkVersion != nil {
		ko.Status.VPCLinkVersion = resp.VpcLinkVersion
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateVpcLinkInput, error) {
	res := &svcsdk.CreateVpcLinkInput{}

	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f1 := []*string{}
		for _, f1iter := range r.ko.Spec.SecurityGroupIDs {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetSecurityGroupIds(f1)
	}
	if r.ko.Spec.SubnetIDs != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.SubnetIDs {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetSubnetIds(f2)
	}
	if r.ko.Spec.Tags != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range r.ko.Spec.Tags {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		res.SetTags(f3)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateVpcLinkWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.VpcLinkId != nil {
		ko.Status.VPCLinkID = resp.VpcLinkId
	}
	if resp.VpcLinkStatus != nil {
		ko.Status.VPCLinkStatus = resp.VpcLinkStatus
	}
	if resp.VpcLinkStatusMessage != nil {
		ko.Status.VPCLinkStatusMessage = resp.VpcLinkStatusMessage
	}
	if resp.VpcLinkVersion != nil {
		ko.Status.VPCLinkVersion = resp.VpcLinkVersion
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateVpcLinkInput, error) {
	res := &svcsdk.UpdateVpcLinkInput{}

	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Status.VPCLinkID != nil {
		res.SetVpcLinkId(*r.ko.Status.VPCLinkID)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteVpcLinkWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteVpcLinkInput, error) {
	res := &svcsdk.DeleteVpcLinkInput{}

	if r.ko.Status.VPCLinkID != nil {
		res.SetVpcLinkId(*r.ko.Status.VPCLinkID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.VPCLink,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

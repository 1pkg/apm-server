// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeSSLPoliciesInput
type DescribeSSLPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `type:"string"`

	// The names of the policies.
	Names []string `type:"list"`

	// The maximum number of results to return with this call.
	PageSize *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DescribeSSLPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSSLPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSSLPoliciesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeSSLPoliciesOutput
type DescribeSSLPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string `type:"string"`

	// Information about the policies.
	SslPolicies []SslPolicy `type:"list"`
}

// String returns the string representation
func (s DescribeSSLPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSSLPolicies = "DescribeSSLPolicies"

// DescribeSSLPoliciesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the specified policies or all policies used for SSL negotiation.
//
// For more information, see Security Policies (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
// in the Application Load Balancers Guide.
//
//    // Example sending a request using DescribeSSLPoliciesRequest.
//    req := client.DescribeSSLPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeSSLPolicies
func (c *Client) DescribeSSLPoliciesRequest(input *DescribeSSLPoliciesInput) DescribeSSLPoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeSSLPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSSLPoliciesInput{}
	}

	req := c.newRequest(op, input, &DescribeSSLPoliciesOutput{})
	return DescribeSSLPoliciesRequest{Request: req, Input: input, Copy: c.DescribeSSLPoliciesRequest}
}

// DescribeSSLPoliciesRequest is the request type for the
// DescribeSSLPolicies API operation.
type DescribeSSLPoliciesRequest struct {
	*aws.Request
	Input *DescribeSSLPoliciesInput
	Copy  func(*DescribeSSLPoliciesInput) DescribeSSLPoliciesRequest
}

// Send marshals and sends the DescribeSSLPolicies API request.
func (r DescribeSSLPoliciesRequest) Send(ctx context.Context) (*DescribeSSLPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSSLPoliciesResponse{
		DescribeSSLPoliciesOutput: r.Request.Data.(*DescribeSSLPoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSSLPoliciesResponse is the response type for the
// DescribeSSLPolicies API operation.
type DescribeSSLPoliciesResponse struct {
	*DescribeSSLPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSSLPolicies request.
func (r *DescribeSSLPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

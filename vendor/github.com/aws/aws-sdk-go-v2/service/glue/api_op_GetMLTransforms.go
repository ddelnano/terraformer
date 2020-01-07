// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMLTransformsInput struct {
	_ struct{} `type:"structure"`

	// The filter transformation criteria.
	Filter *TransformFilterCriteria `type:"structure"`

	// The maximum number of results to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A paginated token to offset the results.
	NextToken *string `type:"string"`

	// The sorting criteria.
	Sort *TransformSortCriteria `type:"structure"`
}

// String returns the string representation
func (s GetMLTransformsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMLTransformsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMLTransformsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sort != nil {
		if err := s.Sort.Validate(); err != nil {
			invalidParams.AddNested("Sort", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMLTransformsOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token, if more results are available.
	NextToken *string `type:"string"`

	// A list of machine learning transforms.
	//
	// Transforms is a required field
	Transforms []MLTransform `type:"list" required:"true"`
}

// String returns the string representation
func (s GetMLTransformsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMLTransforms = "GetMLTransforms"

// GetMLTransformsRequest returns a request value for making API operation for
// AWS Glue.
//
// Gets a sortable, filterable list of existing AWS Glue machine learning transforms.
// Machine learning transforms are a special type of transform that use machine
// learning to learn the details of the transformation to be performed by learning
// from examples provided by humans. These transformations are then saved by
// AWS Glue, and you can retrieve their metadata by calling GetMLTransforms.
//
//    // Example sending a request using GetMLTransformsRequest.
//    req := client.GetMLTransformsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMLTransforms
func (c *Client) GetMLTransformsRequest(input *GetMLTransformsInput) GetMLTransformsRequest {
	op := &aws.Operation{
		Name:       opGetMLTransforms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetMLTransformsInput{}
	}

	req := c.newRequest(op, input, &GetMLTransformsOutput{})
	return GetMLTransformsRequest{Request: req, Input: input, Copy: c.GetMLTransformsRequest}
}

// GetMLTransformsRequest is the request type for the
// GetMLTransforms API operation.
type GetMLTransformsRequest struct {
	*aws.Request
	Input *GetMLTransformsInput
	Copy  func(*GetMLTransformsInput) GetMLTransformsRequest
}

// Send marshals and sends the GetMLTransforms API request.
func (r GetMLTransformsRequest) Send(ctx context.Context) (*GetMLTransformsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMLTransformsResponse{
		GetMLTransformsOutput: r.Request.Data.(*GetMLTransformsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetMLTransformsRequestPaginator returns a paginator for GetMLTransforms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetMLTransformsRequest(input)
//   p := glue.NewGetMLTransformsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetMLTransformsPaginator(req GetMLTransformsRequest) GetMLTransformsPaginator {
	return GetMLTransformsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetMLTransformsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetMLTransformsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetMLTransformsPaginator struct {
	aws.Pager
}

func (p *GetMLTransformsPaginator) CurrentPage() *GetMLTransformsOutput {
	return p.Pager.CurrentPage().(*GetMLTransformsOutput)
}

// GetMLTransformsResponse is the response type for the
// GetMLTransforms API operation.
type GetMLTransformsResponse struct {
	*GetMLTransformsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMLTransforms request.
func (r *GetMLTransformsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAPIUsageHitCountURL generates an URL for the get API usage hit count operation
type GetAPIUsageHitCountURL struct {
	DestinationIPIsNot   []string
	DestinationIPIs      []string
	DestinationPortIsNot []string
	DestinationPortIs    []string
	EndTime              strfmt.DateTime
	HasSpecDiffIs        *bool
	MethodIs             []string
	PathContains         []string
	PathEnd              *string
	PathIsNot            []string
	PathIs               []string
	PathStart            *string
	ShowNonAPI           bool
	SourceIPIsNot        []string
	SourceIPIs           []string
	SpecDiffTypeIs       []string
	SpecContains         []string
	SpecEnd              *string
	SpecIsNot            []string
	SpecIs               []string
	SpecStart            *string
	StartTime            strfmt.DateTime
	StatusCodeGte        *string
	StatusCodeIsNot      []string
	StatusCodeIs         []string
	StatusCodeLte        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAPIUsageHitCountURL) WithBasePath(bp string) *GetAPIUsageHitCountURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAPIUsageHitCountURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAPIUsageHitCountURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apiUsage/hitCount"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var destinationIPIsNotIR []string
	for _, destinationIPIsNotI := range o.DestinationIPIsNot {
		destinationIPIsNotIS := destinationIPIsNotI
		if destinationIPIsNotIS != "" {
			destinationIPIsNotIR = append(destinationIPIsNotIR, destinationIPIsNotIS)
		}
	}

	destinationIPIsNot := swag.JoinByFormat(destinationIPIsNotIR, "")

	if len(destinationIPIsNot) > 0 {
		qsv := destinationIPIsNot[0]
		if qsv != "" {
			qs.Set("destinationIP[isNot]", qsv)
		}
	}

	var destinationIPIsIR []string
	for _, destinationIPIsI := range o.DestinationIPIs {
		destinationIPIsIS := destinationIPIsI
		if destinationIPIsIS != "" {
			destinationIPIsIR = append(destinationIPIsIR, destinationIPIsIS)
		}
	}

	destinationIPIs := swag.JoinByFormat(destinationIPIsIR, "")

	if len(destinationIPIs) > 0 {
		qsv := destinationIPIs[0]
		if qsv != "" {
			qs.Set("destinationIP[is]", qsv)
		}
	}

	var destinationPortIsNotIR []string
	for _, destinationPortIsNotI := range o.DestinationPortIsNot {
		destinationPortIsNotIS := destinationPortIsNotI
		if destinationPortIsNotIS != "" {
			destinationPortIsNotIR = append(destinationPortIsNotIR, destinationPortIsNotIS)
		}
	}

	destinationPortIsNot := swag.JoinByFormat(destinationPortIsNotIR, "")

	if len(destinationPortIsNot) > 0 {
		qsv := destinationPortIsNot[0]
		if qsv != "" {
			qs.Set("destinationPort[isNot]", qsv)
		}
	}

	var destinationPortIsIR []string
	for _, destinationPortIsI := range o.DestinationPortIs {
		destinationPortIsIS := destinationPortIsI
		if destinationPortIsIS != "" {
			destinationPortIsIR = append(destinationPortIsIR, destinationPortIsIS)
		}
	}

	destinationPortIs := swag.JoinByFormat(destinationPortIsIR, "")

	if len(destinationPortIs) > 0 {
		qsv := destinationPortIs[0]
		if qsv != "" {
			qs.Set("destinationPort[is]", qsv)
		}
	}

	endTimeQ := o.EndTime.String()
	if endTimeQ != "" {
		qs.Set("endTime", endTimeQ)
	}

	var hasSpecDiffIsQ string
	if o.HasSpecDiffIs != nil {
		hasSpecDiffIsQ = swag.FormatBool(*o.HasSpecDiffIs)
	}
	if hasSpecDiffIsQ != "" {
		qs.Set("hasSpecDiff[is]", hasSpecDiffIsQ)
	}

	var methodIsIR []string
	for _, methodIsI := range o.MethodIs {
		methodIsIS := methodIsI
		if methodIsIS != "" {
			methodIsIR = append(methodIsIR, methodIsIS)
		}
	}

	methodIs := swag.JoinByFormat(methodIsIR, "")

	if len(methodIs) > 0 {
		qsv := methodIs[0]
		if qsv != "" {
			qs.Set("method[is]", qsv)
		}
	}

	var pathContainsIR []string
	for _, pathContainsI := range o.PathContains {
		pathContainsIS := pathContainsI
		if pathContainsIS != "" {
			pathContainsIR = append(pathContainsIR, pathContainsIS)
		}
	}

	pathContains := swag.JoinByFormat(pathContainsIR, "")

	if len(pathContains) > 0 {
		qsv := pathContains[0]
		if qsv != "" {
			qs.Set("path[contains]", qsv)
		}
	}

	var pathEndQ string
	if o.PathEnd != nil {
		pathEndQ = *o.PathEnd
	}
	if pathEndQ != "" {
		qs.Set("path[end]", pathEndQ)
	}

	var pathIsNotIR []string
	for _, pathIsNotI := range o.PathIsNot {
		pathIsNotIS := pathIsNotI
		if pathIsNotIS != "" {
			pathIsNotIR = append(pathIsNotIR, pathIsNotIS)
		}
	}

	pathIsNot := swag.JoinByFormat(pathIsNotIR, "")

	if len(pathIsNot) > 0 {
		qsv := pathIsNot[0]
		if qsv != "" {
			qs.Set("path[isNot]", qsv)
		}
	}

	var pathIsIR []string
	for _, pathIsI := range o.PathIs {
		pathIsIS := pathIsI
		if pathIsIS != "" {
			pathIsIR = append(pathIsIR, pathIsIS)
		}
	}

	pathIs := swag.JoinByFormat(pathIsIR, "")

	if len(pathIs) > 0 {
		qsv := pathIs[0]
		if qsv != "" {
			qs.Set("path[is]", qsv)
		}
	}

	var pathStartQ string
	if o.PathStart != nil {
		pathStartQ = *o.PathStart
	}
	if pathStartQ != "" {
		qs.Set("path[start]", pathStartQ)
	}

	showNonAPIQ := swag.FormatBool(o.ShowNonAPI)
	if showNonAPIQ != "" {
		qs.Set("showNonApi", showNonAPIQ)
	}

	var sourceIPIsNotIR []string
	for _, sourceIPIsNotI := range o.SourceIPIsNot {
		sourceIPIsNotIS := sourceIPIsNotI
		if sourceIPIsNotIS != "" {
			sourceIPIsNotIR = append(sourceIPIsNotIR, sourceIPIsNotIS)
		}
	}

	sourceIPIsNot := swag.JoinByFormat(sourceIPIsNotIR, "")

	if len(sourceIPIsNot) > 0 {
		qsv := sourceIPIsNot[0]
		if qsv != "" {
			qs.Set("sourceIP[isNot]", qsv)
		}
	}

	var sourceIPIsIR []string
	for _, sourceIPIsI := range o.SourceIPIs {
		sourceIPIsIS := sourceIPIsI
		if sourceIPIsIS != "" {
			sourceIPIsIR = append(sourceIPIsIR, sourceIPIsIS)
		}
	}

	sourceIPIs := swag.JoinByFormat(sourceIPIsIR, "")

	if len(sourceIPIs) > 0 {
		qsv := sourceIPIs[0]
		if qsv != "" {
			qs.Set("sourceIP[is]", qsv)
		}
	}

	var specDiffTypeIsIR []string
	for _, specDiffTypeIsI := range o.SpecDiffTypeIs {
		specDiffTypeIsIS := specDiffTypeIsI
		if specDiffTypeIsIS != "" {
			specDiffTypeIsIR = append(specDiffTypeIsIR, specDiffTypeIsIS)
		}
	}

	specDiffTypeIs := swag.JoinByFormat(specDiffTypeIsIR, "")

	if len(specDiffTypeIs) > 0 {
		qsv := specDiffTypeIs[0]
		if qsv != "" {
			qs.Set("specDiffType[is]", qsv)
		}
	}

	var specContainsIR []string
	for _, specContainsI := range o.SpecContains {
		specContainsIS := specContainsI
		if specContainsIS != "" {
			specContainsIR = append(specContainsIR, specContainsIS)
		}
	}

	specContains := swag.JoinByFormat(specContainsIR, "")

	if len(specContains) > 0 {
		qsv := specContains[0]
		if qsv != "" {
			qs.Set("spec[contains]", qsv)
		}
	}

	var specEndQ string
	if o.SpecEnd != nil {
		specEndQ = *o.SpecEnd
	}
	if specEndQ != "" {
		qs.Set("spec[end]", specEndQ)
	}

	var specIsNotIR []string
	for _, specIsNotI := range o.SpecIsNot {
		specIsNotIS := specIsNotI
		if specIsNotIS != "" {
			specIsNotIR = append(specIsNotIR, specIsNotIS)
		}
	}

	specIsNot := swag.JoinByFormat(specIsNotIR, "")

	if len(specIsNot) > 0 {
		qsv := specIsNot[0]
		if qsv != "" {
			qs.Set("spec[isNot]", qsv)
		}
	}

	var specIsIR []string
	for _, specIsI := range o.SpecIs {
		specIsIS := specIsI
		if specIsIS != "" {
			specIsIR = append(specIsIR, specIsIS)
		}
	}

	specIs := swag.JoinByFormat(specIsIR, "")

	if len(specIs) > 0 {
		qsv := specIs[0]
		if qsv != "" {
			qs.Set("spec[is]", qsv)
		}
	}

	var specStartQ string
	if o.SpecStart != nil {
		specStartQ = *o.SpecStart
	}
	if specStartQ != "" {
		qs.Set("spec[start]", specStartQ)
	}

	startTimeQ := o.StartTime.String()
	if startTimeQ != "" {
		qs.Set("startTime", startTimeQ)
	}

	var statusCodeGteQ string
	if o.StatusCodeGte != nil {
		statusCodeGteQ = *o.StatusCodeGte
	}
	if statusCodeGteQ != "" {
		qs.Set("statusCode[gte]", statusCodeGteQ)
	}

	var statusCodeIsNotIR []string
	for _, statusCodeIsNotI := range o.StatusCodeIsNot {
		statusCodeIsNotIS := statusCodeIsNotI
		if statusCodeIsNotIS != "" {
			statusCodeIsNotIR = append(statusCodeIsNotIR, statusCodeIsNotIS)
		}
	}

	statusCodeIsNot := swag.JoinByFormat(statusCodeIsNotIR, "")

	if len(statusCodeIsNot) > 0 {
		qsv := statusCodeIsNot[0]
		if qsv != "" {
			qs.Set("statusCode[isNot]", qsv)
		}
	}

	var statusCodeIsIR []string
	for _, statusCodeIsI := range o.StatusCodeIs {
		statusCodeIsIS := statusCodeIsI
		if statusCodeIsIS != "" {
			statusCodeIsIR = append(statusCodeIsIR, statusCodeIsIS)
		}
	}

	statusCodeIs := swag.JoinByFormat(statusCodeIsIR, "")

	if len(statusCodeIs) > 0 {
		qsv := statusCodeIs[0]
		if qsv != "" {
			qs.Set("statusCode[is]", qsv)
		}
	}

	var statusCodeLteQ string
	if o.StatusCodeLte != nil {
		statusCodeLteQ = *o.StatusCodeLte
	}
	if statusCodeLteQ != "" {
		qs.Set("statusCode[lte]", statusCodeLteQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAPIUsageHitCountURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAPIUsageHitCountURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAPIUsageHitCountURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAPIUsageHitCountURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAPIUsageHitCountURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAPIUsageHitCountURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

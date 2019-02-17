/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

type serviceMutator struct {
	client  client.Client
	decoder types.Decoder
}

var _ admission.Handler = &serviceMutator{}

func (a *serviceMutator) Handle(ctx context.Context, req types.Request) types.Response {
	service := &corev1.Service{}

	err := a.decoder.Decode(req, service)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}

	err = a.mutateServicesFn(ctx, service)
	if err != nil {
		return admission.ErrorResponse(http.StatusInternalServerError, err)
	}

	marshaledService, err := json.Marshal(service)
	if err != nil {
		return admission.ErrorResponse(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.AdmissionRequest.Object.Raw, marshaledService)
}

func (a *serviceMutator) mutateServicesFn(ctx context.Context, service *corev1.Service) error {
	if  service.Annotations["loadbalancer-source-ranger-mutating-admission-webhook"] != "false" &&
		service.Spec.Type == corev1.ServiceTypeLoadBalancer {

		ipAddrs := strings.Fields(os.Getenv("ALLOW_IP_ADDRESSES"))
		service.Spec.LoadBalancerSourceRanges = ipAddrs
		service.Annotations["loadbalancer-source-ranger-mutating-admission-webhook"] = "true"
	}

	return nil
}

var _ inject.Client = &serviceMutator{}

func (v *serviceMutator) InjectClient(c client.Client) error {
	v.client = c
	return nil
}

var _ inject.Decoder = &serviceMutator{}

func (v *serviceMutator) InjectDecoder(d types.Decoder) error {
	v.decoder = d
	return nil
}

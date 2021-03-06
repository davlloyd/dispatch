///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package clients

// NO TEST

import (
	entitystore "github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/service-manager/entities"
)

// BrokerClient defines the Service Broker interface.  This interface very closely resembles Open Service Broker API
type BrokerClient interface {
	ListServiceClasses() ([]entitystore.Entity, error)
	ListServiceInstances() ([]entitystore.Entity, error)
	ListServiceBindings() ([]entitystore.Entity, error)
	CreateService(*entities.ServiceClass, *entities.ServiceInstance) error
	CreateBinding(*entities.ServiceInstance, *entities.ServiceBinding) error
	DeleteService(*entities.ServiceInstance) error
	DeleteBinding(*entities.ServiceBinding) error
}

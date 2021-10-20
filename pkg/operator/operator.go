package operator

import (
	"context"
	"fmt"
	poctrl "rhobs/monitoring-stack-operator/pkg/controllers/prometheus-operator"

	"sigs.k8s.io/controller-runtime/pkg/client"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Operator struct {
	manager manager.Manager
}

func New(metricsAddr string, poOpts poctrl.Options) (*Operator, error) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             NewScheme(),
		MetricsBindAddress: metricsAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create manager: %w", err)
	}

	if err := poctrl.RegisterWithManager(mgr, poOpts); err != nil {
		return nil, fmt.Errorf("unable to register prometheus-operator controller: %w", err)
	}

	return &Operator{
		manager: mgr,
	}, nil
}

func (o *Operator) Start(ctx context.Context) error {
	if err := o.manager.Start(ctx); err != nil {
		return fmt.Errorf("unable to start manager: %w", err)
	}

	return nil
}

func (o *Operator) GetClient() client.Client {
	return o.manager.GetClient()
}

package manifests

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/resolved"
)

func TestNewApplicationSyncPolicyEnabled(t *testing.T) {
	app := resolved.Application{
		Application: catalog.Application{
			Name: "payments",
		},
		Environment: catalog.Environment{
			Project:        "platform",
			TargetRevision: "main",
			Namespace:      "payments",
			SyncPolicy:     "enabled",
		},
	}

	manifest := NewApplication(app)

	if manifest.Spec.SyncPolicy == nil {
		t.Fatal("expected sync policy")
	}

	if manifest.Spec.SyncPolicy.Automated == nil {
		t.Fatal("expected automated sync policy")
	}

	if !manifest.Spec.SyncPolicy.Automated.Prune {
		t.Fatal("expected prune to be enabled")
	}

	if !manifest.Spec.SyncPolicy.Automated.SelfHeal {
		t.Fatal("expected selfHeal to be enabled")
	}
}

func TestNewApplicationSyncPolicyDisabled(t *testing.T) {
	app := resolved.Application{
		Application: catalog.Application{
			Name: "payments",
		},
		Environment: catalog.Environment{
			SyncPolicy: "disabled",
		},
	}

	manifest := NewApplication(app)

	if manifest.Spec.SyncPolicy != nil {
		t.Fatal("expected no sync policy")
	}
}

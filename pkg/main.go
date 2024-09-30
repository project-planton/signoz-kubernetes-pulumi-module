package pkg

import (
	signozkubernetesv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/apis/provider/kubernetes/signozkubernetes/v1"
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *signozkubernetesv1.SignozKubernetesStackInput) error {
	//create kubernetes-provider from the credential in the stack-input
	_, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	return nil
}

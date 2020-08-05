package image

import (
	"errors"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/worker"
)

var ErrUnsupportedResourceType = errors.New("unsupported resource type")

type imageFactory struct {
}

func NewImageFactory() worker.ImageFactory {
	return &imageFactory{}
}

func (f *imageFactory) GetImage(
	logger lager.Logger,
	worker worker.Worker,
	volumeClient worker.VolumeClient,
	imageSpec worker.ImageSpec,
	teamID int,
) (worker.Image, error) {
	if imageSpec.ImageArtifactSource != nil {
		artifactVolume, existsOnWorker, err := imageSpec.ImageArtifactSource.ExistsOn(logger, worker)
		if err != nil {
			logger.Error("failed-to-check-if-volume-exists-on-worker", err)
			return nil, err
		}

		if existsOnWorker {
			return &imageProvidedByPreviousStepOnSameWorker{
				artifactVolume: artifactVolume,
				imageSpec:      imageSpec,
				teamID:         teamID,
				volumeClient:   volumeClient,
			}, nil
		}

		return &imageProvidedByPreviousStepOnDifferentWorker{
			imageSpec:    imageSpec,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	if imageSpec.BaseResourceType != "" {
		return &imageFromBaseResourceType{
			worker:           worker,
			resourceTypeName: imageSpec.BaseResourceType,
			teamID:           teamID,
			volumeClient:     volumeClient,
		}, nil
	}

	if imageSpec.ImageURL != "" {
		return &imageFromRootfsURI{
			url: imageSpec.ImageURL,
		}, nil
	}

	return nil, errors.New("no image specified")
}

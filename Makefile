PROJECT := github.com/raphael-trzpit/poll

prepare-deploy:
	mkdir -p src/$(PROJECT)
	cp -r vendor/* src/ | true
	cp -r internal/ src/$(PROJECT)/internal/
	cp -r cmd/test/* src/

clean:
	rm -rf src tmp

build-make-builder:
	mkdir tmp
	cd tmp && git clone https://github.com/GoogleCloudPlatform/cloud-builders-community
	cd tmp/cloud-builders-community/make && gcloud builds submit --config cloudbuild.yaml .
	rm -rf tmp

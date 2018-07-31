PROJECT := github.com/raphael-trzpit/poll

build-deploy:
	mkdir -p tmp-release/$(PROJECT)
	cp -r vendor/* tmp-release/
	cp -r internal/ tmp-release/$(PROJECT)/internal/
	cp cmd/test/* tmp-release/
	cd tmp-release ; zip -r ../release.zip .
	rm -rf tmp-release/

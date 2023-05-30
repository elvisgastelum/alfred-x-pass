build: dist dist/pass-x-alfred.alfredworkflow

dist:
	[ -d dist ] || mkdir -p dist

dist/pass-x-alfred.alfredworkflow: info.plist assets/key-icon.png assets/not-found-icon.png build/filter scripts/pass-show.sh scripts/pass-generate.sh scripts/pass-otp.sh
	zip $@ $^

build/filter: src/alfred-x-pass.go
	go build -o $@ $^

zip: info.plist build/filter scripts/pass-show.sh scripts/pass-generate.sh scripts/pass-otp.sh
	zip dist/pass-x-alfred.zip $^

dev: build
	open dist/pass-x-alfred.alfredworkflow

clean:
	rm -rf dist build

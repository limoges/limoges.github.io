
build:
	hugo

publish: clean build 

clean:
	rm -fdr ./docs

module github.com/songjiang951130/gostudy

go 1.12

require (
	github.com/nsf/gocode v0.0.0-20190302080247-5bee97b48836 // indirect
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	golang.org/x/crypto v0.0.0-20190418165655-df01cb2cc480 // indirect
	golang.org/x/net v0.0.0-20190419010253-1f3472d942ba // indirect
	golang.org/x/sys v0.0.0-20190418153312-f0ce4c0180be // indirect
	golang.org/x/tools v0.0.0-20190418235243-4796d4bd3df0 // indirect
	rsc.io/quote v1.5.2
)

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.37.4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net => github.com/golang/net v0.0.0-20190419010253-1f3472d942ba
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190418153312-f0ce4c0180be
	golang.org/x/text => github.com/golang/text v0.3.0

	google.golang.org/appengine => github.com/golang/appengine v1.5.0
)

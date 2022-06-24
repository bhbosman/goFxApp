module github.com/bhbosman/goFxApp

go 1.18

require (
	github.com/bhbosman/goUi     v0.0.0-00010101000000-000000000000
	github.com/bhbosman/gocommon v0.0.0-20220621055214-3b04298a9d45
	github.com/cskr/pubsub v1.0.2
	github.com/golang/mock v1.4.4
	go.uber.org/fx v1.14.2
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/bhbosman/goerrors v0.0.0-20210201065523-bb3e832fa9ab // indirect
	github.com/icza/gox v0.0.0-20220321141217-e2d488ab2fbc // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.12.0 // indirect
	golang.org/x/sys v0.0.0-20220318055525-2edf467146b5 // indirect
)

replace github.com/bhbosman/gocommon => ../gocommon
replace github.com/bhbosman/goUi    ==> ../goUi
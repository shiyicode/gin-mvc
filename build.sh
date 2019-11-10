#!/bin/bash

workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace

app=gin-framework
cfg=./cfg

function build() {
	export GOPATH="$GOPATH:$(pwd)"

	echo "`go version`"
	go build -o $app
	if [[ $running -gt 0 ]]; then
        	echo "failure to go build"
        	exit 0
	fi

}

function make_output {
	local output="./output"
	rm -rf $output &>/dev/null
	mkdir -p $output &>/dev/null
	(
		cp -vrf $app $output &&       # 拷贝 二进制文件 至output目录
		cp -vrf $cfg $output &&       # 拷贝 cfg配置文件目录 至output目录
		cp -vrf control.sh $output && # 拷贝 启停脚本 至output目录
		cp -vrf scripts $output &&    # 拷贝 脚本目录 至output目录
		echo -e "make output ok"
	) || { echo -e "make output error"; rm -rf "./output"; exit 2; }
}

build

make_output

echo "build down"
exit 0

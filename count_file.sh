#!/bin/bash

function count_files() {
	cur_dir=$1
	count=0
	files=`ls $cur_dir`
	for element in $files
	do
		if [ -f $cur_dir/$element ];then
			count=`expr $count + 1`
		elif [ -d $cur_dir/$element ];then
			subCount=`count_files $cur_dir/$element`
			count=`expr $count + $subCount`
		fi
	done
	echo $count
}

count_files $1
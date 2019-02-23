#!/bin/bash

function count_files() {
	count=0
	files=`ls $1`
	for element in $files 
	do
		if [ -f $element ];then
			count=`expr $count + 1`
		elif [ -d $element ];then
			subCount=`count_files ./$element`
			count=`expr $count + $subCount`
		fi
	done
	echo $count
}


count_files ./
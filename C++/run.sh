if [ $# -gt 0 ]
then
  day=$(printf "%02d" $1)
else
  day=-1
fi

if test -f aoc2025;
then
	rm -rf aoc2025
fi

if [ $day -gt 0 ]
then
	if test -f Day$day.cpp;
	then
		echo "compiling single day..."
		g++ *.cpp -O2 -o aoc2025 -std=c++20 -DDAY$day -Wall -Wextra -Wunused-variable
		echo "done."
		./aoc2025
		rm -rf aoc2025
	else
		echo "Day" $day "does not exist"
	fi
else
	missing=""
	COMPILERSTRING=" "
	for i in $(seq -f "%02g" 1 12)
	do
		if test -f Day$i.cpp;
		then
			COMPILERSTRING+="-DDAY$i "
		else
			if [ "$missing" = "" ]
			then
				missing=$i
			else
				missing=$missing","$i
			fi
		fi
	done
	g++ *.cpp -O2 -o aoc2025 -std=c++20 $COMPILERSTRING
	./aoc2025
	rm -rf aoc2025
	if [ "$missing" != "" ]
	then
		echo "Missing days = ["$missing"]"
	fi
fi

for f in *.txt
do
	echo "Processing $f file..."
	cat $f | tr -d '[:punct:]' | tr -d '[:digit:]' | tr -d '[«»]' | tr -s ' ' > c_$f
done

